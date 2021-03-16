module github.com/lightstep/ls-examples/go/opentelemetry/client

go 1.14

require (
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.18.0
	go.opentelemetry.io/contrib/propagators v0.18.0
	go.opentelemetry.io/otel v0.18.0
	go.opentelemetry.io/otel/exporters/otlp v0.18.0
	go.opentelemetry.io/otel/sdk v0.18.0
	google.golang.org/grpc v1.36.0
)
