import os

import grpc
from opentelemetry import trace
from opentelemetry.exporter.otlp.proto.grpc.trace_exporter import OTLPSpanExporter
from opentelemetry.propagate import set_global_textmap
from opentelemetry.propagators.composite import CompositeHTTPPropagator
from opentelemetry.sdk.trace.export import BatchSpanProcessor
from opentelemetry.propagators.b3 import B3Format


def get_otlp_exporter():

    return OTLPSpanExporter(
        headers=(("lightstep-access-token", os.environ.get("LS_ACCESS_TOKEN")),),
    )


def get_otel_tracer():

    set_global_textmap(CompositeHTTPPropagator([B3Format()]))
    span_exporter = get_otlp_exporter()

    trace.get_tracer_provider().add_span_processor(BatchSpanProcessor(span_exporter))
    return trace.get_tracer(__name__)


def get_tracer():
    return get_otel_tracer()
