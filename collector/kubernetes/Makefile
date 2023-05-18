
.PHONY: run-cluster
run:
	kind create cluster --name kubernetes

	helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
	helm install kube-state-metrics prometheus-community/kube-state-metrics -n kube-system --version 5.6.2
	helm install prometheus-node-exporter prometheus-community/prometheus-node-exporter -n kube-system --version 4.17.2

	kubectl apply -f collector-configmap.yaml
	kubectl apply -f collector-rbac.yaml
	kubectl create secret generic ls --from-literal=access_token=${LS_ACCESS_TOKEN} -n collector
	kubectl apply -f collector.yaml

.PHONY: delete-cluster
stop:
	kind delete cluster --name=kubernetes

.PHONY: get-logs
get-logs:
	kubectl logs -n collector collector
