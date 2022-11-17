# Java Autoinstrumentation with the Collector Operator

This installs a sample Java spring boot application and instruments it automatically using the collector operator.

### Requirements

* A locally running k8s test cluster (minikube, kind)
* A Lightstep project and access token

### Add Helm Repos

```
helm repo add springboot https://josephrodriguez.github.io/springboot-helm-charts
helm repo add jetstack https://charts.jetstack.io
helm repo add prometheus https://prometheus-community.github.io/helm-charts
helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
helm repo add lightstep https://lightstep.github.io/prometheus-k8s-opentelemetry-collector
```

Then update:

```
    helm repo update
```

### Install Helm Charts

```
helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --create-namespace \
  --version v1.8.0 \
  --set installCRDs=true

helm install \
  opentelemetry-operator open-telemetry/opentelemetry-operator \
  -n opentelemetry-operator \
  --create-namespace
```

### Apply Auto-Instrumentation Configuration and Run Collector

```
export LS_TOKEN=<your-access-token>
kubectl create secret generic otel-collector-secret -n default --from-literal=LS_TOKEN=$LS_TOKEN

# Also sends k8s metrics to Lightstep
helm install kube-otel-stack --namespace default --set autoinstrumentation.enabled=true --set tracesCollector.enabled=true

### Run Java Demo App with Annotation
helm install springboot-starterkit-svc springboot/springboot-starterkit-svc -f values.yaml

```

