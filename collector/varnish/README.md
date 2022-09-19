---
# Ingest metrics using the Varnish integration

The OTel Collector has a variety of [third party receivers](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/master/receiver) that provide integration with a wide variety of metric sources.

Please note that not all metrics receivers available for the OpenTelemetry Collector have been tested by Lightstep Observability, and there may be bugs or unexpected issues in using these contributed receivers with Lightstep Observability metrics. File any issues with the appropriate OpenTelemetry community.
{: .callout}

## Requirements

* OpenTelemetry Collector Contrib v0.51.0+

## Prerequisites

You must have a Lightstep Observability [access token](/docs/create-and-manage-access-tokens) for the project to report metrics to.

## Running the Example

You can run this example with `docker-compose up` in this directory.

## Configuration

Installation of the OpenTelemetry Collector varies, please refer to the [collector documentation](https://opentelemetry.io/docs/collector/) for more information.

Varnish integration requires using 3rd party plugin [prometheus varnish exporter](https://github.com/jonnenauha/prometheus_varnish_exporter) as shown in the Dockerfile. Varnish metrics described [here](https://docs.varnish-software.com/tutorials/monitoring/#counters).

The example configuration, used for this project shows using processors to add metrics with Lightstep Observability, add the following to your collector's configuration file:

``` yaml
# add the receiver configuration for your integration
receivers:
  prometheus:
    config:
      scrape_configs:
        - job_name: otel-varnish
          static_configs:
            - targets: [varnish:9131]
  nginx/appsrv:
    endpoint: 'http://nginx_appsrv:1080/status'
    collection_interval: 10s

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
    metrics/varnish:
      receivers: [prometheus]
      processors: [batch]
      exporters: [logging, otlp/public]
    metrics/appsrv:
      receivers: [nginx/appsrv]
      processors: [batch]
      exporters: [logging, otlp/public]
```

