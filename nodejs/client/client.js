'use strict';

const ACCESS_TOKEN = process.env.LS_ACCESS_TOKEN;
const COMPONENT_NAME =
  process.env.LIGHTSTEP_COMPONENT_NAME || 'js-lstrace-client';
const SERVICE_VERSION = process.env.LIGHTSTEP_SERVICE_VERSION || '0.0.1';
const TARGET_URL = process.env.TARGET_URL || 'http://localhost:8080/ping';

const tracer = require('ls-trace').init({
  experimental: {
    b3: true,
  },
  tags: `lightstep.service_name:${COMPONENT_NAME},lightstep.access_token:${ACCESS_TOKEN},service.version:${SERVICE_VERSION}`,
});
const scope = tracer.scope();

const http = require('http');

setInterval(() => {
  const span = tracer.startSpan('client.ping');
  console.log('send: ping');
  scope.activate(span, () => {
    http.get(TARGET_URL, resp => {
      let data = '';
      resp.on('data', chunk => (data += chunk));
      resp.on('end', () => console.log(`recv: ${data}`));
      resp.on('error', err => console.log('Error: ' + err.message));
    });
  });
  span.finish();
}, 500);
