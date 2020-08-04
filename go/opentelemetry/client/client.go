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
	"net/url"
	"os"
	"strconv"
	"time"

	ls "github.com/lightstep/opentelemetry-exporter-go/lightstep"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/sdk/trace"
)

var (
	lsToken        = os.Getenv("LS_ACCESS_TOKEN")
	lsMetricsURL   = os.Getenv("LS_METRICS_URL")
	targetURL      = os.Getenv("TARGET_URL")
	componentName  = os.Getenv("LIGHTSTEP_COMPONENT_NAME")
	serviceVersion = os.Getenv("LIGHTSTEP_SERVICE_VERSION")
)

func initLightstepTracer() {
	u, err := url.Parse(lsMetricsURL)

	host := "ingest.lightstep.com"
	port := 443
	plaintext := false

	if err == nil {
		host = u.Hostname()
		port, _ = strconv.Atoi(u.Port())
		if u.Scheme == "http" {
			plaintext = true
		}
	}

	if len(componentName) == 0 {
		componentName = "test-go-client"
	}
	if len(serviceVersion) == 0 {
		serviceVersion = "0.0.0"
	}

	exporter, err := ls.NewExporter(
		ls.WithAccessToken(lsToken),
		ls.WithHost(host),
		ls.WithPort(port),
		ls.WithPlainText(plaintext),
		ls.WithServiceName(componentName),
		ls.WithServiceVersion(serviceVersion),
	)
	if err != nil {
		log.Fatal(err)
	}
	tp, err := trace.NewProvider(trace.WithConfig(trace.Config{DefaultSampler: trace.AlwaysSample()}),
		trace.WithSyncer(exporter))
	if err != nil {
		log.Fatal(err)
	}
	global.SetTraceProvider(tp)
}

func makeRequest() {
	tracer := global.Tracer("otel-example/client")
	tracer.WithSpan(context.Background(), "makeRequest", func(ctx context.Context) error {
		contentLength := mathrand.Intn(2048)
		url := fmt.Sprintf("%s/content/%d", targetURL, contentLength)
		res, err := http.Get(url)
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
	initLightstepTracer()
	if len(targetURL) == 0 {
		targetURL = "http://localhost:8081"
	}
	for {
		makeRequest()
		time.Sleep(1 * time.Second)
	}

}
