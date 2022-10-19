---
# Ingest metrics using the PowerDNS Recursor integration

The OTel Collector has a variety of [third party receivers](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/master/receiver) that provide integration with a wide variety of metric sources.

Please note that not all metrics receivers available for the OpenTelemetry Collector have been tested by Lightstep Observability, and there may be bugs or unexpected issues in using these contributed receivers with Lightstep Observability metrics. File any issues with the appropriate OpenTelemetry community.
{: .callout}

## Prerequisites

You must have a Lightstep Observability [access token](/docs/create-and-manage-access-tokens) for the project to report metrics to.

## Running the Example

You can run this example with `docker-compose up` in this directory.

## Configuration

Installation of the OpenTelemetry Collector varies, please refer to the [collector documentation](https://opentelemetry.io/docs/collector/) for more information.

PowerDNS Recursor integration requires enabling webserver and allowing external access to it:
```
webserver=yes
webserver-address=0.0.0.0
webserver-allow-from=0.0.0.0/0
```

PowerDNS Recursor metrics described [here](https://doc.powerdns.com/recursor/metrics.html#gathered-information).

The example collector's configuration, used for this project shows using processors to add metrics with Lightstep Observability:

``` yaml
receivers:
  prometheus:
    config:
      scrape_configs:
        - job_name: otel-powerdns
          static_configs:
            - targets: [powerdns:8082]

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
    metrics/powerdns:
      receivers: [prometheus]
      processors: [batch]
      exporters: [logging, otlp/public]
```
