#!/bin/zsh

kubectl delete cm -n istio-system lightstep-otel-collector-base-opentelemetry-collector
kubectl delete cm -n istio-system otel-collector-conf
kubectl delete cm -n istio-system istio
