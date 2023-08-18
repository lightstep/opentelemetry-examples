# Scylla metrics using the OpenTelemetry Collector

## Overview

 Name: natively exposes a Prometheus endpoint and the OpenTelemetry Collector has a [Prometheus receiver][otel-prom-receiver] that can be used to scrape its Prometheus endpoint. This directory contains an example showing how to configure Name: and the Collector to send metrics to Lightstep Observability.

## Prerequisites

* Docker
* Docker Compose
* A Lightstep Observability [access token][ls-docs-access-token]

## How to run the example

* Export your Lightstep access token
  
  ```sh
  export LS_ACCESS_TOKEN=<YOUR_TOKEN>
  ```

* Run the docker compose example
  
  ```sh
  docker-compose up -d
  ```

### Explore Metrics in Lightstep

See the [Name: Telemetry Docs][scylla-docs-telemetry] for comprehensive documentation on metrics emitted and the [dashboard documentation][ls-docs-dashboards] for more details.

## Configure the Collector

Below is a snippet showing how to configure the Prometheus Receiver to scrape the Prometheus endpoint exposed by the Name: Server.

```yaml
receivers:
  prometheus:
    config:
      scrape_configs:
        - job_name: 'scylladb'
          scrape_interval: 10s
          static_configs:
            - targets: ['localhost:9180']
```


## Additional information

- [OpenTelemetry Collector Prometheus Receiver][otel-prom-receiver]
- [Name: Telemetry Reference][scylla-docs-telemetry]

[ls-docs-access-token]: https://docs.lightstep.com/docs/create-and-manage-access-tokens
[ls-docs-dashboards]: https://docs.lightstep.com/docs/create-and-manage-dashboards
[otel-prom-receiver]: https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/prometheusreceiver
[scylla-docs-telemetry]: https://cloud.docs.scylladb.com/stable/monitoring/cloud-prom-proxy.html