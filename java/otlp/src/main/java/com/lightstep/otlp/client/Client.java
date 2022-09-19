package com.lightstep.otlp.client;

import io.opentelemetry.api.GlobalOpenTelemetry;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.SpanKind;
import io.opentelemetry.api.trace.Tracer;
import io.opentelemetry.context.Context;
import io.opentelemetry.context.Scope;
import io.opentelemetry.instrumentation.annotations.WithSpan;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;

public class Client {

  private static final Tracer tracer = GlobalOpenTelemetry.getTracer("LightstepExample");

  @WithSpan
  public static void main(String[] args) {

    String targetURL = System.getenv("DESTINATION_URL");
    if (targetURL == null || targetURL.length() == 0) {
      targetURL = "http://127.0.0.1:8083/ping";
    }

    while (true) {
      doWork(tracer, targetURL);
      try {
        Thread.sleep(1000);
      } catch (InterruptedException ignore) {
      }
    }
  }

  private static void doWork(Tracer tracer, String targetURL) {
    Span doWorkSpan = tracer.spanBuilder("doWork").setSpanKind(SpanKind.CLIENT).startSpan();
    doWorkSpan.setAttribute("Attribute 1", "Value 1");
    doWorkSpan.addEvent("Started doWork");

    // Create a new span. Note that it is possible to set the parent manually.
    Span makeRequestSpan = tracer.spanBuilder("makeRequest")
                          .setParent(Context.current().with(doWorkSpan))
                          .startSpan();

    // Inject the current Span into the Request.
    try (Scope scope = makeRequestSpan.makeCurrent()) {

      OkHttpClient client = new OkHttpClient();
      Request.Builder reqBuilder = new Request.Builder();
  
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
      
    } finally {
      makeRequestSpan.end();  
    }

    doWorkSpan.addEvent("Finished doWork");
    doWorkSpan.end();  

  }
}