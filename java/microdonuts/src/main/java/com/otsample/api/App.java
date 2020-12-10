package com.otsample.api;

import com.lightstep.opentelemetry.launcher.OpenTelemetryConfiguration;
import io.opentelemetry.opentracingshim.OpenTracingShim;
import io.opentracing.Tracer;
import io.opentracing.util.GlobalTracer;
import org.eclipse.jetty.server.Handler;
import org.eclipse.jetty.server.Server;
import org.eclipse.jetty.server.handler.ContextHandler;
import org.eclipse.jetty.server.handler.ContextHandlerCollection;
import org.eclipse.jetty.server.handler.ResourceHandler;

public class App {
  public static void main(String[] args)
      throws Exception {
    if (!configureGlobalTracer()) {
      throw new Exception("Could not configure the global tracer");
    }

    ResourceHandler filesHandler = new ResourceHandler();
    filesHandler.setWelcomeFiles(new String[]{"./index.html"});
    filesHandler.setResourceBase("./client");

    ContextHandler fileCtxHandler = new ContextHandler();
    fileCtxHandler.setHandler(filesHandler);

    ContextHandlerCollection handlers = new ContextHandlerCollection();
    handlers.setHandlers(new Handler[]{
        fileCtxHandler,
        new ApiContextHandler(),
        new KitchenContextHandler(),
    });
    Server server = new Server(10001);
    server.setHandler(handlers);

    server.start();
    server.dumpStdErr();
    server.join();
  }

  static boolean configureGlobalTracer() {
    OpenTelemetryConfiguration.newBuilder().install();
    Tracer tracer = OpenTracingShim.createTracerShim();

    GlobalTracer.registerIfAbsent(tracer);
    return true;
  }
}