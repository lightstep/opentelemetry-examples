# Monit Metrics to Lightstep via Telegraf

## Setup Monit

Monit isn't typically run in containers, so we'll walk through running monit and Telegraf without containers.

Both tools have packages for all major platforms and distros, along with convenient binary downloaders.

## Configure Telegraf

### Configure Telegraf: Net Response

`monit start all` and `monit monitor all` will have monit gathering many metrics.

You can see them in the monit web UI on the host at `localhost:2812`. 

In this example we'll configure the collector as if monit is listening on the default port and you haven't changed the default credentials.

```
[[inputs.monit]]
  address = "http://localost:2812"
  username = "admin"
  password = "monit"
```

### Configure Telegraf: Output OTLP

You can use Telegraf's OpenTelemetry output plugin to send OTLP over gRPC to Lightstep with configuration similar to this.

```
[[outputs.opentelemetry]]
  service_address = "ingest.lightstep.com:443"
  insecure_skip_verify = true

  [outputs.opentelemetry.headers]
    lightstep-access-token = "$LS_ACCESS_TOKEN"
```

## View the Results in Lightstep

For the example that we configured metrics should be appearing in your Lightstep project with the name `net_response_response_time` and `net_response_result_code`.

