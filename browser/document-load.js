'use strict';

import { WebTracerProvider } from '@opentelemetry/web';
import { ConsoleSpanExporter, SimpleSpanProcessor } from '@opentelemetry/tracing';
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
tracerProvider.addSpanProcessor(new SimpleSpanProcessor(new ConsoleSpanExporter()));
tracerProvider.addSpanProcessor(new SimpleSpanProcessor(new LightstepExporter({
  token: 'WA1hHti46U7aknMCn42ar/mt4ExmJirNBdrhvKt7JOU1to1Ot6FrolCpD5AzHD4+5sODLtg2lT1p5/+2BPzaNbPNKg6AXVa6vUo+Y2eP'
})));

// Register the tracer
tracerProvider.register({
  contextManager: new ZoneContextManager().enable(),
});