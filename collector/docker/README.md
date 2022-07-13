# Ingest Docker metrics using OTEL Collector's Prometheus receiver

This example illustrates how you can ingest Docker metrics using the OTEL collector's Prometheus receiver. Docker exposes a Prometheus compatible metrics endpoint.

## Requirements

* OpenTelemetry Collector Contrib v0.51.0+

## Prerequisites

You must have a Lightstep Observability [access token](/docs/create-and-manage-access-tokens) for the project to report metrics to.

## Running the Example

You can run this example with `docker-compose up` in this directory. Using `docker-compose --profile loadgen up` also creates an instance to send requests to the NGINX service. You'll want to view this in Lightstep with a dashboard. 

## Configuration

Installation of the OpenTelemetry Collector varies, please refer to the [collector documentation](https://opentelemetry.io/docs/collector/) for more information.

The example configuration used for this project shows how to configure OTEL's prometheus receiver to collect metrics from a Docker endpoint. Note that Docker provides metrics at a custom path of `/stats/prometheus` instead of the usual `/metrics` endpoint.

``` yaml
receivers:
  otlp:
    protocols:
      http:
      grpc:
  prometheus/front-proxy:
    config:
      scrape_configs:
        - job_name: otel-docker-eg
          scrape_interval: 5s
          metrics_path: /stats/prometheus
          static_configs:
            - targets: ["docker:8001"]

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

* This example is based on [dockerProxy's front-proxy sandbox](https://www.dockerproxy.io/docs/docker/latest/start/sandboxes/front_proxy). See [dockerProxy's sandboxes](https://www.dockerproxy.io/docs/docker/latest/start/sandboxes/) for alternative examples.
