const meter = require('./factory')('counter');
const testCases = require('./test_cases');

const labels = { 'A': 'B' };

const metric = meter.createCounter('Counter', {
  description: 'description',
  unit: '1',
});

for (let i = 0, j = testCases.length; i < j; i++) {
  metric.add(testCases[i], labels);
}

// shutdown to export metrics before script exits;
meter.shutdown().then(() => {
  console.log('finished Counter');
});
