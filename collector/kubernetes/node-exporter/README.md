---
## node-exporter - resources

* [Monitoring Linux Host Metrics with the Node Exporter](https://prometheus.io/docs/guides/node-exporter/)

* [Using Prometheus Operator to Run Node Exporter](https://www.cloudforecast.io/blog/node-exporter-and-kubernetes/)

* [Setup Guide: kubectl create on manifests ](https://devopscube.com/node-exporter-kubernetes/)

    kubectl apply -k collector/ 
    kubectl apply -f node-exporter/service.yaml node-exporter/daemonset.yaml

    TODO - migrate away from style used for node-exporter right
            set fewer variables, but keep good enough practices
