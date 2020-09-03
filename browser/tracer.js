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
import { CollectorTraceExporter } from '@opentelemetry/exporter-collector';

// Create a provider for activating and tracking spans
const tracerProvider = new WebTracerProvider();

// Configure a span processor and exporter for the tracer
tracerProvider.addSpanProcessor(
  new SimpleSpanProcessor(new ConsoleSpanExporter())
);
tracerProvider.addSpanProcessor(new SimpleSpanProcessor(new CollectorTraceExporter({
  url: 'https://ingest.lightstep.com:443/api/v2/otel/trace',
  headers: {
    'Lightstep-Access-Token': 'YOUR_TOKEN'
  }
})));

tracerProvider.addSpanProcessor(
  new SimpleSpanProcessor(
    new LightstepExporter({
      collectorUrl: 'YOUR_SATELLITE_URL',
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
