package com.lightstep.launcher.server;

import io.opentelemetry.OpenTelemetry;
import io.opentelemetry.context.Scope;
import io.opentelemetry.trace.Span;
import io.opentelemetry.trace.Tracer;
import io.opentelemetry.trace.TracingContextUtils;
import java.io.IOException;
import java.io.PrintWriter;
import java.util.Random;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import org.eclipse.jetty.servlet.ServletContextHandler;
import org.eclipse.jetty.servlet.ServletHolder;

public class ApiContextHandler extends ServletContextHandler {
  public ApiContextHandler() {
    addServlet(new ServletHolder(new ApiServlet()), "/content");
  }

  static final class ApiServlet extends HttpServlet {
    static final String LETTERS = "abcdefghijklmnopqrstuvwxyz";
    final Random rand = new Random();

    @Override
    public void doGet(HttpServletRequest req, HttpServletResponse res)
        throws ServletException, IOException {

      io.grpc.Context context = OpenTelemetry.getPropagators().getHttpTextFormat().extract(
          io.grpc.Context.current(), req, HttpServletRequest::getHeader);

      Span.Builder builder = tracer
          .spanBuilder("/");
      Span parent = TracingContextUtils.getSpanWithoutDefault(context);
      if (parent != null) {
        builder.setParent(parent);
      }
      Span span = builder.startSpan();

      try (Scope scope = tracer.withSpan(span)) {
        // Set the Semantic Convention
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

    // OTel API
    private static Tracer tracer =
        OpenTelemetry.getTracer("io.opentelemetry.example.http.HttpServer");

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
