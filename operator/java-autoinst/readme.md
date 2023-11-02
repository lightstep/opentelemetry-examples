# Java Autoinstrumentation with the Collector Operator

This installs a sample Java spring boot application and instruments it automatically using the collector operator.

### Requirements

- A locally running k8s test cluster (minikube, kind)
- A Cloud Observability project and access token

### Add Helm Repos

```sh
helm repo add springboot https://josephrodriguez.github.io/springboot-helm-charts
helm repo add jetstack https://charts.jetstack.io
helm repo add prometheus https://prometheus-community.github.io/helm-charts
helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
helm repo add lightstep https://lightstep.github.io/otel-collector-charts
```

Then update:

```sh
helm repo update
```

### Install Helm Charts

```sh
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

```sh
NAMESPACE="your-namespace"
LS_TOKEN="<your-access-token>"
kubectl create secret generic otel-collector-secret -n ${NAMESPACE} --from-literal=LS_TOKEN=${LS_TOKEN}

# Also sends k8s metrics to Cloud Observability
helm install kube-otel-stack lightstep/kube-otel-stack -n ${NAMESPACE} --set autoinstrumentation.enabled=true --set tracesCollector.enabled=true

### Run Java Demo App with Annotation
helm install springboot-starterkit-svc springboot/springboot-starterkit-svc -n ${NAMESPACE} -f values.yaml
```
