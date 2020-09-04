package com.lightstep.launcher.client;

import com.lightstep.opentelemetry.common.VariablesConverter;
import com.lightstep.opentelemetry.launcher.OpenTelemetryConfiguration;
import io.grpc.Context;
import io.opentelemetry.OpenTelemetry;
import io.opentelemetry.exporters.otlp.OtlpGrpcSpanExporter;
import io.opentelemetry.sdk.OpenTelemetrySdk;
import io.opentelemetry.sdk.trace.export.SimpleSpanProcessor;
import io.opentelemetry.trace.Span;
import io.opentelemetry.trace.Span.Kind;
import io.opentelemetry.trace.Tracer;
import io.opentelemetry.trace.TracingContextUtils;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;

public class Client {
  public static void main(String[] args) {
    String targetURL = System.getenv("TARGET_URL");
    if (targetURL == null || targetURL.length() == 0) {
      targetURL = "http://127.0.0.1:8084";
    }

    final String spanEndpoint = System.getenv("OTEL_EXPORTER_OTLP_SPAN_ENDPOINT");
    final String lsToken = System.getenv("LS_ACCESS_TOKEN");
    final String serviceName = System.getenv("LS_SERVICE_NAME");
    final String serviceVersion = System.getenv("LS_SERVICE_VERSION");

    OtlpGrpcSpanExporter exporter = OpenTelemetryConfiguration.newBuilder()
        .setServiceName(serviceName)
        .setServiceVersion(serviceVersion)
        .setAccessToken(lsToken)
        .setSpanEndpoint(spanEndpoint)
        .buildExporter();

    OpenTelemetrySdk.getTracerProvider()
        .addSpanProcessor(SimpleSpanProcessor.newBuilder(exporter).build());

    Tracer tracer =
        OpenTelemetry.getTracerProvider().get("LightstepExample");

    while (true) {
      doWork(tracer, targetURL);
      try {
        Thread.sleep(1000);
      } catch (InterruptedException ignore) {
      }
    }
  }

  private static void doWork(Tracer tracer, String targetURL) {
    Span span = tracer.spanBuilder("start example").setSpanKind(Kind.CLIENT).startSpan();
    span.setAttribute("Attribute 1", "Value 1");
    span.addEvent("Event 0");

    OkHttpClient client = new OkHttpClient();
    Request.Builder reqBuilder = new Request.Builder();

    // Inject the current Span into the Request.
    Context withSpanContext = TracingContextUtils.withSpan(span, Context.current());
    OpenTelemetry.getPropagators().getHttpTextFormat()
        .inject(withSpanContext, reqBuilder, Request.Builder::addHeader);

    Request req = reqBuilder
        .url(targetURL + "/content")
        .build();

    try (Response res = client.newCall(req).execute()) {
      String retval = res.body().string();
      System.out.println(String.format("Request to %s, got %s bytes",
          targetURL, retval.length()));
    } catch (Exception e) {
      System.out.println(String.format("Request to %s failed: %s",
          targetURL, e));
    }
    span.addEvent("Event 1");
    span.end();
  }
}