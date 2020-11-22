module github.com/lightstep/opentelemetry-examples/go/launcher/client

go 1.14

replace github.com/lightstep/otel-launcher-go => ../../../../otel-launcher-go

require (
	github.com/benbjohnson/clock v1.0.3 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/lightstep/otel-launcher-go v0.0.0-20200724154648-7aba65d4ed0f
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/stretchr/testify v1.6.1 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.14.0
	go.opentelemetry.io/otel v0.14.0
	go.opentelemetry.io/otel/exporters/otlp v0.14.0
)
