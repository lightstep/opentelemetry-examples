#!/usr/bin/env python
#

from os import getenv
from random import randint
from time import sleep

from requests import get

from opentracing import set_global_tracer, global_tracer

from opentelemetry.trace import get_tracer_provider, set_tracer_provider
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.instrumentation.opentracing_shim import create_tracer
from opentelemetry.launcher import configure_opentelemetry

set_tracer_provider(TracerProvider())
configure_opentelemetry()
shim = create_tracer(get_tracer_provider())


def send_requests(target):
    integrations = ["pymongo", "redis", "sqlalchemy"]
    with global_tracer().start_active_span("client") as client_scope_shim:

        client_scope_shim.span.set_baggage_item("key_client", "value_client")

        print(
            "client shim key_client: {}".format(
                client_scope_shim.span.get_baggage_item("key_client")
            )
        )

        for i in integrations:

            with global_tracer().start_active_span(f"{i}") as (
                integration_scope_shim
            ):

                print(
                    "integration shim key_client: {}".format(
                        integration_scope_shim.span.get_baggage_item(
                            "key_client"
                        )
                    )
                )

                url = f"{target}/{i}/{randint(1,1024)}"
                try:
                    res = get(url)
                    print(f"Request to {url}, got {len(res.content)} bytes")
                except Exception as e:
                    print(f"Request to {url} failed {e}")
                    pass


if __name__ == "__main__":
    target = getenv("DESTINATION_URL", "http://localhost:5000")

    set_global_tracer(shim)

    while True:
        send_requests(target)
        sleep(5)
