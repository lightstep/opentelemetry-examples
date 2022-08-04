# Ingest Envoy metrics and traces with the OpenTelemetry Collector

This example illustrates how you can ingest Envoy metrics and traces. Envoy exposes a Prometheus compatible metrics endpoint and has built-in (experimental as of v1.23.0) support for generating OpenTelemetry traces.

## Requirements

* OpenTelemetry Collector Contrib v0.51.0+
* Envoy v1.23.0+

## Prerequisites

You must have a Lightstep Observability [access token](/docs/create-and-manage-access-tokens) for the project to report metrics to.

## Running the Example

You can run this example with `docker-compose up` in this directory. Using `docker-compose --profile loadgen up` also creates an instance to send requests to the NGINX service. You'll want to view this in Lightstep with a dashboard. 

```
  $ export LS_ACCESS_TOKEN=<your-lightstep-access-token>
  $ docker-compose up

  # make some requests
  $ curl http://localhost:8080/service/1
  $ curl http://localhost:8080/service/2
```

## Configuration: Metrics

Installation of the OpenTelemetry Collector varies, please refer to the [collector documentation](https://opentelemetry.io/docs/collector/) for more information.

The example configuration used for this project shows how to configure the collector's prometheus receiver to collect metrics from an Envoy endpoint. Note that Envoy provides metrics at a custom path of `/stats/prometheus` instead of the usual `/metrics` endpoint.

``` yaml
receivers:
  otlp:
    protocols:
      http:
      grpc:
  prometheus/front-proxy:
    config:
      scrape_configs:
        - job_name: otel-envoy-eg
          scrape_interval: 5s
          metrics_path: /stats/prometheus
          static_configs:
            - targets: ["front-envoy:8001"]

exporters:
  logging:
    loglevel: debug
  otlp/public:
    endpoint: ingest.lightstep.com:443
    headers:
        "lightstep-access-token": "${LS_ACCESS_TOKEN}"

processors:
  batch:

service:
  pipelines:
    metrics:
      receivers: [otlp, prometheus/front-proxy]
      processors: [batch]
      exporters: [otlp/public]

```

## Configuration: Traces

This example forwards traces from Envoy to a locally-running collector using gRPC. See `service-envoy.yaml` for configuration details.

## Additional resources

* Envoy v1.23.0 [OpenTelemetry tracer](https://www.envoyproxy.io/docs/envoy/v1.23.0/api-v3/config/trace/v3/opentelemetry.proto.html?highlight=opentelemetry) docs.
* This example is based on [Envoy's front-proxy sandbox](https://www.envoyproxy.io/docs/envoy/latest/start/sandboxes/front_proxy). See [Envoy's sandboxes](https://www.envoyproxy.io/docs/envoy/latest/start/sandboxes/) for alternative examples.
