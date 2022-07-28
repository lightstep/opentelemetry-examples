---
# Monitoring Cilium & Hubble

This example demonstrates monitoring Cilium and Hubbles with the OTEL Collector's Prometheus Receiver. The example configuration deploys Cilium's Kubernetes operator with Hubble enabled.

### Prerequisites

To run the example you'll need to put your Lightstep Access Token in a file at `collector/.patch.token.yaml`. That file should look exactly like `collector/secret.yaml` execept that it will include your actual Lightstep access token where indicated. There's a `kustomization.yaml` file that references the secret properly and is applied with a rule to deploy the OTEL Collector in the Makefile.

#### kind

You can use any approach to managing your cluster, but the Makefile builds a cluster in `kind` and theirs a configuration file to deploy a 3 node cluster. It's important you have 3 nodes to minimize additional settings required.

#### helm

We use helm charts to install everything in this example.

#### Linux with eBPF

Cilium uses eBPF which is in the Linux kernel. It also may require special kernel parameters depending on your distribution. This was developed on Debian 11.

## Steps

### 1. Create a cluster

First you'll need to create a cluster by a method of your choice. `kind create cluster` works well on Linux.

### 2. Add helm repos

The OTEL Collector operator will require that we install cert-manager. So we add repos for cert-manager, otel-collector, and cilium.

### 3. Install Cilium

It's important that we install the Cilium operator in the cluster before our other requirements. If we didn't then we'd need restart the other pods.

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

#### b. Collector installation

Installing the Collector is straightforward. Assuming you have already added the chart repo, you can install it like this...

Or we can use the helm charts like ... 

```sh
helm install your-release-name -n your-collector-operator-namespace --create-namespace
```

3. Install the Cilium Operator

Cilium can run as an operator. Note that Hubble is not turned on by default and we have to enable metrics for each of 3 endpoints explicitly.

```
TODO show the helm install variation
```

5. Deploy the Collector instance

The first thing you need to do is configure the OTEL Collector's Prometheus Receiver to look for metrics at the 3 endpoints we exposed. Note that the ports are also available for configuration, but we're using default ports for each.

TODO: update the static config part of this and job_name to match what's in collector/values.yaml

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

This command uses the kustomize flag (`-k`) to override the access token with the real value. To make it work you'll need to make a file at `collector/.patch.token.yaml`, which is just a copy of `collector/secret.yaml` with the place indicated replaced by your actual Lightstep access token. This arrangement is mostly to simplify keeping secrets out of version control during development. But the example would also work if you delete the file at `collector/kustomization.yaml` and use the `-f` flag in place of `-k` in the command above, assuming you you have another mechanism to get the access token variable into the environment.

6. Make sure everything installed in a good state 

At this point we expect to see the metrics sent to our account in Lightstep.

We should also be able to see that our deployments are in good health with a command like `kubectl get all --all-namespaces`.

7. Cleanup example work

If you used kind to run this example then the simplest way to delete the resources is to delete the cluster with `kind delete cluster` or `kind delete clusters name-of-my-cluster`.

If you can't delete your cluster then it's simplest to begin by deleting the namespaces.

```sh
kubectl delete namespace my-example
kubectl delete namespace cert-manager
kubectl delete namespace my-otel-collector-operator-system-namespace
helm uninstal cilium-release-name
```

## Additional Resources

TODO links to the three main pages referenced in the final cut.
