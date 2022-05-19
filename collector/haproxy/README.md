# Ingest HAProxy metrics using OTel Collector's Prometheus receiver

The OpenTelemetry Collector has a [Prometheus receiver][otel-prom-receiver] and HAProxy exposes metrics via a Prometheus compatible endpoint. This example shows how to configure the Collector and HAProxy to export metrics to Lightstep Observability.

## Requirements

* OpenTelemetry Collector Contrib v0.51.0+

## Prerequisites

You must have a Lightstep Observability [access token](/docs/create-and-manage-access-tokens) for the project to report metrics to.

## Running the Example

This example assumes you have exported your access token as `LS_ACCESS_TOKEN`. You can run this example with `docker-compose up` in this directory. HAProxy is mapped to port 8080 on the host machine and requests to `http://localhost:8080/*`, where `*` can be any path, will be routed to one of three echo server instances.

If you would like to generate load automatically, run this example using `docker-compose --profile loadgen up`.

### Charting the data

You can see the metrics emitted by HAProxy by inspecting its Prometheus end point at: http://localhost:8404/metrics and [build dashboards][ls-docs-dashboards] using these metrics in Lightstep.

## Configuration

You will need a build of the OpenTelemery collector that includes the Prometheus receiver. This example uses the [collector-contrib][docker-collector-contrib] images published to dockerhub. For other ways to install and run the collector see [collector documentation](https://opentelemetry.io/docs/collector/) for more information.

### Collector Configuration

Below is a snippet showing how to configure the Prometheus receiver to scrape the HAProxy Prometheus endpoint.

```yaml
receivers:
  prometheus:
      config:
        scrape_configs:
          - job_name: 'haproxy'
            scrape_interval: 10s
            static_configs:
              - targets: ['haproxy:8404']
```

### HAProxy Configuration

HAProxy exposes metrics via a Prometheus compatible endpoint that needs to be configured in `haproxy.cfg`. See the snippet below for an example, the key line is `http-request use-service prometheus-exporter if { path /metrics }`. For more information on the `http-request use-service` directive, see the documentation [here][haproxy-use-service-docs]. For more information on HAProxy's Prometheus endpoint see [this blog post][haproxy-prom-blog].


```
frontend stats
  bind *:8404
  http-request use-service prometheus-exporter if { path /metrics }
  stats enable
  stats uri /
  stats refresh 10s
```

[otel-prom-receiver]: https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/prometheusreceiver
[ls-docs-dashboards]: https://docs.lightstep.com/docs/create-and-manage-dashboards
[docker-collector-contrib]: https://hub.docker.com/r/otel/opentelemetry-collector-contrib
[haproxy-use-service-docs]: https://www.haproxy.com/documentation/hapee/latest/onepage/#4.2-http-request%20use-service
[haproxy-prom-blog]: https://www.haproxy.com/blog/haproxy-exposes-a-prometheus-metrics-endpoint/