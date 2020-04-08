const opentelemetry = require('@opentelemetry/api');
const { NodeTracerProvider } = require('@opentelemetry/node');
const {
  SimpleSpanProcessor,
  ConsoleSpanExporter,
} = require('@opentelemetry/tracing');

// Create an exporter for sending span data
const exporter = new ConsoleSpanExporter({
  serviceName: 'demo-service',
});

// Create a provider for activating and tracking spans
const tracerProvider = new NodeTracerProvider({
  plugins: {
    express: {
      enabled: true,
      path: '@opentelemetry/plugin-express',
    },
    http: {
      enabled: true,
      path: '@opentelemetry/plugin-http',
    },
  },
});

// Configure a span processor for the tracer
tracerProvider.addSpanProcessor(new SimpleSpanProcessor(exporter));

// Register the tracer
tracerProvider.register();

const tracer = opentelemetry.trace.getTracer();

const express = require('express');

// --- Simple Express app setup

const port = 3000;
const app = express();

// Attach a mock middleware, automatically picked up by the express-plugin
app.use(function mockMiddleware(req, res, next) {
  console.log('Mock middleware');
  next();
});

// Mount our demo route
app.get('/demo', (req, res) => {
  const span = tracer.startSpan('handler');

  // Annotate our span to capture metadata about our operation
  span.addEvent('doing work');

  mockAdditionalWork(span).then(() => {
    // Be sure to end the span!
    span.end();

    res.send('OpenTelemetry Fun 🎉');
  });
});

// Start the server
app.listen(port, () => console.log(`Example app listening on port ${port}!`));

// --- Mock function, demonstrates creating child spans

async function mockAdditionalWork(parentSpan) {
  // Start a child span using the passed parent span to maintain context
  const span = tracer.startSpan('mockAdditionalWork()', {
    parent: parentSpan,
  });

  // Additional metadata can be attached to spans
  span.setAttribute('key', 'value');
  span.addEvent(
    'Additional work happening, eg calling a database or making a request to another service'
  );

  return new Promise((resolve, reject) => {
    setTimeout(() => {
      // Be sure to end the child span!
      span.end();
      resolve();
    }, 500);
  });
}
