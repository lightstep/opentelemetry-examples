# Ingest Cockroachdb metrics using the OpenTelemetry Collector

## Overview

 Cockroachdb natively exposes a Prometheus endpoint and the OpenTelemetry Collector has a [Prometheus receiver][otel-prom-receiver] that can be used to scrape its Prometheus endpoint. This directory contains an example showing how to configure Cockroachdb and the Collector to send metrics to Lightstep Observability.

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
* Start the SQL shell in the first container
  ```
  docker exec -it roach1 ./cockroach sql --insecure
  ```
  * Run CockroachDB SQL statements
  ```
  CREATE DATABASE bank;
  ```
  ```
  CREATE TABLE bank.accounts (id INT PRIMARY KEY, balance DECIMAL);
  ```
  ```
  INSERT INTO bank.accounts VALUES (1, 1000.50);
  ```
  ```
  SELECT * FROM bank.accounts;
  ```
    * Exit shell
  ```
  \q
  ```

* Stop the cluster
  ```
  docker-compose down`
  ```

### Explore Metrics in Lightstep

See the [Cockroachdb Telemetry Docs][cockroachdb-docs-telemetry] for comprehensive documentation on metrics emitted and the [dashboard documentation][ls-docs-dashboards] for more details.

### Explore the Cockroachdb Example

* Access the DB Console [http://localhost:8080](http://localhost:8080).


## Configure Cockroachdb

CockroachDB generates detailed time series metrics for each node in a cluster or a single node.

## Configure the Collector

Below is a snippet showing how to configure the Prometheus Receiver to scrape the Prometheus endpoint exposed by the Cockroachdb Server.

```yaml
​​receivers:
  prometheus:
    config:
      scrape_configs:
        - job_name: 'cockroachdb'
          scrape_interval: 3s
          metrics_path: '/_status/vars'
          params:
            format: ['prometheus']
          scheme: 'http'
          tls_config:
            insecure_skip_verify: true
          static_configs:
            - targets: [locahost:8080']

```



## Additional information

- [OpenTelemetry Collector Prometheus Receiver][otel-prom-receiver]
- [Cockroachdb Telemetry Reference][cockroachdb-docs-telemetry]

[ls-docs-access-token]: https://docs.lightstep.com/docs/create-and-manage-access-tokens
[ls-docs-dashboards]: https://docs.lightstep.com/docs/create-and-manage-dashboards
[otel-prom-receiver]: https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/prometheusreceiver
[cockroachdb-docs-telemetry]: https://www.cockroachlabs.com/docs/v22.1/monitor-cockroachdb-with-prometheus.html/
[learn-cockroachdb-repo]: https://github.com/Cockroachdb/Cockroachdb/blob/master/docker/server/README.md