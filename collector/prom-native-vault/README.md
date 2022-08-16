# Ingest Hashicorp Vault metrics using the OpenTelemetry Collector

## Overview

 Hashicorp Vault natively exposes a Prometheus endpoint and the OpenTelemetry Collector has a [Prometheus receiver][otel-prom-receiver] that can be used to scrape its Prometheus endpoint. This directory contains an example showing how to configure Hashicorp Vault and the Collector to send metrics to Lightstep Observability.

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
  docker-compose down`
  ```
* Remove Vault data
  ```
  rm -rf config/data/* 
  ```
* Remove Consul data
  ```
  rm -rf consul/config/data/*
  ``` 

### Explore Metrics in Lightstep

See the [Hashicorp Vault Telemetry Docs][hashicorp-vault-docs-telemetry] for comprehensive documentation on metrics emitted and the [dashboard documentation][ls-docs-dashboards] for more details.

### Explore the Hashicorp Vault Example

* Access Vault UI [http://localhost:8280](http://localhost:8280).


## Configure Hashicorp Vault

- Prometheus metrics are not enabled by default; setting the prometheus_retention_time to a non-zero value enables them.

`/vault/config/server.hcl`
```sh
...
telemetry {
  disable_hostname = true
  prometheus_retention_time = "12h"
}
```

Define Prometheus ACL Policy

- The Vault /sys/metrics endpoint is authenticated. Prometheus requires a Vault token with sufficient capabilities to successfully consume metrics from the endpoint.

```sh
vault policy write prometheus-metrics - << EOF
path "/sys/metrics" {
  capabilities = ["read"]
}
EOF
```

## Configure the Collector

Below is a snippet showing how to configure the Prometheus Receiver to scrape the Prometheus endpoint exposed by the Hashicorp Vault Server.

```yaml
receivers:
  prometheus:
    config:
      scrape_configs:
        - job_name: 'vault-server'
          scrape_interval: 10s
          metrics_path: '/v1/sys/metrics'
          params:
            format: ['prometheus']
          static_configs:
            - targets: ['localhost:8200']

```



## Additional information

- [OpenTelemetry Collector Prometheus Receiver][otel-prom-receiver]
- [Hashicorp Vault Telemetry Reference][hashicorp-vault-docs-telemetry]

[ls-docs-access-token]: https://docs.lightstep.com/docs/create-and-manage-access-tokens
[ls-docs-dashboards]: https://docs.lightstep.com/docs/create-and-manage-dashboards
[otel-prom-receiver]: https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/prometheusreceiver
[hashicorp-vault-docs-telemetry]: https://www.vaultproject.io/docs/internals/telemetry
[learn-consul-repo]: https://github.com/hashicorp/learn-consul-docker