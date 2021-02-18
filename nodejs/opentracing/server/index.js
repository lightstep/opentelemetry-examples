'use strict';

const ACCESS_TOKEN = process.env.LS_ACCESS_TOKEN;
const COMPONENT_NAME = process.env.LS_SERVICE_NAME || 'lightstep-js-server';
const SERVICE_VERSION = process.env.LS_SERVICE_VERSION || '0.0.1';
const PORT = process.env.PORT || 8080;

const opentracing = require('opentracing');
const lightstep = require('lightstep-tracer');
const { default: middleware } = require('express-opentracing');

const tracer = new lightstep.Tracer({
  access_token: ACCESS_TOKEN,
  component_name: COMPONENT_NAME,
  nodejs_instrumentation: true,
  verbosity: 4,
});

opentracing.initGlobalTracer(tracer);

const express = require('express');
const app = express();
app.use(middleware({ trace: tracer }));
app.use(express.json());

app.get('/', (req, res) => {
  res.send('running...');
});

app.get('/ping', (req, res) => {
  res.send('pong');
});

app.listen(PORT);
console.log(`Running on ${PORT}`);
