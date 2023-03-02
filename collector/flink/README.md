# Ingest Flink metrics using the OpenTelemetry Collector

## Overview

This example is inspired by https://github.com/mbode/flink-prometheus-example .

In order to expose Prometheus endpoint, you have to modify flink configuration and also add some Jar libraries. Checkout provided Dockerfile.

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

* Explore Metrics in Lightstep

* Stop the cluster
  ```
  docker-compose down`
  ```

## Configure the Collector

Below is a snippet showing how to configure the Prometheus Receiver to scrape the Prometheus endpoint exposed by the Flink.

```yaml
receivers:
  prometheus:
    config:
      scrape_configs:
        - job_name: 'flink'
          scrape_interval: 1s
          params:
            format: [ 'prometheus' ]
          scheme: 'http'
          tls_config:
            insecure_skip_verify: true
          static_configs:
            - targets: [ 'job-cluster:9249', 'taskmanager1:9249', 'taskmanager2:9249' ]

```



## Additional information

- [OpenTelemetry Collector Prometheus Receiver][otel-prom-receiver]
- [Flink and Prometheus: Cloud-native monitoring of streaming applications][flink-and-prometheus]

[ls-docs-access-token]: https://docs.lightstep.com/docs/create-and-manage-access-tokens
[otel-prom-receiver]: https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/prometheusreceiver
[flink-and-prometheus]: https://flink.apache.org/features/2019/03/11/prometheus-monitoring.html
