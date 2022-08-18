from sys import argv

from requests import get

from opentelemetry import trace
from opentelemetry.propagate import inject

tracer = trace.get_tracer_provider().get_tracer(__name__)

assert len(argv) == 2

with tracer.start_as_current_span("client"):

    with tracer.start_as_current_span("client-server"):
        headers = {}
        inject(headers)
        requested = get(
            "http://localhost:8082/rolldice",
            params={"param": argv[1]},
            headers=headers,
        )

        assert requested.status_code == 200