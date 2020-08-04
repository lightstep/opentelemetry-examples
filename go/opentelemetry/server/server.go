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
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"

	ls "github.com/lightstep/opentelemetry-exporter-go/lightstep"
	muxtrace "go.opentelemetry.io/contrib/instrumentation/gorilla/mux"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/sdk/trace"
)

var (
	lsToken        = os.Getenv("LS_ACCESS_TOKEN")
	lsMetricsURL   = os.Getenv("LS_METRICS_URL")
	componentName  = os.Getenv("LIGHTSTEP_COMPONENT_NAME")
	serviceVersion = os.Getenv("LIGHTSTEP_SERVICE_VERSION")
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
		componentName = "test-go-server"
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

func main() {
	initLightstepTracer()
	fmt.Printf("Starting server on http://localhost:8081\n")
	r := mux.NewRouter()
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
