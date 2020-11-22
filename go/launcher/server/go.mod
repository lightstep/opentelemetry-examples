module github.com/lightstep/ls-examples/go/launcher/server

go 1.14

replace github.com/lightstep/otel-launcher-go => ../../../../otel-launcher-go

require (
	github.com/gorilla/mux v1.8.0
	github.com/lightstep/otel-launcher-go v0.0.0-20200724154648-7aba65d4ed0f
	go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux v0.14.0
)
