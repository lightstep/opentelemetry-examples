.DEFAULT_GOAL:=help

<<<<<<< HEAD
<<<<<<< HEAD
# --------------------------
.PHONY: setup keystore certs all elastic build down stop restart rm logs

keystore:		## Setup Elasticsearch Keystore, by initializing passwords, and add credentials defined in `keystore.sh`.
	docker-compose -f docker-compose.setup.yml run --rm keystore
=======
COMPOSE_ALL_FILES := -f docker-compose.yml -f docker-compose.monitor.yml -f docker-compose.tools.yml -f docker-compose.nodes.yml -f docker-compose.logs.yml
COMPOSE_MONITORING := -f docker-compose.yml -f docker-compose.monitor.yml
COMPOSE_LOGGING := -f docker-compose.yml -f docker-compose.logs.yml
COMPOSE_TOOLS := -f docker-compose.yml -f docker-compose.tools.yml
COMPOSE_NODES := -f docker-compose.yml -f docker-compose.nodes.yml
ELK_SERVICES   := elasticsearch logstash kibana
ELK_LOG_COLLECTION := filebeat
ELK_MONITORING := elasticsearch-exporter logstash-exporter filebeat-cluster-logs
ELK_TOOLS  := rubban
ELK_NODES := elasticsearch-1 elasticsearch-2
ELK_MAIN_SERVICES := ${ELK_SERVICES} ${ELK_MONITORING} ${ELK_TOOLS}
ELK_ALL_SERVICES := ${ELK_MAIN_SERVICES} ${ELK_NODES} ${ELK_LOG_COLLECTION}

=======
>>>>>>> 43056f9 (simplify example)
# --------------------------
.PHONY: setup keystore certs all elastic build down stop restart rm logs

keystore:		## Setup Elasticsearch Keystore, by initializing passwords, and add credentials defined in `keystore.sh`.
	docker-compose -f docker-compose.setup.yml run --rm keystore
<<<<<<< HEAD

>>>>>>> cbfd42f (Refactor docker-compose and .env)
=======
>>>>>>> b10d96e ([es] Remove unused example parts)
certs:		    ## Generate Elasticsearch SSL Certs.
	docker-compose -f docker-compose.setup.yml run --rm certs

setup:		    ## Generate Elasticsearch SSL Certs and Keystore.
	@make certs
	@make keystore

<<<<<<< HEAD
<<<<<<< HEAD
elastic:		    ## Start Elastic.
	docker-compose up -d --build

up:
<<<<<<< HEAD
<<<<<<< HEAD
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
=======
all:		    ## Start Elk and all its component (ELK, Monitoring, and Tools).
	docker-compose ${COMPOSE_ALL_FILES} up -d --build ${ELK_MAIN_SERVICES}

elk:		    ## Start ELK.
=======
elastic:		    ## Start Elastic.
>>>>>>> 43056f9 (simplify example)
	docker-compose up -d --build

up:
=======
	if [ ! -d "secrets" ]; then echo "Did not find secrets directory. Starting setup."; \
=======
	@if [ ! -d "secrets" ]; then echo "Did not find secrets directory. Starting setup."; \
>>>>>>> 93728d9 (update rules and names)
		make setup; echo "Setup complete."; fi
>>>>>>> b10d96e ([es] Remove unused example parts)
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
<<<<<<< HEAD
<<<<<<< HEAD
	@docker volume prune -f --filter label=com.docker.compose.project=elastic
>>>>>>> cbfd42f (Refactor docker-compose and .env)
=======
	@docker volume prune -f --filter label=com.docker.compose.project=Elastic
>>>>>>> 43056f9 (simplify example)
=======
	@docker volume prune -f --filter label=com.docker.compose.project=elasticsearch

clean:
<<<<<<< HEAD
	rm -rf secrets
>>>>>>> b10d96e ([es] Remove unused example parts)
=======
	@make rm
>>>>>>> 93728d9 (update rules and names)

help:       	## Show this help.
	@echo "Make Application Docker Images and Containers using Docker-Compose files in 'docker' Dir."
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m (default: help)\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
