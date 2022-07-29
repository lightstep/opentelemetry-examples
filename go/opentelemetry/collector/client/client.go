//
// example code to illustrate sending OTel traces to Lightstep via the OTel Collector
//
// NOTE: Requires that you run a Collector instance
// (see collector/vanilla/readme.md for details on how to run the
// Collector locally)
//
// usage:
//   go run client.go

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
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
	targetURL      = os.Getenv("DESTINATION_URL")
)

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
		serviceName = "test-go-client-collector"
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
				attribute.String("environment", lsEnvironment),
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

func makeRequest(ctx context.Context) {
	ctx, span := tracer.Start(ctx, "makeRequest")
	defer span.End()

	if len(targetURL) == 0 {
		targetURL = "http://localhost:8081/ping"
		log.Printf("Using default targetURL %s", targetURL)
	}

	span.AddEvent("Making a request")
	res, err := otelhttp.Get(ctx, targetURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("Body : %s", body)
	fmt.Printf("Request to %s, got %d bytes\n", targetURL, res.ContentLength)

	span.SetAttributes(
		attribute.String("response", string(body)),
	)

	span.AddEvent("Made a request", trace.WithAttributes(attribute.String("greeting", "Hello"), attribute.String("farewell", "Bye")))
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

	// Register context and baggage propagation.
	// Although not strictly necessary, for this sample,
	// it is required for distributed tracing.
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	tracer = tp.Tracer(serviceName, trace.WithInstrumentationVersion(serviceVersion))

	for {
		makeRequest(ctx)
		time.Sleep(1 * time.Second)
	}

}
