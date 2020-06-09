'use strict';

import { WebTracerProvider } from '@opentelemetry/web';
import {
  ConsoleSpanExporter,
  SimpleSpanProcessor,
} from '@opentelemetry/tracing';
import { LightstepExporter } from 'lightstep-opentelemetry-exporter';
import { DocumentLoad } from '@opentelemetry/plugin-document-load';
import { ZoneContextManager } from '@opentelemetry/context-zone';

// Create a provider for activating and tracking spans
const tracerProvider = new WebTracerProvider({
  plugins: [
    new DocumentLoad(),
  ],
});

// Configure a span processor and exporter for the tracer
tracerProvider.addSpanProcessor(
  new SimpleSpanProcessor(new ConsoleSpanExporter()),
);
tracerProvider.addSpanProcessor(
  new SimpleSpanProcessor(
    new LightstepExporter({
      // if testing on staging
      collectorUrl: 'https://collector-staging.lightstep.com/api/v2/reports',
      token: 'YOUR_TOKEN',
    }),
  ),
);

// Register the tracer
tracerProvider.register({
  contextManager: new ZoneContextManager().enable(),
});
