package com.lightstep.launcher.server;

import com.lightstep.opentelemetry.launcher.OpenTelemetryConfiguration;
import io.opentelemetry.api.GlobalOpenTelemetry;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.SpanKind;
import io.opentelemetry.api.trace.Tracer;
import java.util.concurrent.TimeUnit;
import org.eclipse.jetty.server.Handler;
import org.eclipse.jetty.server.Server;
import org.eclipse.jetty.server.handler.ContextHandlerCollection;

import io.opentelemetry.instrumentation.annotations.WithSpan;

public class ExampleServer {

  private static final Tracer tracer = GlobalOpenTelemetry.getTracer("LightstepExample");


  @WithSpan
  public static void main(String[] args) throws Exception {
    // OpenTelemetryConfiguration.newBuilder().install();

    // Tracer tracer = GlobalOpenTelemetry.getTracer("LightstepExample");

    // Span span = tracer.spanBuilder("start example").setSpanKind(SpanKind.CLIENT).startSpan();
    Span span = Span.current();
    span.setAttribute("Attribute 1", "Value 1");
    span.addEvent("Event 0");
    // execute my use case - here we simulate a wait
    doWork();
    span.addEvent("Event 1");
    span.end();

    ContextHandlerCollection handlers = new ContextHandlerCollection();
    handlers.setHandlers(new Handler[]{
        new ApiContextHandler(),
    });
    Server server = new Server(8084);
    server.setHandler(handlers);

    server.start();
    server.dumpStdErr();
    server.join();
  }

  @WithSpan
  private static void doWork() {
    try {
      TimeUnit.SECONDS.sleep(1);
    } catch (InterruptedException ignore) {
    }
  }
}