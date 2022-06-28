# Ingest HAProxy metrics using OTel Collector's Prometheus receiver

HAProxy emits metrics via a Prometheus compatible endpoint and [OpenTelemetry Collector Contrib][otel-collector-contrib] has a [Prometheus receiver][otel-prom-receiver] that can be used to scrape those metrics. The examples in this repo show how to configure HAProxy and the Collector to send metrics to Lightstep Observability.

## Requirements

* HAProxy v2.0+
* OpenTelemetry Collector Contrib v0.51.0+

## Prerequisites

You must have a Lightstep Observability [access token][ls-docs-access-token] for the project to report metrics to.

## Running the Example

**Set LS_ACCESS_TOKEN as an environment variable**

The `docker-compose.yml` assumes your access token has been set as an environment variable named `LS_ACCESS_TOKEN`. Set the environment variable using the method of your choosing, for example:

```
export LS_ACCESS_TOKEN=<YOUR-TOKEN>
```

**Example wihout load generation**

```
docker-compose up
```

Note: This example sets up HAProxy to listen on port 8080 and it will forward requests to any path to a pool of echo servers. You can drive load by making requests such as `curl http://localhost:8080/foo`, `curl http://localhost:8080/bar`, `curl http://localhost:8080/any/path`.

**Example with load generation**

You can run the example from the previous step with load generation using the following command. Load is generated using [wrk](https://github.com/wg/wrk).

```
docker-compose --profile loadgen up
```

### Charting the data

You can see the metrics emitted by HAProxy by inspecting its Prometheus end point at: http://localhost:8404/metrics and [build dashboards][ls-docs-dashboards] using these metrics in Lightstep.

## Configuration

Below you will find relevant snippets of configuration needed to configure the HAProxy Prometheus endpoint and how to configure the Colllector's Prometheus receiver to scrape it. Look at the [docker-compose.yml](docker-compose.yml) and [haproxy.cfg](haproxy.cfg) files in this directory for more details.

### HAProxy Configuration

HAProxy has native support for Prometheus, but it must be enabled in your HAProxy configuration file. The line `http-request use-service rometheus-exporter if { path /metrics }` directive needs to be added to your existing `frontend stats` stats section. See below for an example:

~~~
frontend stats
  bind *:8404
  http-request use-service prometheus-exporter if { path /metrics }
  stats enable
  stats uri /
  stats refresh 10s
~~~

See the official HAProxy [blog post][haproxy-prom-blog] for more details about its Prometheus endpoint.

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

The OpenTelemetry repo provides additional details and options for [Prometheus receiver configuration][otel-prom-receiver].

[otel-collector-contrib]: https://github.com/open-telemetry/opentelemetry-collector-contrib
[otel-prom-receiver]: https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/prometheusreceiver
[ls-docs-access-token]: https://docs.lightstep.com/docs/create-and-manage-access-tokens
[ls-docs-dashboards]: https://docs.lightstep.com/docs/create-and-manage-dashboards
[docker-collector-contrib]: https://hub.docker.com/r/otel/opentelemetry-collector-contrib
[haproxy-prom-blog]: https://www.haproxy.com/blog/haproxy-exposes-a-prometheus-metrics-endpoint/