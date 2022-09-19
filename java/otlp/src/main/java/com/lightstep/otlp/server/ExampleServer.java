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
import io.opentelemetry.instrumentation.annotations.WithSpan;
import io.opentelemetry.context.Scope;
import io.opentelemetry.context.Context;

public class ExampleServer {
  
  private static final Tracer tracer = GlobalOpenTelemetry.getTracer("LightstepExample");

  @WithSpan
  public static void main(String[] args) throws Exception {
  
    // Span span = tracer.spanBuilder("start example").setSpanKind(SpanKind.CLIENT).startSpan();
    Span mainSpan = Span.current();
    mainSpan.setAttribute("Attribute 1", "Value 1");
    mainSpan.addEvent("Started main server span");
  
    Span serveSpan = tracer.spanBuilder("serve")
                        .setParent(Context.current().with(mainSpan))
                        .startSpan();

      // execute my use case - here we simulate a wait
      doWork();


    try (Scope scope = serveSpan.makeCurrent()) {                        
      Span span = Span.current();

      span.addEvent("Did some stuff, yo!");
      // span.end();

      ContextHandlerCollection handlers = new ContextHandlerCollection();
      handlers.setHandlers(new Handler[]{
          new ApiContextHandler(),
      });
      Server server = new Server(8083);
      server.setHandler(handlers);

      server.start();
      server.dumpStdErr();
      server.join();

    } finally {
      serveSpan.end();  
    }

    mainSpan.addEvent("Finished main server span");
    mainSpan.end();  

  }

  private static void doWork() {
    try {
      Thread.sleep(1000);
    } catch (InterruptedException ignore) {
    }
  }
}