#!/usr/bin/env python
#
# perform requests to a local kong service
#
#   export DESTINATION_HOST=my-host-address
#   export COLLECTOR_HOST=my-collector-host
#   python client.py

import os
import time
import requests
import sys

def config_kong(host, collector):
    while True:
        try:
            requests.get(f"http://{host}:8001/")
            break
        except Exception as e:
            print("Kong API is not running yet, sleeping")
            time.sleep(1)
            continue

    try:
        # Configure the OpenTelemetry plugin
        res = requests.post(f"http://{host}:8001/plugins", json={
            'name': 'opentelemetry',
            'config': {
                'endpoint': f"http://{collector}:4318/v1/traces",
                'resource_attributes': { 'service.name': 'kong-dev' },
            },
        })
        print(res.text)
        # Create the service
        res = requests.post(f"http://{host}:8001/services", json={
            'name': 'servicenow_service',
            'url': 'http://www.servicenow.com',
        })
        print(res.text)
        # Create the route
        res = requests.post(f"http://{host}:8001/services/servicenow_service/routes", json={
            'paths': ['/servicenow'],
            'name': 'servicenow_route',
        })
        print(res.text)
    except Exception as e:
        print(f"Initial configuration failed: {e}")
        sys.exit(1)

def send_requests(url):
    try:
        res = requests.get(url)
        print(f"Request to {url}, got {len(res.content)} bytes")
    except Exception as e:
        print(f"Request to {url} failed {e}")

if __name__ == "__main__":
    host = os.getenv("DESTINATION_HOST", "localhost")
    collector = os.getenv("COLLECTOR_HOST", "localhost")
    config_kong(host, collector)

    target = f"http://{host}:8000/servicenow/company.html"
    while True:
        send_requests(target)
        time.sleep(.5)
