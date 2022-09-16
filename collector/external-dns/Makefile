
all: setup install-operators

setup: create-cluster add-repositories create-namespaces
install-operators: install-etcd install-coredns install-external-dns install-ingress install-collector

add-repositories:
	helm repo add coredns https://coredns.github.io/helm
	helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts

create-namespaces:
	kubectl create namespace etcd

create-cluster:
	minikube start
	minikube addons enable ingress

delete-cluster:
	minikube delete

install-etcd:
	kubectl create -f etcd-cluster.yaml -n etcd

install-coredns:
	helm install my-coredns coredns/coredns -f values-coredns.yaml

install-external-dns:
	kubectl apply -f external-dns.yaml

install-ingress:
	sleep 15
	kubectl apply -f ingress.yaml

install-collector:
	sed -e "s~LS_ACCESS_TOKEN~${LS_ACCESS_TOKEN}~" values-collector.tpl > values-collector.yaml
	helm install my-collector open-telemetry/opentelemetry-collector -f values-collector.yaml
