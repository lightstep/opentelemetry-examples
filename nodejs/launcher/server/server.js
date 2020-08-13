'use strict';

const {
  lightstep,
  opentelemetry,
} = require('lightstep-opentelemetry-launcher-node');

const PORT = process.env.PORT || 8080;

const sdk = lightstep.configureOpenTelemetry();

sdk.start().then(() => {
  const express = require('express');
  const app = express();
  app.use(express.json());

  app.get('/', (req, res) => {
    res.send('running...');
  });

  app.get('/ping', (req, res) => {
    console.log(req.rawHeaders);
    res.send('pong');
  });

  app.listen(PORT);
  console.log(`Running on ${PORT}`);
});
