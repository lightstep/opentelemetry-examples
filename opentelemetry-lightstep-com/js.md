Get up and running  with OpenTelemetry in just a few quick steps! The setup process consists of two phases--getting OpenTelemetry installed and configured, and then validating that configuration to ensure that data is being sent as expected. This guide explains how to download, install, and run OpenTelemetry in Node.js and the browser.

## Install and run OpenTelemetry JS: Node.js

To install OpenTelemetry, we recommend our handy [OTel-Launcher](https://github.com/lightstep/otel-launcher-node), which simplifies the process.

```bash
npm install lightstep-opentelemetry-launcher-node --save
```

Once you've downloaded the launcher, you can run OpenTelemetry using the following basic configuration.

The full list of configuration options can be found in the [README](https://github.com/lightstep/otel-launcher-node).

### Requirements

* Node.js version 12 or newer
* An app to add OpenTelemetry to. You can use [this example application](https://github.com/lightstep/opentelemetry-examples/tree/main/nodejs) or bring your own.
* A [free Lightstep account](https://app.lightstep.com/signup/developer?signup_source=otelnode) account, or another OpenTelemetry backend. 

> LS Note: When connecting to Lightstep, a project [Access Token](https://docs.lightstep.com/paths/gs-lightstep-path/step-three#create-an-access-token) is required.

```js
const { lightstep, opentelemetry } = require('lightstep-opentelemetry-launcher-node');

const sdk = lightstep.configureOpenTelemetry({
  accessToken: 'YOUR_ACCESS_TOKEN',
  serviceName: 'example-service',
});

sdk.start().then(() => {
  // All of your application code and any imports that should leverage
  // OpenTelemetry automatic instrumentation must go here.
});
```

### Validate installation by checking for traces

With your application running, you can now verify that you’ve installed OpenTelemetry correctly by confirming that telemetry data is being reported to your observability backend. <br>

To do this, you need to make sure that your application is actually generating data. Applications will generally not produce traces unless they are being interacted with, and opentelemetry will often buffer data before sending it. So it may take some amount of time and interaction before your application data begins to appear in your backend..

> __Validate your traces in Lightstep:__
> 1. Trigger an action in your app that generates a web request.
> 2. In Lightstep, click on the Explorer in the sidebar.
> 3. Refresh your query until you see traces.
> 4. View the traces and verify that important aspects of your application are captured by the trace data.

## Troubleshooting your JavaScript installation

Having trouble with installation? Check the following tips.

### Logging to the console

Set an environment variable to run the OpenTelemetry launcher in debug mode, where it logs details about the configuration and emitted spans:

```shell
export OTEL_LOG_LEVEL=debug
```

The output may be very verbose with some benign errors. Early in the console output, look for logs about the configuration and check that your access token is correct. Next, look for lines like the ones below, which are emitted when spans are emitted to Lightstep.

```json
{
  "traceId": "985b66d592a1299f7d12ebca56ca1fe3",
  "parentId": "8d62a70aa335a227",
  "name": "bar",
  "id": "17ada85c3d55376a",
  "kind": 0,
  "timestamp": 1685674607399000,
  "duration": 299,
  "attributes": {},
  "status": { "code": 0 },
  "events": []
}
{
  "traceId": "985b66d592a1299f7d12ebca56ca1fe3",
  "name": "foo",
  "id": "8d62a70aa335a227",
  "kind": 0,
  "timestamp": 1585130342183948,
  "duration": 315,
  "attributes": {
    "name": "value"
  },
  "status": { "code": 0 },
  "events": [
    {
      "name": "event in foo",
      "time": [1585130342, 184213041]
    }
  ]
}
```

### Running short applications (Lambda/Serverless/etc)
If your application exits quickly after startup, you may need to explicitly shutdown the tracer to ensure that all spans are flushed:

```
opentelemetry.trace.getTracer('your_tracer_name').getActiveSpanProcessor().shutdown()
```

## Install and run OpenTelemetry JS in browsers

This section will show you how to use [OpenTelemetry](https://opentelemetry.io) in your browser to:

- Configure a tracer
- Generate trace data
- Propagate context over HTTP
- Export the trace data to the console and to the Lightstep
- Enable auto instrumentation for document load
- Enable auto instrumentation for button any XMLHttpRequest

The full code for the example in this guide can be found [here](https://github.com/lightstep/opentelemetry-examples/tree/master/browser).

### Requirements 

* An up to date modern browser.
* An app to add OpenTelemetry to. You can use [this example application](https://github.com/lightstep/opentelemetry-examples/tree/main/browser) or bring your own.
* A Lightstep account, or another OpenTelemetry backend. 

> Need an account? [Create a free Lightstep account here](https://app.lightstep.com/signup/developer?signup_source=otelnode).

To use OpenTelemetry, you need to install the API, SDK, span processor and exporter packages. The version of the SDK and API used in this guide is **0.16.0**, the most current version as of writing.

```bash
npm install @opentelemetry/api @opentelemetry/web @opentelemetry/tracing --save
```

Once you've downloaded the launcher, you can run OpenTelemetry using the following basic configuration.

> LS Note: When connecting to Lightstep, a project Access Token is required.

1. Import OpenTelemetry and create a tracer configured to send data to the console, saving it as `tracer.js`.

```js
// this will be needed to get a tracer
import opentelemetry from '@opentelemetry/api';
// tracer provider for web
import { WebTracerProvider } from '@opentelemetry/web';
// and an exporter with span processor
import {  
  SimpleSpanProcessor,
} from '@opentelemetry/tracing';
import { CollectorTraceExporter } from '@opentelemetry/exporter-collector';

// Create a provider for activating and tracking spans
const tracerProvider = new WebTracerProvider();

// Connect to Lightstep by configuring the exporter with your endpoint and access token.
tracerProvider.addSpanProcessor(new SimpleSpanProcessor(new CollectorTraceExporter({
  url: 'https://ingest.lightstep.com:443/api/v2/otel/trace',
  headers: {
    'Lightstep-Access-Token': 'YOUR_TOKEN'
  }
})));

// Register the tracer
tracerProvider.register();

```

2. Load the tracer into your HTML document.

```html
<script type="text/javascript" src="tracer.js"></script>
```

### Validate installation by checking for traces

With your application running, you can now verify that you’ve installed OpenTelemetry correctly by confirming that telemetry data is being reported to your observability backend. <br>

To do this, you need to make sure that your application is actually generating data. Applications will generally not produce traces unless they are being interacted with, and opentelemetry will often buffer data before sending it. So it may take some amount of time and interaction before your application data begins to appear in your backend..

> __Validate your traces in Lightstep:__
> 1. Trigger an action in your app that generates a web request.
> 2. In Lightstep, click on the Explorer in the sidebar.
> 3. Refresh your query until you see traces.
> 4. View the traces and verify that important aspects of your application are captured by the trace data.

## OpenTracing support for JavaScript

The OpenTracing shim allows existing OpenTracing instrumentation to report to the OpenTelemetry SDK. OpenTracing support is not enabled by default. Instructions for enabling the shim can be found in the README.

* [JS OpenTracing shim](https://github.com/open-telemetry/opentelemetry-js/tree/master/packages/opentelemetry-shim-opentracing) 

Read more about upgrading to OpenTelemetry in our [OpenTracing Migration Guide](/migrating-from-opentracing).

### Library and framework support

OpenTelemetry automatically provides instrumentation for a large number of libraries and frameworks, right out of the box.

The full list of supported plugins can be found in the [README](https://github.com/open-telemetry/opentelemetry-js/#plugins).

## OpenTelemetry JavaScript Resources

Telemetry data is indexed by service. In OpenTelemetry, services are described by **resources**, which are set when the OpenTelemetry SDK is initialized during program startup.

We want our data to be normalized, so we can compare apples to apples. OpenTelemetry defines a schema for the keys and values which describe common service resources such as hostname, region, version, etc. 
These standards are called Semantic Conventions, and are defined in the [OpenTelemetry Specification](https://github.com/open-telemetry/opentelemetry-specification/tree/master/specification/resource/semantic_conventions#resource-semantic-conventions).

We recommend that, at minimum, the following resources be applied to every service:

| Attribute  | Description  | Example  | Required? |
|---|---|---|---|
| service.name | Logical name of the service. <br/> MUST be the same for all instances of horizontally scaled services. | `shoppingcart` | Yes |
| service.namespace | A namespace for `service.name`. A string value having a meaning that helps to distinguish a group of services. | `Shop` | No |
| service.instance.id | The string ID of the service instance. | `627cc493-f310-47de-96bd-71410b7dec09` | No |
| service.version | The version string of the service API or implementation as defined in [Version Attributes](#version-attributes). | `semver:2.0.0` | No |

## Node.js configuration

At this time, resources for NodeJS can only be set from the command line via the `OTEL_RESOURCE_ATTRIBUTES` environment variable.

```bash
OTEL_RESOURCE_ATTRIBUTES=service.name:myservice,service.version:1.2.3
```

The format is a comma-separated list of attributes, e.g. `key1:val1,key2:val2`.

### Semantic conventions

Standardizing the format of your data is a critical part of using OpenTelemetry. OpenTelemetry provides a schema for describing common resources, so that backends can easily parse and identify relevant information.  

It is important to understand these conventions when writing instrumentation, in order to normalize your data and increase its utility.
The semantic conventions for resources can be found in the specification.

The following types of resources are currently defined:
* [Service](https://github.com/open-telemetry/opentelemetry-specification/tree/master/specification/resource/semantic_conventions#service)
* [Host](https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/resource/semantic_conventions/host.md)
* [Telemetry](https://github.com/open-telemetry/opentelemetry-specification/tree/master/specification/resource/semantic_conventions#telemetry-sdk)
* [Container](https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/resource/semantic_conventions/container.md)
* [Function as a Service](https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/resource/semantic_conventions/faas.md)
* [Process](https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/resource/semantic_conventions/process.md)
* [Kubernetes](https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/resource/semantic_conventions/k8s.md)
* [Operating System](https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/resource/semantic_conventions/os.md)
* [Cloud IaaS](https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/resource/semantic_conventions/cloud.md)

## OpenTelemetry Spans in JavaScript

OpenTelemetry comes with many instrumentation plugins for libraries and frameworks. This should be enough detail to get started with tracing in production. 

As great as that is, you will still want to add additional spans to your application code, in order to break down larger operations and gain more detailed insights into where your application is spending its time.

> When you create a new span to measure a subcomponent, that span is added to the current trace as the `child` of the current span, and then becomes the current span itself.

## JavaScript Spans and the Tracer API

### Accessing the tracer

In order to interact with traces, you must first acquire a handle to a `Tracer`. 

By convention, Tracers are named after the component they are instrumenting; usually a library, a package, or a class.

```js
const tracer = opentelemetry.trace.getTracer('my-package-name');
```

> Note that there is no need to "set" a tracer by name before getting it. The name you provide is to help identify which component generated which spans, and to potentially disable tracing for individual components.

 We recommend calling `getTracer` once per component during initialization and retaining a handle to the tracer, rather than calling `getTracer` repeatedly.

### Accessing the current span
Ideally, when tracing application code, spans are created and managed in the application framework.

If your application framework is supported, a trace will automatically be created for each request, and your application code will already be wrapped in a span, which can be used for adding application specific attributes and events.

To access the currently active span, call `getSpan(context.active())`:

```js
import { context, getSpan } from '@opentelemetry/api';
const span = getSpan(context.active());
```

### Setting a new current span
Let’s demonstrate creating a new span by example. Imagine you have an automated kitchen, and you want to time how long the robot chef takes to bake a cake. 
The naive way to do this would be to just start a span, call your method, then end the span:

```js
// make a child span to measure how long it takes to bake a cake.
import { context, getSpan } from '@opentelemetry/api';
const cakeSpan = tracer.startSpan('bake-cake');

chef.bakeCake()
cakeSpan.end()
```

The above example will work just fine, but with one big problem: the `bakeCake` method itself has no access to this new `bake-cake` span. 
That means there would be no way to add attributes and events to this span from within the bakeCake method. 
Even worse, `get_current_span` would return the parent of "bake-cake", since that span is still set as current. 

What should we do instead? Replace the current span with `bake-cake.` To do this first set the desired span to be an active span in current context `setSpan(context.active(), cakeSpan)` and then call `context.with` to make a closure around the `bakeCake` method. Within this closure, the `getSpan` method will now return `bake-cake`.

```js
import { context, setSpan } from '@opentelemetry/api';
// Replace the current active span with a new child span
const cakeSpan = tracer.startSpan("bake-cake");
context.with(setSpan(context.active(), cakeSpan), () => {
  // now returns the bake-cake span
  console.log(getSpan(context.active()));
  chef.bake_cake()
});
```

This pattern of wrapping method calls is important, because we always want application code to be able to assume that the current span is correct.

## Attributes in OpenTelemetry JavaScript

When performing root cause analysis, span attributes are an important tool for pinpointing the source of performance issues. 

### Setting attributes

> Note that it is only possible to set attributes, not to get them.

Much like how resources are used to describe your services, attributes are used to describe your spans. 
Here is an example of setting attributes to correctly define an HTTP client request:

```js
import * as api from '@opentelemetry/api';

const parentSpan = tracer.startSpan('parent');
api.setSpan(api.context.active(), parentSpan);

const span = tracer.startSpan('handleRequest', {
    kind: api.SpanKind.CLIENT, // server
    attributes: {
      "http.method": "GET",
      "http.flavor": "1.1",
      "http.url": "https://example.com:8080/project/123/list/?page=2",
      "net.peer.ip": "192.0.2.5",
      "http.status_code": 200,
      "http.status_text": "OK"},
  });
  
// In addition to the standard attributes, custom attributes can be added as well.
span.setAttribute("list.page_number", 2);

// To avoid collisions, always namespace your attribute keys using dot notation.
span.setAttribute("project.id", 2);

// attributes can be added to a span at any time before the span is finished.
span.end()
```

### Conventions

Spans represent specific operations in and between systems. Many operations represent well-known protocols like HTTP or database calls. Like with resources, OpenTelemetry defines a schema for the attributes which describe these common operations. These standards are called Semantic Conventions, and are defined in the [OpenTelemetry Specification](https://github.com/open-telemetry/opentelemetry-specification/tree/master/specification/trace/semantic_conventions).

OpenTelemetry provides a schema for describing common attributes so that backends can easily parse and identify relevant information.  It is important to understand these conventions when writing instrumentation, in order to normalize your data and increase its utility.

The following semantic conventions are defined for tracing:

* [General](https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/trace/semantic_conventions/span-general.md): General semantic attributes that may be used in describing different kinds of operations.
* [HTTP](https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/trace/semantic_conventions/http.md): Spans for HTTP client and server.
* [Database](https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/trace/semantic_conventions/database.md): Spans for SQL and NoSQL client calls.
* [RPC/RMI](https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/trace/semantic_conventions/rpc.md): Spans for remote procedure calls (e.g., gRPC).
* [Messaging](https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/trace/semantic_conventions/messaging.md): Spans for interaction with messaging systems (queues, publish/subscribe, etc.).
* [FaaS](https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/trace/semantic_conventions/faas.md): Spans for Function as a Service (e.g., AWS Lambda).

## Events in OpenTelemetry Javascript

The finest-grained tracing tool is the event system. 

Span events are a form of structured logging. Each event has a name, a timestamp, and a set of attributes. 
When events are added to a span, they inherit the span's context. 
This additional context allows events to be searched, filtered, and grouped by trace ID and other span attributes. 

> Span context is one of the key differences between distributed tracing and traditional logging.

### Adding events
Events are automatically timestamped when they are added to a span. Timestamps can also be set manually if the events are being added after the fact.

For example, enqueuing an item might be recorded as an event.

```js
// Get the current span
const span = api.getSpan(api.context.active());

// Perform the action
queue.enqueue(myItem);

// Record the action
span.addEvent( "enqueued item", {
    "item.id": myItem.ID(),
    "queue.id": queue.ID(),
    "queue.length": queue.length(),
})
```

> Spans should be created for recording course-grained operations, and events should be created for recording fine-grained operations.

### Recording exceptions

Many of the tracing conventions can apply to event attributes as well as span attributes. 
The most important event-specific convention is [recording exceptions](https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/trace/semantic_conventions/exceptions.md).

```js
const span = api.getSpan(api.context.active());

// recordException converts the error into a span event. 
span.recordException(err);

// If the exception means the operation results in an 
// error state, you can also use it to update the span status.
span.setStatus({ code: api.SpanStatusCode.ERROR });
```

> Marking the span as an error is independent of recordings exceptions. To mark the entire span as an error, and have it count against error rates, set the SpanStatus to any value other than OK.

StatusCode definitions can be found in the [OpenTelemetry specification](https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/trace/api.md#statuscanonicalcode). If no status code directly maps to the type of error you are recording, set the status code to `UNKNOWN` for common errors, and `INTERNAL` for serious errors.
