# Ingest Hashicorp Nomad metrics using the OpenTelemetry Collector

## Overview

 Hashicorp Nomad natively exposes a Prometheus endpoint and the OpenTelemetry Collector has a [Prometheus receiver][otel-prom-receiver] that can be used to scrape its Prometheus endpoint. This directory contains an example showing how to configure Hashicorp Nomad and the Collector to send metrics to Lightstep Observability.

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
* Stop the cluster
  ```
  docker-compose down
  ```

### Explore Metrics in Lightstep

See the [Hashicorp Nomad Telemetry Docs][hashicorp-nomad-docs-telemetry] for comprehensive documentation on metrics emitted and the [dashboard documentation][ls-docs-dashboards] for more details.

### Explore the Hashicorp Nomad Example

* Access Nomad UI [http://localhost:4646/ui](http://localhost:4646/ui).


## Configure Hashicorp Nomad

- Nomad operator deploys Prometheus to collect metrics from a Nomad cluster. The operator enables telemetry on the Nomad servers and clients as well as configure Prometheus.

#### Enable telemetry on Nomad servers and clients

`./nomad/config/local.json`
```json
...
"telemetry": {
        "publish_allocation_metrics": true,
        "publish_node_metrics": true,
        "collection_interval": "1s",
        "disable_hostname": true,
        "prometheus_metrics": true
    }
```

These telemetry parameters apply to Prometheus.

`prometheus_metrics (bool: false)` - Specifies whether the agent should make Prometheus formatted metrics available at /v1/metrics?format=prometheus

curl -s ['localhost:4646/v1/metrics?format=prometheus']('localhost:4646/v1/metrics?format=prometheus')

## Configure the Collector

Below is a snippet showing how to configure the Prometheus Receiver to scrape the Prometheus endpoint exposed by the Hashicorp Nomad Server.

```yaml
receivers:
  prometheus:
    config:
      scrape_configs:
        - job_name: 'nomad-server'
          scrape_interval: 10s
          metrics_path: '/v1/metrics'
          params:
            format: ['prometheus']
          static_configs:
            - targets: ['nomad-server:4646']

```



## Additional information

- [OpenTelemetry Collector Prometheus Receiver][otel-prom-receiver]
- [Hashicorp Nomad Telemetry Reference][hashicorp-Nomad-docs-telemetry]

[ls-docs-access-token]: https://docs.lightstep.com/docs/create-and-manage-access-tokens
[ls-docs-dashboards]: https://docs.lightstep.com/docs/create-and-manage-dashboards
[otel-prom-receiver]: https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/prometheusreceiver
[hashicorp-nomad-docs-telemetry]: https://www.nomadproject.io/docs/configuration/telemetry.html
[learn-consul-repo]: https://github.com/hashicorp/learn-consul-docker