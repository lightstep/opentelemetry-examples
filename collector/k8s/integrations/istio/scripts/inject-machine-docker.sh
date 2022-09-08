#!/bin/zsh

eval $(minikube docker-env)
cd cmd/machine && docker build -f Dockerfile -t machine .
