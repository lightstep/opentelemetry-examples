# Ingest Clickhouse metrics using the OpenTelemetry Collector

## Overview

 Clickhouse natively exposes a Prometheus endpoint and the OpenTelemetry Collector has a [Prometheus receiver][otel-prom-receiver] that can be used to scrape its Prometheus endpoint. This directory contains an example showing how to configure Clickhouse and the Collector to send metrics to Lightstep Observability.

## Prerequisites

* Docker
* Docker Compose
* A Lightstep Observability [access token][ls-docs-access-token]

## How to run the example

* Export your Lightstep access token
  ```
  export LS_ACCESS_TOKEN=<YOUR_TOKEN>
  ```
* Run the docker compose example
  ```
  docker-compose up -d
  ```
* Run clickhouse client
  ```
  make run-client
  ```
  * Test DB
  ```
  CREATE DATABASE IF NOT EXISTS tutorial
  CREATE TABLE tutorial.xyz (a UInt8, d Date) ENGINE = MergeTree() ORDER BY (a) PARTITION BY toYYYYMM(d);
  INSERT INTO tutorial.xyz values(8,'2022-07-25');
  SELECT * FROM tutorial.xyz;
  ```
* Clean up
  ```
  docker-compose down`
  ```

### Explore Metrics in Lightstep

See the [Clickhouse Telemetry Docs][clickhouse-docs-telemetry] for comprehensive documentation on metrics emitted and the [dashboard documentation][ls-docs-dashboards] for more details.

### Explore the clickhouse Example

* The Clickhouse UI is available at [http://127.0.0.1:8123/play?user=default](http://127.0.0.1:8123/play?user=default).


## Configure clickhouse

clickhouse has native support for Prometheus, but you need to set the `prometheus_retention_time` configuration option to a number greater than zero for it to be enabled. The documentation also recommends setting `disable_hostname` to `true` to prevent metrics from being prefixed with hostname. See the clickhouse Server configuration example below in HCL format. Adapt as needed for other formats (json, yaml, etc).

```hcl
telemetry {
  prometheus_retention_time = "60s"
  disable_hostname = true
}
```

## Configure the Collector

Below is a snippet showing how to configure the Prometheus Receiver to scrape the Prometheus endpoint exposed by the Clickhouse Server.

```yaml
receivers:
  prometheus:
    config:
      scrape_configs:
        - job_name: 'clickhouse-server'
          metrics_path: '/v1/agent/metrics'
          params:
            format: ['prometheus']
          static_configs:
            - targets: ['localhost:8123']
```



## Additional information

- [OpenTelemetry Collector Prometheus Receiver][otel-prom-receiver]
- [Clickhouse Telemetry Reference][clickhouse-docs-telemetry]

[ls-docs-access-token]: https://docs.lightstep.com/docs/create-and-manage-access-tokens
[ls-docs-dashboards]: https://docs.lightstep.com/docs/create-and-manage-dashboards
[otel-prom-receiver]: https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/prometheusreceiver
[clickhouse-docs-telemetry]: https://clickhouse.com/docs/en/operations/opentelemetry/
[learn-clickhouse-repo]: https://github.com/ClickHouse/ClickHouse/blob/master/docker/server/README.md