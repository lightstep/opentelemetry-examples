'use strict';

const {
  lightstep,
  opentelemetry,
} = require('lightstep-opentelemetry-launcher-node');

const { OTTracePropagator } = require('@opentelemetry/propagator-ot-trace');

const PORT = process.env.PORT || 8080;

const sdk = lightstep.configureOpenTelemetry({
  serviceName: 'otel-js-server (ot-ctx)',
  textMapPropagator: new OTTracePropagator(),
  spanEndpoint: 'https://ingest.lightstep.com/traces/otlp/v0.6',
});

sdk.start().then(() => {
  console.log(opentelemetry.trace.getTracer('default'));

  const express = require('express');
  const app = express();
  app.use(express.json());

  app.get('/', (req, res) => {
    res.send('running...');
  });

  app.get('/ping', (req, res) => {
    res.send('pong');
  });

  app.listen(PORT);
  console.log(`Running on ${PORT}`);
});
