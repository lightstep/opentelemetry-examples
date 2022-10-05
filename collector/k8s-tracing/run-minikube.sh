#!/bin/sh


mkdir -p ~/.minikube/files/etc/ssl/certs/
mkdir -p ~/.minikube/files/etc/crio
mkdir -p ~/.minikube/files/etc/containerd
cp apiserver-tracing.yaml ~/.minikube/files/etc/ssl/certs/apiserver-tracing.yaml 
cp kubelet-tracing.yaml ~/.minikube/files/etc/ssl/certs/kubelet-tracing.yaml 
cp crio.conf ~/.minikube/files/etc/crio/crio.conf 
cp containerd.toml ~/.minikube/files/etc/containerd/config.toml 

SPAN_INGEST_ADDR=192.168.1.253:4317

# requires minikube v1.26.1 or greater

minikube start --kubernetes-version=v1.25.0-rc.1 \
    --feature-gates=APIServerTracing=true \
    --extra-config=apiserver.feature-gates=APIServerTracing=true \
    --extra-config=apiserver.tracing-config-file=/etc/ssl/certs/apiserver-tracing.yaml \
    --extra-config=kubelet.config=/etc/ssl/certs/kubelet-tracing.yaml \
    --extra-config=etcd.experimental-enable-distributed-tracing=true \
    --extra-config=etcd.experimental-distributed-tracing-address=$SPAN_INGEST_ADDR