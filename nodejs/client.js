const opentelemetry = require('@opentelemetry/api');
const { NodeTracerProvider } = require('@opentelemetry/node');
const { CollectorTraceExporter } = require('@opentelemetry/exporter-collector');
const {
  SimpleSpanProcessor,
  ConsoleSpanExporter,
} = require('@opentelemetry/tracing');

// --- Setup the tracer for the client

const tracerProvider = new NodeTracerProvider({
  plugins: {
    http: {
      enabled: true,
      path: '@opentelemetry/plugin-http',
    },
  },
});

tracerProvider.addSpanProcessor(new SimpleSpanProcessor(new ConsoleSpanExporter({ serviceName: 'demo-client' })));
tracerProvider.addSpanProcessor(new SimpleSpanProcessor(new CollectorTraceExporter({
  url: 'https://ingest.lightstep.com:443/api/v2/otel/trace',
  headers: {
    'Lightstep-Access-Token': 'YOUR_TOKEN',
  },
})));

tracerProvider.register();

// --- Make a request to the example service

const api = require('@opentelemetry/api');
const axios = require('axios');

const tracer = opentelemetry.trace.getTracer('node-opentelemetry-example');

function clientDemoRequest() {
  console.log('Starting client demo request');

  const span = tracer.startSpan('clientDemoRequest()', {
    parent: tracer.getCurrentSpan(),
    kind: api.SpanKind.CLIENT,
  });

  tracer.withSpan(span, async () => {
    await axios.get('http://localhost:3000/demo');
    span.setStatus({ code: api.CanonicalCode.OK });

    span.end();

    // The process must remain alive for the duration of the exporter flush
    // timeout or spans might be dropped
    console.log('Client request complete, waiting to ensure spans flushed...');
    setTimeout(() => {
      console.log('Done 🎉');
    }, 2000);
  });
}

clientDemoRequest();
