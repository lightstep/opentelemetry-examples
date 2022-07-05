---

# Install NGINX Ingress Helm Operator

Generally instructions are from https://github.com/nginxinc/nginx-ingress-helm-operator#readme
1. Deploy the Operator and associated resources
	a. clone the operator
```sh
	git clone https://github.com/nginxinc/nginx-ingress-helm-operator/
	cd nginx-ingress-helm-operator/
	git checkout v1.0.0
```
	b. deploy it the operator to the k8s environment
```sh
	make deploy IMG=nginx/nginx-ingress-operator:1.0.0
```

2. Deploy the NGINX Ingress Controller using the Operator
	- There's a sample config at https://github.com/nginxinc/nginx-ingress-helm-operator/blob/main/config/samples/charts_v1alpha1_nginxingress.yaml
	- Command is like `kubectl create -n my-nginx-ingress -f nginx-ingress-controller.yaml`

3. Check that resources were deployed
	- `kubectl -n my-nginx-ingress get all`

4. Delete the Ingress Controller
	- `kubectl delete -f nginx-ingress-controller.yaml`

5. Delete the namespace
	- `kubectl delete namespace my-nginx-ingress`

# Ingest NGINX Ingress Controller metrics with the OTEL Operator

Here's a potentially (probably not) relevant tutorial for running opentracing: https://github.com/opentracing-contrib/nginx-opentracing/blob/master/doc/Tutorial.md.

* Docs for Ingress Controllers (k8s objects): https://kubernetes.github.io/ingress-nginx/
* Tutorial example by NGINX on running NGINX Ingress Controller: https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/complete-example
* Docs for OTEL Operator: https://github.com/open-telemetry/opentelemetry-operator) 
- start by installing cert manager
- then install the other part 

In your helm chart you'll need to set `prometheus.create` to true.

You can set variables when you initialize helm like this...

helm install my-release nginx-stable/nginx-ingress --set prometheus.create=true

Here are the variables you can configure with helm...
https://docs.nginx.com/nginx-ingress-controller/installation/installation-with-helm

prometheus.create           Expose NGINX or NGINX Plus metrics in the Prometheus format.                    false
prometheus.portConfigures   the port to scrape the metrics.                                                 9113
prometheus.schemeConfigures the HTTP scheme that requests must use to connect to the Prometheus endpoint.   http
prometheus.secret           Specifies the namespace/name of a Kubernetes TLS secret which can be used to 
                            establish a secure HTTPS connection with the Prometheus endpoint.               owq
