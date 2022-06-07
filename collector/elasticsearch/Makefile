.DEFAULT_GOAL:=help

# --------------------------
.PHONY: setup keystore certs all elastic build down stop restart rm logs

keystore:		## Setup Elasticsearch Keystore, by initializing passwords, and add credentials defined in `keystore.sh`.
	docker-compose -f docker-compose.setup.yml run --rm keystore
certs:		    ## Generate Elasticsearch SSL Certs.
	docker-compose -f docker-compose.setup.yml run --rm certs

setup:		    ## Generate Elasticsearch SSL Certs and Keystore.
	@make certs
	@make keystore

elastic:		    ## Start Elastic.
	docker-compose up -d --build

up:
	@if [ ! -d "secrets" ]; then echo "Did not find secrets directory. Starting setup."; \
		make setup; echo "Setup complete."; fi
	@make elastic

build:			## Build ElasticSearch and all its extra components.
	@docker-compose build 

ps:				## Show all running containers.
	@docker-compose ps

down:			## Down ElasticSearch and all its extra components.
	@docker-compose down

stop:			## Stop Elasticsearch and all its extra components.
	@docker-compose stop 
	
restart:		## Restart ElasticSearch and all its extra components.
	@docker-compose restart 

rm:				## Remove ElasticSearch and all its extra components containers.
	@docker-compose rm
	@rm -rf secrets
	@rm -rf statsout

logs:			## Tail all logs with -n 1000.
	@docker-compose logs --follow --tail=1000 

images:			## Show all Images of ElasticSearch and all its extra components.
	@docker-compose images 

prune:			## Remove ElasticSearch Containers and Delete ElasticSearch-related Volume Data (the Elastic_Elasticsearch-data volume)
	@make stop && make rm
	@docker volume prune -f --filter label=com.docker.compose.project=elasticsearch

clean:
	@make rm

help:       	## Show this help.
	@echo "Make Application Docker Images and Containers using Docker-Compose files in 'docker' Dir."
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m (default: help)\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
