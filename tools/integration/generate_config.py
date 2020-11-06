#!/usr/bin/env python3
#
# This script generates a yaml config file by reading in the docker-compose
# file, then outputting client, services and endpoints which are then
# used by the integration-test, demo-client-lstrace and demo-client-otlp
# containers

import os
import yaml


config_path = os.environ.get("DOCKER_COMPOSE_PATH", "/docker-compose.yml")
output_path = os.environ.get("OUTPUT_PATH", "/config/integration.yml")

config = None
servers = []
clients = []
endpoints = []

with open(config_path) as config_file:
    config = yaml.load(config_file, Loader=yaml.FullLoader)

for service, data in config.get("services", {}).items():
    environment = data.get("environment")
    if environment:
        for var in environment:
            if var.startswith("DESTINATION_URL"):
                endpoints.append(var.split("=")[-1])

    if service.endswith("server"):
        servers.append(service)
        continue

    if service.endswith("client"):
        clients.append(service)
        continue

with open(output_path, "w") as output_file:
    yaml.dump(
        {"services": servers, "clients": clients, "endpoints": endpoints}, output_file
    )

