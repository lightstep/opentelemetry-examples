# OpenTelemetry Collector Integration for Memcached Example

This directory contains a docker-compose based example demonstrating the usage of the OpenTelemetry Collector Memcached Reciever and Lightstep.

## Run the example

1. Export your [Lightstep access token][ls-docs-access-token]:

    `export LIGHTSTEP_ACCESS_TOKEN=<YOUR_TOKEN>`

2. Run the example using docker-compose:

    `docker compose up`

3. [Create a dashboard][ls-docs-dashboards] to view the data at Lightstep Observability.


## Configuration

Below is a snippet taken from the `collector.yml` in this example that shows how to configure the OpenTelemetry Collector Memcached Receiver to scrape statistics from a Memcached instance.

```yaml
receivers:
  memcached:
    endpoint: "memcache:11211"
```

For more information about the Memcached Receiver and additional configuration options see [the documentation][otel-memcached-receiver] at the [Collector Contrib][otel-collector-contrib] repo.

[otel-collector-contrib]: https://github.com/open-telemetry/opentelemetry-collector-contrib
[otel-memcached-receiver]: https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/memcachedreceiver
[ls-docs-access-token]: https://docs.lightstep.com/docs/create-and-manage-access-tokens
[ls-docs-dashboards]: https://docs.lightstep.com/docs/create-and-manage-dashboards
