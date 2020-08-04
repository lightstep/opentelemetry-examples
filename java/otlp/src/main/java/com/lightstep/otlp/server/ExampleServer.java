package com.lightstep.otlp.server;

import static io.grpc.Metadata.ASCII_STRING_MARSHALLER;
import org.eclipse.jetty.server.Server;
import org.eclipse.jetty.server.Handler;
import org.eclipse.jetty.server.handler.ContextHandlerCollection;

import io.grpc.CallOptions;
import io.grpc.Channel;
import io.grpc.ClientCall;
import io.grpc.ClientInterceptor;
import io.grpc.ForwardingClientCall.SimpleForwardingClientCall;
import io.grpc.ManagedChannelBuilder;
import io.grpc.Metadata;
import io.grpc.Metadata.Key;
import io.grpc.MethodDescriptor;
import io.grpc.Status;
import io.opentelemetry.OpenTelemetry;
import io.opentelemetry.exporters.otlp.OtlpGrpcSpanExporter;
import io.opentelemetry.sdk.OpenTelemetrySdk;
import io.opentelemetry.sdk.trace.export.SimpleSpanProcessor;
import io.opentelemetry.trace.Span;
import io.opentelemetry.trace.Span.Kind;
import io.opentelemetry.trace.Tracer;

public class ExampleServer {
  private static final Key<String> ACCESS_TOKEN_HEADER = Key
      .of("lightstep-access-token", ASCII_STRING_MARSHALLER);

  public static void main(String[] args) throws Exception{
    final String satelliteURL = System.getenv("LS_SATELLITE_URL");
    final String lsToken = System.getenv("LS_ACCESS_TOKEN");

    final OtlpGrpcSpanExporter exporter = OtlpGrpcSpanExporter.newBuilder()
        .setDeadlineMs(60_000)
        .readSystemProperties()
        .readEnvironmentVariables()
        .setChannel(
            ManagedChannelBuilder
                .forTarget(satelliteURL)
                .useTransportSecurity()
                .intercept(new ClientInterceptor() {
                  @Override
                  public <ReqT, RespT> ClientCall<ReqT, RespT> interceptCall(
                      MethodDescriptor<ReqT, RespT> method, CallOptions callOptions, Channel next) {
                    return new SimpleForwardingClientCall<ReqT, RespT>(
                        next.newCall(method, callOptions)
                    ) {
                      @Override
                      public void start(final Listener<RespT> responseListener, Metadata headers) {
                        headers.put(ACCESS_TOKEN_HEADER, lsToken);
                        super.start(new Listener<RespT>() {
                          @Override
                          public void onHeaders(Metadata headers) {
                            responseListener.onHeaders(headers);
                          }

                          @Override
                          public void onMessage(RespT message) {
                            responseListener.onMessage(message);
                          }

                          @Override
                          public void onClose(Status status, Metadata trailers) {
                            System.out.println(status);
                            responseListener.onClose(status, trailers);
                          }

                          @Override
                          public void onReady() {
                            responseListener.onReady();
                          }
                        }, headers);
                      }
                    };
                  }
                })
                .build()
        ).build();

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
    handlers.setHandlers(new Handler[] {
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