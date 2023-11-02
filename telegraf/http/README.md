---
# Monitoring HTTP Services in Cloud Observability with Telegraf

## Setup an HTTP Server to Test

There's a fairly minimal Go metrics service in the `app` directory of this example.

It publishes its metrics on an HTTP endpoint in JSON.

## Configure Telegraf

### Configure Telegraf: HTTP Input

In our Telegraf configuration we configure the input to reference the host at the default port for Mosquitto.

```
[[inputs.http]]
  urls = [
    "http://demosvc:8080/heapbasics"
  ]
  timeout = "10s"
  data_format = "json"
  [inputs.http.json_v2]
    timestamp_key = "time"
    timestamp_format = "unix"
```

Telegraf will get the metrics from the endpoint on a configurable scrape interval. 

### Configure Telegraf: Output OTLP

You can use Telegraf's OpenTelemetry output plugin to send OTLP over gRPC to Cloud Observability with configuration similar to this.

```
[[outputs.opentelemetry]]
  service_address = "ingest.lightstep.com:443"
  insecure_skip_verify = true

  [outputs.opentelemetry.headers]
    lightstep-access-token = "$LS_ACCESS_TOKEN"
```

## View the Results in Cloud Observability

For the example that we configured metrics should appear in your Cloud Observability project with the name `http_value`.

