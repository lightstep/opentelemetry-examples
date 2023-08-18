# User must set GOOS, GOARCH, GOPATH, LIGHTSTEP_API_KEY in their environment for full use of this Makefile
# GOOS, GOARCH, GOPATH may be automatically set with env depending on your platform - if not run go env
# and export the given vars
#
# otelcontrib_$(GOOS)_$(GOARCH) must exist in ${GOPATH}/bin

SHELL = /bin/sh

API_KEY:=$(shell echo ${LIGHTSTEP_API_KEY})
GOOS:=$(shell echo ${GOOS})
GOARCH:=$(shell echo ${GOARCH})
GOPATH:=$(shell echo ${GOPATH})

.PHONY: run-collector-from-bin
run-collector-from-bin:
ifndef RUN_CONFIG
	$(error RUN_CONFIG variable was not defined)
endif
	$(GOPATH)/bin/otelcontribcol_$(GOOS)_$(GOARCH) --config ${RUN_CONFIG} ${RUN_ARGS}

.PHONY: reset_run
reset_run:
ifndef PORT
	$(error PORT variable was not defined)
endif
ifndef RUN_CONFIG
	$(error RUN_CONFIG variable was not defined)
endif
	kill -9 $$(lsof -nP -iTCP:$(PORT) -sTCP:LISTEN | awk 'FNR == 2 {print $$2}')
	$(GOPATH)/bin/otelcontribcol_$(GOOS)_$(GOARCH) --config ${RUN_CONFIG} ${RUN_ARGS}

.PHONY: docker-component-run
docker-component-run:
ifndef COMPONENT
	$(error COMPONENT variable was not defined)
endif
	(cd collector/$(COMPONENT) && docker compose up)

.PHONY: docker-component-down
docker-component-down:
ifndef COMPONENT
	$(error COMPONENT variable was not defined)
endif
	(cd receiver/$(COMPONENT) && docker compose down)

.PHONY: new-component
new-component:
ifndef COMPONENT
	$(error COMPONENT variable was not defined)
else
	mkdir collector/$(COMPONENT)
	cp tools/templates/docker-compose.yaml collector/$(COMPONENT)
	cp tools/templates/config.yaml collector/$(COMPONENT)
	sed -i '' -e 's/component/$(COMPONENT)/g' collector/$(COMPONENT)/config.yaml
	sed -i '' -e 's/component/$(COMPONENT)/g' collector/$(COMPONENT)/docker-compose.yaml
endif

# new-prometheus-component creates a new component that uses a prometheus receiver
.PHONY: new-component-prometheus
new_component_prometheus:
ifndef COMPONENT
	$(error COMPONENT variable was not defined)
else
	mkdir collector/$(COMPONENT)
	cp tools/templates/docker-compose.yaml collector/$(COMPONENT)
	cp tools/templates/config-prometheus.yaml collector/$(COMPONENT)
	sed -i '' -e 's/component/$(COMPONENT)/g' collector/$(COMPONENT)/config.yaml
	sed -i '' -e 's/component/$(COMPONENT)/g' collector/$(COMPONENT)/docker-compose.yaml
endif

.PHONY: replace-api-key
replace-api-key:
ifndef API_KEY
	$(error API_KEY is not defined)
endif
ifndef COMPONENT
	$(error COMPONENT is not defined)
endif
	sed -i '' s/${LIGHTSTEP_ACCESS_TOKEN}/$(API_KEY)/g' collector/$(COMPONENT)/config.yaml

.PHONY: run-component-project-api-key
run-component-project-api-key: replace-api-key run-collector-from-bin docker-component-run

.PHONY: run-component-project
run-component-project: run-collector-from-bin docker-component-run

.PHONY: run-logs-for-collector-examples
run-logs-for-collector-examples:
	# needs to run for all `collector` examples
	# python tools/run-logs.py

.PHONY: parse-logs-for-metrics
parse-logs-for-metrics:
	go run tools/logs2metricstables.go

