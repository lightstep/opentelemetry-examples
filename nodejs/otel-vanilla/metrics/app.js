const opentelemetry = require('@opentelemetry/api');
const {
  MeterProvider,
  PeriodicExportingMetricReader,
} = require('@opentelemetry/sdk-metrics-base');
const {
  OTLPMetricExporter,
} = require('@opentelemetry/exporter-metrics-otlp-proto');
const { Resource } = require('@opentelemetry/resources');
const {
  SemanticResourceAttributes,
} = require('@opentelemetry/semantic-conventions');

const token = process.env.LS_ACCESS_TOKEN;
const exportUrl =
  process.env.OTEL_EXPORTER_OTLP_METRICS_ENDPOINT ||
  'https://ingest.lightstep.com/metrics/otlp/v0.9';
const serviceName = process.env.LS_SERVICE_NAME || 'otel-js-demo';

const collectorOptions = {
  url: exportUrl,
  headers: { 'Lightstep-Access-Token': token },
};
const metricExporter = new OTLPMetricExporter(collectorOptions);
const resource = new Resource({
  [SemanticResourceAttributes.SERVICE_NAME]: serviceName,
});
const meterProvider = new MeterProvider({
  resource,
});

opentelemetry.diag.setLogger(
  new opentelemetry.DiagConsoleLogger(),
  opentelemetry.DiagLogLevel.ALL
);

meterProvider.addMetricReader(
  new PeriodicExportingMetricReader({
    exporter: metricExporter,
    exportIntervalMillis: 1000,
  })
);

const prefix = process.env.METRIC_PREFIX || 'demo';

const meter = meterProvider.getMeter('example-meter');
const counter = meter.createCounter(`${prefix}.counter`);
const updownCounter = meter.createUpDownCounter(`${prefix}.updowncounter`);

setInterval(() => {
  counter.add(1);
  updownCounter.add(-1);
}, 1000);

const startTime = Date.now();

const observableCounter = meter.createObservableCounter(
  `${prefix}.observablecounter`
);
observableCounter.addCallback(async (observableResult) => {
  const elapsedSeconds = Math.floor((Date.now() - startTime) / 1000);
  observableResult.observe(elapsedSeconds);
});

const observableUpDownCounter = meter.createObservableUpDownCounter(
  `${prefix}.observableupdowncounter`
);
observableUpDownCounter.addCallback(async (observableResult) => {
  const elapsedSeconds = Math.floor((Date.now() - startTime) / 1000);
  observableResult.observe(-elapsedSeconds);
});

const observableGauge = meter.createObservableGauge(
  `${prefix}.observablegauge`
);
observableGauge.addCallback(async (observableResult) => {
  const value = 50 + Math.random() * 50;
  observableResult.observe(value);
});

const sineWave = meter.createObservableGauge(`${prefix}.sine_wave`);
sineWave.addCallback(async (observableResult) => {
  const secs = Math.floor(Date.now() / 1000);

  observableResult.observe(Math.sin(secs / (50 * Math.PI)), {
    period: 'fastest',
  });
  observableResult.observe(Math.sin(secs / (200 * Math.PI)), {
    period: 'fast',
  });
  observableResult.observe(Math.sin(secs / (1000 * Math.PI)), {
    period: 'regular',
  });
  observableResult.observe(Math.sin(secs / (5000 * Math.PI)), {
    period: 'slow',
  });
});
