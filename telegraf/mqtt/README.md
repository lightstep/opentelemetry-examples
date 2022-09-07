# MQTT to Lightstep via Telegraf

## Configure Mosquitto

We can use a minimal configuration like the following to get started with tests using Mosquitto.

```
persistence true

listener 1883
allow_anonymous true

log_dest file /mosquitto/log/mosquitto.log
log_dest stdout
```

## Configure Telegraf

### Configure Telegraf: Consume MQTT 

In our Telegraf configuration we configure the input to reference the host at the default port for Mosquitto.

```
[[inputs.mqtt_consumer]]
  servers = [
    "tcp://broker:1883"
  ]
  topics = [
    "test/topic"
  ]
  data_format = "json"
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

## Simulate metrics messages

Given the configuration thus far we can simulate messages with `docker compose exec` using the `mosquitto_pub` cli.

The `-h` flag is the hostname where we sending the message. It will go to port 1883 by default. Then we provide a topic. We use test/topic which we also configured as one of the topics for Telegraf's MQTT consumer input plugin. And finally the payload which is an object or array since we configured Telegraf to use that data_format.

```
docker compose exec client mosquitto_pub -h broker -t test/topic -m '[{"key1": 9, "key2": 13}]'
```

## View the Results in Lightstep

For a message like what the demo illustrated in the last message we should find a metric in Lightstep named `mqtt_consumer_key1`. It will have keys for host which is the container id, topic which we set as "test/topic", and instrumentation.name.

