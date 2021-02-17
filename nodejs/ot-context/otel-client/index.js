'use strict';

const {
  lightstep,
  opentelemetry,
} = require('lightstep-opentelemetry-launcher-node');

const { OTTracePropagator } = require('@opentelemetry/propagator-ot-trace');

const sdk = lightstep.configureOpenTelemetry({
  serviceName: 'otel-js-client (ot-ctx)',
  textMapPropagator: new OTTracePropagator(),
});

const DESTINATION_URL =
  process.env.DESTINATION_URL || 'http://localhost:8080/ping';

sdk.start().then(() => {
  const http = require('http');
  setInterval(() => {
    const tracer = opentelemetry.trace.getTracer('otel-client-example');
    const span = tracer.startSpan('client.ping');
    console.log('send: ping');
    tracer.withSpan(span, () => {
      http.get(DESTINATION_URL, resp => {
        let data = '';
        resp.on('data', chunk => (data += chunk));
        resp.on('end', () => console.log(`recv: ${data}`));
        resp.on('error', err => console.log('Error: ' + err.message));
      });
    });
    span.end();
  }, 500);
});
