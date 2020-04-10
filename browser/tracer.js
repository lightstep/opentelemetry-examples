'use strict';

import opentelemetry from '@opentelemetry/api';
import { WebTracerProvider } from '@opentelemetry/web';
import {
  ConsoleSpanExporter,
  SimpleSpanProcessor,
} from '@opentelemetry/tracing';
import { LightstepExporter } from 'lightstep-opentelemetry-exporter';

// Create a provider for activating and tracking spans
const tracerProvider = new WebTracerProvider();

// Configure a span processor and exporter for the tracer
tracerProvider.addSpanProcessor(
  new SimpleSpanProcessor(new ConsoleSpanExporter())
);

// Register the tracer
tracerProvider.register();

const tracer = opentelemetry.trace.getTracer('lightstep-web-example');

const span = tracer.startSpan('foo');
span.setAttribute('name', 'value');
span.addEvent('event in foo');

const childSpan = tracer.startSpan('bar', {
  parent: span,
});

childSpan.end();
span.end();
