
RELEASE=otel-example

add-repository : 	## add repository where we can get the nginx ingress controller
	helm repo add nginx-stable https://helm.nginx.com/stable
	helm repo update

install-ic : add-repository ## install ingress controller and expose prometheus metrics
	helm install $(RELEASE) nginx-stable/nginx-ingress --set prometheus.create=true

install-cert-manager :  	## install the cert manager that otel-operator depends upon
	kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.8.2/cert-manager.yaml

install-otel-operator : install-cert-manager 	## install the otel-operator
	kubectl apply -f https://github.com/open-telemetry/opentelemetry-operator/releases/latest/download/opentelemetry-operator.yaml

create-lightstep-secret : 	## add your lightstep-access-token for the environment variable
	kubectl apply -f lightstep-secret.yml

up :  ## simply creates from the collector spec. You should run `create` first.
	kubectl apply -f collector-spec.yml

create : install-otel-operator install-ic collector-up create-lightstep-secret

uninstall-ic :
	helm uninstall $(RELEASE)

upgrade-ic : 	## upgrade ingress-controller
	helm upgrade $(RELEASE) nginx-stable/nginx-ingress

status-ic :  	## status of ingress-controller
	helm status $(RELEASE)

