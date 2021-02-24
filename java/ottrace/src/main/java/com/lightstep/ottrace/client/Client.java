package com.lightstep.ottrace.client;

import com.lightstep.opentelemetry.launcher.OpenTelemetryConfiguration;
import com.lightstep.opentelemetry.launcher.Propagator;
import io.opentelemetry.api.GlobalOpenTelemetry;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.SpanKind;
import io.opentelemetry.api.trace.Tracer;
import io.opentelemetry.context.Scope;
import java.util.concurrent.TimeUnit;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;

public class Client {
  public static void main(String[] args) {
    String targetURL = System.getenv("DESTINATION_URL");
    if (targetURL == null || targetURL.length() == 0) {
      targetURL = "http://127.0.0.1:8084/ping";
    }

    OpenTelemetryConfiguration.newBuilder()
        .setPropagator(Propagator.OT_TRACE)
        .install();

    Tracer tracer = GlobalOpenTelemetry.getTracer("LightstepExample");

    while (true) {
      doWork(tracer, targetURL);
      try {
        TimeUnit.SECONDS.sleep(1);
      } catch (InterruptedException ignore) {
        break;
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