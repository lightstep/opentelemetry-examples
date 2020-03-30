'use strict';

import opentelemetry from '@opentelemetry/api';
// import { UserInteractionPlugin } from '@opentelemetry/plugin-user-interaction';
import { XMLHttpRequestPlugin } from '@opentelemetry/plugin-xml-http-request';
import { WebTracerProvider } from '@opentelemetry/web';
import { ConsoleSpanExporter, SimpleSpanProcessor } from '@opentelemetry/tracing';

// Create a provider for activating and tracking spans
const tracerProvider = new WebTracerProvider({
  plugins: [
    // new UserInteractionPlugin(),
    // new XMLHttpRequestPlugin(),
  ],
});

// Configure a span processor and exporter for the tracer
tracerProvider.addSpanProcessor(new SimpleSpanProcessor(new ConsoleSpanExporter()));

// Register the tracer
tracerProvider.register();

const tracer = opentelemetry.trace.getTracer('lightstep-web-example');

const span = tracer.startSpan('my first span');
span.setAttribute('my first attribute', 'otel is great!');
span.addEvent('something has happened, save it');
span.end();
