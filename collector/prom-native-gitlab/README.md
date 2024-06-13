# Gitlab metrics using the OpenTelemetry Collector

## Overview

 Gitlab natively exposes a Prometheus endpoint and the OpenTelemetry Collector has a [Prometheus receiver][otel-prom-receiver] that can be used to scrape its Prometheus endpoint. This directory contains an example showing how to configure Gitlab and the Collector to send metrics to Lightstep Observability.

## Prerequisites

* Docker
* Docker Compose
* A Lightstep Observability [access token][ls-docs-access-token]

## How to run the example

* Export your Lightstep access token
  ```
  export LS_ACCESS_TOKEN=<YOUR_TOKEN>
  ```
* Run the docker compose example
  ```
  docker-compose up -d
  ```
* Stop the container
  ```
  docker-compose down
  ```

### Explore Metrics in Lightstep

See the [Gitlab monitoring Docs][https://docs.gitlab.com/ee/administration/monitoring/prometheus/] for comprehensive documentation on metrics emitted and the [dashboard documentation][ls-docs-dashboards] for more details.


## Configure Gitlab

- To enable the GitLab Prometheus endpoint to be scraped by the OpenTelemetry Prometheus receiver over HTTP instead of HTTPS, you can configure GitLab to expose the Prometheus endpoint on port 9090 over HTTP using the following settings in the `/etc/gitlab/gitlab.rb` file.

#### Enable prometheus on Gitlab and clients

* Run the docker compose first time
  ```
  docker-compose up -d
  ```
* Then, stop the container
  ```
  docker-compose down
  ```

Docker mount the gitlab path as 

```
volumes:
    - ./srv/gitlab/config:/etc/gitlab
    - ./srv/gitlab/logs:/var/log/gitlab
    - ./srv/gitlab/data:/var/opt/gitlab
```

Then remove all comments on this path, `./etc/gitlab/gitlab.rb` and add the below config.

`./etc/gitlab/gitlab.rb`
```bash
# Enable HTTP for the Prometheus endpoint
prometheus['enable'] = true
prometheus['listen_address'] = 'localhost:9090'
prometheus['ssl_enabled'] = false

# Allow Prometheus to access the metrics endpoint
gitlab_workhorse['prometheus_listen_addr'] = "localhost:9229"
```

## Configure the Collector

Below is a snippet showing how to configure the Prometheus Receiver to scrape the Prometheus endpoint exposed by the Gitlab Server.

```yaml
receivers:
  prometheus:
    config:
      scrape_configs:
        - job_name: 'gitlab'
          scrape_interval: 10s
          metrics_path: "/"
          scheme: http
          static_configs:
            - targets: ['gitlab:9090']

```



## Additional information

- [OpenTelemetry Collector Prometheus Receiver][otel-prom-receiver]
- [Gitlab Prometheus Reference][gitlab-docs-prometheus]

[ls-docs-access-token]: https://docs.lightstep.com/docs/create-and-manage-access-tokens
[ls-docs-dashboards]: https://docs.lightstep.com/docs/create-and-manage-dashboards
[otel-prom-receiver]: https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/prometheusreceiver
[gitlab-docs-prometheus]: https://docs.gitlab.com/ee/administration/monitoring/prometheus/
[learn-consul-repo]: https://github.com/hashicorp/learn-consul-docker