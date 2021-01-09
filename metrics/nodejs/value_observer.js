const meter = require('./factory')('value-observer');
const testCases = require('./test_cases');

const labels = { 'A': 'B' };

let count = 0;

meter.createValueObserver('ValueObserver', {
  description: 'description',
  unit: '1',
}, (observerResult) => {
  if (count < testCases.length) {
    observerResult.observe(testCases[count], labels);
  }
  count++;
});

const promises = [];
for (let i = 0, j = testCases.length; i < j; i++) {
  promises.push(meter.collect());
}

Promise.all(promises).then(() => {
  // shutdown to export metrics before script exits;
  meter.shutdown().then(() => {
    console.log('finished ValueObserver');
  });
});
