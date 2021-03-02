//
// example code to test lightstep/opentelemetry-exporter-go
//
// usage:
//   LS_ACCESS_TOKEN=${SECRET_TOKEN} \
//   LS_SERVICE_NAME=demo-client-go \
//   LS_SERVICE_VERSION=0.1.8 \
//   go run client.go

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/contrib/propagators/ot"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpgrpc"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc/credentials"
)

var (
	componentName  = os.Getenv("LS_SERVICE_NAME")
	serviceVersion = os.Getenv("LS_SERVICE_VERSION")
	lsToken        = os.Getenv("LS_ACCESS_TOKEN")
	collectorURL   = os.Getenv("LS_SATELLITE_URL")
	targetURL      = os.Getenv("DESTINATION_URL")
	insecure       = os.Getenv("LS_INSECURE")
)

func initExporter(url string, token string) *otlp.Exporter {
	headers := map[string]string{
		"lightstep-access-token": token,
	}

	secureOption := otlpgrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	if len(insecure) > 0 {
		secureOption = otlpgrpc.WithInsecure()
	}

	exporter, err := otlp.NewExporter(
		context.Background(),
		otlpgrpc.NewDriver(
			secureOption,
			otlpgrpc.WithEndpoint(url),
			otlpgrpc.WithHeaders(headers),
		),
	)

	if err != nil {
		log.Fatal(err)
	}
	return exporter
}

func initTracer() {
	otPropagator := ot.OT{}
	// Register the OT propagator globally.
	otel.SetTextMapPropagator(otPropagator)
	if len(collectorURL) == 0 {
		collectorURL = "localhost:55680"
	}

	if len(componentName) == 0 {
		componentName = "test-go-client"
	}
	if len(serviceVersion) == 0 {
		serviceVersion = "0.0.0"
	}

	exporter := initExporter(collectorURL, lsToken)

	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			label.String("service.name", componentName),
			label.String("service.version", serviceVersion),
			label.String("library.language", "go"),
			label.String("library.version", "1.2.3"),
		),
	)
	if err != nil {
		log.Printf("Could not set resources: ", err)
	}
	tp := trace.NewTracerProvider(
		trace.WithConfig(trace.Config{DefaultSampler: trace.AlwaysSample()}),
		trace.WithSyncer(exporter),
		trace.WithResource(resources),
	)
	otel.SetTracerProvider(tp)
}

func makeRequest() {
	httpClient := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	tracer := otel.Tracer("otel-example/client")
	_, span := tracer.Start(context.Background(), "makeRequest")
	defer span.End()

	req, _ := http.NewRequest("GET", targetURL, nil)
	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("Request to %s, got %d bytes\n", targetURL, res.ContentLength)
}

func main() {
	initTracer()
	if len(targetURL) == 0 {
		targetURL = "http://localhost:8081"
	}
	for {
		makeRequest()
		time.Sleep(1 * time.Second)
	}

}
