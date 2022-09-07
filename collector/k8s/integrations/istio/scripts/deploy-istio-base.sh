#!/bin/zsh
# deploy-istio-base.sh deploys istio CRUDs (base) and istiod

helm install istio-base istio/base -n istio-system
helm install istiod istio/istiod -n istio-system --wait
