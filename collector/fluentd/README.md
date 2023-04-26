---
# Ingest metrics using the Fluentd ingress integration

The OTel Collector has a variety of [third party receivers](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/master/receiver) that provide integration with a wide variety of metric sources.

Please note that not all metrics receivers available for the OpenTelemetry Collector have been tested by Lightstep Observability, and there may be bugs or unexpected issues in using these contributed receivers with Lightstep Observability metrics. File any issues with the appropriate OpenTelemetry community.
{: .callout}

## Prerequisites for local installation

You must have a Lightstep Observability [access token](/docs/create-and-manage-access-tokens) for the project to report metrics to.

#### kind

Provided example based on kind kubernetes cluster.

## Running the Example

### 1. Create a cluster

First you'll need to create a cluster.

```bash
kind create cluster --name fluentd
```

### 2. Install Fluentd

Adding config map for Fluentd. In this example we're going to spin up Fluentd as daemon set, which will run pods on every nodes in the cluster. And for test case, we're setting export logs to the file. Also we need to enable exporting metrics by using prometheus format.

```bash
kubectl apply -f fluentd-configmap.yaml
kubectl apply -f fluentd-rbac.yaml
kubectl apply -f fluentd.yaml
``` 

### 3. Install Collector

As we have Fluentd instances at every node in the cluster, we need to collect metrics from all of the instances as well, that's why we're going to spin up Collectors as Daemon Set as well adn each collector is going to connect to Fluentd instance from the same node.

```bash
kubectl apply -f collector-configmap.yaml
```

Assuming we have LightStep Access Token in the environment variable at the host machine, we need to create secret in the cluster.

```bash
kubectl create secret generic ls --from-literal=access_token=$LS_ACCESS_TOKEN -n collector
```

Adding collector Daemon Set.

```bash
kubectl apply -f collector.yaml
```

## Configuration

Installation of the OpenTelemetry Collector varies, please refer to the [collector documentation](https://opentelemetry.io/docs/collector/) for more information.

Detailed description of available [Fluentd metrics](https://docs.fluentd.org/monitoring-fluentd/monitoring-prometheus).

Collector Prometheus receiver has to be pointed to the Fluentd metrics endpoint.

The following example configuration collects metrics from Kong and send them to Lightstep Observability:

```yaml
receivers:
  prometheus:
    config:
      scrape_configs:
        - job_name: otel-fluentd
          static_configs:
            - targets: ["${FLUENTD_HOST}:24224"]
exporters:
  logging:
    loglevel: debug
  otlp/public:
    endpoint: ingest.lightstep.com:443
    headers:
        "lightstep-access-token": "${LS_ACCESS_TOKEN}"
processors:
  batch:
service:
  pipelines:
    metrics/fluentd:
      receivers: [prometheus]
      processors: [batch]
      exporters: [logging, otlp/public]
```

