# Ingest Consul metrics using the OpenTelemetry Collector

## Overview

 Consul natively exposes a Prometheus endpoint and the OpenTelemetry Collector has a [Prometheus receiver][otel-prom-receiver] that can be used to scrape its Prometheus endpoint. This directory contains an example showing how to configure Consul and the Collector to send metrics to Lightstep Observability.

 This example is based on the [service mesh][consul-service-mesh-example-repo] example from the [learn Consul repository][learn-consul-repo]. See the [repository][consul-service-mesh-example-repo], or this [Consul service mesh tutorial][consul-service-mesh-example-docs] for more about the underlying Consul setup in the example.

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
* Install services
  ```
  chmod 755 service-install.sh
  ./service-install.sh
  ```
* Explore the example
* Clean up
  ```
  docker-compose down --rmi all`
  ```

### Explore Metrics in Lightstep

See the [Consul Telemetry Docs][consul-docs-telemetry] for comprehensive documentation on metrics emitted. Note that the the metrics collected via prometheus will be renamed to conform to prometheus conventions; dots will be replaced with underscores (e.g. consul.kvs.apply will be renamed to consul_kvs_apply). Consul metrics will be prefixed with "consul_" and you can build dashboards with them in Lightstep Observability. See the [dashboard documentation][ls-docs-dashboards] for more details.

### Explore the Consul Example

* The Consul UI is available at [http://localhost:8500/ui](http://localhost:8500/ui/).
* See the Counting Service at [http://localhost:9002](http://localhost:9002)
* Inspect the Prometheus endpoint at [http://localhost:8500/v1/agent/metrics?format=prometheus](http://localhost:8500/v1/agent/metrics?format=prometheus)


## Configure Consul

Consul has native support for Prometheus, but you need to set the `prometheus_retention_time` configuration option to a number greater than zero for it to be enabled. The documentation also recommends setting `disable_hostname` to `true` to prevent metrics from being prefixed with hostname. See the Consul Server configuration example below in HCL format. Adapt as needed for other formats (json, yaml, etc).

```hcl
telemetry {
  prometheus_retention_time = "60s"
  disable_hostname = true
}
```

## Configure the Collector

Below is a snippet showing how to configure the Prometheus Receiver to scrape the Prometheus endpoint exposed by the Consul Server.

```yaml
receivers:
  prometheus:
    config:
      scrape_configs:
        - job_name: 'consul-server'
          metrics_path: '/v1/agent/metrics'
          params:
            format: ['prometheus']
          static_configs:
            - targets: ['localhost:8500']
```



## Additional information

- [OpenTelemetry Collector Prometheus Receiver][otel-prom-receiver]
- [Consul Telemetry Reference][consul-docs-telemetry]

[ls-docs-access-token]: https://docs.lightstep.com/docs/create-and-manage-access-tokens
[ls-docs-dashboards]: https://docs.lightstep.com/docs/create-and-manage-dashboards
[otel-prom-receiver]: https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/prometheusreceiver
[consul-docs-telemetry]: https://www.consul.io/docs/agent/telemetry#key-metrics
[consul-service-mesh-example-docs]: https://learn.hashicorp.com/tutorials/consul/service-mesh-with-envoy-proxy
[consul-service-mesh-example-repo]: https://github.com/hashicorp/learn-consul-docker/tree/main/datacenter-deploy-service-mesh/config-entries
[learn-consul-repo]: https://github.com/hashicorp/learn-consul-docker
