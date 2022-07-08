---

## Running this Example

### Prerequisites

To run the example you'll need to put your Lightstep Access Token in a file at `collector/.patch.token.yaml`. That file should look exactly like `collector/secret.yaml` execept that it will include your actual Lightstep access token where indicated. You can run `make copy-otel-secret-patch` which is just a rule to execute `cp collector/secret.yaml collector/.patch.token.yaml`. There's already a `kustomization.yaml` file that references this configuration. 


## Steps

### 1. Create a cluster

First you'll need to create a cluster by a method of your choice. `kind create cluster` works well on Linux and is satisfactory for our purposes on MacOS.

### 2. Install the OTEL collector operator 

#### a. Installation prerequisite 

The OTEL collector operator depends on cert-manager, so we install that first.

TODO: link to install instructions for OTEL Collector Operator and write command of this step

Since the Collector operator depends on the condition of cert-manager, so lets wait on the prerequisites before we proceed.

```
kubectl wait deployment -n cert-manager cert-manager --for condition=Available=True --timeout=90s 
kubectl wait deployment -n cert-manager cert-manager-caininjector --for condition=Available=True --timeout=90s 
kubectl wait deployment -n cert-manager cert-manager-webhook --for condition=Available=True --timeout=90s 
```

#### b. Collector installation

Installing the Collector is straightforward. As we usually do when installing Helm charts, we'll start by adding the repo.

TODO: Link the Collector operator/helm install instructions

```sh
TODO: add commands for add/update repository
```

```sh
TODO: add command installing the OTEL collector
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

```
TODO: add command to deploy ingress controller
```

5. Deploy the Collector instance

```
TODO: add command to deploy ingress controller
```

This command illustrates using the kustomize option (`-k`) which is what the repo is presently configured for. To make it work you'll need to make a file at `collector/patch.token.yaml`, which is just a copy of `collector/secret.yaml` with your actual Lightstep access token. This arrangement is mostly to simplify keeping secrets out of version control during development. But it would also be fine to delete the file at `collector/kustomization.yaml` and use the `-f` flag in place of `-k` in the command above.

6. Make sure everything installed in a good state 

At this point we expect to see the metrics sent to our account in Lightstep.

7. Delete the Resources 

If you used Kind to run this example then the simplest way to delete the resources is to delete the cluster with `kind delete cluster` or `kind delete clusters name-of-my-cluster`.

If you will be keeping your cluster then it's simplest to begin by deleting the namespaces.

```sh
kubectl delete namespace my-example
kubectl delete namespace cert-manager
kubectl delete namespace nginx-ingress-operator-system 
```

Then you can proceed to delete any individual resources that may be in the default namespace. Look over it with `kubectl get all` and delete accordingly.


