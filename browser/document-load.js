'use strict';

import { WebTracerProvider } from '@opentelemetry/web';
import { ConsoleSpanExporter, SimpleSpanProcessor } from '@opentelemetry/tracing';
import { CollectorTraceExporter } from '@opentelemetry/exporter-collector';
import { DocumentLoad } from '@opentelemetry/plugin-document-load';
import { ZoneContextManager } from '@opentelemetry/context-zone';

// Create a provider for activating and tracking spans
const tracerProvider = new WebTracerProvider({
  plugins: [
    new DocumentLoad(),
  ],
});

// Configure a span processor and exporter for the tracer
tracerProvider.addSpanProcessor(new SimpleSpanProcessor(new ConsoleSpanExporter()));
tracerProvider.addSpanProcessor(new SimpleSpanProcessor(new CollectorTraceExporter({
  url: 'https://ingest.lightstep.com:443/api/v2/otel/trace',
  headers: {
    'Lightstep-Access-Token': 'YOUR_TOKEN'
  }
})));

// Register the tracer
tracerProvider.register({
  contextManager: new ZoneContextManager(),
});
