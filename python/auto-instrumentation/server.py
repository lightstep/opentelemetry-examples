#!/usr/bin/env python3
# server.py
import os
from collections import defaultdict

import requests

from flask import Flask, request
from opentelemetry import trace
from opentelemetry.ext.lightstep import LightstepSpanExporter
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchExportSpanProcessor

# The preferred tracer implementation must be set, as the opentelemetry-api
# defines the interface with a no-op implementation.
trace.set_tracer_provider(TracerProvider())

# SpanExporter receives the spans and send them to the target location.
exporter = LightstepSpanExporter(
    name="auto-instrument-example",
    service_version="0.7.0",
    token=os.environ.get("LIGHTSTEP_ACCESS_TOKEN"),
)
span_processor = BatchExportSpanProcessor(exporter)
trace.get_tracer_provider().add_span_processor(span_processor)

app = Flask(__name__)
CACHE = defaultdict(int)


@app.route("/")
def fetch():
    url = request.args.get("url")
    CACHE[url] += 1
    resp = requests.get(url)
    return resp.content


@app.route("/cache")
def cache():
    keys = CACHE.keys()
    return "{}".format(keys)


if __name__ == "__main__":
    app.run()
