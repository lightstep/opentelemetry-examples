---
# Ingest metrics using the PHP-FPM integration

The OTel Collector has a variety of [third party receivers](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/master/receiver) that provide integration with a wide variety of metric sources.

Please note that not all metrics receivers available for the OpenTelemetry Collector have been tested by Lightstep Observability, and there may be bugs or unexpected issues in using these contributed receivers with Lightstep Observability metrics. File any issues with the appropriate OpenTelemetry community.
{: .callout}

## Prerequisites

You must have a Lightstep Observability [access token](/docs/create-and-manage-access-tokens) for the project to report metrics to.

## Running the Example

You can run this example with `docker-compose up` in this directory.

## Configuration

Installation of the OpenTelemetry Collector varies, please refer to the [collector documentation](https://opentelemetry.io/docs/collector/) for more information.

PHP-FPM integration requires enabling status page and the use of [PHP-FPM Exporter](https://github.com/hipages/php-fpm_exporter).

PHP-FPM metrics are described [here](https://github.com/hipages/php-fpm_exporter#metrics-collected).

The example collector's configuration, used for this project shows using processors to add metrics with Lightstep Observability:

``` yaml
receivers:
  prometheus:
    config:
      scrape_configs:
        - job_name: otel-php-fpm
          static_configs:
            - targets: [php-fpm-exporter:9253]

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
    metrics/php-fpm:
      receivers: [prometheus]
      processors: [batch]
      exporters: [logging, otlp/public]
```
