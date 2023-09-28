//
// example code to illustrate sending OTel traces to Cloud Observability directly via OTLP gRPC
//
// usage:
//   export LS_ACCESS_TOKEN=<YOUR_LS_ACCESS_TOKEN>
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
	lsToken        = os.Getenv("LS_ACCESS_TOKEN")
	lsEnvironment  = os.Getenv("LS_ENVIRONMENT")
	targetURL      = os.Getenv("DESTINATION_URL")
)

func newExporter(ctx context.Context) (*otlptrace.Exporter, error) {

	if len(endpoint) == 0 {
		endpoint = "ingest.lightstep.com:443"
		log.Printf("Using default LS endpoint %s", endpoint)
	}

	if len(lsToken) == 0 {
		log.Fatalf("Cloud Observability token missing. Please set environment variable LS_ACCESS_TOKEN")
	}

	var headers = map[string]string{
		"lightstep-access-token": lsToken,
	}

	client := otlptracegrpc.NewClient(
		otlptracegrpc.WithHeaders(headers),
		otlptracegrpc.WithEndpoint(endpoint),
	)
	return otlptrace.New(ctx, client)
}

func newTraceProvider(exp *otlptrace.Exporter) *sdktrace.TracerProvider {

	if len(serviceName) == 0 {
		serviceName = "test-go-client-grpc"
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

	span.AddEvent("Did some cool stuff")
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

	span.AddEvent("Cancelled wait due to external signal", trace.WithAttributes(attribute.Int("pid", 4328), attribute.String("signal", "SIGHUP")))
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

	for {
		makeRequest(ctx)
		time.Sleep(1 * time.Second)
	}

}
