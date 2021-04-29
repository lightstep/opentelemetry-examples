//
// example code to test lightstep/opentelemetry-exporter-go
//
// usage:
//   LS_ACCESS_TOKEN=${SECRET_TOKEN} \
//   LS_SERVICE_NAME=demo-server-go \
//   LS_SERVICE_VERSION=0.1.8 \
//   go run server.go

package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	muxtrace "go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpgrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc/credentials"
)

var (
	componentName  = os.Getenv("LS_SERVICE_NAME")
	serviceVersion = os.Getenv("LS_SERVICE_VERSION")
	lsToken        = os.Getenv("LS_ACCESS_TOKEN")
	collectorURL   = os.Getenv("LS_SATELLITE_URL")
	insecure       = os.Getenv("LS_INSECURE")
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

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
	b3 := b3.B3{}
	// Register the B3 propagator globally.
	otel.SetTextMapPropagator(b3)

	if len(collectorURL) == 0 {
		collectorURL = "localhost:4317"
	}

	if len(componentName) == 0 {
		componentName = "test-go-server"
	}

	if len(serviceVersion) == 0 {
		serviceVersion = "0.0.0"
	}

	exporter := initExporter(collectorURL, lsToken)

	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", componentName),
			attribute.String("service.version", serviceVersion),
			attribute.String("library.language", "go"),
			attribute.String("library.version", "1.2.3"),
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

func main() {
	initTracer()
	fmt.Printf("Starting server on http://localhost:8081\n")
	r := mux.NewRouter()
	// re-enable once the new version of otel-go and otel-go-contrib is released
	r.Use(muxtrace.Middleware(componentName))
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		length := rand.Intn(1024)
		log.Printf("%s %s %s", r.Method, r.URL.Path, r.Proto)
		fmt.Fprintf(w, randString(length))
	})
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
