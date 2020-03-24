// first_tracer.go
package main

import (
	"context"
	"log"

	"go.opentelemetry.io/otel/api/core"
	"go.opentelemetry.io/otel/api/global"
	apitrace "go.opentelemetry.io/otel/api/trace"
	"go.opentelemetry.io/otel/exporters/trace/stdout"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func initTracer() {
	exporter, err := stdout.NewExporter(stdout.Options{PrettyPrint: true})
	if err != nil {
		log.Fatal(err)
	}
	tp, err := sdktrace.NewProvider(sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
		sdktrace.WithSyncer(exporter))
	if err != nil {
		log.Fatal(err)
	}
	global.SetTraceProvider(tp)
}

func main() {
	initTracer()
	tracer := global.TraceProvider().Tracer("ex.com/basic")
	ctx, span := tracer.Start(context.Background(), "foo")
	span.SetAttributes(core.KeyValue{Key: "platform", Value: core.String("osx")})
	span.SetAttributes(core.KeyValue{Key: "version", Value: core.String("1.2.3")})
	span.AddEvent(ctx, "event in foo", core.KeyValue{Key: "name", Value: core.String("foo1")})

	attributes := []core.KeyValue{
		core.KeyValue{Key: "platform", Value: core.String("osx")},
		core.KeyValue{Key: "version", Value: core.String("1.2.3")},
	}

	ctx, child := tracer.Start(ctx, "baz", apitrace.WithAttributes(attributes...))

	child.End()
	span.End()

	_ = tracer.WithSpan(ctx, "foo",
		func(ctx context.Context) error {
			tracer.WithSpan(ctx, "bar",
				func(ctx context.Context) error {
					tracer.WithSpan(ctx, "baz",
						func(ctx context.Context) error {
							return nil
						},
					)
					return nil
				},
			)
			return nil
		},
	)
}
