package com.lightstep.otlp.server;

import io.opentelemetry.api.GlobalOpenTelemetry;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.SpanKind;
import io.opentelemetry.api.trace.Tracer;
// import io.opentelemetry.context.propagation.ContextPropagators;
// import io.opentelemetry.exporter.otlp.trace.OtlpGrpcSpanExporter;
// import io.opentelemetry.extension.trace.propagation.B3Propagator;
// import io.opentelemetry.sdk.OpenTelemetrySdk;
// import io.opentelemetry.sdk.trace.SdkTracerProvider;
// import io.opentelemetry.sdk.trace.export.BatchSpanProcessor;
import java.util.concurrent.TimeUnit;
import org.eclipse.jetty.server.Handler;
import org.eclipse.jetty.server.Server;
import org.eclipse.jetty.server.handler.ContextHandlerCollection;

public class ExampleServer {
  
  private static final Tracer tracer = GlobalOpenTelemetry.getTracer("LightstepExample");

  public static void main(String[] args) throws Exception {
  
    Span span = tracer.spanBuilder("start example").setSpanKind(SpanKind.CLIENT).startSpan();
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
    Server server = new Server(8083);
    server.setHandler(handlers);

    server.start();
    server.dumpStdErr();
    server.join();
  }

  private static void doWork() {
    try {
      Thread.sleep(1000);
    } catch (InterruptedException ignore) {
    }
  }
}