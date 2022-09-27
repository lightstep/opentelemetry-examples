#!/usr/bin/env python
#
# example code to test opentelemetry
#
# usage:
#  OTEL_EXPORTER_OTLP_TRACES_HEADERS="lightstep-access-token=<LS_ACCESS_TOKEN>"
#   opentelemetry-instrument \
#         --traces_exporter console,otlp \
#       --metrics_exporter console,otlp \
#       --service_name test-py-auto-client \
#       --exporter_otlp_endpoint "ingest.lightstep.com:443" \
#       python client.py test
#
# See README.md for more details.


import os
import time
import requests

# from common import get_tracer
from opentelemetry import trace

# tracer = get_tracer()
tracer = trace.get_tracer_provider().get_tracer(__name__)


def send_requests(url):
    with tracer.start_as_current_span("client operation"):
        try:
            res = requests.get(url)
            print(f"Request to {url}, got {len(res.content)} bytes")
            # print(f"Response returned: {res.json()}")
        except Exception as e:
            print(f"Request to {url} failed {e}")


if __name__ == "__main__":
    target = os.getenv("DESTINATION_URL", "http://localhost:8081/ping")
    while True:
        send_requests(target)
        time.sleep(5)
