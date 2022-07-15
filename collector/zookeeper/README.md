# Ingest ZooKeeper metrics using OTEL Collector's Prometheus receiver

This example illustrates how you can ingest ZooKeeper metrics using the OTEL collector's Prometheus receiver. ZooKeeper exposes a Prometheus compatible metrics endpoint.

## Requirements

* OpenTelemetry Collector Contrib v0.51.0+

## Prerequisites

You must have a Lightstep Observability [access token](/docs/create-and-manage-access-tokens) for the project to report metrics.

## Running the Example

You can run this example with `docker compose up` or `make up` in this directory.

## Configuration

Installation of the OpenTelemetry Collector varies, please refer to the [collector documentation](https://opentelemetry.io/docs/collector/) for more information.

The example configuration used for this project shows how to configure OTEL's prometheus receiver to collect metrics from an ZooKeeper endpoint. Note that the Prometheus receiver in the OTEL project provides more configuration options, but we can use defaults for most of them with ZooKeeper.

``` yaml
receivers:
  prometheus/zookeeper:
    config:
      scrape_configs:
        - job_name: otel-zookeeper-eg
          static_configs:
            - targets: ["zookeeper:7000"]

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
      receivers: [prometheus/zookeeper]
      processors: [batch]
      exporters: [otlp/public, logging]

```

You must enable the Prometheus MetricsProvider in Zookeeper by setting 
```bash
metricsProvider.className=org.apache.zookeeper.metrics.prometheus.PrometheusMetricsProvider
```
in the zoo.cfg.

## Additional resources

* For information about configuring Zookeeper to provide monitoring compatible with the OTEL Collector's Prometheus receiver see the [Zookeeper Monitor Guide](https://zookeeper.apache.org/doc/r3.7.0/zookeeperMonitor.html#Prometheus).
