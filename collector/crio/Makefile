PLATFORMS=linux/amd64
OUTPUT=--load
PROGRESS=auto
IMAGE=kindnode/crio:1.20.9

.PHONY: build
build:
	docker build -t kindest/node:crio-build .
	kind build node-image --base-image kindest/node:crio-build --image kindest/node:crio

.PHONY: run-cluster
run-cluster:
	kind create cluster --config kind-config.yaml

	kubectl apply -f collector-configmap.yaml
	kubectl apply -f collector-rbac.yaml
	kubectl create secret generic ls --from-literal=access_token=${LS_ACCESS_TOKEN} -n collector
	kubectl apply -f collector.yaml

.PHONY: delete-cluster
delete-cluster:
	kind delete cluster --name=crio-demo

.PHONY: get-logs
get-logs:
	kubectl logs -n collector collector
