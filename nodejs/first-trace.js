const opentelemetry = require('@opentelemetry/api');
const { NodeTracerProvider } = require('@opentelemetry/node');
const {
  SimpleSpanProcessor,
  ConsoleSpanExporter,
} = require('@opentelemetry/tracing');

// Create an exporter for sending span data
const exporter = new ConsoleSpanExporter();

// Create a provider for activating and tracking spans
const tracerProvider = new NodeTracerProvider();

// Configure a span processor for the tracer
tracerProvider.addSpanProcessor(new SimpleSpanProcessor(exporter));

// Register the tracer
tracerProvider.register();

const tracer = opentelemetry.trace.getTracer();

const span = tracer.startSpan('foo');
span.setAttribute('platform', 'osx');
span.setAttribute('version', '1.2.3');
span.addEvent('event in foo');

const childSpan = tracer.startSpan('bar', {
  parent: span,
});

childSpan.end();
span.end();
