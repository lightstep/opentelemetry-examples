Get up and running  with OpenTelemetry in just a few quick steps! The setup process consists of two phases--getting OpenTelemetry installed and configured, and then validating that configuration to ensure that data is being sent as expected. This guide explains how to download, install, and run OpenTelemetry in Python.

### Requirements
* Python 3.5 or newer
* An app to add OpenTelemetry to. You can use [this example application](https://github.com/tedsuo/otel-python-basics) or bring your own.
* A [free Lightstep account](https://app.lightstep.com/signup/developer?signup_source=otelpython) account, or another OpenTelemetry backend. 

> LS Note: When connecting to Lightstep, a project [Access Token](https://docs.lightstep.com/paths/gs-lightstep-path/step-three#create-an-access-token) is required.

## Install and run OpenTelemetry Python
To install OpenTelemetry, we recommend our handy [OTel-Launcher](https://github.com/lightstep/otel-launcher-python), which simplifies the process.

Installing the launcher using pip:

```bash
pip install opentelemetry-launcher
```

It is also important to install instrumentation for your frameworks and libraries. Once the launcher is installed, you can run [opentelemetry-bootstrap](https://github.com/open-telemetry/opentelemetry-python/tree/master/opentelemetry-instrumentation#opentelemetry-bootstrap) to automatically detect supported libraries and install the associated instrumentation.

```
opentelemetry-bootstrap --action=install
```

Once you've installed the launcher and instrumentation packages, you can now  add OpenTelemetry to your service.

When auto-instrumenting, all configuration parameters must be passed as environment variables. The following configuration options are required.

```
export LS_ACCESS_TOKEN="<ACCESS TOKEN>"
export LS_SERVICE_NAME="<SERVICE NAME>"
export OTEL_PROPAGATORS="b3,baggage"
export OTEL_PYTHON_TRACER_PROVIDER=sdk_tracer_provider
```

The full list of configuration options can be found in the [README](https://github.com/lightstep/otel-launcher-python).

Once you've set up your environent, you can run your application via the [opentelemetry-instrument](https://github.com/open-telemetry/opentelemetry-python/tree/master/opentelemetry-instrumentation#opentelemetry-instrument) command.

```
opentelemetry-instrument python my_service.py
```

`opentelemetry-instrument` automatically integrates the opentelemetry launcher and library instrumentation into your application.

### Validate installation by checking for traces

With your application running, you can now verify that you’ve installed OpenTelemetry correctly by confirming that telemetry data is being reported to your observability backend. <br>

To do this, you need to make sure that your application is actually generating data. Applications will generally not produce traces unless they are being interacted with, and opentelemetry will often buffer data before sending it. So it may take some amount of time and interaction before your application data begins to appear in your backend.

> __Validate your traces in Lightstep:__
> 1. Trigger an action in your app that generates a web request.
> 2. In Lightstep, click on the Explorer in the sidebar.
> 3. Refresh your query until you see traces.
> 4. View the traces and verify that important aspects of your application are captured by the trace data.

## OpenTracing Support for Python
The OpenTracing shim allows existing OpenTracing instrumentation to report to the OpenTelemetry SDK. OpenTracing support is not enabled by default. Instructions for enabling the shim can be found in the README.

```bash
pip install opentelemetry-opentracing-shim
```

* [Python OpenTracing shim](https://github.com/open-telemetry/opentelemetry-python/tree/main/shim/opentelemetry-opentracing-shim) 

Read more about upgrading to OpenTelemetry in our [OpenTracing Migration Guide](/migrating-from-opentracing).

## Troubleshooting your Python installation

Having trouble with installation? Check the following tips.

### Run the launcher in debug mode

If spans are not being reported to Lightstep, try running in debug mode by setting `OTEL_LOG_LEVEL=debug`:

~~~bash
OTEL_LOG_LEVEL=debug opentelemetry-instrument python my_service.py
~~~

The debug log level will print out the configuration information.  Check to ensure that your access token looks correct.  It will also emit every span to the console, which should look something like:

   ~~~shell
{
    "name": "HTTP GET",
    "context": {
        "trace_id": "0x38b3983f92fdd6080facb6a72c09740e",
        "span_id": "0xfc9f1ec29b8b9cf1",
        "trace_state": "[]"
    },
    "kind": "SpanKind.CLIENT",
    "parent_id": null,
    "start_time": "2021-02-08T21:47:34.317562Z",
    "end_time": "2021-02-08T21:47:34.589803Z",
    "status": {
        "status_code": "UNSET"
    },
    "attributes": {
        "component": "http",
        "http.method": "GET",
        "http.url": "https://lightstep.com",
        "http.status_code": 200,
        "http.status_text": "OK"
    },
    "events": [],
    "links": [],
    "resource": {
        "telemetry.sdk.language": "python",
        "telemetry.sdk.version": "0.17b0",
        "service.name": "thing-client",
        "host.name": "myhost"
    }
}
   ~~~

### Ensure gcc and gcc-c++ are installed

Compiling the grpc packages requires gcc and gcc-c++, which may need to be installed manually:

```bash
$ yum -y install python3-devel
$ yum -y install gcc-c++
```

## OpenTelemetry Python Resources

Telemetry data is indexed by service. In OpenTelemetry, services are described by **resources**, which are set when the OpenTelemetry SDK is initialized during program startup.
We want our data to be normalized, so we can compare apples to apples. OpenTelemetry defines a schema for the keys and values which describe common service resources such as hostname, region, version, etc. These standards are called Semantic Conventions, and are defined in the [OpenTelemetry Specification](https://github.com/open-telemetry/opentelemetry-specification/tree/master/specification/resource/semantic_conventions#resource-semantic-conventions).

We recommend that, at minimum, the following resources be applied to every service:

| Attribute  | Description  | Example  | Required? |
|---|---|---|---|
| service.name | Logical name of the service. <br/> MUST be the same for all instances of horizontally scaled services. | `shoppingcart` | Yes |
| service.version | The version string of the service API or implementation as defined in [Version Attributes](#version-attributes). | `semver:2.0.0` | No |
| host.hostname | Contains what the hostname command would return on the host machine. | `server1.mydomain.com,` | No |

## Python Configuration

Resources can be set during launcher configuration at program startup.

```python
import socket

from opentelemetry.launcher import configure_opentelemetry

configure_opentelemetry(
    access_token="my-token",
    service_name="service-123",
    service_version="1.2.3",
    resource_attributes={
      "host.hostname":  socket.gethostname(),
		  "container.name": "my-container-name",
      "cloud.region": "us-central1",    
    }
)
```

### Semantic Conventions

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

The Python constants for these conventions can be found [here](https://javadoc.io/doc/io.opentelemetry/opentelemetry-sdk/latest/io/opentelemetry/sdk/resources/ResourceConstants.html).

## OpenTelemetry Python Spans

OpenTelemetry comes with many instrumentation plugins for libraries and frameworks. This should be enough detail to get started with tracing in production. 

As great as that is, you will still want to add additional spans to your application code, in order to break down larger operations and gain more detailed insights into where your application is spending its time.

> When you create a new span to measure a subcomponent, that span is added to the current trace as the `child` of the current span, and then becomes the current span itself.

## Python Spans and the Tracer API

### Accessing the tracer

In order to interact with traces, you must first acquire a handle to a `Tracer`. 

By convention, Tracers are named after the component they are instrumenting; usually a library, a package, or a class.

```python
from opentelemetry import trace

tracer = trace.get_tracer(__name__)
```

> Note that there is no need to "set" a tracer by name before getting it. The name you provide is to help identify which component generated which spans, and to potentially disable tracing for individual components.

 We recommend calling `getTracer` once per component during initialization and retaining a handle to the tracer, rather than calling `getTracer` repeatedly.

### Accessing the current span
Ideally, when tracing application code, spans are created and managed in the application framework.

Assuming that your application framework is [supported](python/setup/auto-instrumentation), a trace will automatically be created for each request, and your application code will already be wrapped in a span, which can be used for adding application specific attributes and events.

To access the currently active span, call `getCurrentSpan`

```python
from opentelemetry import trace

span = trace.get_current_span()
```

### Setting a new current span
Let’s demonstrate creating a new span by example. Imagine you have an automated kitchen, and you want to time how long the robot chef takes to bake a cake. 
The naive way to do this would be to just start a span, call your method, then end the span:

```python
# make a child span to measure how long it takes to bake a cake.
span = trace.get_current_span()

cake_span = tracer.start_span(name="bake-cake")
chef.bake_cake()
cake_span.end()
```

The above example will work just fine, but with one big problem: the `bakeCake` method itself has no access to this new `bake-cake` span. That means there would be no way to add attributes and events to this span from within the bakeCake method. Even worse, `get_current_span` would return the parent of “bake-cake,” since that span is still set as current. 

What should we do instead? Replace the current span with `bake-cake.` To do this, call `withSpan` to make a closure around the `bakeCake` method. Within this closure, the `getCurrentSpan` method will now return `bake-cake`.

```python
from opentelemetry import trace

tracer = trace.get_tracer(__name__)

# Replace the current active span with a new child span
with tracer.start_as_current_span("bake-cake"):
  # now returns the bake-cake span
  trace.get_current_span()
  # bake your cake!
  chef.bake_cake()
```

This pattern of wrapping method calls is important, because we always want application code to be able to assume that the current span is correct.

When performing root cause analysis, span attributes are an important tool for pinpointing the source of performance issues. 

## Attributes in OpenTelemetry Python

### Setting Attributes

> Note that it is only possible to set attributes, not to get them.

Much like how resources are used to describe your services, attributes are used to describe your spans. Here is an example of setting attributes to correctly define an HTTP client request:

```python
from opentelemetry.trace import SpanKind

span = tracer.start_span(
  "/project/:project-id/list",
  kind=SpanKind.CLIENT,
  attributes={
    "http.method": "GET",
    "http.flavor": "1.1",
    "http.url": "https://example.com:8080/project/123/list/?page=2",
    "net.peer.ip": "192.0.2.5",
    "http.status_code": 200,
    "http.status_text": "OK"
  },
)

# In addition to the standard attributes, custom attributes can be added as well.
span.set_attribute("list.page_number", 2);

# To avoid collisions, always namespace your attribute keys using dot notation.
span.set_attribute("project.id", 2);

# attributes can be added to a span at any time before the span is finished.
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

## Events in OpenTelemetry Python

The finest-grained tracing tool is the event system. 

Span events are a form of structured logging. Each event has a name, a timestamp, and a set of attributes. When events are added to a span, they inherit the span's context. This additional context allows events to be searched, filtered, and grouped by trace ID and other span attributes. 

> Span context is one of the key differences between distributed tracing and traditional logging.

### Adding events
Events are automatically timestamped when they are added to a span. Timestamps can also be set manually if the events are being added after the fact.

For example, enqueuing an item might be recorded as an event.

```python
from opentelemetry import trace

# Get the current span
span = trace.get_current_span()

# Perform the action
queue.enqueue(item)

# Record the action
span.add_event( "enqueued item", {
  "item.id": item.id(),
	"queue.id": queue.id(),
	"queue.length": queue.Length(),
})
```

> Spans should be created for recording course-grained operations, and events should be created for recording fine-grained operations.

### Recording exceptions

Many of the tracing conventions can apply to event attributes as well as span attributes. The most important event-specific convention is [recording exceptions](https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/trace/semantic_conventions/exceptions.md).

```python
from opentelemetry import trace
from opentelemetry.trace.status import Status, StatusCode

span = trace.get_current_span()

# recordException converts the exception into a span event. 
exception = Exception("exception!!!")
span.record_exception(exception)

# If the exception means the operation results in an 
# error state, you can also use it to update the span status.
span.set_status(Status(StatusCode.ERROR, "error happened"))
```

> Marking the span as an error is independent from recordings exceptions. To mark the entire span as an error, and have it count against error rates, set the SpanStatus to ERROR.

StatusCode definitions can be found in the [OpenTelemetry specification](https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/trace/api.md#statuscanonicalcode).
