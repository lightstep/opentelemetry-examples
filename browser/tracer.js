'use strict';

import opentelemetry from '@opentelemetry/api';
import { WebTracerProvider } from '@opentelemetry/web';
import { ConsoleSpanExporter, SimpleSpanProcessor } from '@opentelemetry/tracing';
import { LightstepExporter } from 'lightstep-opentelemetry-exporter';

// Create a provider for activating and tracking spans
const tracerProvider = new WebTracerProvider();

// Configure a span processor and exporter for the tracer
tracerProvider.addSpanProcessor(new SimpleSpanProcessor(new ConsoleSpanExporter()));

tracerProvider.addSpanProcessor(new SimpleSpanProcessor(new LightstepExporter({
  token: 'WA1hHti46U7aknMCn42ar/mt4ExmJirNBdrhvKt7JOU1to1Ot6FrolCpD5AzHD4+5sODLtg2lT1p5/+2BPzaNbPNKg6AXVa6vUo+Y2eP'
})));

// Register the tracer
tracerProvider.register();

const tracer = opentelemetry.trace.getTracer('lightstep-web-example');

const span = tracer.startSpan('my first span');
span.setAttribute('my first attribute', 'otel is great!');
span.addEvent('something has happened, save it');
span.end();
