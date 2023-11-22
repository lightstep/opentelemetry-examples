---
# Ingest metrics using the Kong ingress integration

The OTel Collector has a variety of [third party receivers](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/master/receiver) that provide integration with a wide variety of metric sources.

Please note that not all metrics receivers available for the OpenTelemetry Collector have been tested by Cloud Observability, and there may be bugs or unexpected issues in using these contributed receivers with Cloud Observability metrics. File any issues with the appropriate OpenTelemetry community.
{: .callout}

## Prerequisites for local installation

You must have a Cloud Observability [access token](/docs/create-and-manage-access-tokens) for the project to report metrics to.

#### kind

You can use any approach to managing your cluster, but the Makefile builds a cluster in `kind`.

#### helm

We use helm charts to install all apps in this example.

## Running the Example

You can run this example with `make all` in this directory.
After tests you just need to run `make delete-cluster`.

## Steps

### 1. Create a cluster

First you'll need to create a cluster by a method of your choice. `kind create cluster --config kind-config.yaml` works well for local development on Linux.

```yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
  - role: control-plane
    kubeadmConfigPatches:
      - |
        kind: InitConfiguration
        nodeRegistration:
          kubeletExtraArgs:
            node-labels: "ingress-ready=true"
    extraPortMappings:
      - containerPort: 80
        hostPort: 80
        protocol: TCP
      - containerPort: 443
        hostPort: 443
        protocol: TCP
```

### 2. Required libraries

Kong requires Contour to be installed. You also nedd to add the Helm repo for otel-collector.

### 3. Installation

#### a. Install Contour components

Contour components are required by Kong in kind cluster.

```sh
helm install my-contour bitnami/contour --namespace projectcontour --create-namespace
```

#### c. Install Kong ingress

Kong requires [Prometheus plugin](https://docs.konghq.com/hub/kong-inc/prometheus/#example-config) to expose metrics.
In this example Kong is configured with the plugin through helm chart by providing: `-set metrics.enabled=true` parameter.

```sh
helm install my-kong --set metrics.enabled=true bitnami/kong
```

#### e. Collector installation

Assuming you have already added the chart repo, you can install the Collector it with the helm chart like this.

```sh
helm install my-collector open-telemetry/opentelemetry-collector -f values-collector.yaml
```

## Configuration

Installation of the OpenTelemetry Collector varies, please refer to the [collector documentation](https://opentelemetry.io/docs/collector/) for more information.

Detailed description of available [Kong ingress metrics](https://docs.konghq.com/hub/kong-inc/prometheus/#available-metrics).

Collector Prometheus receiver has to be pointed to the Kong metrics endpoint.

The following example configuration collects metrics from Kong and send them to Cloud Observability:

```yaml
# add the receiver configuration for your integration
receivers:
  prometheus:
    config:
      scrape_configs:
        - job_name: otel-kong
          static_configs:
            - targets: [my-kong-metrics:9119]

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

