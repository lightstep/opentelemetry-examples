#!/usr/bin/env python
#
# example code to test opentelemetry
#
# usage:
#   LS_ACCESS_TOKEN=${SECRET_TOKEN} \
#   OTEL_RESOURCE_ATTRIBUTES=service.name=py-opentelemetry-manual-otlp-client,service.version=10.10.10 \
#   python client.py test \

import os
import time
import requests

from opentelemetry import propagators
from opentelemetry.trace.propagation.tracecontext import TraceContextTextMapPropagator

from common import get_tracer

# Init tracer
tracer = get_tracer()
 
def send_requests(url):
    with tracer.start_as_current_span("client operation"):
        try:
            carrier = {}
            TraceContextTextMapPropagator().inject(carrier)
            header = {"traceparent": carrier["traceparent"]}
            res = requests.get(url, headers=header)
            print(f"Request to {url}, got {len(res.content)} bytes")
        except Exception as e:
            print(f"Request to {url} failed {e}")
            pass


if __name__ == "__main__":
    target = os.getenv("DESTINATION_URL", "http://localhost:8081/ping")
    while True:
        send_requests(target)
        time.sleep(5)
