#!/usr/bin/env python
#
# example code to test opentelemetry
#
# usage:
#   LS_ACCESS_TOKEN=${SECRET_TOKEN} \
#   LS_SERVICE_NAME=demo-python \
#   LS_SERVICE_VERSION=0.0.8 \
#   opentelemetry-instrument python client.py

import os
import time
import requests

from opentelemetry import propagators

from common import get_tracer
# from opentelemetry import trace

tracer = get_tracer()
# tracer = trace.get_tracer_provider().get_tracer(__name__)

carrier = {}

def header_from_carrier(carrier, key):
  header = carrier.get(key)
  return [header] if header else []


def send_requests(url):
    ctx = propagators.extract(header_from_carrier, carrier)
    with tracer.start_as_current_span("client operation", context=ctx):
        try:
            res = requests.get(url)
            print(f"Request to {url}, got {len(res.content)} bytes")
        except Exception as e:
            print(f"Request to {url} failed {e}")
            pass


if __name__ == "__main__":
    target = os.getenv("DESTINATION_URL", "http://localhost:8081/ping")
    while True:
        send_requests(target)
        time.sleep(5)
