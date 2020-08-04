//
// example code to test lightstep/otel-launcher-go/launcher
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
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/lightstep/otel-launcher-go/launcher"
	muxtrace "go.opentelemetry.io/contrib/instrumentation/gorilla/mux"
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

func main() {
	otel := launcher.ConfigureOpentelemetry()
	defer otel.Shutdown()
	fmt.Printf("Starting server on http://localhost:8081\n")
	r := mux.NewRouter()
	r.Use(muxtrace.Middleware(os.Getenv("LS_SERVICE_NAME")))
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
