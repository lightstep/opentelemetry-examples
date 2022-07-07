.DEFAULT_GOAL:=help

# --------------------------
.PHONY: couchbase logs clean help

couchbase:		    ## Start Couchbase.
	docker compose up -d --build

logs:			## Tail all logs with -n 1000.
	@docker compose logs --follow --tail=1000 

clean:			## Remove Couchbase Containers
	@docker compose stop && docker compose rm
	@docker volume prune -f --filter label=com.docker.compose.project=couchbase

help:       	## Show this help.
	@echo "Make Application Docker Images and Containers using Docker-Compose files in 'docker' Dir."
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m (default: help)\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
