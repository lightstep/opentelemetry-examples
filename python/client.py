#!/usr/bin/env python3
# client.py
import requests

from opentelemetry import trace
from opentelemetry.ext import http_requests
from opentelemetry.sdk.trace import Tracer
from opentelemetry.sdk.trace.export import ConsoleSpanExporter
from opentelemetry.sdk.trace.export import BatchExportSpanProcessor


exporter = ConsoleSpanExporter()
trace.set_preferred_tracer_implementation(lambda T: Tracer())
tracer = trace.tracer()
span_processor = BatchExportSpanProcessor(exporter)
tracer.add_span_processor(span_processor)

http_requests.enable(tracer)
response = requests.get(url="http://127.0.0.1:5000/")
span_processor.shutdown()
