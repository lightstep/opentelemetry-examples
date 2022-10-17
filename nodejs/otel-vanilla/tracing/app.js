'use strict';
const opentelemetry = require('@opentelemetry/api');
const tracer = opentelemetry.trace.getTracer('otel-js-demo');
let count = 0;

setInterval(() => {
  // start a trace by starting a new span
  tracer.startActiveSpan('parent', (parent) => {
    // set an attribute
    parent.setAttribute('count', count);
    // record an event
    parent.addEvent(`message: ${count}`);

    // create a child span
    const child1 = tracer.startSpan('child-1');
    child1.end();

    // create a second child span
    const child2 = tracer.startSpan('child-2');
    // record an error status on a span
    const err = new Error('there was a problem');
    child2.setStatus({code: opentelemetry.SpanStatusCode.ERROR, message: err.message});
    // record the err as an exception (event) on the span
    child2.recordException(err);
    child2.end();

    // end the trace
    parent.end();
    count++;
  });
}, 10000);
