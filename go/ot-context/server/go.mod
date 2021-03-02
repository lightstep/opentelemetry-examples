module github.com/codeboten/ls-examples/go/server

go 1.14

require (
	github.com/gorilla/mux v1.8.0
	go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux v0.17.0
	go.opentelemetry.io/contrib/propagators v0.17.0
	go.opentelemetry.io/otel v0.17.0
	go.opentelemetry.io/otel/exporters/otlp v0.17.0
	go.opentelemetry.io/otel/sdk v0.17.0
	google.golang.org/grpc v1.35.0
)
