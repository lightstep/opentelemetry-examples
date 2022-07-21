//
// example code to test lightstep/opentelemetry-exporter-go
//
// usage:
//	 export OTEL_LOG_LEVEL=debug
//	 export LS_ACCESS_TOKEN="<your_access_token>"
//   go run client.go

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	// "net/http"
	"log"
	"os"
	"time"

	"github.com/lightstep/otel-launcher-go/launcher"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.opentelemetry.io/otel/trace"
	// "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	// "go.opentelemetry.io/otel"
	// "go.opentelemetry.io/otel/attribute"
	// "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	// "go.opentelemetry.io/otel/propagation"
	// "go.opentelemetry.io/otel/sdk/resource"
	// sdktrace "go.opentelemetry.io/otel/sdk/trace"
	// semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	// "go.opentelemetry.io/otel/trace"
)

var (
	tracer         trace.Tracer
	serviceName           = os.Getenv("LS_SERVICE_NAME")
	serviceVersion        = os.Getenv("LS_SERVICE_VERSION")
	endpoint              = os.Getenv("LS_SATELLITE_URL")
	lsToken               = os.Getenv("LS_ACCESS_TOKEN")
	targetURL      string = "http://localhost:8081/ping"
)

func newLauncher() launcher.Launcher {
	if len(endpoint) == 0 {
		endpoint = "ingest.lightstep.com:443"
		// endpoint = "0.0.0.0:4317"	// Use for Collector
		log.Printf("Using default LS endpoint %s", endpoint)
	}

	if len(serviceName) == 0 {
		serviceName = "test-go-client-launcher"
		log.Printf("Using default service name %s", serviceName)
	}

	if len(serviceVersion) == 0 {
		serviceVersion = "0.1.0"
		log.Printf("Using default service version %s", serviceVersion)
	}

	// if len(lsToken) == 0 {
	// 	log.Fatalf("Lightstep token missing. Please set environment variable LS_ENVIRONMENT")
	// }

	otelLauncher := launcher.ConfigureOpentelemetry(
		launcher.WithServiceName(serviceName),
		launcher.WithServiceVersion(serviceVersion),
		// launcher.WithAccessToken(lsToken),
		// launcher.WithSpanExporterInsecure(true),		// Use for Collector
		launcher.WithSpanExporterEndpoint(endpoint),
		launcher.WithMetricExporterEndpoint(endpoint),
		// launcher.WithMetricExporterInsecure(true),	// Use for Collector
		launcher.WithPropagators([]string{"tracecontext", "baggage"}),
		launcher.WithResourceAttributes(map[string]string{
			string(semconv.ContainerNameKey): "my-container-name",
		}),
	)

	return otelLauncher
}

func makeRequest(ctx context.Context) {
	ctx, span := tracer.Start(ctx, "makeRequest")
	defer span.End()

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
	otelLauncher := newLauncher()
	defer otelLauncher.Shutdown()

	tracer = otel.Tracer(serviceName)

	// ctx := context.Background()

	// tp := newTraceProvider(ctx)
	// defer func() { _ = tp.Shutdown(ctx) }()

	// otel.SetTracerProvider(tp)

	// // Register context and baggage propagation.
	// // Although not strictly necessary, for this sample,
	// // it is required for distributed tracing.
	// otel.SetTextMapPropagator(
	// 	propagation.NewCompositeTextMapPropagator(
	// 		propagation.TraceContext{},
	// 		propagation.Baggage{},
	// 	),
	// )

	// tracer = tp.Tracer(serviceName, trace.WithInstrumentationVersion(serviceVersion))

	for {
		makeRequest(ctx)
		time.Sleep(1 * time.Second)
	}

}
