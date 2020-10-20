#!/usr/bin/env python
#
# example code to test ls-trace-py
#
# usage:
#   LS_ACCESS_TOKEN=${SECRET_TOKEN} \
#   LS_SERVICE_NAME=demo-python \
#   LS_SERVICE_VERSION=0.0.8 \
#   ls-trace-run python client.py

import os
import time

from ddtrace import tracer
import requests

from common import get_tracer

get_tracer()


def send_requests(url):
    with tracer.trace("client operation"):
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
