'use strict';

const PORT = process.env.PORT || 8080;
const ACCESS_TOKEN = process.env.LS_ACCESS_TOKEN;
const COMPONENT_NAME =
  process.env.LIGHTSTEP_COMPONENT_NAME || 'js-lstrace-server';
const SERVICE_VERSION = process.env.LIGHTSTEP_SERVICE_VERSION || '0.0.1';

const express = require('express');
const tracer = require('ls-trace').init({
  experimental: {
    b3: true,
  },
  tags: `lightstep.service_name:${COMPONENT_NAME},lightstep.access_token:${ACCESS_TOKEN},service.version:${SERVICE_VERSION}`,
});

const app = express();
app.get('/', (req, res) => {
  res.send('running...');
});

app.get('/ping', (req, res) => {
  res.send('pong');
});

app.listen(PORT);
console.log(`Running on ${PORT}`);
