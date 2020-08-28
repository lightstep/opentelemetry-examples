import os

from ddtrace import tracer
from ddtrace.propagation.b3 import B3HTTPPropagator


def get_ls_tracer():
    tracer.configure(http_propagator=B3HTTPPropagator, settings={})
    tracer.set_tags(
        {
            "lightstep.service_name": os.getenv("LS_SERVICE_NAME"),
            "service.version": os.getenv("LS_SERVICE_VERSION"),
            "lightstep.access_token": os.getenv("LS_ACCESS_TOKEN"),
        }
    )


def get_tracer():
    return get_ls_tracer()
