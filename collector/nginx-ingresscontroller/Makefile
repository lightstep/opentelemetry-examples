NGINX_RELEASE := my-ingress-release
OTEL_RELEASE := my-collector-release
CERT_MANAGER_RELEASE := my-cert-manager-release

OTEL_OPERATOR_NAMESPACE := otel-collector-operator-system
NGINX_OPERATOR_NAMESPACE := nginx-ingress-operator-system
CERT_MANAGER_NAMESPACE := cert-manager

EXAMPLE_NAMESPACE := my-example

all : | setup install-operators deploy-instances

setup : add-repositories create-namespaces
install-operators : | install-cert-manager install-nginx-ingress-operator install-otel-operator
deploy-instances : deploy-collector deploy-nginx

##@ convenience

clean :
	@echo ** Deleting namespaces. **
	@kubectl delete namespace $(EXAMPLE_NAMESPACE)
	@kubectl delete namespace $(OTEL_OPERATOR_NAMESPACE)
	@kubectl delete namespace $(NGINX_OPERATOR_NAMESPACE)
	@kubectl delete namespace $(CERT_MANAGER_NAMESPACE)
	@echo ** Namespaces have been deleted. **

##@ deployment:

_wait-for-nginx-operator :
	@echo ** Waiting for NGINX Operator. **
	@kubectl wait deployment -n $(NGINX_OPERATOR_NAMESPACE) nginx-ingress-operator-controller-manager --for condition=Available=True --timeout=120s
	@echo ** NGINX Operator running. **

_wait-for-collector-operator :
	@echo ** Waiting for OTEL Operator. **
	@kubectl wait deployment -n $(OTEL_OPERATOR_NAMESPACE) opentelemetry-operator-controller-manager --for condition=Available=True --timeout=120s
	@echo ** Otel Operator running is. **

deploy-nginx :
	$(MAKE) _wait-for-nginx-operator 
	@kubectl apply -n $(EXAMPLE_NAMESPACE) -f ingress/

deploy-collector :
	$(MAKE) _wait-for-collector-operator
	@kubectl apply -n $(EXAMPLE_NAMESPACE) -k collector/

create-namespaces :
	@echo ** Creating namespace: $(EXAMPLE_NAMESPACE) **
	@kubectl create namespace $(EXAMPLE_NAMESPACE) --dry-run=client -o yaml | kubectl apply -f -
	@echo ** Namespaces created **

install-cert-manager :  	## install cert manager. otel-operator requires this - in namespace=cert-manager.
	@echo ** Installing cert-manager in cluster. **
	@helm install $(CERT_MANAGER_RELEASE) jetstack/cert-manager \
		--namespace $(CERT_MANAGER_NAMESPACE) \
		--create-namespace \
		--version v1.8.0 \
		--set installCRDs=true
	@echo ** cert-manager install done. **

_wait-for-cert-manager :
	@echo ** Waiting for cert-manager deployments to become available. **
	@kubectl wait deployment -n $(CERT_MANAGER_NAMESPACE) $(CERT_MANAGER_RELEASE) --for condition=Available=True --timeout=90s 
	@echo ** cert-manager deployments are now avaliable. **

undeploy-instances : 	## delete instances
	kubectl delete -n $(EXAMPLE_NAMESPACE) -k collector/
	kubectl delete -n $(EXAMPLE_NAMESPACE) -f ingress/

##@ monitoring

status : 	## check status of controllers
	helm status $(OTEL_RELEASE) -n $(EXAMPLE_NAMESPACE)
	helm status $(NGINX_RELEASE) -n $(EXAMPLE_NAMESPACE)

get-values :
	helm get values $(OTEL_RELEASE) -n $(EXAMPLE_NAMESPACE)
	helm get values $(NGINX_RELEASE) -n $(EXAMPLE_NAMESPACE)

##@ cluster prep

_create-cluster : ## create kind cluster. Kind also configures kubectl to use this cluster. 
	kind create cluster --name $(EXAMPLE_NAMESPACE)-cluster

_delete-cluster : ## delete the kind cluster with this name
	kind delete cluster --name $(EXAMPLE_NAMESPACE)-cluster

add-repositories : 	## add repositories for helm charts
	echo 'Adding repositories'
	@helm repo add jetstack https://charts.jetstack.io
	@helm repo add nginx-stable https://helm.nginx.com/stable
	@helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
	echo 'Repositories added'
	@helm  repo update
	echo 'Repositories updated'

prereqs : _create-cluster create-namespaces

##@ operators

install-otel-operator :
	$(MAKE) _wait-for-cert-manager # otel-operator requires cert-manager
	helm install -n $(OTEL_OPERATOR_NAMESPACE) --create-namespace $(OTEL_RELEASE) open-telemetry/opentelemetry-operator 

install-nginx-ingress-operator :
	if [ ! -d nginx-ingress-helm-operator ]; then git clone https://github.com/nginxinc/nginx-ingress-helm-operator; fi
	cd nginx-ingress-helm-operator; git checkout v1.0.0; $(MAKE) deploy IMG=nginx/nginx-ingress-operator:1.0.0;

uninstall-operators :
	helm uninstall $(OTEL_RELEASE)
	helm uninstall $(NGINX_RELEASE)
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

