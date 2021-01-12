module github.com/lightstep/ls-examples/go/launcher/server

go 1.14

require (
	github.com/gorilla/mux v1.8.0
	github.com/lightstep/otel-launcher-go v0.15.0
	github.com/opentracing-contrib/go-gorilla v0.0.0-20190110000444-ced666783644 // indirect
	go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux v0.15.1
	go.opentelemetry.io/otel/bridge/opentracing v0.15.0 // indirect
)
