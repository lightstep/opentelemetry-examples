package com.lightstep.launcher.server;

import com.lightstep.opentelemetry.launcher.OpenTelemetryConfiguration;
import io.opentelemetry.OpenTelemetry;
import io.opentelemetry.exporters.otlp.OtlpGrpcSpanExporter;
import io.opentelemetry.sdk.OpenTelemetrySdk;
import io.opentelemetry.sdk.trace.export.SimpleSpanProcessor;
import io.opentelemetry.trace.Span;
import io.opentelemetry.trace.Span.Kind;
import io.opentelemetry.trace.Tracer;
import org.eclipse.jetty.server.Handler;
import org.eclipse.jetty.server.Server;
import org.eclipse.jetty.server.handler.ContextHandlerCollection;

public class ExampleServer {

  public static void main(String[] args) throws Exception {
    OtlpGrpcSpanExporter exporter = OpenTelemetryConfiguration.newBuilder()
        .buildExporter();

    OpenTelemetrySdk.getTracerProvider()
        .addSpanProcessor(SimpleSpanProcessor.newBuilder(exporter).build());

    Tracer tracer =
        OpenTelemetry.getTracerProvider().get("LightstepExample");

    Span span = tracer.spanBuilder("start example").setSpanKind(Kind.CLIENT).startSpan();
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

  private static void doWork() {
    try {
      Thread.sleep(1000);
    } catch (InterruptedException ignore) {
    }
  }
}