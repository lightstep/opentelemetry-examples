// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Example using the OTLP exporter + collector + third-party backends. For
// information about using the exporter, see:
// https://pkg.go.dev/go.opentelemetry.io/otel/exporters/otlp?tab=doc#example-package-Insecure
package main

import (
	"context"
	"log"
	"time"
	"fmt"

	"google.golang.org/grpc"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric/controller/push"
	"go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
	"go.opentelemetry.io/otel/sdk/resource"
)


func main() {
	log.Printf("Waiting for connection...")

	ctx := context.Background()

	exp, err := otlp.NewExporter(
		ctx,
		[]otlp.ExporterOption{
			otlp.WithInsecure(),
			otlp.WithAddress("127.0.0.1:7001"),
			otlp.WithGRPCDialOption(grpc.WithBlock()),
		}...,
	)

	handleErr(err, "failed to create exporter")

	res, _ := resource.TelemetrySDK{}.Detect(ctx)

	pusher := push.New(
		basic.New(
			simple.NewWithExactDistribution(),
			exp,
		),
		exp,
		push.WithPeriod(2*time.Second),
		push.WithResource(res),
	)

	otel.SetMeterProvider(pusher.MeterProvider())
	pusher.Start()

	shutdown := func () {
		pusher.Stop()
		handleErr(exp.Shutdown(ctx), "failed to stop exporter")
	}
	defer shutdown()

	meter := otel.Meter("name", metric.WithInstrumentationVersion("version"))

	labels := []label.KeyValue{
		label.String("A", "B"),
	}

	counter := metric.Must(meter).NewInt64Counter(
		"counter",
		metric.WithDescription("description"),
		metric.WithUnit("1"),
	)

	updowncounter := metric.Must(meter).NewInt64UpDownCounter(
		"updowncounter",
		metric.WithDescription("description"),
		metric.WithUnit("1"),
	)

	testValues := [5]int64{-1, 4, 3, 6, -5}

	_ = metric.Must(meter).NewInt64SumObserver(
		"sumobserver",
		func(_ context.Context, result metric.Int64ObserverResult) {
			for _, testValue := range testValues {
				result.Observe(testValue, labels...)
			}
		},
		metric.WithDescription("description"),
		metric.WithUnit("1"),
	)

	_ = metric.Must(meter).NewInt64UpDownSumObserver(
		"updownsumobserver",
		func(_ context.Context, result metric.Int64ObserverResult) {
			for _, testValue := range testValues {
				result.Observe(testValue, labels...)
			}
		},
		metric.WithDescription("description"),
		metric.WithUnit("1"),
	)

	_ = metric.Must(meter).NewInt64UpDownSumObserver(
		"updownsumobserver",
		func(_ context.Context, result metric.Int64ObserverResult) {
			for _, testValue := range testValues {
				result.Observe(testValue, labels...)
			}
		},
		metric.WithDescription("description"),
		metric.WithUnit("1"),
	)

	_ = metric.Must(meter).NewInt64ValueObserver(
		"valueobserver",
		func(_ context.Context, result metric.Int64ObserverResult) {
			for _, testValue := range testValues {
				result.Observe(testValue, labels...)
			}
		},
		metric.WithDescription("description"),
		metric.WithUnit("1"),
	)

	for _, testValue := range testValues {
		counter.Add(ctx, testValue, labels...)
		updowncounter.Add(ctx, testValue, labels...)
	}

}

func handleErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
