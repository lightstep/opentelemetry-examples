import os

import grpc
from opentelemetry import trace
from opentelemetry.exporter.otlp.trace_exporter import OTLPSpanExporter
from opentelemetry.propagators import set_global_textmap
from opentelemetry.propagators.composite import CompositeHTTPPropagator
from opentelemetry.sdk.trace import Resource
from opentelemetry.sdk.trace.export import BatchExportSpanProcessor
from opentelemetry.sdk.trace.propagation.b3_format import B3Format


def get_otlp_exporter():

    if os.getenv("OTEL_EXPORTER_OTLP_INSECURE", False):
        credentials = None
    else:
        credentials = grpc.ssl_channel_credentials()

    return OTLPSpanExporter(
        credentials=credentials,
        headers=(("lightstep-access-token", os.environ.get("LS_ACCESS_TOKEN")),),
    )


def get_otel_tracer():

    set_global_textmap(CompositeHTTPPropagator([B3Format()]))
    span_exporter = get_otlp_exporter()

    trace.get_tracer_provider().add_span_processor(
        BatchExportSpanProcessor(span_exporter)
    )
    trace.get_tracer_provider().resource = Resource(
        {
            "service.name": os.getenv("LS_SERVICE_NAME"),
            "service.version": os.getenv("LS_SERVICE_VERSION"),
        }
    )
    return trace.get_tracer(__name__)


def get_tracer():
    return get_otel_tracer()
