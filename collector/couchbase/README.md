# Ingest Couchbase metrics using OTEL Collector's Prometheus receiver

This example illustrates how you can ingest Couchbase metrics using the OTEL collector's Prometheus receiver. Couchbase exposes a Prometheus compatible metrics endpoint.

## Requirements

* OpenTelemetry Collector Contrib v0.51.0+

## Prerequisites

You must have a Lightstep Observability [access token](/docs/create-and-manage-access-tokens) for the project to report metrics.

## Running the Example

You can run this example with `docker compose up` or `make up` in this directory.

## Configuration

The example configuration used for this project shows how to configure OTEL's prometheus receiver to collect metrics from a Couchbase endpoint. Note that the Prometheus receiver in the OTEL project provides more configuration options, but we can use defaults for most of them with Couchbase.

``` yaml
receivers:
  otlp:
    protocols:
      http:
      grpc:
  prometheus/couchbase:
    config:
      scrape_configs:
        - job_name: otel-couchbase-eg
          static_configs:
            - targets: ["couchbase:8091"]

exporters:
  logging:
    loglevel: debug
  otlp/public:
    endpoint: ingest.lightstep.com:443
    headers:
        "lightstep-access-token": "${LS_ACCESS_TOKEN}"
  file:
    path: statsout/collector-out.json

processors:
  batch:

service:
  pipelines:
    metrics:
      receivers: [otlp, prometheus/couchbase]
      processors: [batch]
      exporters: [otlp/public, file]
```


## Additional resources

* For information about configuring Couchbase to provide monitoring compatible with the OTEL Collector's Prometheus receiver see the [Couchbase Documentation](https://docs.couchbase.com/operator/current/howto-prometheus.html).
