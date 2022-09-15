package com.lightstep.otlp.client;

import io.opentelemetry.api.GlobalOpenTelemetry;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.SpanKind;
import io.opentelemetry.api.trace.Tracer;
import io.opentelemetry.context.Scope;
import io.opentelemetry.context.propagation.ContextPropagators;
import io.opentelemetry.exporter.otlp.trace.OtlpGrpcSpanExporter;
import io.opentelemetry.extension.trace.propagation.B3Propagator;
import io.opentelemetry.sdk.OpenTelemetrySdk;
import io.opentelemetry.sdk.autoconfigure.OpenTelemetrySdkAutoConfiguration;
import io.opentelemetry.sdk.trace.SdkTracerProvider;
import io.opentelemetry.sdk.trace.export.BatchSpanProcessor;
import java.util.concurrent.TimeUnit;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;

public class Client {
  private static final String ACCESS_TOKEN_HEADER = "lightstep-access-token";

  public static void main(String[] args) {

    // final String satelliteURL = "https://" + System.getenv("LS_SATELLITE_URL");
    // final String lsToken = System.getenv("LS_ACCESS_TOKEN");
    String targetURL = System.getenv("DESTINATION_URL");
    if (targetURL == null || targetURL.length() == 0) {
      targetURL = "http://127.0.0.1:8080/ping";
    }

    // final OtlpGrpcSpanExporter exporter = OtlpGrpcSpanExporter.builder()
    //     .setTimeout(60_000, TimeUnit.MILLISECONDS)
    //     .addHeader(ACCESS_TOKEN_HEADER, lsToken)
    //     .setEndpoint(satelliteURL).build();

    // SdkTracerProvider sdkTracerProvider = SdkTracerProvider.builder()
    //     .addSpanProcessor(BatchSpanProcessor.builder(exporter).build())
    //     .setResource(OpenTelemetrySdkAutoConfiguration.getResource())
    //     .build();

    // OpenTelemetrySdk.builder()
    //     .setTracerProvider(sdkTracerProvider)
    //     .setPropagators(
    //         ContextPropagators.create(B3Propagator.injectingMultiHeaders()))
    //     .buildAndRegisterGlobal();

    Tracer tracer = GlobalOpenTelemetry.getTracer("LightstepExample");

    while (true) {
      doWork(tracer, targetURL);
      try {
        Thread.sleep(1000);
      } catch (InterruptedException ignore) {
      }
    }
  }

  private static void doWork(Tracer tracer, String targetURL) {
    Span span = tracer.spanBuilder("start example").setSpanKind(SpanKind.CLIENT).startSpan();
    span.setAttribute("Attribute 1", "Value 1");
    span.addEvent("Event 0");

    OkHttpClient client = new OkHttpClient();
    Request.Builder reqBuilder = new Request.Builder();

    // Inject the current Span into the Request.
    try (Scope scope = span.makeCurrent()) {
      GlobalOpenTelemetry.getPropagators().getTextMapPropagator()
          .inject(io.opentelemetry.context.Context.current(), reqBuilder,
              Request.Builder::addHeader);
    }

    Request req = reqBuilder
        .url(targetURL)
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