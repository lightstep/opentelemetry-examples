'use strict';

const ACCESS_TOKEN = process.env.LS_ACCESS_TOKEN;
const COMPONENT_NAME = process.env.LS_SERVICE_NAME || 'lightstep-js-client';
const SERVICE_VERSION = process.env.LS_SERVICE_VERSION || '0.0.1';
const TARGET_URL = process.env.TARGET_URL || 'http://localhost:8080/ping';

const opentracing = require('opentracing');
const lightstep = require('lightstep-tracer');
const tracer = new lightstep.Tracer({
  access_token: ACCESS_TOKEN,
  component_name: COMPONENT_NAME,
  nodejs_instrumentation: true,
});

opentracing.initGlobalTracer(tracer);

const http = require('http');

setInterval(() => {
  console.log('send: ping');
  http.get(TARGET_URL, resp => {
    let data = '';
    resp.on('data', chunk => (data += chunk));
    resp.on('end', () => console.log(`recv: ${data}`));
    resp.on('error', err => console.log('Error: ' + err.message));
  });
}, 500);
