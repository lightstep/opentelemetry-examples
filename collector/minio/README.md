# Ingest Minio metrics using the OpenTelemetry Collector

## Overview

 Minio natively exposes a Prometheus endpoint and the OpenTelemetry Collector has a [Prometheus receiver][otel-prom-receiver] that can be used to scrape its Prometheus endpoint. This directory contains an example showing how to configure Minio and the Collector to send metrics to Lightstep Observability.

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
* Access to minio console from the following link: http://localhost:9001/login

* Explore Metrics in Lightstep

* Stop the cluster
  ```
  docker-compose down`
  ```


## Configure Minio

In this example we configured Minio with a single node. You can also setup a Minio cluster and fill `targets` section in collector config.

## Configure the Collector

Below is a snippet showing how to configure the Prometheus Receiver to scrape the Prometheus endpoint exposed by the Minio Server.

```yaml
receivers:
  prometheus:
    config:
      scrape_configs:
        - job_name: 'minio'
          scrape_interval: 5s
          metrics_path: '/minio/v2/metrics/cluster'
          params:
            format: ['prometheus']
          scheme: 'http'
          tls_config:
            insecure_skip_verify: true
          static_configs:
            - targets: [ 'minio:9000' ]

```



## Additional information

- [OpenTelemetry Collector Prometheus Receiver][otel-prom-receiver]
- [Minio Monitoring and Alerts][minio-monitoring-alerts]

[ls-docs-access-token]: https://docs.lightstep.com/docs/create-and-manage-access-tokens
[otel-prom-receiver]: https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/prometheusreceiver
[minio-monitoring-alerts]: https://min.io/docs/minio/linux/operations/monitoring.html
