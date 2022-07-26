//
// example code to test lightstep/opentelemetry-exporter-go
//
// usage:
//	 export OTEL_LOG_LEVEL=debug
//	 export LS_ACCESS_TOKEN="<your_access_token>"
//   go run server-direct.go

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/lightstep/otel-launcher-go/launcher"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.opentelemetry.io/otel/trace"
)

var (
	tracer         trace.Tracer
	serviceName    = os.Getenv("LS_SERVICE_NAME")
	serviceVersion = os.Getenv("LS_SERVICE_VERSION")
	endpoint       = os.Getenv("LS_SATELLITE_URL")
	lsToken        = os.Getenv("LS_ACCESS_TOKEN")
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func newLauncher() launcher.Launcher {
	if len(endpoint) == 0 {
		endpoint = "ingest.lightstep.com:443"
		log.Printf("Using default LS endpoint %s", endpoint)
	}

	if len(serviceName) == 0 {
		serviceName = "test-go-server-launcher-direct"
		log.Printf("Using default service name %s", serviceName)
	}

	if len(serviceVersion) == 0 {
		serviceVersion = "0.1.0"
		log.Printf("Using default service version %s", serviceVersion)
	}

	if len(lsToken) == 0 {
		log.Fatalf("Lightstep token missing. Please set environment variable LS_ENVIRONMENT")
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

func randString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

// Wrap the handleRollDice so that telemetry data
// can be automatically generated for it
func wrapHandler() {
	handler := http.HandlerFunc(handlePing)
	wrappedHandler := otelhttp.NewHandler(handler, "pingHandler")
	http.Handle("/ping", wrappedHandler)
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	operationName := "ping"
	_, span := tracer.Start(r.Context(), operationName)
	defer span.End()

	length := rand.Intn(1024)
	log.Printf("%s %s %s", r.Method, r.URL.Path, r.Proto)

	pingResult := randString(length)
	span.SetAttributes(
		attribute.String("library.language", "go"),
		attribute.String("library.version", "v1.7.0"),
	)

	// setting span as successful
	span.SetStatus(codes.Ok, "Success")

	// setting span event
	span.AddEvent(fmt.Sprint(r.Header))

	fmt.Fprintf(w, pingResult)
}

func main() {
	otelLauncher := newLauncher()
	defer otelLauncher.Shutdown()

	tracer = otel.Tracer(serviceName)

	wrapHandler()

	fmt.Printf("Starting server on http://localhost:8081\n")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
