package com.lightstep.launcher.server;


import io.opentelemetry.api.GlobalOpenTelemetry;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.Tracer;
import io.opentelemetry.context.Scope;
import io.opentelemetry.context.propagation.TextMapGetter;
import java.io.IOException;
import java.io.PrintWriter;
import java.util.ArrayList;
import java.util.Enumeration;
import java.util.List;
import java.util.Random;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import org.eclipse.jetty.servlet.ServletContextHandler;
import org.eclipse.jetty.servlet.ServletHolder;

public class ApiContextHandler extends ServletContextHandler {
  public ApiContextHandler() {
    addServlet(new ServletHolder(new ApiServlet()), "/ping");
  }

  static final class ApiServlet extends HttpServlet {
    static final String LETTERS = "abcdefghijklmnopqrstuvwxyz";
    final Random rand = new Random();

    @Override
    public void doGet(HttpServletRequest req, HttpServletResponse res)
        throws ServletException, IOException {

      io.opentelemetry.context.Context parentContext = GlobalOpenTelemetry.getPropagators()
          .getTextMapPropagator().extract(
              io.opentelemetry.context.Context.current(), req,
              new TextMapGetter<HttpServletRequest>() {
                @Override
                public Iterable<String> keys(HttpServletRequest carrier) {
                  final Enumeration<String> headerNames = carrier.getHeaderNames();
                  List<String> keys = new ArrayList<>();
                  while (headerNames.hasMoreElements()) {
                    keys.add(headerNames.nextElement());
                  }
                  return keys;
                }

                @Override
                public String get(HttpServletRequest carrier, String key) {
                  return carrier.getHeader(key);
                }
              });

      try (final Scope scope = parentContext.makeCurrent()) {
        Span span = tracer.spanBuilder("/").startSpan();
        try {
          span.setAttribute("component", "http");
          span.setAttribute("http.method", "GET");
          span.setAttribute("http.scheme", "http");
          span.setAttribute("http.target", "/");
          try (PrintWriter writer = res.getWriter()) {
            writer.write(createRandomString());
          }
        } finally {
          // Close the span
          span.end();
        }
      }

    }

    // OTel API
    private static Tracer tracer =
        GlobalOpenTelemetry.getTracer("io.opentelemetry.example.http.HttpServer");

    String createRandomString() {
      int length = rand.nextInt(1023) + 1;
      StringBuilder sb = new StringBuilder(length);

      for (int i = 0; i < length; i++) {
        sb.append(LETTERS.charAt(rand.nextInt(LETTERS.length())));
      }

      return sb.toString();
    }
  }
}
