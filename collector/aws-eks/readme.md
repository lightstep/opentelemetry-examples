# AWS EKS

### Create new EKS environment

We recommend using `eksctl` to easily spin up a testing AWS EKS cluster for testing purposes. Instructions for installing it can be found [here](https://eksctl.io/introduction/#installation).

```sh
eksctl create cluster --name=otel-demo-cluster --nodes=3 --region=us-east-1
```

After `eksctl` completes, verify you are connected to your new cluster with `kubectl cluster-info`.

### Deploy the OpenTelemetry Demo to the new cluster

Edit `lightstep-values.yaml` to use your access token.

```sh
helm repo add jetstack https://charts.jetstack.io
helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
helm repo add lightstep https://lightstep.github.io/otel-collector-charts

helm repo update

helm install \
     cert-manager jetstack/cert-manager \
     --namespace cert-manager \
     --create-namespace \
     --version v1.8.0 \
     --set installCRDs=true

helm install \
     opentelemetry-operator open-telemetry/opentelemetry-operator \
     -n default

helm upgrade my-otel-demo open-telemetry/opentelemetry-demo --install -f lightstep-values.yaml
```

### Collect Kubernetes cluster Metrics

Set LS_TOKEN as secret for collector.

```sh
export LS_TOKEN='<your-token>'
kubectl create secret generic otel-collector-secret -n default --from-literal="LS_TOKEN=$LS_TOKEN"
```

```sh
helm upgrade collector-k8s-noprom lightstep/collector-k8s-noprom --install
```

### Forward Cloudwatch Metrics to your account

Follow instructions in the Cloud Observability UI to forward Amazon Cloudwatch metrics to your account: https://docs.lightstep.com/docs/setup-aws-for-metrics
