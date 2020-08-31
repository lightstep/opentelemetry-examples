#!/usr/bin/env python3
import time

import requests
from environs import Env
from opentelemetry import trace

tracer = trace.get_tracer(__name__)


def send_requests(destinations):
    with tracer.start_as_current_span("send_requests"):
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
    while True:
        send_requests(destinations)
        time.sleep(2)
