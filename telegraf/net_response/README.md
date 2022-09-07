# Net Response to Lightstep via Telegraf

## Setup an Response Service to Test

There's a fairly minimal Go metrics service in the `app` directory of this example.

It listens for "udp" messages and sometimes responds with "up", other timesf with "down".

## Configure Telegraf

### Configure Telegraf: Net Response

We configure the Net Response input plugin to send a particular string to our service.

The options will are slightly different depending on what protocol we use.

```
[[inputs.net_response]]
  protocol = "udp"
  address = "demosvc:9876"
  send = "yolo"
  expect = "up"
  timeout = "2s"
```

Telegraf will send the message and record the measurements depending on the response that it receives.

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

