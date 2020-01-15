package main

import (
	"io"
	"log"
	"net/http"

	"go.opentelemetry.io/otel/api/distributedcontext"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/trace"
	"go.opentelemetry.io/otel/exporter/trace/stdout"
	"go.opentelemetry.io/otel/plugin/httptrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func initTracer() {
	// Create stdout exporter to be able to retrieve
	// the collected spans.
	exporter, err := stdout.NewExporter(stdout.Options{PrettyPrint: true})
	if err != nil {
		log.Fatal(err)
	}

	// For the demonstration, use sdktrace.AlwaysSample sampler to sample all traces.
	// In a production application, use sdktrace.ProbabilitySampler with a desired probability.
	tp, err := sdktrace.NewProvider(sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
		sdktrace.WithSyncer(exporter))
	if err != nil {
		log.Fatal(err)
	}
	global.SetTraceProvider(tp)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	tracer := global.TraceProvider().Tracer("example/server")

	// Extracts the conventional HTTP span attributes,
	// distributed context tags, and a span context for
	// tracing this request.
	attrs, entries, spanCtx := httptrace.Extract(req.Context(), req)

	// Apply the distributed context tags to the request
	// context.
	req = req.WithContext(distributedcontext.WithMap(req.Context(), distributedcontext.NewMap(distributedcontext.MapUpdate{
		MultiKV: entries,
	})))

	// Start the server-side span, passing the remote
	// child span context explicitly.
	_, span := tracer.Start(
		req.Context(),
		"hello",
		trace.WithAttributes(attrs...),
		trace.ChildOf(spanCtx),
	)
	defer span.End()

	io.WriteString(w, "Hello, world!\n")
}

func main() {
	initTracer()

	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		panic(err)
	}
}
