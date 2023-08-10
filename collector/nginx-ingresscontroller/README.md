---
# Monitoring NGINX Ingress Controller

This example demonstrates monitoring the NGINX Ingress Controller via the Prometheus metrics endpoint with the OTEL Collector. The example configuration deploys the NGINX Ingress Controller and the OTEL Collector via their Kubernetes operators, each of which we deploy using helm charts. 

### Prerequisites

To run the example you'll need to put your Cloud Observability Access Token in a file at `collector/.patch.token.yaml`. That file should look exactly like `collector/secret.yaml` execept that it will include your actual Cloud Observability access token where indicated. You can run `make copy-otel-secret-patch` which is just a rule to execute `cp collector/secret.yaml collector/.patch.token.yaml`. There's already a `kustomization.yaml` file that references this configuration. 


## Steps

### 1. Create a cluster

First you'll need to create a cluster by a method of your choice. `kind create cluster` works well on Linux and is satisfactory for our purposes on MacOS.

### 2. Install the OTEL collector operator 

#### a. Installation prerequisite 

The OTEL collector operator depends on cert-manager, so we install that first.

```sh
helm repo add jetstack https://charts.jetstack.io
helm repo update
helm install \
  my-cert-manager-release jetstack/cert-manager \
  --namespace cert-manager \
  --create-namespace \
  --version v1.8.2 \
  --set installCRDs=true
```

Since the Collector operator depends on the condition of cert-manager, we'll wait on the prerequisites before we proceed.

```
kubectl wait deployment -n cert-manager cert-manager --for condition=Available=True --timeout=90s 
```

#### b. Collector installation

Installing the Collector is straightforward. As we usually do when installing Helm charts, we'll start by adding the chart collection to our helm repo and updating.

Or we can use the helm charts like ... 

```sh
helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
helm repo update
helm install your-release-name -n your-collector-operator-namespace --create-namespace
```

3. Install the [NGINX Ingress Controller Operator](https://github.com/nginxinc/nginx-ingress-helm-operator#readme).

The NGINX Ingress Controller Operator is a Helm based operator created with the [Operator Framework](https://sdk.operatorframework.io/). 

Installing the NGINX Ingress Controller Operator requires some manual steps at this time. First we need a copy of the repo. Then we can use the Makefile in that repository root to complete the installation. We do that in this sequence of commands.

```sh
	git clone https://github.com/nginxinc/nginx-ingress-helm-operator/
	cd nginx-ingress-helm-operator/
	git checkout v1.0.0
	make deploy IMG=nginx/nginx-ingress-operator:1.0.0
```

This action is in this repo's Makefile by the rule `install-nginx-ingress-operator`.

5. Deploy an NGINX Ingress Controller instance

The file at `ingress/values.yaml` tells the NGINX Ingress Controller Operator how to operate our instance. And `ingress/default-server-secret.yaml` provides a TLS cert for the purpose of the demo. 

```sh
kubectl apply -f ingress/
```

For purposes of monitoring via the Prometheus endpoint, we need to tend to two places in our configuration. In the first place we need to add customPorts to our service. This needs to map the port where Prometheus metrics are exposed to a port where it will be exposed in the Kubernetes Service.

```yaml
    service:
      create: true
      type: LoadBalancer
      customPorts:
        - name: prometheus
          port: 9113
          targetPort: 9113
```

Now we need to enable Prometheus metrics. We do that in the Operator by configuring a prometheus section in the spec. The port we choose here needs to match what we map from in customPorts. And we need to be sure that create is true. We could configure security by adding a TLS cert as a secret for Prometheus. See the [NGINX docs](https://docs.nginx.com/nginx-ingress-controller/installation/installation-with-helm) for details of configuration with the helm chart.

```yaml  
  prometheus:
    create: true
    port: 9113
    scheme: http
```

5. Deploy the Collector instance

The first thing you need to do is configure the OTEL Collector's Prometheus Receiver to look for metrics at the endpoint where we've exposed by the NGINX Ingress Controller. You can approach this a variety of ways, but we're going to illustrate by embedding the configuration into the pod definition. In the spec you provide the Operator for the OpenTelemetryCollector kind you need to add a key for config, end the line with the | which indicates that a multi-line value follows.

```yaml
spec:
  ...
  config: |
    receivers:
      prometheus:
        config:
          scrape_configs:
            - job_name: otel-nginx-eg
              static_configs:
                - targets: ["my-nginx-nginx-ingress:9113"]
  ...
    exporters:
      otlp/public:
        endpoint: ingest.lightstep.com:443
        headers:
          "lightstep-access-token": "${LS_ACCESS_TOKEN}"
  ...
```

This example provides a method for injecting the access token secret by `kustomize` merging a hidden at `collector/.patch.token.yaml`, which is the same as `collector/secret.yaml`, but it contains the real access token. You'll need to establish something consistent with your practices for managing secret.

```sh
kubectl apply -k collector/
```

This command uses the kustomize flag (`-k`) to override the access token with the real value. To make it work you'll need to make a file at `collector/.patch.token.yaml`, which is just a copy of `collector/secret.yaml` with the place indicated replaced by your actual Cloud Observability access token. This arrangement is mostly to simplify keeping secrets out of version control during development. But the example would also work if you delete the file at `collector/kustomization.yaml` and use the `-f` flag in place of `-k` in the command above, assuming you you have another mechanism to get the access token variable into the environment.

6. Make sure everything installed in a good state 

At this point we expect to see the metrics sent to our account in Cloud Observability.

We should also be able to see that our deployments are in good health with a command like `kubectl get all -n my-example`.

7. Cleanup example work

If you used Kind to run this example then the simplest way to delete the resources is to delete the cluster with `kind delete cluster` or `kind delete clusters name-of-my-cluster`.

If you can't delete your cluster then it's simplest to begin by deleting the namespaces.

```sh
kubectl delete namespace my-example
kubectl delete namespace cert-manager
kubectl delete namespace nginx-ingress-operator-system 
kubectl delete namespace my-otel-collector-operator-system-namespace
```

Then you can proceed to delete any individual resources that may be in the default namespace. Look over it with `kubectl get all` and delete accordingly.

## Additional Resources

* NGINX demonstrates collecting Ingress Controller metrics with the OTEL Collector and sending to Cloud Observability in their reference architecture. See more at the [NGINX blog](https://www.nginx.com/blog/integrating-opentelemetry-modern-apps-reference-architecture-progress-report/#metrics-collection).
* For additional detail on collecting metrics with the OTEL Collector documentation please see ["Replace Prometheus with an OpenTelemetry Collector"](https://docs.lightstep.com/docs/replace-prometheus-with-an-otel-collector-on-kubernetes).
