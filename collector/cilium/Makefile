CILIUM_RELEASE := cilium-release
OTEL_RELEASE := collector-release
CERT_MANAGER_RELEASE := cert-manager-release

OTEL_OPERATOR_NAMESPACE := otel-collector-operator-system
CILIUM_OPERATOR_NAMESPACE := cilium-operator-system
CERT_MANAGER_NAMESPACE := cert-manager


DEMO_NAMESPACE := cilium-demo
CLUSTER_NAME := $(DEMO_NAMESPACE)-cluster

CILIUM_URL="https://github.com/cilium/cilium-cli/releases/latest/download/cilium-linux-amd64.tar.gz"

HUBBLE_VERSION := $(shell curl -s https://raw.githubusercontent.com/cilium/hubble/master/stable.txt)
HUBBLE_URL="https://github.com/cilium/hubble/releases/download/$(HUBBLE_VERSION)/hubble-linux-amd64.tar.gz"

# the API Address values are accessible with `kubectl context-info`
# API_SERVER_IP := 127.0.0.1
# API_SERVER_PORT := 45811 

all : | setup install-operators deploy-instances

setup : | create-cluster add-repositories create-namespaces load-cilium-image
install-operators : | install-cilium-operator wait-for-cilium install-cert-manager install-otel-operator
deploy-instances : deploy-collector 


# TODO either make this work or delete it
# it should be making a comparison and bailing if the command results aren't the same
check-kind-nodes-have-cgroup :
	KIND_CGROUP:=$(shell docker inspect -f '{{.State.Pid}}' cilium-demo-cluster-control-plane)
	if [ "`sudo ls -al /proc/$(KIND_CGROUP)/ns/cgroup`" != "`sudo ls -al /proc/self/ns/cgroup`" ]; then \
		echo "ERROR: kind nodes do not have their own cgroup"; \
		echo "See: note at https://docs.cilium.io/en/stable/gettingstarted/kind/"; \
		exit 1; \
	fi

##@ tools

install-cilium-cli-linux:
	curl -L --remote-name-all $(CILIUM_VERSION){,.sha256sum}
	sha256sum --check cilium-linux-amd64.tar.gz.sha256sum
	sudo tar xzvfC cilium-linux-amd64.tar.gz /usr/local/bin
	rm cilium-linux-amd64.tar.gz*

install-hubble-cli-linux:
	curl -L --remote-name-all $(HUBBLE_URL){,.sha256sum}
	sha256sum --check hubble-linux-amd64.tar.gz.sha256sum
	sudo tar xzvfC hubble-linux-amd64.tar.gz /usr/local/bin
	rm hubble-linux-amd64.tar.gz*

##@ setup

# TODO ask what can actually accomplish this - taint all nodes
taint-nodes:
	kubectl taint nodes --all taint=node.cilium.io/agent-not-ready NoExecute

load-cilium-image:
	docker pull quay.io/cilium/cilium:v1.12.0
	kind load docker-image --name $(CLUSTER_NAME) quay.io/cilium/cilium:v1.12.0

add-repositories : 	## add repositories for helm charts
	# echo 'Adding repositories'
	helm repo add jetstack https://charts.jetstack.io
	helm repo add cilium https://helm.cilium.io/
	helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
	# echo 'Repositories added'
	helm  repo update
	# echo 'Repositories updated'


##@ install operators

install-cilium-operator: upgrade-cilium-operator

## NOTE: if you say kubeProxyReplacement=strict then you evict pods if they can't replace kubeproxy
## we use partial. See: https://docs.cilium.io/en/stable/gettingstarted/kubeproxy-free/#kubeproxy-free

# agent port: 9962
# operator port: 9964
# hubble port: 
upgrade-cilium-operator: 
	helm upgrade cilium cilium/cilium --version 1.12.0 \
		--install \
		--wait \
		--namespace kube-system \
		--set kubeProxyReplacement=partial \
		--set socketLB.enabled=false \
		--set externalIPs.enabled=true \
		--set nodePort.enabled=true \
		--set hostPort.enabled=true \
		--set bpf.masquerade=false \
		--set image.pullPolicy=IfNotPresent \
		--set ipam.mode=kubernetes \
		--set hubble.relay.enabled=true \
		--set hubble.ui.enabled=true \
		--set prometheus.enabled=true \
		--set operator.prometheus.enabled=true \
		--set hubble.prometheus.enabled=true \
		--set hubble.metrics.enabled="{dns,drop,tcp,flow,icmp,http}"

wait-for-cilium:
	kubectl wait deployment -n kube-system cilium-operator --for condition=Available=True

# TODO fix whatever makes this one broken in the awk 
restart-unmanaged-pods:
	kubectl get pods --all-namespaces -o custom-columns=NAMESPACE:.metadata.namespace,NAME:.metadata.name,HOSTNETWORK:.spec.hostNetwork --no-headers=true | grep '<none>' | awk '{print -n $1 $2}' | xargs -L 1 -r kubectl delete pod

clean :
	@echo ** Deleting namespaces. **
	@kubectl delete namespace $(DEMO_NAMESPACE)
	@kubectl delete namespace $(OTEL_OPERATOR_NAMESPACE)
	@kubectl delete namespace $(CILIUM_OPERATOR_NAMESPACE)
	@kubectl delete namespace $(CERT_MANAGER_NAMESPACE)
	@echo ** Namespaces have been deleted. **

##@ deployment:

_wait-for-collector-operator :
	@echo ** Waiting for OTEL Operator. **
	@kubectl wait deployment -n $(OTEL_OPERATOR_NAMESPACE) opentelemetry-operator-controller-manager --for condition=Available=True --timeout=120s
	@echo ** Otel Operator running is. **

deploy-collector :
	$(MAKE) _wait-for-collector-operator
	@kubectl apply -n $(DEMO_NAMESPACE) -k collector/

create-namespaces :
	@echo ** Creating namespace: $(DEMO_NAMESPACE) **
	@kubectl create namespace $(DEMO_NAMESPACE) --dry-run=client -o yaml | kubectl apply -f -
	@echo ** Namespaces created **

install-cert-manager :  	## install cert manager. otel-operator requires this - in namespace=cert-manager.
	helm install $(CERT_MANAGER_RELEASE) jetstack/cert-manager \
		--wait \
		--namespace $(CERT_MANAGER_NAMESPACE) \
		--create-namespace \
		--version v1.8.0 \
		--set installCRDs=true

_wait-for-cert-manager :
	@echo ** Waiting for cert-manager deployments to become available. **
	@kubectl wait deployment -n $(CERT_MANAGER_NAMESPACE) $(CERT_MANAGER_RELEASE) --for condition=Available=True --timeout=90s 
	@echo ** cert-manager deployments are now avaliable. **

##@ monitoring

status : 	## check status of controllers
	helm status $(OTEL_RELEASE) -n $(DEMO_NAMESPACE)
	helm status $(CILIUM_RELEASE) -n kube-system

values :
	helm get values $(OTEL_RELEASE) -n $(DEMO_NAMESPACE)
	helm get values $(CILIUM_RELEASE) -n kube-system

logs :
	kubectl logs -n $(DEMO_NAMESPACE) otel

##@ cluster management

create-cluster:
	# this is just a demo cluster. What's important in the config is that you not
	# try to demo this with a single-node test cluster.
	kind create cluster --name $(CLUSTER_NAME) --config kind-config.yaml 

delete-cluster : ## delete the kind cluster with this name
	kind delete cluster --name $(CLUSTER_NAME)

##@ operators

install-otel-operator :
	$(MAKE) _wait-for-cert-manager # don't have to use it, because we run install with --wait.. but it's good extra
	helm install $(OTEL_RELEASE) open-telemetry/opentelemetry-operator \
		--wait \
		-n $(OTEL_OPERATOR_NAMESPACE) \
		--create-namespace

uninstall-operators :
	helm uninstall $(OTEL_RELEASE)
	helm uninstall $(CILIUM_RELEASE)
	helm uninstall $(CERT_MANAGER_RELEASE)

## ========================================================================================================================

##@ general

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

