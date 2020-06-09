'use strict';

// this will be needed to get a tracer
import opentelemetry from '@opentelemetry/api';
// tracer provider for web
import { WebTracerProvider } from '@opentelemetry/web';
// and an exporter with span processor
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

tracerProvider.addSpanProcessor(
  new SimpleSpanProcessor(
    new LightstepExporter({
      // if testing on staging
      collectorUrl: 'https://collector-staging.lightstep.com/api/v2/reports',
      serviceName: 'browser-demo',
      token: 'YOUR_TOKEN',
    })
  )
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
