//
// example code to test lightstep/opentelemetry-exporter-go
//
// usage:
//   LS_ACCESS_TOKEN=${SECRET_TOKEN} \
//   LIGHTSTEP_COMPONENT_NAME=demo-server-go \
//   LIGHTSTEP_SERVICE_VERSION=0.1.8 \
//   go run server.go

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/grpc/credentials"

	// re-enable once the new version of otel-go and otel-go-contrib is released
	muxtrace "go.opentelemetry.io/contrib/instrumentation/gorilla/mux"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/kv"
	"go.opentelemetry.io/otel/exporters/otlp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
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

func main() {
	initTracer()
	fmt.Printf("Starting server on http://localhost:8081\n")
	r := mux.NewRouter()
	// re-enable once the new version of otel-go and otel-go-contrib is released
	r.Use(muxtrace.Middleware(componentName))
	r.HandleFunc("/content/{length:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		length, err := strconv.Atoi(vars["length"])
		if err != nil {
			length = 10
		}

		log.Printf("%s %s %s", r.Method, r.URL.Path, r.Proto)
		fmt.Fprintf(w, randString(length))
	})
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
