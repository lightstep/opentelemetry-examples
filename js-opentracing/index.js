const {
  lightstep,
  opentelemetry,
} = require('lightstep-opentelemetry-launcher-node');
const opentracing = require('opentracing');
const axios = require('axios');

// TraceShim is needed to join OpenTelemetry with OpenTracing
const { TracerShim } = require('@opentelemetry/shim-opentracing');

const sdk = lightstep.configureOpenTelemetry({
  accessToken: 'YOUR_TOKEN',
  serviceName: 'CollectorTraceExporter with OpenTracing',
});

sdk.start().then(async () => {
  const otelTracer = opentelemetry.trace.getTracer('node-opentelemetry-opentracing-example');

// Init OpenTracing Global Tracer with OpenTelemetry Tracer
  opentracing.initGlobalTracer(new TracerShim(otelTracer));
// Now get the Open Tracing global tracer
  const tracer = opentracing.globalTracer();

// do a simple OpenTracing test
  const span = tracer.startSpan('OpenTracing Span');
  const headers = {};
  tracer.inject(span, opentracing.FORMAT_HTTP_HEADERS, headers);
  await axios.get('https://raw.githubusercontent.com/open-telemetry/opentelemetry-js/master/package.json', headers).then(response => {
    span.logEvent('package.json loaded', response.data.description);
    span.setTag('name', response.data.name);
    span.setTag('version', response.data.version);
    span.finish();
  }, console.log);
  opentelemetry.trace.getTracerProvider().getActiveSpanProcessor().shutdown();
});

