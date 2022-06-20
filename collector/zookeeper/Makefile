.DEFAULT_GOAL:=help

# --------------------------
.PHONY: all zookeeper build down stop restart rm logs

zookeeper:		    ## Start Zookeeper.
	mkdir -p statsout
	chmod 777 -R statsout
	touch statsout/collector-out.json
	chmod 777 statsout/collector-out.json
	docker compose up -d --build

build:			## Build Zookeeper and all its extra components.
	@docker compose build 

ps:				## Show all running containers.
	@docker compose ps

down:			## Down Zookeeper and all its extra components.
	@docker compose down

stop:			## Stop Zookeepersearch and all its extra components.
	@docker compose stop 
	
restart:		## Restart Zookeeper and all its extra components.
	@docker compose restart 

rm:				## Remove Zookeeper and all its extra components containers.
	@docker compose rm
	@rm -rf statsout

logs:			## Tail all logs with -n 1000.
	@docker compose logs --follow --tail=1000 

images:			## Show all Images of Zookeeper and all its extra components.
	@docker compose images 

prune:			## Remove Zookeeper Containers and Delete Zookeeper-related Volume Data 
	@make stop && make rm
	@docker volume prune -f --filter label=com.docker.compose.project=zookeeper

up:
	@make zookeeper

clean:
	rm -rf statsout
	@make rm

help:       	## Show this help.
	@echo "Make Application Docker Images and Containers using Docker-Compose files in 'docker' Dir."
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m (default: help)\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
