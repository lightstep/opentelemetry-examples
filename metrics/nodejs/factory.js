const { MeterProvider } = require('@opentelemetry/metrics');
const { CollectorMetricExporter } = require('@opentelemetry/exporter-collector-grpc');

const exporter = new CollectorMetricExporter({
  serviceName: 'test-nodejs-metrics',
  url: '127.0.0.1:7001',
});

function createMeter(meterName) {
  const meter = new MeterProvider({
    exporter,
    interval: 1000,
  }).getMeter('meter-nodejs' + (meterName ? ('-' + meterName) : ''));

  return meter;
}

module.exports = createMeter;
