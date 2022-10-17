'use strict';

const { NodeSDK } = require('@opentelemetry/sdk-node');
const { OTLPTraceExporter } = require('@opentelemetry/exporter-trace-otlp-proto');
const { Resource } = require('@opentelemetry/resources');
const { SemanticResourceAttributes } = require('@opentelemetry/semantic-conventions');
const { getNodeAutoInstrumentations } = require('@opentelemetry/auto-instrumentations-node');
const { diag, DiagLogLevel, DiagConsoleLogger} = require('@opentelemetry/api');

const token = process.env.LS_ACCESS_TOKEN;
const exportUrl =
  process.env.OTEL_EXPORTER_OTLP_TRACES_ENDPOINT ||
  'https://ingest.lightstep.com/traces/otlp/v0.9';
const serviceName = process.env.LS_SERVICE_NAME || 'otel-js-demo';

const collectorOptions = {
  url: exportUrl,
  headers: { 'Lightstep-Access-Token': token },
};

const traceExporter = new OTLPTraceExporter(collectorOptions);

const sdk = new NodeSDK({
  resource: new Resource({
    [SemanticResourceAttributes.SERVICE_NAME]: 'otel-js-demo',
  }),
  traceExporter,
  instrumentations: [getNodeAutoInstrumentations()],
});

diag.setLogger(new DiagConsoleLogger(), DiagLogLevel.ALL)

sdk.start().then(
    require('./app')
);
