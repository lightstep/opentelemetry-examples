#!/usr/bin/env python
#
# example code to test opentelemetry
#
# usage:
#   LS_ACCESS_TOKEN=${SECRET_TOKEN} \
#   python client.py test \

import os
import time
import requests

from opentelemetry import propagators
from opentelemetry.trace.propagation.tracecontext import TraceContextTextMapPropagator

from common import get_tracer

tracer = get_tracer()

def set_header_into_requests_request(request: requests,
                                        key: str, value: str):
    
    # request.header[key] = value
    # print("Blah")
    return {key: value}

def send_requests(url):
    with tracer.start_as_current_span("client operation"):
        try:
            carrier = {}
            TraceContextTextMapPropagator().inject(carrier)
            header = set_header_into_requests_request(requests.request, "traceparent", carrier["traceparent"])
            print(f"header {header}")
            # print(f"Request = {requests.header}")
            res = requests.get(url, headers=header)
            print(f"Request to {url}, got {len(res.content)} bytes")
            # print(f"*** Header: {res.headers}")
        except Exception as e:
            print(f"Request to {url} failed {e}")
            pass


if __name__ == "__main__":
    target = os.getenv("DESTINATION_URL", "http://localhost:8081/rolldice")
    while True:
        send_requests(target)
        time.sleep(5)
