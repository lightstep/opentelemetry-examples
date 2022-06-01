SHELL = /bin/sh

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
.PHONY: new_component
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
