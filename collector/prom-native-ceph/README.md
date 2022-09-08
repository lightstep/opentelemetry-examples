# Ingest Ceph metrics using the OpenTelemetry Collector

## Overview

 Ceph natively exposes a Prometheus endpoint and the OpenTelemetry Collector has a [Prometheus receiver][otel-prom-receiver] that can be used to scrape its Prometheus endpoint. This directory contains an example showing how to configure Ceph and the Collector to send metrics to Lightstep Observability.

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
  docker compose up -d
  ```
* Enabling Prometheus
  ```
  docker exec mon1 ceph --cluster ceph mgr module enable prometheus
  ```
* Disabling Prometheus
  ```
  docker exec mon1 ceph --cluster ceph mgr module disable prometheus
  ```
* Check Ceph Health
  ```
  docker exec mon1 ceph --cluster ceph -s
  ```
* Stop the cluster
  ```
  docker compose down
  ```

### Explore Metrics in Lightstep

See the [Ceph Telemetry Docs][ceph-docs-prometheus] for comprehensive documentation on metrics emitted and the [dashboard documentation][ls-docs-dashboards] for more details.

### Explore the Ceph Example

* Access Vault UI [http://localhost:8280](http://localhost:8280).


## Configure Ceph

- Prometheus Module provides a Prometheus exporter to pass on Ceph performance counters from the collection point in ceph-mgr. Ceph-mgr receives MMgrReport messages from all MgrClient processes (mons and OSDs, for instance) with performance counter schema data and actual counter data, and keeps a circular buffer of the last N samples. This module creates an HTTP endpoint (like all Prometheus exporters) and retrieves the latest sample of every counter when scraped.

```sh
$ sudo ceph mgr module enable prometheus
```
```sh
$ sudo ceph mgr module disable prometheus
```

## Enabling Prometheus Output  

The Prometheus manager module needs to be restarted for configuration changes to be applied.

- By default the module will accept HTTP requests on port 9283 on all IPv4 and IPv6 addresses on the host. The port and listen address are both configurable with ceph config set, with keys mgr/prometheus/server_addr and mgr/prometheus/server_port. 

```sh
$ sudo ceph config set mgr mgr/prometheus/server_addr 0.0.0.0
```
```sh
$ sudo ceph config set mgr mgr/prometheus/server_port 9283
```

## Configure the Collector

Below is a snippet showing how to configure the Prometheus Receiver to scrape the Prometheus endpoint exposed by the Ceph Server.

```yaml
 prometheus:
    config:
      scrape_configs:
        - job_name: 'ceph-mgr'
          scrape_interval: 15s
          metrics_path: '/metrics'
          static_configs:
            - targets: ['mgr1:9283']

```

## Additional information

- [OpenTelemetry Collector Prometheus Receiver][otel-prom-receiver]
- [Ceph Telemetry Reference][ceph-docs-prometheus]

[ls-docs-access-token]: https://docs.lightstep.com/docs/create-and-manage-access-tokens
[ls-docs-dashboards]: https://docs.lightstep.com/docs/create-and-manage-dashboards
[otel-prom-receiver]: https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/prometheusreceiver
[ceph-docs-prometheus]: https://docs.ceph.com/en/quincy/mgr/prometheus/