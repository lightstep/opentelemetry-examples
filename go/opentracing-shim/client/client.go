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
	otShim "go.opentelemetry.io/otel/bridge/opentracing"
	ot "github.com/opentracing/opentracing-go"
)

var (
	destinationURL = os.Getenv("DESTINATION_URL")
)

func makeRequest() {
	httpClient := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	// create OTEL tracer which we  then bind to the OT shim
	otelTracer := otel.Tracer("otel-ot-bridge-example/client")
	bridge, wrapper := otShim.NewTracerPair(otelTracer)

	// register BridgeTracer with opentracing and WrapperTracerProvider with opentelemetry.
	otel.SetTracerProvider(wrapper)
	ot.SetGlobalTracer(bridge)

	// start a span using OT which will pass through to the OTEL tracer
	span, _ := ot.StartSpanFromContext(context.Background(), "makeRequest")
	defer span.Finish()

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
