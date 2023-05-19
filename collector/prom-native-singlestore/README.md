# Ingest SingleStoreDB (MemSQL) metrics using the OpenTelemetry Collector

## Overview

SingleStore is a distributed SQL database that provides a Prometheus endpoint for collecting metrics data. The OpenTelemetry Collector's Prometheus receiver can scrape this data to monitor and analyze SingleStore's performance and usage, such as query latency, throughput, and resource utilization. Sending these metrics to Lightstep Observability enables you to correlate them with traces and logs for deeper insights into your application behavior and performance. This comprehensive view allows you to quickly identify and resolve issues before they become major problems, making it an essential part of a modern DevOps toolkit.

## Prerequisites

* Docker
* Docker Compose
* A Lightstep Observability [access token][ls-docs-access-token]

## How to run the example

* Obtain a `LICENSE_KEY` by [singlestore](https://www.singlestore.com) and export it
  ```
  export LICENSE_KEY=<YOUR_LICENSE>
  ```
* Export your Lightstep access token
  ```
  export LS_ACCESS_TOKEN=<YOUR_TOKEN>
  ```
* Run the docker compose example
  ```
  docker-compose up -d
  ```
* Stop the cluster
  ```
  docker-compose down`
  ```


### Explore Metrics in Lightstep

See the [SingleStore Telemetry Docs][single-store-docs-telemetry] for comprehensive documentation on metrics emitted and the [dashboard documentation][ls-docs-dashboards] for more details.

### Explore the SingleStore Example

* Access SingleStore UI [http://localhost:8080](http://localhost:8080).


## Configure SingleStore

- Prometheus metrics are not enabled by default; Look at init.sql to find out how to enable exporter.


```mysql
/* The memsql-exporter process (or simply “the exporter”) collects data about a running cluster. The user that starts
   the exporter (other than the SingleStore DB root user) must have the following permissions at a minimum:
*/

#GRANT CLUSTER on *.* to <user>
#GRANT SHOW METADATA on *.* to <user>
  #GRANT SELECT on *.* to <user>

/* HTTP */
SET GLOBAL exporter_user = root;
SET GLOBAL exporter_password = 'password_here';
SET GLOBAL exporter_port= 9104;

/* HTTPS */
#SET GLOBAL exporter_user = root;
#SET GLOBAL exporter_password = '<secure-password>';
#SET GLOBAL exporter_use_https= true;
#SET GLOBAL exporter_ssl_cert = '/path/to/server-cert.pem';
#SET GLOBAL exporter_ssl_key = '/path/to/server-key.pem';
#SET GLOBAL exporter_ssl_key_passphrase= '<passphrase>';

/* Use an engine variable to stop the exporter process by setting the port to 0. */
# SET GLOBAL exporter_port = 0;
```

## Configure the Collector

Below is a snippet showing how to configure the Prometheus Receiver to scrape the Prometheus endpoint exposed by the Hashicorp Vault Server.

```yaml
receivers:
  prometheus:
    config:
      scrape_configs:
        - job_name: 'node'
          scrape_interval: 10s
          metrics_path: '/metrics'
          params:
            format: ['prometheus']
          static_configs:
            - targets: ['singlestore:9104']
        - job_name: 'cluster'
          scrape_interval: 10s
          metrics_path: '/cluster-metrics'
          params:
            format: ['prometheus']
          static_configs:
            - targets: ['singlestore:9104']

```



## Additional information

- [OpenTelemetry Collector Prometheus Receiver][otel-prom-receiver]
- [SingleStore Telemetry Reference][single-store-docs-telemetry]

[ls-docs-access-token]: https://docs.lightstep.com/docs/create-and-manage-access-tokens
[ls-docs-dashboards]: https://docs.lightstep.com/docs/create-and-manage-dashboards
[otel-prom-receiver]: https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/prometheusreceiver
[single-store-docs-telemetry]: https://docs.singlestore.com/db/v8.1/en/user-and-cluster-administration/cluster-health-and-performance/configure-monitoring.html
