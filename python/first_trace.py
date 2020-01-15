#!/usr/bin/env python3
# first_trace.py
from opentelemetry import trace
from opentelemetry.sdk.trace import Tracer
from opentelemetry.sdk.trace.export import ConsoleSpanExporter
from opentelemetry.sdk.trace.export import BatchExportSpanProcessor

trace.set_preferred_tracer_implementation(lambda T: Tracer())
tracer = trace.tracer()
span_processor = BatchExportSpanProcessor(ConsoleSpanExporter())
tracer.add_span_processor(span_processor)

span = tracer.start_span('foo')
span.set_attribute("platform", "osx")
span.set_attribute("version", "1.2.3")
span.add_event("event in foo", {"name": "foo1"})

attributes = {
  "platform": "osx",
  "version": "1.2.3",
}
child_span = tracer.start_span('baz', parent=span, attributes=attributes)

child_span.end()
span.end()


span = tracer.start_span('foo', attributes=attributes)
with tracer.use_span(span):
    child_span = tracer.start_span('baz', attributes=attributes)
child_span.end()
span.end()

with tracer.start_as_current_span('foo', attributes=attributes):
    with tracer.start_as_current_span('bar', attributes=attributes):
        print("hello")
