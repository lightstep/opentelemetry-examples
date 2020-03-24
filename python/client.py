#!/usr/bin/env python3
# client.py
import requests

from opentelemetry import trace
from opentelemetry.ext import http_requests
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import ConsoleSpanExporter
from opentelemetry.sdk.trace.export import BatchExportSpanProcessor


exporter = ConsoleSpanExporter()
trace.set_tracer_provider(TracerProvider())
span_processor = BatchExportSpanProcessor(exporter)
trace.get_tracer_provider().add_span_processor(span_processor)

http_requests.enable(trace.get_tracer_provider())
response = requests.get(url="http://127.0.0.1:5000/")
span_processor.shutdown()
