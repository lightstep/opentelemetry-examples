module github.com/codeboten/ls-examples/go/server

go 1.13

require (
	github.com/gorilla/mux v1.7.4
	go.opentelemetry.io/contrib/instrumentation/gorilla/mux v0.7.0
	go.opentelemetry.io/otel v0.9.0
	go.opentelemetry.io/otel/exporters/otlp v0.9.0
	google.golang.org/grpc v1.30.0
)
