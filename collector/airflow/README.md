---
# Ingest metrics using the Airflow integration

The OTel Collector has a variety of [third party receivers](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/master/receiver) that provide integration with a wide variety of metric sources.

{: .callout}

## Prerequisites

You must have a Lightstep Observability [access token](/docs/create-and-manage-access-tokens) for the project to report metrics to.

## Running the Example

You can run this example with `docker-compose up` in this directory.

## Configuration

Installation of the OpenTelemetry Collector varies, please refer to the [collector documentation](https://opentelemetry.io/docs/collector/) for more information.

Airflow integration requires enabling statsd metrics in airflow.cfg:
```
[metrics]
statsd_on = True
statsd_host = otel-collector
statsd_port = 8125
statsd_prefix = airflow
```
or by providing following enivronment variables:
```
AIRFLOW__METRICS__STATSD_ON=True
AIRFLOW__METRICS__STATSD_HOST=otel-collector
AIRFLOW__METRICS__STATSD_PORT=8125
AIRFLOW__METRICS__STATSD_PREFIX=airflow
```

Airflow metrics described [here](https://airflow.apache.org/docs/apache-airflow/stable/administration-and-deployment/logging-monitoring/metrics.html).

The example collector's configuration, used for this project shows using processors to add metrics with Lightstep Observability:

```yaml
receivers:
  statsd:
    endpoint: "otel-collector:8125"
    aggregation_interval: 60s
    is_monotonic_counter: true
    timer_histogram_mapping:
      - statsd_type: "histogram"
        observer_type: "summary"
      - statsd_type: "timing"
        observer_type: "summary"

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
      receivers: [statsd]
      processors: [batch]
      exporters: [logging, otlp/public]
```
