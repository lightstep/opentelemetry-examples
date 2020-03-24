package main

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	"net/http"

	"google.golang.org/grpc/codes"

	"go.opentelemetry.io/otel/api/correlation"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/key"
	"go.opentelemetry.io/otel/api/trace"
	"go.opentelemetry.io/otel/exporters/trace/stdout"
	"go.opentelemetry.io/otel/plugin/httptrace"
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

	client := http.DefaultClient
	ctx := correlation.NewContext(context.Background(),
		key.String("username", "donuts"),
	)

	var body []byte

	tracer := global.TraceProvider().Tracer("example/client")
	err := tracer.WithSpan(ctx, "client hello",
		func(ctx context.Context) error {
			req, _ := http.NewRequest("GET", "http://localhost:7777/hello", nil)

			ctx, req = httptrace.W3C(ctx, req)
			httptrace.Inject(ctx, req)
			res, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			body, err = ioutil.ReadAll(res.Body)
			_ = res.Body.Close()
			trace.SpanFromContext(ctx).SetStatus(codes.OK, "OK")
			return err
		})

	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second * 5)
}
