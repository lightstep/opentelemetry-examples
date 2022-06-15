//
// example code to test lightstep/opentelemetry-exporter-go
//
// usage:
//   go run client.go

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	// "net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.opentelemetry.io/otel/trace"
)

var (
	tracer         trace.Tracer
	serviceName    string = "test-go-client"
	serviceVersion string = "0.1.0"
	collectorAddr  string = "localhost:4318" // HTTP endpoint for collector
	targetURL      string = "http://localhost:8081/ping"
)

func newTraceProvider(ctx context.Context) *sdktrace.TracerProvider {
	exporter, err :=
		otlptracehttp.New(ctx,
			// WithInsecure lets us use http instead of https.
			// This is just for local development.
			otlptracehttp.WithInsecure(),
			otlptracehttp.WithEndpoint(collectorAddr),
		)

	if err != nil {
		panic(err)
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
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource),
	)
}

func makeRequest(ctx context.Context) {
	res, err := otelhttp.Get(ctx, targetURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("Body : %s", body)
	fmt.Printf("Request to %s, got %d bytes\n", targetURL, res.ContentLength)

	_, span := tracer.Start(ctx, "makeRequest")
	defer span.End()

	span.SetAttributes(
		attribute.String("response", string(body)),
	)

}

func main() {
	ctx := context.Background()

	tp := newTraceProvider(ctx)
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
