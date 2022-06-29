---
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
