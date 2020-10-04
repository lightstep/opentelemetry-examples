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
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/instrumentation/httptrace"
)

var (
	destinationURL = os.Getenv("DESTINATION_URL")
)

func makeRequest() {
	client := http.DefaultClient
	tracer := global.Tracer("otel-example/client")
	tracer.WithSpan(context.Background(), "makeRequest", func(ctx context.Context) error {
		req, _ := http.NewRequest("GET", destinationURL, nil)
		httptrace.Inject(ctx, req)
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		defer res.Body.Close()
		fmt.Printf("Request to %s, got %d bytes\n", destinationURL, res.ContentLength)
		return nil
	})
}

func main() {
	otel := launcher.ConfigureOpentelemetry()
	defer otel.Shutdown()
	if len(destinationURL) == 0 {
		destinationURL = "http://localhost:8081"
	}
	for {
		makeRequest()
		time.Sleep(1 * time.Second)
	}

}
