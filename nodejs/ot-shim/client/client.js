'use strict';

const {
  lightstep,
  opentelemetry,
} = require('lightstep-opentelemetry-launcher-node');
const opentracing = require('opentracing');
const { TracerShim } = require('@opentelemetry/shim-opentracing');

const TARGET_URL = process.env.TARGET_URL || 'http://localhost:8080/ping';
const sdk = lightstep.configureOpenTelemetry();

// development purposes
// const sdk = lightstep.configureOpenTelemetry({
//   serviceName: 'js-ot-shim-client',
//   accessToken: 'YOUR TOKEN',
// });

sdk.start().then(() => {
  const axios = require('axios');

  const otelTracer = opentelemetry.trace.getTracer('nodejs-ot-shim');
  // Init OpenTracing Global Tracer with OpenTelemetry Tracer
  opentracing.initGlobalTracer(new TracerShim(otelTracer));
  // Now get the Open Tracing global tracer
  const tracer = opentracing.globalTracer();

  setInterval(() => {
    const span = tracer.startSpan('client.ping');
    // const span = otelTracer.startSpan('client.ping');
    const headers = {};
    tracer.inject(span, opentracing.FORMAT_HTTP_HEADERS, headers);

    console.log('send: ping', span);

    axios.get(TARGET_URL, { timeout: 100, headers: headers }).then(resp => {
      span.setTag('response', resp.data);
      span.finish();
      // make optimistic flush without waiting for batch processor
      otelTracer.getActiveSpanProcessor().forceFlush();
      console.log(`recv: ${resp.data}`);
    }, (err) => {
      span.setTag('response', err.message);
      span.finish();
      // make optimistic flush without waiting for batch processor
      otelTracer.getActiveSpanProcessor().forceFlush();
      console.log('error', err.message);
    });
  }, 1000);
});
