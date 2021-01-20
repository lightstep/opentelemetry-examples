package com.lightstep.metric;

import io.opentelemetry.api.OpenTelemetry;
import io.opentelemetry.api.common.Labels;
import io.opentelemetry.api.metrics.LongCounter;
import io.opentelemetry.api.metrics.LongUpDownCounter;
import io.opentelemetry.api.metrics.Meter;
import io.opentelemetry.exporter.otlp.OtlpGrpcMetricExporter;
import io.opentelemetry.sdk.OpenTelemetrySdk;
import io.opentelemetry.sdk.metrics.export.MetricProducer;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.atomic.AtomicInteger;

public class Main {
  public static void main(String[] args) throws Exception {
    if (args.length == 0) {
      System.err.println("metric name is provided");
      System.exit(-1);
    }

    final OtlpGrpcMetricExporter metricExporter = OtlpGrpcMetricExporter.builder()
        .setEndpoint("127.0.0.1:7001")
        .setUseTls(false)
        .build();

    final MetricProducer metricProducer = OpenTelemetrySdk.getGlobalMeterProvider()
        .getMetricProducer();

    final Labels labels = Labels.of("A", "B");
    final int[] testCases = {-1, 4, 3, 6, -5};

    Meter meter = OpenTelemetry.getGlobalMeter("name", "version");
    if (args[0].equalsIgnoreCase("counter")) {
      LongCounter longCounter = meter.longCounterBuilder("counter").setDescription("description")
          .setUnit("1").build();
      longCounter.add(1, labels);
      metricExporter.export(metricProducer.collectAllMetrics());
    }

    if (args[0].equalsIgnoreCase("updown_counter")) {
      LongUpDownCounter longUpDownCounter = meter.longUpDownCounterBuilder("updowncounter")
          .setDescription("description")
          .setUnit("1").build();
      for (int testCase : testCases) {
        longUpDownCounter.add(testCase, labels);
      }
      metricExporter.export(metricProducer.collectAllMetrics());
    }

    if (args[0].equalsIgnoreCase("sum_observer")) {
      final AtomicInteger counter = new AtomicInteger();
      meter.longSumObserverBuilder("sumobserver")
          .setDescription("description")
          .setUnit("1").setCallback(result -> {
        result.observe(testCases[counter.get()], labels);
      }).build();
      for (int i = 0; i < testCases.length; i++) {
        metricExporter.export(metricProducer.collectAllMetrics());
        counter.incrementAndGet();
      }
    }

    if (args[0].equalsIgnoreCase("updown_sum_observer")) {
      AtomicInteger counter = new AtomicInteger();
      meter.longUpDownSumObserverBuilder("updownsumobserver")
          .setDescription("description")
          .setUnit("1").setCallback(result -> {
        result.observe(testCases[counter.get()], labels);
      }).build();
      for (int i = 0; i < testCases.length; i++) {
        metricExporter.export(metricProducer.collectAllMetrics());
        counter.incrementAndGet();
      }
    }

    if (args[0].equalsIgnoreCase("value_observer")) {
      AtomicInteger counter = new AtomicInteger();
      meter.longValueObserverBuilder("valueobserver")
          .setDescription("description")
          .setUnit("1").setCallback(result -> {
        result.observe(testCases[counter.get()], labels);
      }).build();
      for (int i = 0; i < testCases.length; i++) {
        metricExporter.export(metricProducer.collectAllMetrics());
        counter.incrementAndGet();
      }
    }

    TimeUnit.SECONDS.sleep(1);
  }
}
