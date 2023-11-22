
all: setup install-all

setup: create-cluster add-repositories
install-all: install-contour install-kong install-collector

add-repositories:
	helm repo add bitnami https://charts.bitnami.com/bitnami
	helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts

create-cluster:
	kind create cluster --config kind-config.yaml

delete-cluster:
	kind delete cluster

install-contour:
	helm upgrade --install my-contour bitnami/contour --namespace projectcontour --create-namespace

install-kong:
	helm upgrade --install my-kong --set metrics.enabled=true bitnami/kong

install-collector:
	./install_collector.sh
