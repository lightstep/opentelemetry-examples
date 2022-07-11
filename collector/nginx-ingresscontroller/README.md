---

## Running this Example

The NGINX Ingress Controller Operator is a Helm based operator created with the [Operator Framework](https://sdk.operatorframework.io/).


### Prerequisites

To run the example you'll need to put your Lightstep Access Token in a file at `collector/.patch.token.yaml`. That file should look exactly like `collector/secret.yaml` execept that it will include your actual Lightstep access token where indicated. You can run `make copy-otel-secret-patch` which is just a rule to execute `cp collector/secret.yaml collector/.patch.token.yaml`. There's already a `kustomization.yaml` file that references this configuration. 


## Steps

### 1. Create a cluster

First you'll need to create a cluster by a method of your choice. `kind create cluster` works well on Linux and is satisfactory for our purposes on MacOS.

### 2. Install the OTEL collector operator 

#### a. Installation prerequisite 

The OTEL collector operator depends on cert-manager, so we install that first.

```sh
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.8.2/cert-manager.yaml
```

Or we can add the OTEL helm charts and install the operator from there with something like this...

```sh
helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
helm repo update
helm install your-release-name -n your-otel-operator-namespace open-telemetry/opentelemetry-operator
```

Since the Collector operator depends on the condition of cert-manager, we'll wait on the prerequisites before we proceed.

```
kubectl wait deployment -n cert-manager cert-manager --for condition=Available=True --timeout=90s 
kubectl wait deployment -n cert-manager cert-manager-caininjector --for condition=Available=True --timeout=90s 
kubectl wait deployment -n cert-manager cert-manager-webhook --for condition=Available=True --timeout=90s 
```

#### b. Collector installation

Installing the Collector is straightforward. As we usually do when installing Helm charts, we'll start by adding the chart collection to our helm repo and updating.

Now we can either install the OpenTelemetry Operator by applying the manifest like...

```sh
kubectl apply -f https://github.com/open-telemetry/opentelemetry-operator/releases/latest/download/opentelemetry-operator.yaml
```

Or we can use the helm charts like ... 

```sh
helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
helm repo update
helm install your-release-name -n your-collector-operator-namespace --create-namespace
```

3. Install the [NGINX Ingress Controller Operator](https://github.com/nginxinc/nginx-ingress-helm-operator#readme).

Installing the NGINX Ingress Controller Operator requires some manual steps at this time. First we need a copy of the repo. Then we can use the Makefile in that repository root to complete the installation. We do that in this sequence of commands.

```sh
	git clone https://github.com/nginxinc/nginx-ingress-helm-operator/
	cd nginx-ingress-helm-operator/
	git checkout v1.0.0
	make deploy IMG=nginx/nginx-ingress-operator:1.0.0
```

This action is in this repo's Makefile by the rule `install-nginx-ingress-operator`.

4. Deploy an NGINX Ingress Controller instance

```sh
kubectl apply -f ingress/
```

5. Deploy the Collector instance

```sh
kubectl apply -k collector/
```

This command illustrates uses the kustomize flag (`-k`) to add the secret we need in the Collector environment to the manifest. To make it work you'll need to make a file at `collector/.patch.token.yaml`, which is just a copy of `collector/secret.yaml` with the place indicated replaced by your actual Lightstep access token. This arrangement is mostly to simplify keeping secrets out of version control during development. But it would also work if you delete the file at `collector/kustomization.yaml` and use the `-f` flag in place of `-k` in the command above assuming you want to store the access token in the file.

6. Make sure everything installed in a good state 

At this point we expect to see the metrics sent to our account in Lightstep.

TODO: put commands to view resources

7. Delete the Resources 

If you used Kind to run this example then the simplest way to delete the resources is to delete the cluster with `kind delete cluster` or `kind delete clusters name-of-my-cluster`.

If you will be keeping your cluster then it's simplest to begin by deleting the namespaces.

```sh
kubectl delete namespace my-example
kubectl delete namespace cert-manager
kubectl delete namespace nginx-ingress-operator-system 
```

Then you can proceed to delete any individual resources that may be in the default namespace. Look over it with `kubectl get all` and delete accordingly.

