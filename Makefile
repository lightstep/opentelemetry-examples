
SHELL = /bin/sh

.PHONY: run_from_bin
run_from_bin:
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

.PHONY: new_component
new_component:
ifndef COMPONENT
	$(error COMPONENT variable was not defined)
else
	mkdir collector/$(COMPONENT)
	cp tools/templates/docker-compose.yaml collector/$(COMPONENT)
	cp tools/templates/config.yaml collector/$(COMPONENT)
	sed -i '' -e 's/component/$(COMPONENT)/g' collector/$(COMPONENT)/config.yaml
	sed -i '' -e 's/component/$(COMPONENT)/g' collector/$(COMPONENT)/docker-compose.yaml
endif

# new_component creates a new component that uses a prometheus receiver
.PHONY: new_component_prometheus
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
