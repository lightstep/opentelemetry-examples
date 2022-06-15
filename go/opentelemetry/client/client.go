//
// example code to test lightstep/opentelemetry-exporter-go
//
// usage:
//   LS_ACCESS_TOKEN=${SECRET_TOKEN} \
//   LS_SERVICE_NAME=demo-client-go \
//   LS_SERVICE_VERSION=0.1.8 \
//   go run client.go

package main

import (
	"context"
	"fmt"
	// "log"
	"io/ioutil"
	"net/http"
	// "os"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	// "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	// "go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.opentelemetry.io/otel/trace"
	// "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	// "go.opentelemetry.io/contrib/propagators/b3"
	// "go.opentelemetry.io/otel"
	// "go.opentelemetry.io/otel/attribute"
	// "go.opentelemetry.io/otel/exporters/otlp"
	// "go.opentelemetry.io/otel/exporters/otlp/otlpgrpc"
	// "go.opentelemetry.io/otel/propagation"
	// "go.opentelemetry.io/otel/sdk/resource"
	// sdktrace "go.opentelemetry.io/otel/sdk/trace"
	// semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	// "google.golang.org/grpc/credentials"
)

// var (
// 	componentName  = os.Getenv("LS_SERVICE_NAME")
// 	serviceVersion = os.Getenv("LS_SERVICE_VERSION")
// 	// lsToken        = os.Getenv("LS_ACCESS_TOKEN")
// 	collectorURL = os.Getenv("OTEL_COLLECTOR_URL")
// 	targetURL    = os.Getenv("DESTINATION_URL")
// 	// insecure       = os.Getenv("LS_INSECURE")
// )

var (
	tracer         trace.Tracer
	serviceName    string = "test-go-client"
	serviceVersion string = "0.1.0"
	collectorAddr  string = "localhost:4318" // HTTP endpoint for collector
	targetURL      string = "http://localhost:8081/ping"
)

func newTraceProvider(ctx context.Context) *sdktrace.TracerProvider {
	exporter, err :=
		otlptracehttp.New(ctx,
			// WithInsecure lets us use http instead of https.
			// This is just for local development.
			otlptracehttp.WithInsecure(),
			otlptracehttp.WithEndpoint(collectorAddr),
		)

	if err != nil {
		panic(err)
	}

	// This includes the following resources:
	//
	// - sdk.language, sdk.version
	// - service.name, service.version, environment
	//
	// Including these resources is a good practice because it is commonly
	// used by various tracing backends to let you more accurately
	// analyze your telemetry data.
	resource, rErr :=
		resource.Merge(
			resource.Default(),
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(serviceName),
				semconv.ServiceVersionKey.String(serviceVersion),
				attribute.String("environment", "getting-started"),
			),
		)

	if rErr != nil {
		panic(rErr)
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource),
	)
}

// func initExporter(url string, token string) *otlp.Exporter {
// 	headers := map[string]string{
// 		"lightstep-access-token": token,
// 	}

// 	secureOption := otlpgrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
// 	if len(insecure) > 0 {
// 		secureOption = otlpgrpc.WithInsecure()
// 	}

// 	exporter, err := otlp.NewExporter(
// 		context.Background(),
// 		otlpgrpc.NewDriver(
// 			secureOption,
// 			otlpgrpc.WithEndpoint(url),
// 			otlpgrpc.WithHeaders(headers),
// 		),
// 	)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return exporter
// }

// func initTracer() {

// 	if len(collectorURL) == 0 {
// 		collectorURL = "localhost:4317"
// 	}

// 	if len(componentName) == 0 {
// 		componentName = "test-go-client"
// 	}
// 	if len(serviceVersion) == 0 {
// 		serviceVersion = "0.0.0"
// 	}

// 	ctx := context.Background()
// 	res, err := resource.New(ctx,
// 		resource.WithAttributes(
// 			// Service name to be use by observability tool
// 			semconv.ServiceNameKey.String(componentName)))
// 	// Checking for errors
// 	if err != nil {
// 		fmt.Printf("Error adding %v to the tracer engine: %v", "applicationName", err)
// 	}

// 	traceExporter, err := otlptracegrpc.New(ctx,
// 		otlptracegrpc.WithInsecure(),
// 		otlptracegrpc.WithEndpoint(collectorURL),
// 	)

// 	// Check for errors
// 	if err != nil {
// 		fmt.Printf("Error initializing the trace exporter: %v", err)
// 	}

// 	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
// 	tp := sdktrace.NewTracerProvider(
// 		sdktrace.WithResource(res),
// 		sdktrace.WithSpanProcessor(bsp),
// 	)
// 	defer tp.Shutdown(context.Background())
// 	otel.SetTracerProvider(tp)
// 	otel.SetTextMapPropagator(propagation.TraceContext{})
// }

// func initTracerOld() {
// 	b3 := b3.B3{}
// 	// Register the B3 propagator globally.
// 	otel.SetTextMapPropagator(b3)
// 	if len(collectorURL) == 0 {
// 		collectorURL = "localhost:4317"
// 	}

// 	if len(componentName) == 0 {
// 		componentName = "test-go-client"
// 	}
// 	if len(serviceVersion) == 0 {
// 		serviceVersion = "0.0.0"
// 	}

// 	exporter := initExporter(collectorURL, lsToken)

// 	resources, err := resource.New(
// 		context.Background(),
// 		resource.WithAttributes(
// 			attribute.String("service.name", componentName),
// 			attribute.String("service.version", serviceVersion),
// 			attribute.String("library.language", "go"),
// 			attribute.String("library.version", "1.2.3"),
// 		),
// 	)
// 	if err != nil {
// 		log.Printf("Could not set resources: ", err)
// 	}
// 	tp := trace.NewTracerProvider(
// 		trace.WithConfig(trace.Config{DefaultSampler: trace.AlwaysSample()}),
// 		trace.WithSyncer(exporter),
// 		trace.WithResource(resources),
// 	)
// 	otel.SetTracerProvider(tp)
// }

func makeRequest() {
	httpClient := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	tracer := otel.Tracer("otel-example/client")
	_, span := tracer.Start(context.Background(), "makeRequest")
	defer span.End()

	req, _ := http.NewRequest("GET", targetURL, nil)
	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("Body : %s", body)
	fmt.Printf("Request to %s, got %d bytes\n", targetURL, res.ContentLength)
}

func main() {
	// initTracer()

	ctx := context.Background()

	tp := newTraceProvider(ctx)
	defer func() { _ = tp.Shutdown(ctx) }()

	otel.SetTracerProvider(tp)

	// if len(targetURL) == 0 {
	// 	targetURL = "http://localhost:8081/ping"
	// }
	for {
		makeRequest()
		time.Sleep(1 * time.Second)
	}

}
