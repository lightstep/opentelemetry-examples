# Ingest Micrometer metrics using the OpenTelemetry Collector

## Overview

 Micrometer natively exposes a Prometheus endpoint and the OpenTelemetry Collector has a [Prometheus receiver][otel-prom-receiver] that can be used to scrape its Prometheus endpoint. This directory contains an example showing how to configure Micrometer and the Collector to send metrics to Lightstep Observability.

## Prerequisites

* Docker
* Docker Compose
* A Lightstep Observability [access token][ls-docs-access-token]

### Java application with Spring framework and a Postgres database

Example structure:
```
.
├── backend
│   ├── Dockerfile
│   └── ...
├── db
│   └── password.txt
├── docker-compose.yaml
├── collector.yaml
└── README.md

```

## How to run the example

* Export your Lightstep access token
  ```
  export LS_ACCESS_TOKEN=<YOUR_TOKEN>
  ```
* Run the docker compose example
  ```
  docker-compose up -d
  ```
* Stop the cluster
  ```
  docker-compose down`
  ```

### Explore Metrics in Lightstep

See the [Micrometer Telemetry Docs][micrometer-prometheus-docs] for comprehensive documentation on metrics emitted and the [dashboard documentation][ls-docs-dashboards] for more details.

### Explore the Micrometer Example

* After the application starts, navigate to [http://localhost:8080](http://localhost:8080).


## Configure Micrometer

Micrometer generates detailed time series metrics for each node in a cluster or a single node.

## Configure the Collector

Below is a snippet showing how to configure the Prometheus Receiver to scrape the Prometheus endpoint exposed by the Micrometer Server.

```yaml
receivers:
  prometheus:
    config:
      scrape_configs:
        - job_name: 'micrometer-demo'
          scrape_interval: 10s
          scrape_timeout: 10s
          metrics_path: '/actuator/prometheus'
          scheme: 'http'
          tls_config:
            insecure_skip_verify: true
          static_configs:
            - targets: ['backend:8080']
```



## Additional information

- [OpenTelemetry Collector Prometheus Receiver][otel-prom-receiver]
- [Micrometer Telemetry Reference][micrometer-prometheus-docs]

[ls-docs-access-token]: https://docs.lightstep.com/docs/create-and-manage-access-tokens
[ls-docs-dashboards]: https://docs.lightstep.com/docs/create-and-manage-dashboards
[otel-prom-receiver]: https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/prometheusreceiver
[micrometer-prometheus-docs]: https://micrometer.io/docs/registry/prometheus/
[learn-Micrometer-repo]: https://github.com/Micrometer/Micrometer/blob/master/docker/server/README.md