# MicroDonuts: An OpenTracing + OpenTelemetry Shim with Lighstep Launcher

Welcome to MicroDonuts! This is a sample application instrumented with OpenTracing.
It uses OpenTelemetry Shim with Lighstep Launcher.

## Step 0: Setup MicroDonuts

### Getting it
Build the jar file (for this, Maven must be installed):

```
mvn package
```

### Running

MicroDonuts has two server components, `API` and `Kitchen`, which
communicate each other over HTTP - they are, however, part of
the same process:

```
mvn package exec:exec
```

#### Accessing

In your web browser, navigate to http://127.0.0.1:10001 and order yourself some
µ-donuts.


#### Lightstep Configuration

If you have access to [LightStep](https://app.lightstep.com]), you will need your access token. 