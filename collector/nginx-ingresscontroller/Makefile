# https://docs.nginx.com/nginx-ingress-controller/installation/installation-with-helm/

RELEASE=otel-example

add-repository :
	helm repo add nginx-stable https://helm.nginx.com/stable
	helm repo update

install : add-repository 	## install ingress controller and expose prometheus metrics
	helm install $(RELEASE) nginx-stable/nginx-ingress --set prometheus.create=true

uninstall :
	helm uninstall $(RELEASE)

upgrade :
	helm upgrade $(RELEASE) nginx-stable/nginx-ingress

status-ic :  	## status of ingress-controller
	helm status $(RELEASE)

collector-up :
	kubectl apply -f collector-spec.yml


