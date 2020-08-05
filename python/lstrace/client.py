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
import random
import time

from ddtrace import tracer
import requests

from common import get_tracer

get_tracer()


def send_requests(target):
    integrations = ["pymongo", "redis", "sqlalchemy"]
    with tracer.trace("client operation"):
        for i in integrations:
            url = f"{target}/{i}/{random.randint(1,1024)}"
            try:
                res = requests.get(url)
                print(f"Request to {url}, got {len(res.content)} bytes")
            except Exception as e:
                print(f"Request to {url} failed {e}")
                pass


if __name__ == "__main__":
    target = os.getenv("TARGET_URL", "http://localhost:8081")
    while True:
        send_requests(target)
        time.sleep(5)
