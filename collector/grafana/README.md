# Ingest Grafana metrics using the OpenTelemetry Collector

## Overview

Grafana typically uses Prometheus as its primary data source. The OpenTelemetry Collector has a [Prometheus receiver][otel-prom-receiver]   that can be used to scrape the Grafana server's Prometheus endpoint. This directory contains an example showing how to configure Grafana and the Collector to send metrics to Lightstep Observability.

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
* Access to the Grafana dashboard from the following link: http://localhost:3000/

* Explore Metrics in Lightstep

* Stop the cluster
  ```
  docker-compose down`
  ```


## Configure Grafana

In this example, we configured Grafana with a single node. You can also set up a Grafana cluster and fill the targets section in the collector config.

## Configure the Collector

Below is a snippet showing how to configure the Prometheus Receiver to scrape the Prometheus endpoint exposed by the Grafana Server.

```yaml
receivers:
  prometheus:
    config:
      scrape_configs:
        - job_name: 'grafana_metrics'
          scrape_interval: 15s
          scrape_timeout: 5s
          params:
            format: [ 'prometheus' ]
          scheme: 'http'
          tls_config:
            insecure_skip_verify: true
          static_configs:
            - targets: [ 'grafana:3000' ]
```



## Additional information

- [OpenTelemetry Collector Prometheus Receiver][otel-prom-receiver]
- [Grafana documentation][grafana-docs]

[ls-docs-access-token]: https://docs.lightstep.com/docs/create-and-manage-access-tokens
[otel-prom-receiver]: https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/prometheusreceiver
[grafana-docs]: https://grafana.com/docs/
