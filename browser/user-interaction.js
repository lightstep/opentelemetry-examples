'use strict';

import opentelemetry from '@opentelemetry/api';
import { WebTracerProvider } from '@opentelemetry/web';
import {
  ConsoleSpanExporter,
  SimpleSpanProcessor,
} from '@opentelemetry/tracing';
import { LightstepExporter } from 'lightstep-opentelemetry-exporter';
import { ZoneContextManager } from '@opentelemetry/context-zone';
import { UserInteractionPlugin } from '@opentelemetry/plugin-user-interaction';
import { XMLHttpRequestPlugin } from '@opentelemetry/plugin-xml-http-request';

// Create a provider for activating and tracking spans
const tracerProvider = new WebTracerProvider({
  plugins: [
    new UserInteractionPlugin(),
    new XMLHttpRequestPlugin({
      // this is webpack  auto reload - we can ignore it
      ignoreUrls: [/localhost:8091\/sockjs-node/],
      propagateTraceHeaderCorsUrls: '*',
    }),
  ],
});

// Configure a span processor and exporter for the tracer
tracerProvider.addSpanProcessor(
  new SimpleSpanProcessor(new ConsoleSpanExporter())
);
tracerProvider.addSpanProcessor(
  new SimpleSpanProcessor(
    new LightstepExporter({
      collectorUrl: 'YOUR_SATELLITE_URL',
      token: 'YOUR_TOKEN',
    })
  )
);

// Register the tracer
tracerProvider.register({
  contextManager: new ZoneContextManager().enable(),
});

function getData(url, resolve) {
  return new Promise(async (resolve, reject) => {
    const req = new XMLHttpRequest();
    req.open('GET', url, true);
    req.setRequestHeader('Content-Type', 'application/json');
    req.setRequestHeader('Accept', 'application/json');
    req.send();
    req.onload = function() {
      resolve();
    };
  });
}
const tracer = opentelemetry.trace.getTracer('lightstep-web-example');

window.addEventListener('load', () => {
  const btnAdd = document.getElementById('btn');
  btnAdd.addEventListener('click', () => {
    tracer.getCurrentSpan().addEvent('starting ...');
    getData('https://httpbin.org/get?a=1').then(() => {
      tracer.getCurrentSpan().addEvent('first file downloaded');
      getData('https://httpbin.org/get?a=1').then(() => {
        tracer.getCurrentSpan().addEvent('second file downloaded');
      });
    });
  });
});
