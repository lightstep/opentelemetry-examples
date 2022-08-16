import os

from opentelemetry import trace
from opentelemetry.exporter.otlp.proto.grpc.trace_exporter import OTLPSpanExporter
from opentelemetry.sdk.resources import SERVICE_NAME, Resource
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor

def get_otlp_exporter():
    ls_access_token = os.environ.get("LS_ACCESS_TOKEN")
    print(f"Using access token: '{ls_access_token}'")
    return OTLPSpanExporter(
        endpoint="ingest.lightstep.com:443",
        headers=(("lightstep-access-token", ls_access_token),),
    )


def get_tracer():
    span_exporter = get_otlp_exporter()
    
    # Service name is required for most backends
    resource = Resource(attributes={
        SERVICE_NAME: "test-python-server-grpc"
    })

    provider = TracerProvider(resource=resource)
    processor = BatchSpanProcessor(span_exporter)
    provider.add_span_processor(processor)
    trace.set_tracer_provider(provider)    

    return trace.get_tracer(__name__)
