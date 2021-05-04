//
// example code to test lightstep/otel-launcher-go/launcher
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
	"net/http"
	"os"
	"time"

	"github.com/lightstep/otel-launcher-go/launcher"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
)

var (
	destinationURL = os.Getenv("DESTINATION_URL")
)

func makeRequest() {
	httpClient := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	tracer := otel.Tracer("otel-example/client")
	_, span := tracer.Start(context.Background(), "makeRequest")
	defer span.End()
	req, _ := http.NewRequest("GET", destinationURL, nil)

	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("Request to %s, got %d bytes\n", destinationURL, res.ContentLength)
}

func main() {
	otel := launcher.ConfigureOpentelemetry(
		launcher.WithPropagators([]string{"ottrace"}),
	)
	defer otel.Shutdown()
	if len(destinationURL) == 0 {
		destinationURL = "http://localhost:8081"
	}
	for {
		makeRequest()
		time.Sleep(1 * time.Second)
	}

}
