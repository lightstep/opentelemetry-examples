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
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/propagation"
	"go.opentelemetry.io/otel/exporters/otlp"
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
		componentName = "test-go-server"
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
