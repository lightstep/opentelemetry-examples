//
// example code to test lightstep-tracer-go
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
	mathrand "math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/lightstep/lightstep-tracer-go"
	"github.com/opentracing/opentracing-go"
)

var lsToken = os.Getenv("LS_ACCESS_TOKEN")
var lsMetricsURL = os.Getenv("LS_METRICS_URL")
var targetURL = os.Getenv("TARGET_URL")

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

	componentName := os.Getenv("LIGHTSTEP_COMPONENT_NAME")
	if len(componentName) == 0 {
		componentName = "test-go-client"
	}
	serviceVersion := os.Getenv("LIGHTSTEP_SERVICE_VERSION")
	if len(serviceVersion) == 0 {
		serviceVersion = "0.0.0"
	}
	endpoint := lightstep.Endpoint{Host: host, Port: port, Plaintext: plaintext}
	opentracing.InitGlobalTracer(lightstep.NewTracer(lightstep.Options{
		AccessToken: lsToken,
		Collector:   endpoint,
		UseHttp:     true,
		Tags: opentracing.Tags{
			"lightstep.component_name": componentName,
			"service.version":          serviceVersion,
		},
		SystemMetrics: lightstep.SystemMetricsOptions{
			Endpoint: endpoint,
		},
	}))
}

func makeRequest() {
	trivialSpan, _ := opentracing.StartSpanFromContext(context.Background(), "makeRequest")
	defer trivialSpan.Finish()

	contentLength := mathrand.Intn(2048)
	url := fmt.Sprintf("%s/content/%d", targetURL, contentLength)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("Request to %s, got %d bytes\n", url, res.ContentLength)
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
