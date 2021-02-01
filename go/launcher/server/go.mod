module github.com/lightstep/ls-examples/go/launcher/server

go 1.14

require (
	github.com/gorilla/mux v1.8.0
	github.com/lightstep/otel-launcher-go v0.16.1
	go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux v0.16.0
)
