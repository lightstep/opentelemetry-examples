//
// example code to illustrate sending OTel traces to Lightstep via the OTel Collector
//
// NOTE: Requires that you run a Collector instance
// (see collector/vanilla/readme.md for details on how to run the
// Collector locally)
//
// usage:
//   go run server.go

package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.opentelemetry.io/otel/trace"
)

var (
	tracer         trace.Tracer
	serviceName    = os.Getenv("LS_SERVICE_NAME")
	serviceVersion = os.Getenv("LS_SERVICE_VERSION")
	endpoint       = os.Getenv("LS_SATELLITE_URL")
	lsEnvironment  = os.Getenv("LS_ENVIRONMENT")
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func newExporter(ctx context.Context) (*otlptrace.Exporter, error) {
	if len(endpoint) == 0 {
		endpoint = "localhost:4317"
		log.Printf("Using default LS endpoint %s", endpoint)
	}

	exporter, err :=
		otlptracegrpc.New(ctx,
			// WithInsecure lets us use http instead of https.
			// This is just for local development.
			otlptracegrpc.WithInsecure(),
			otlptracegrpc.WithEndpoint(endpoint),
		)

	return exporter, err
}

func newTraceProvider(exp *otlptrace.Exporter) *sdktrace.TracerProvider {

	if len(serviceName) == 0 {
		serviceName = "test-go-server-collector"
		log.Printf("Using default service name %s", serviceName)
	}

	if len(serviceVersion) == 0 {
		serviceVersion = "0.1.0"
		log.Printf("Using default service version %s", serviceVersion)
	}

	if len(lsEnvironment) == 0 {
		lsEnvironment = "dev"
		log.Printf("Using default environment %s", lsEnvironment)
	}

	// This includes the following resources:
	//
	// - sdk.language, sdk.version
	// - service.name, service.version, environment
	//
	// Including these resources is a good practice because it is commonly
	// used by various tracing backends to let you more accurately
	// analyze your telemetry data.
	resource, rErr :=
		resource.Merge(
			resource.Default(),
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(serviceName),
				semconv.ServiceVersionKey.String(serviceVersion),
				attribute.String("environment", "getting-started"),
			),
		)

	if rErr != nil {
		panic(rErr)
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource),
	)
}

func randString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

// Wrap the handleRollDice so that telemetry data
// can be automatically generated for it
func wrapHandler() {
	handler := http.HandlerFunc(handlePing)
	wrappedHandler := otelhttp.NewHandler(handler, "pingHandler")
	http.Handle("/ping", wrappedHandler)
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	operationName := "ping"
	_, span := tracer.Start(r.Context(), operationName)
	defer span.End()

	length := rand.Intn(1024)
	log.Printf("%s %s %s", r.Method, r.URL.Path, r.Proto)

	pingResult := randString(length)
	span.SetAttributes(
		attribute.String("library.language", "go"),
		attribute.String("library.version", "v1.7.0"),
	)

	// setting span as successful
	span.SetStatus(codes.Ok, "Success")

	// setting span event
	span.AddEvent(fmt.Sprint(r.Header))

	fmt.Fprintf(w, pingResult)
}

func main() {
	ctx := context.Background()

	exp, err := newExporter(ctx)
	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
	}

	tp := newTraceProvider(exp)
	defer func() { _ = tp.Shutdown(ctx) }()

	otel.SetTracerProvider(tp)

	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	tracer = tp.Tracer(serviceName, trace.WithInstrumentationVersion(serviceVersion))

	wrapHandler()

	fmt.Printf("Starting server on http://localhost:8081\n")
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
