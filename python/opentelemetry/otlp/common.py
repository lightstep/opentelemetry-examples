import os

from opentelemetry import trace
from opentelemetry.exporter.otlp.proto.grpc.trace_exporter import OTLPSpanExporter
from opentelemetry.sdk.resources import SERVICE_NAME, Resource
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor

def get_otlp_exporter():

    return OTLPSpanExporter(
        endpoint="ingest.lightstep.com:443",
        headers=(("lightstep-access-token", os.environ.get("LS_ACCESS_TOKEN")),),
    )


def get_tracer():
    span_exporter = get_otlp_exporter()
    
    # Service name is required for most backends
    resource = Resource(attributes={
        SERVICE_NAME: "test-python-trace"
    })

    provider = TracerProvider(resource=resource)
    processor = BatchSpanProcessor(span_exporter)
    provider.add_span_processor(processor)
    trace.set_tracer_provider(provider)    
    # span_exporter = get_otlp_exporter()

    # trace.get_tracer_provider().add_span_processor(BatchSpanProcessor(span_exporter))
    return trace.get_tracer(__name__)
