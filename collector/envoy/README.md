# Ingest Envoy metrics using OTEL Collector's Prometheus receiver

This example illustrates how you can ingest Envoy metrics using the OTEL collector's Prometheus receiver. Envoy exposes a Prometheus compatible metrics endpoint.

## Requirements

* OpenTelemetry Collector Contrib v0.51.0+

## Prerequisites

You must have a Lightstep Observability [access token](/docs/create-and-manage-access-tokens) for the project to report metrics to.

## Running the Example

You can run this example with `docker-compose up` in this directory.

## Configuration

Installation of the OpenTelemetry Collector varies, please refer to the [collector documentation](https://opentelemetry.io/docs/collector/) for more information.

The example configuration used for this project shows how to configure OTEL's prometheus receiver to collect metrics from an Envoy endpoint. Note that Envoy provides metrics at a custom path of `/stats/prometheus` instead of the usual `/metrics` endpoint.

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

## Additional resources

* This example is based on [EnvoyProxy's front-proxy sandbox](https://www.envoyproxy.io/docs/envoy/latest/start/sandboxes/front_proxy). See [EnvoyProxy's sandboxes](https://www.envoyproxy.io/docs/envoy/latest/start/sandboxes/) for alternative examples.
