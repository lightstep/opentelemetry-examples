---
# Ingest metrics using the Kubernetes integration

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
kind create cluster --config kind-config.yaml
```

### 2. Install Prometheus Node exporter and Kube state metrics
`
```bash
helm install kube-state-metrics prometheus-community/kube-state-metrics -n kube-system --version 5.6.2
helm install prometheus-node-exporter prometheus-community/prometheus-node-exporter -n kube-system --version 4.17.2
```

### 3. Install Collector

As we have Fluentd instances at every node in the cluster, we need to collect metrics from all of the instances as well, that's why we're going to spin up Collectors as Daemon Set as well adn each collector is going to connect to Fluentd instance from the same node.

```bash
kubectl apply -f collector-configmap.yaml
```

For DEV env only apply role binding to get access to kubelet metrics.
```bash
kubectl apply -f collector-rbac.yaml
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

Detailed description of available Kubernetes metrics: [Kubelet metrics](https://docs.fluentd.org/monitoring-fluentd/monitoring-prometheus), [Node metrics](https://github.com/lightstep/opentelemetry-examples/blob/main/collector/kubernetes#L2), [Pods metrics](https://github.com/kubernetes/kube-state-metrics/blob/main/docs/pod-metrics.md).

Collector Prometheus receiver has to be pointed to the Kubernetes Prometheus metrics endpoints.

The following example configuration collects metrics from Kong and send them to Lightstep Observability:

```yaml
receivers:
  prometheus:
    config:
      scrape_configs:
        - job_name: otel-kubelet
          scrape_interval: 10s
          scheme: https
          tls_config:
            insecure_skip_verify: true
          bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token
          static_configs:
            - targets: ["${K8S_NODE_NAME}:10250"]
        - job_name: otel-nodes
          scrape_interval: 10s
          static_configs:
            - targets: ["prometheus-node-exporter.kube-system:9100"]
        - job_name: otel-pods
          scrape_interval: 10s
          static_configs:
            - targets: ["kube-state-metrics.kube-system:8080"]
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
    metrics:
      receivers: [prometheus]
      processors: [batch]
      exporters: [logging, otlp/public]
```

