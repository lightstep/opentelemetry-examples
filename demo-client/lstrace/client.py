#!/usr/bin/env python3
import time

import requests
from environs import Env

from ddtrace import tracer
from ddtrace.propagation.b3 import B3HTTPPropagator


tracer.configure(http_propagator=B3HTTPPropagator, settings={})


def send_requests(destinations):
    with tracer.trace("send_requests"):
        for url in destinations:
            try:
                if "/order" in url:
                    res = requests.post(url, data='{"donuts":[{"flavor":"cinnamon","quantity":1}]}')
                else:
                    res = requests.get(url)
                print(f"Request to {url}, got {len(res.content)} bytes")
            except Exception as e:
                print(f"Request to {url} failed {e}")


if __name__ == "__main__":
    env = Env()
    env.read_env()
    destinations = env.list("DESTINATIONS")
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


curl -X POST -d '{"donuts":[{"flavor":"cinnamon","quantity":1}]}' 