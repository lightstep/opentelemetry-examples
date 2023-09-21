---
# Ingest metrics using the Docker integration

The OTel Collector has a variety of [third party receivers](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/master/receiver) that provide integration with a wide variety of metric sources.

## Prerequisites for local installation

You must have a Lightstep Observability [access token](/docs/create-and-manage-access-tokens) for the project to report metrics to.

## Running the Example

Run the docker compose example.

```bash
docker-compose up
```

## Configuration

Installation of the OpenTelemetry Collector varies, please refer to the [collector documentation](https://opentelemetry.io/docs/collector/) for more information.

Detailed description of available [Docker metrics](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/receiver/dockerstatsreceiver/metadata.yaml#L45).

Collector Prometheus receiver has to be pointed to the Kubernetes Prometheus metrics endpoints.

> :warning: **This configuration is for illustration purposes only**: Mounting Docker socket in containers isn't recommended. It's used here for simplicity in illustration.

The following example configuration collects metrics from Kong and send them to Lightstep Observability:

```yaml
receivers:
  docker_stats:
    endpoint: "unix:///var/run/docker.sock"
    metrics:
      container.cpu.usage.percpu:
        enabled: true

exporters:
  logging:
    loglevel: debug
  otlp/public:
    endpoint: ingest.lightstep.com:443
    headers:
        "lightstep-access-token": "${LS_ACCESS_TOKEN}"

processors:
  batch:

service:
  pipelines:
    metrics:
      receivers: [docker_stats]
      processors: [batch]
      exporters: [logging, otlp/public]
  telemetry:
    logs:
      level: debug
```


