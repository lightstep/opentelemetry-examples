#!/usr/bin/env python3
import os
import time

import requests
import yaml

from ddtrace import tracer
from ddtrace.propagation.b3 import B3HTTPPropagator


tracer.configure(http_propagator=B3HTTPPropagator, settings={})


def send_requests(destinations):
    with tracer.trace("send_requests"):
        for url in destinations:
            try:
                if "/order" in url:
                    res = requests.post(
                        url, data='{"donuts":[{"flavor":"cinnamon","quantity":1}]}'
                    )
                else:
                    res = requests.get(url)
                print(f"Request to {url}, got {len(res.content)} bytes")
            except Exception as e:
                print(f"Request to {url} failed {e}")


if __name__ == "__main__":
    config_file = os.environ.get("INTEGRATION_CONFIG_FILE")
    if not config_file:
        raise Exception("Config file not specified!!")

    config_data = {}
    with open(config_file) as f:
        config_data = yaml.load(f, Loader=yaml.FullLoader)

    destinations = config_data.get("endpoints")
    tracer.set_tags(
        {
            "lightstep.service_name": env.str("LS_SERVICE_NAME"),
            "service.version": env.str("LS_SERVICE_VERSION"),
            "lightstep.access_token": env.str("LS_ACCESS_TOKEN"),
        }
    )
    while True:
        send_requests(destinations)
        time.sleep(2)
