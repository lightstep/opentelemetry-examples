#!/usr/bin/env python3
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
    if service.endswith("server"):
        servers.append(service)
        continue

    if service.endswith("client"):
        clients.append(service)
        environment = data.get("environment")
        print(environment)
        if environment:
            for var in environment:
                if var.startswith("DESTINATION_URL"):
                    endpoints.append(var.split("=")[-1])
        continue

with open(output_path, "w") as output_file:
    yaml.dump(
        {"services": servers, "clients": clients, "endpoints": endpoints}, output_file
    )

