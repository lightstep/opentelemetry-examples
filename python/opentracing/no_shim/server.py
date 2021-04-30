# Set the following environment variables first before running this server:
#
# LS_ACCESS_TOKEN
# LS_SERVICE_NAME
#
# Spans will be exported to https://ingest.staging.lightstep.com:443
# NoShim

from os import environ
from time import sleep

import flask

import opentracing
import lightstep
from opentracing.propagation import Format

opentracing.tracer = lightstep.Tracer(
    collector_host="https://ingest.staging.lightstep.com",
    collector_port=443,
    component_name=environ["LS_SERVICE_NAME"],
    access_token=environ["LS_ACCESS_TOKEN"],
)


app = flask.Flask(__name__)


@app.route("/ping")
def ping():

    context = opentracing.tracer.extract(
        Format.HTTP_HEADERS,
        {key: value for key, value in flask.request.headers.items()}
    )

    random_int = int(context.baggage["random-int"])
    name = "server {}".format(random_int)

    # with opentracing.tracer.start_active_span(name, child_of=context):
    with opentracing.tracer.start_active_span(name):
        sleep(random_int / 10000)
    opentracing.tracer.flush()
    return name

if __name__ == "__main__":
    app.run(host="0.0.0.0")
