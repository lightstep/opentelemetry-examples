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

	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/propagation"
	"go.opentelemetry.io/otel/exporters/otlp"

	"go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace"
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

	secureOption := otlp.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	if len(insecure) > 0 {
		secureOption = otlp.WithInsecure()
	}

	exporter, err := otlp.NewExporter(
		secureOption,
		otlp.WithAddress(url),
		otlp.WithHeaders(headers),
	)

	if err != nil {
		log.Fatal(err)
	}
	return exporter
}

func initTracer() {
	b3 := b3.B3{}
	// Register the B3 propagator globally.
	global.SetPropagators(propagation.New(
		propagation.WithExtractors(b3),
		propagation.WithInjectors(b3),
	))
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

	resources := resource.New(
		label.String("service.name", componentName),
		label.String("service.version", serviceVersion),
		label.String("library.language", "go"),
		label.String("library.version", "1.2.3"),
	)
	tp := trace.NewTracerProvider(
		trace.WithConfig(trace.Config{DefaultSampler: trace.AlwaysSample()}),
		trace.WithSyncer(exporter),
		trace.WithResource(resources),
	)
	global.SetTracerProvider(tp)
}

func makeRequest() {
	client := http.DefaultClient
	tracer := global.Tracer("otel-example/client")
	ctx, span := tracer.Start(context.Background(), "makeRequest")
	defer span.End()

	req, _ := http.NewRequest("GET", targetURL, nil)
	ctx, req = otelhttptrace.W3C(ctx, req)
	otelhttptrace.Inject(ctx, req)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
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
