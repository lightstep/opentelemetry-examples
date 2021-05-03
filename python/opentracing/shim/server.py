# Set the following environment variables first before running this server:
#
# LS_ACCESS_TOKEN
# LS_SERVICE_NAME
#
# Spans will be exported to https://ingest.staging.lightstep.com:443

from os import environ
from time import sleep

import flask

from opentelemetry.propagate import extract, set_global_textmap
from opentelemetry.trace import (
    get_tracer, set_tracer_provider, get_tracer_provider
)
from opentelemetry.propagators.ot_trace import OTTracePropagator
from opentelemetry.sdk.trace import TracerProvider, Resource
from opentelemetry.launcher.tracer import LightstepOTLPSpanExporter
from opentelemetry.sdk.trace.export import BatchSpanProcessor
from grpc import ssl_channel_credentials

set_tracer_provider(TracerProvider())

set_global_textmap(OTTracePropagator())

get_tracer_provider().add_span_processor(
    BatchSpanProcessor(
        LightstepOTLPSpanExporter(
            endpoint="https://ingest.staging.lightstep.com:443",
            credentials=ssl_channel_credentials(),
            headers=(("lightstep-access-token", environ["LS_ACCESS_TOKEN"]),)
        )
    )
)
get_tracer_provider()._resource = Resource(
    {"service.name": environ["LS_SERVICE_NAME"], "service.version": "1.2.3"}
)


tracer = get_tracer(__name__)
app = flask.Flask(__name__)
tracer = get_tracer(__name__)


@app.route("/ping")
def ping():

    context = extract(flask.request.headers)

    random_int = int(context["baggage"]["Random-Int"])
    name = "server {}".format(random_int)

    with tracer.start_as_current_span(name, context=context):
        sleep(random_int / 10000)
        return name

if __name__ == "__main__":
    app.run(host="0.0.0.0")
