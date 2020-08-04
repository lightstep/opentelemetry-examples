//
// example code to test lightstep-tracer-go
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

	"github.com/lightstep/lightstep-tracer-go"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
)

var lsToken = os.Getenv("LS_ACCESS_TOKEN")
var lsMetricsURL = os.Getenv("LS_METRICS_URL")

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

	componentName := os.Getenv("LIGHTSTEP_COMPONENT_NAME")
	if len(componentName) == 0 {
		componentName = "test-go-server"
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

func main() {
	initLightstepTracer()
	fmt.Printf("Starting server on http://localhost:8081")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		length, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/content/"))
		if err != nil {
			length = 10
		}

		log.Printf("%s %s %s", r.Method, r.URL.Path, r.Proto)
		fmt.Fprintf(w, randString(length))
	})

	log.Fatal(http.ListenAndServe(
		":8081",
		nethttp.Middleware(opentracing.GlobalTracer(), http.DefaultServeMux)),
	)
}
