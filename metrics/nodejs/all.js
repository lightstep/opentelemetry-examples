const { ValueType } = require('@opentelemetry/api');
const { MeterProvider } = require('@opentelemetry/metrics');
const { CollectorMetricExporter } = require('@opentelemetry/exporter-collector-grpc');

const exporter = new CollectorMetricExporter({
  serviceName: 'test-nodejs-metrics',
  url: '127.0.0.1:7001',
});

const meter = new MeterProvider({
  exporter,
  interval: 1000,
}).getMeter('name', "version");


const testValues = require('./test_cases');

const labels = { 'A': 'B' };

// Instruments start here

const counter= meter.createCounter('counter', {
  description: 'description',
  unit: '1',
  valueType: ValueType.INT,
});

for (let i = 0, j = testValues.length; i < j; i++) {
  counter.add(testValues[i], labels);
}

const updowncounter = meter.createUpDownCounter('updowncounter', {
  description: 'description',
  unit: '1',
  valueType: ValueType.INT,
});

for (let i = 0, j = testValues.length; i < j; i++) {
  updowncounter.add(testValues[i], labels);
}

let count = 0;
meter.createSumObserver('sumobserver', {
  description: 'description',
  unit: '1',
  valueType: ValueType.INT,
}, (observerResult) => {
  if (count < testValues.length) {
    observerResult.observe(testValues[count], labels);
  }
  count++;
});

promises = [];
for (let i = 0, j = testValues.length; i < j; i++) {
  promises.push(meter.collect());
}

count = 0;

meter.createUpDownSumObserver('updownsumobserver', {
  description: 'description',
  unit: '1',
  valueType: ValueType.INT,
}, (observerResult) => {
  if (count < testValues.length) {
    observerResult.observe(testValues[count], labels);
  }
  count++;
});

promises = [];
for (let i = 0, j = testValues.length; i < j; i++) {
  promises.push(meter.collect());
}


count = 0;

meter.createValueObserver('valueobserver', {
  description: 'description',
  unit: '1',
  valueType: ValueType.INT,
}, (observerResult) => {
  if (count < testValues.length) {
    observerResult.observe(testValues[count], labels);
  }
  count++;
});

promises = [];
for (let i = 0, j = testValues.length; i < j; i++) {
  promises.push(meter.collect());
}
Promise.all(promises).then(() => {
  // shutdown to export metrics before script exits;
  meter.shutdown().then(() => {
    console.log('finished SumObserver');
  });
});

Promise.all(promises).then(() => {
  // shutdown to export metrics before script exits;
  meter.shutdown().then(() => {
    console.log('finished UpDownSumObserver');
  });
});
Promise.all(promises).then(() => {
  // shutdown to export metrics before script exits;
  meter.shutdown().then(() => {
    console.log('finished ValueObserver');
  });
});
meter.shutdown().then(() => {
  console.log('finished Counter');
});
meter.shutdown().then(() => {
  console.log('finished UpDownCounter');
});
