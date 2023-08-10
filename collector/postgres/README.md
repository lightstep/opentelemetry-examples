---
# Ingest metrics using the POSTGRE SQL integration

The OTel Collector has a variety of [third party receivers](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/master/receiver) that provide integration with a wide variety of metric sources.

Please note that not all metrics receivers available for the OpenTelemetry Collector have been tested by Cloud Observability Observability, and there may be bugs or unexpected issues in using these contributed receivers with Cloud Observability Observability metrics. File any issues with the appropriate OpenTelemetry community.
{: .callout}

## Prerequisites

You must have a Cloud Observability Observability [access token](/docs/create-and-manage-access-tokens) for the project to report metrics to.

## Running the Example

You can run this example with `docker-compose up` in this directory.

## Configuration

Installation of the OpenTelemetry Collector varies, please refer to the [collector documentation](https://opentelemetry.io/docs/collector/) for more information.

The list of Postgre SQL metrics are described [here](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/receiver/postgresqlreceiver/metadata.yaml#L76).

The example collector's configuration, used for this project shows using processors to add metrics with Cloud Observability Observability:

``` yaml
receivers:
  postgresql:
    endpoint: postgres:5432
    username: postgres
    password: postgres
    databases:
      - postgres
    collection_interval: 5s
    tls:
      insecure: true

processors:
  batch:

exporters:
  logging:
    loglevel: debug
  otlp:
    endpoint: ingest.lightstep.com:443
    headers:
             "lightstep-access-token": "${LS_ACCESS_TOKEN}"

service:
  telemetry:
    logs:
      level: "debug"
  pipelines:
    metrics:
     receivers: [postgresql]
     processors: [batch]
     exporters: [logging,otlp]
```
