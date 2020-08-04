//
// example code to test lightstep/opentelemetry-exporter-go
//
// usage:
//   LS_ACCESS_TOKEN=${SECRET_TOKEN} \
//   LIGHTSTEP_COMPONENT_NAME=demo-client-go \
//   LIGHTSTEP_SERVICE_VERSION=0.1.8 \
//   go run client.go

package main

import (
	"context"
	"fmt"
	"log"
	mathrand "math/rand"
	"net/http"
	"os"
	"time"

	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/kv"
	"go.opentelemetry.io/otel/exporters/otlp"
	"go.opentelemetry.io/otel/instrumentation/httptrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc/credentials"
)

var (
	componentName  = os.Getenv("LS_SERVICE_NAME")
	serviceVersion = os.Getenv("LS_SERVICE_VERSION")
	lsToken        = os.Getenv("LS_ACCESS_TOKEN")
	collectorURL   = os.Getenv("LS_SATELLITE_URL")
	targetURL      = os.Getenv("TARGET_URL")
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
		kv.String("service.name", componentName),
		kv.String("service.version", serviceVersion),
		kv.String("library.language", "go"),
		kv.String("library.version", "1.2.3"),
	)
	tp, err := trace.NewProvider(
		trace.WithConfig(trace.Config{DefaultSampler: trace.AlwaysSample()}),
		trace.WithSyncer(exporter),
		trace.WithResource(resources),
	)
	if err != nil {
		log.Fatal(err)
	}
	global.SetTraceProvider(tp)
}

func makeRequest() {
	client := http.DefaultClient
	tracer := global.Tracer("otel-example/client")
	tracer.WithSpan(context.Background(), "makeRequest", func(ctx context.Context) error {
		contentLength := mathrand.Intn(2048)
		url := fmt.Sprintf("%s/content/%d", targetURL, contentLength)
		req, _ := http.NewRequest("GET", url, nil)
		ctx, req = httptrace.W3C(ctx, req)
		httptrace.Inject(ctx, req)
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		defer res.Body.Close()
		fmt.Printf("Request to %s, got %d bytes\n", url, res.ContentLength)
		return nil
	})
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
