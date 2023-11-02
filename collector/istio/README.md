
---
# Ingest metrics using the Istio integration

The OTel Collector has a variety of [third party receivers](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/master/receiver) that provide integration with a wide variety of metric sources.

Please note that not all metrics receivers available for the OpenTelemetry Collector have been tested by Cloud Observability, and there may be bugs or unexpected issues in using these contributed receivers with Cloud Observability metrics. File any issues with the appropriate OpenTelemetry community.
{: .callout}

## Prerequisites

You must have a Cloud Observability [access token](/docs/create-and-manage-access-tokens) for the project to report metrics to.

## Running the Example

You can run this example with `docker-compose up` in this directory.

## Configuration

Installation of the OpenTelemetry Collector varies, please refer to the [collector documentation](https://opentelemetry.io/docs/collector/) for more information.

The example collector's configuration, used for this project shows using processors to add metrics with Cloud Observability:

``` yaml
receivers:
      prometheus:
        config:
          scrape_configs:
            - job_name: 'otel-collector'
              scrape_interval: 5s
              static_configs:
                - targets: ['0.0.0.0:8888']
            - job_name: "istio"
              scrape_interval: 5s
              metrics_path: "/stats/prometheus"
              kubernetes_sd_configs:
                - role: "pod"
              relabel_configs:
                - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_scrape]
                  action: keep
                  regex: true
                - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_path]
                  action: replace
                  target_label: __metrics_path__
                  regex: (.+)
                - source_labels: [__meta_kubernetes_pod_ip, __meta_kubernetes_pod_container_port_number]
                  action: replace
                  target_label: __address__
                  regex: ([^:]+);(\d+)
                  replacement: ${1}:${2}
                - action: labelmap
                  regex: __meta_kubernetes_pod_label_(.+)
                - source_labels: [__meta_kubernetes_namespace]
                  action: replace
                  target_label: kubernetes_namespace
                - source_labels: [__meta_kubernetes_pod_name]
                  action: replace
                  target_label: kubernetes_pod_name

    processors:
      batch:

    exporters:
      logging:
        loglevel: debug
      otlp:
        endpoint: ingest.lightstep.com:443
        headers:
                "lightstep-access-token": "{LS_ACCESS_TOKEN}"

    service:
      telemetry:
        logs:
          level: "debug"
      pipelines:
        metrics:
          receivers: [prometheus]
          processors: [batch]
          exporters: [logging,otlp]
```

## otel-collector-istio setup

Before we start, make sure you have the following installed:

Go: https://golang.org/doc/install
Docker: https://docs.docker.com/get-docker/
Kubernetes CLI (kubectl): https://kubernetes.io/docs/tasks/tools/install-kubectl/
kind (Kubernetes in Docker): https://kind.sigs.k8s.io/docs/user/quick-start/
Istio: https://istio.io/latest/docs/setup/getting-started/


#### Build the Docker image:

```sh
docker build -t go-istio-demo .
```

#### Create the cluster

```sh
kind create cluster --name go-istio-cluster --config kind-config.yaml
```

#### Install Istio

* Label the default namespace to enable automatic sidecar injection for Istio

```sh
istioctl install --set profile=demo -y
kubectl label namespace default istio-injection=enabled
```

#### Load the Docker image into the kind cluster

```sh
kind load docker-image go-istio-demo:latest --name go-istio-cluster
```

#### Deploy the application

```sh
kubectl apply -f go-istio-demo.yaml
```

#### Apply the deployment to your Kubernetes cluster

```sh
kubectl apply -f otel-collector-deployment.yaml
```

* Verify that the OpenTelemetry Collector is running:

```sh
kubectl get pods -l app=otel-collector
```

#### Apply the ConfigMap to your Kubernetes cluster

```sh
kubectl apply -f otel-collector-configmap.yaml
```

#### Verify that the ConfigMap has been created

```sh
kubectl get configmap otel-collector-conf
```


#### Apply the Secret to your Kubernetes cluste

```sh
kubectl apply -f lightstep-secret.yaml
```

#### Test service

```sh
kubectl -n go-istio-demo port-forward svc/go-istio-demo 8080:80
```

```sh
curl http://localhost:8080
Hello, Golang with Istio!
```

#### Checks

* Check the status of the pods in the go-istio-demo namespace

```sh
kubectl -n go-istio-demo get pods
```

```sh
kubectl port-forward svc/golang-demo 8080:80
```

* Check the pod's metrics logs and events
  
    * Find the pod name of the OpenTelemetry Collector deployment
  
        ```sh
        kubectl get pods -l app=otel-collector
        kubectl logs -f <otel-collector-pod-name>
        ```

