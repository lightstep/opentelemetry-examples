.DEFAULT_GOAL:=help

# --------------------------
.PHONY: zookeeper up build down stop restart rm logs

zookeeper:		    ## Start Zookeeper.
	@make up

up:                 ## Start hte 
	@docker-compose up -d --build

build:			## Build Zookeeper and all its extra components.
	@docker-compose build 

ps:				## Show all running containers.
	@docker-compose ps

down:			## Down Zookeeper and all its extra components.
	@docker-compose down

stop:			## Stop Zookeeper and all its extra components.
	@docker-compose stop 
	
restart:		## Restart Zookeeper and all its extra components.
	@docker-compose restart 

rm:				## Remove Zookeeper and all its extra components containers.
	@docker-compose rm
	@rm -rf secrets
	@rm -rf statsout

logs:			## Tail all logs with -n 1000.
	@docker-compose logs --follow --tail=1000 

images:			## Show all Images of Zookeeper and all its extra components.
	@docker-compose images 

prune:			## Remove Zookeeper Containers 
	@make stop && make rm

clean:
	rm -rf statsout
	@make rm

help:       	## Show this help.
	@echo "Make Application Docker Images and Containers using Docker-Compose files in 'docker' Dir."
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m (default: help)\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
