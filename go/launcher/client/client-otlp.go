//
// example code to illustrate sending OTel traces to Lightstep directly via OTLP
// using the Go Launcher
//
// usage:
//	 export OTEL_LOG_LEVEL=debug
//   export LS_ACCESS_TOKEN=<YOUR_LS_ACCESS_TOKEN>
//   go run client-otlp.go

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
	serviceName    = os.Getenv("LS_SERVICE_NAME")
	serviceVersion = os.Getenv("LS_SERVICE_VERSION")
	endpoint       = os.Getenv("LS_SATELLITE_URL")
	lsToken        = os.Getenv("LS_ACCESS_TOKEN")
	targetURL      = os.Getenv("DESTINATION_URL")
)

func newLauncher() launcher.Launcher {
	if len(endpoint) == 0 {
		endpoint = "ingest.lightstep.com:443"
		log.Printf("Using default LS endpoint %s", endpoint)
	}

	if len(serviceName) == 0 {
		serviceName = "test-go-client-launcher-direct"
		log.Printf("Using default service name %s", serviceName)
	}

	if len(serviceVersion) == 0 {
		serviceVersion = "0.1.0"
		log.Printf("Using default service version %s", serviceVersion)
	}

	if len(lsToken) == 0 {
		log.Fatalf("Lightstep token missing. Please set environment variable LS_ACCESS_TOKEN")
	}

	otelLauncher := launcher.ConfigureOpentelemetry(
		launcher.WithServiceName(serviceName),
		launcher.WithServiceVersion(serviceVersion),
		launcher.WithAccessToken(lsToken),
		launcher.WithSpanExporterEndpoint(endpoint),
		launcher.WithMetricExporterEndpoint(endpoint),
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
	otelLauncher := newLauncher()
	defer otelLauncher.Shutdown()

	tracer = otel.Tracer(serviceName)

	for {
		makeRequest(ctx)
		time.Sleep(1 * time.Second)
	}

}
