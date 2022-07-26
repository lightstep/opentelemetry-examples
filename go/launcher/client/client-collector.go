//
// example code to test lightstep/opentelemetry-exporter-go
//
// usage:
//	 export OTEL_LOG_LEVEL=debug
//   go run client-collector.go

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/lightstep/otel-launcher-go/launcher"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.opentelemetry.io/otel/trace"
)

var (
	tracer         trace.Tracer
	serviceName           = os.Getenv("LS_SERVICE_NAME")
	serviceVersion        = os.Getenv("LS_SERVICE_VERSION")
	endpoint              = os.Getenv("LS_SATELLITE_URL")
	targetURL      string = "http://localhost:8081/ping"
)

func newLauncher() launcher.Launcher {
	if len(endpoint) == 0 {
		endpoint = "localhost:4317" // Collector endpoint
		log.Printf("Using default LS endpoint %s", endpoint)
	}

	if len(serviceName) == 0 {
		serviceName = "test-go-client-launcher-collector"
		log.Printf("Using default service name %s", serviceName)
	}

	if len(serviceVersion) == 0 {
		serviceVersion = "0.1.0"
		log.Printf("Using default service version %s", serviceVersion)
	}

	otelLauncher := launcher.ConfigureOpentelemetry(
		launcher.WithServiceName(serviceName),
		launcher.WithServiceVersion(serviceVersion),
		launcher.WithSpanExporterInsecure(true), // Use for Collector
		launcher.WithSpanExporterEndpoint(endpoint),
		launcher.WithMetricExporterEndpoint(endpoint),
		launcher.WithMetricExporterInsecure(true), // Use for Collector
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

	for {
		makeRequest(ctx)
		time.Sleep(1 * time.Second)
	}

}
