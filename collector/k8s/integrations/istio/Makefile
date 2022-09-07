start: docker_build deploy_istio_base deploy_application deploy_telemetry run_istio_demo view_metrics

.PHONY:
docker_build:
	./scripts/inject-machine-docker.sh

.PHONY:
deploy_istio_base:
	./scripts/deploy-istio-base.sh

.PHONY:
deploy_application:
	helm install helm-application --generate-name

.PHONY:
deploy_telemetry:
	kubectl delete deployment/lightstep-otel-collector-base-opentelemetry-collector -n=istio-system
	kubectl delete service/lightstep-otel-collector-base-opentelemetry-collector -n=istio-system
	./scripts/deploy-istio-otel.sh
	helm install helm-otel-istio --generate-name

.PHONY:
run_istio_demo:
	./scripts/run-istio-demo.sh

.PHONY:
view_metrics:
	tmux attach-session -t istio-demo.watch-otel-collector

.PHONY:
exit_istio_demo:
	tmux detach-client istio-demo

.PHONY:
exit: exit_istio_demo
