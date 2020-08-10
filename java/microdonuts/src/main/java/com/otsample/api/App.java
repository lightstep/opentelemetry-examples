package com.otsample.api;

import com.lightstep.opentelemetry.launcher.OpenTelemetryConfiguration;
import io.opentelemetry.opentracingshim.TraceShim;
import io.opentracing.Tracer;
import io.opentracing.util.GlobalTracer;
import java.io.FileInputStream;
import java.io.IOException;
import java.net.MalformedURLException;
import java.util.Properties;
import org.eclipse.jetty.server.Handler;
import org.eclipse.jetty.server.Server;
import org.eclipse.jetty.server.handler.ContextHandler;
import org.eclipse.jetty.server.handler.ContextHandlerCollection;
import org.eclipse.jetty.server.handler.ResourceHandler;

public class App {
  public static void main(String[] args)
      throws Exception {
    Properties config = loadConfig(args);
    if (!configureGlobalTracer(config, "MicroDonuts")) {
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
        new ApiContextHandler(config),
        new KitchenContextHandler(config),
    });
    Server server = new Server(10001);
    server.setHandler(handlers);

    server.start();
    server.dumpStdErr();
    server.join();
  }

  static Properties loadConfig(String[] args)
      throws IOException {
    String file = "tracer_config.properties";
    if (args.length > 0) {
      file = args[0];
    }

    FileInputStream fs = new FileInputStream(file);
    Properties config = new Properties();
    config.load(fs);
    return config;
  }

  static boolean configureGlobalTracer(Properties config, String componentName)
      throws MalformedURLException {
    System.setProperty("ls.service.name", componentName);
    System.setProperty("ls.access.token", config.getProperty("ls.access.token"));
    System.setProperty("ls.service.version", config.getProperty("ls.service.version"));
    System.setProperty("otel.exporter.otlp.span.endpoint",
        config.getProperty("ls.collector_host"));
    OpenTelemetryConfiguration.newBuilder().install();
    Tracer tracer = TraceShim.createTracerShim();

    GlobalTracer.registerIfAbsent(tracer);
    return true;
  }
}