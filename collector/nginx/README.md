---
# Ingest metrics using the NGINX integration

The OTEL Collector has a variety of [third party receivers](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/master/receiver) that provide integration with a wide variety of metric sources.

Please note that not all metrics receivers available for the OpenTelemetry Collector have been tested by Lightstep Observability, and there may be bugs or unexpected issues in using these contributed receivers with Lightstep Observability metrics. File any issues with the appropriate OpenTelemetry community.
{: .callout}

## Requirements

* OpenTelemetry Collector Contrib v0.51.0+

## Prerequisites

You must have a Lightstep Observability [access token](/docs/create-and-manage-access-tokens) for the project to report metrics to.

## Running the Example

You can run this example with `docker-compose up` in this directory. Using `docker-compose --profile loadgen up` also creates an instance to send requests to the NGINX service. You'll want to view this in Lightstep with a dashboard. 

## Configuration

Installation of the OpenTelemetry Collector varies, please refer to the [collector documentation](https://opentelemetry.io/docs/collector/) for more information.

The example configuration, used for this project shows using processors to add metrics with Lightstep Observability, add the following to your collector's configuration file:

``` yaml
# add the receiver configuration for your integration
receivers:
  otlp:
    protocols:
      http:
      grpc:
  hostmetrics:
    collection_interval: 10s
    scrapers:
      memory:
      load:
      network:
      paging:
  nginx/proxy:
    endpoint: 'http://nginx_proxy:8080/status'
    collection_interval: 10s
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
  resource/proxy:
    attributes:
    - key: instance.type
      value: "proxy"
      action: insert
  resource/appsrv:
    attributes:
    - key: instance.type
      value: "appsrv"
      action: insert
  batch:

service:
  pipelines:
    metrics/proxy:
      receivers: [otlp, nginx/proxy]
      processors: [resource/proxy]
      exporters: [logging, otlp/public]
    metrics/appsrv:
      receivers: [otlp, nginx/appsrv]
      processors: [resource/appsrv]
      exporters: [logging, otlp/public]

  telemetry:
    logs:
      level: debug
```
