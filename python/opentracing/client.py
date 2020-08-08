#!/usr/bin/env python
#

from os import getenv
from random import randint
from time import sleep

from requests import get

from opentracing import set_global_tracer, global_tracer
from lightstep import Tracer


def send_requests(target):
    integrations = ["pymongo", "redis", "sqlalchemy"]
    with global_tracer.start_active_span("client operation"):
        for i in integrations:
            url = f"{target}/{i}/{randint(1,1024)}"
            try:
                res = get(url)
                print(f"Request to {url}, got {len(res.content)} bytes")
            except Exception as e:
                print(f"Request to {url} failed {e}")
                pass


if __name__ == "__main__":
    target = getenv("DESTINATION_URL", "http://localhost:8081")

    set_global_tracer(
        Tracer(
            component_name="testing01",
            access_token=getenv("LS_ACCESS_TOKEN")
        )
    )

    while True:
        send_requests(target)
        sleep(5)
