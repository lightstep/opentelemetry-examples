---
# Ingest metrics using the Solr integration

The OTEL Collector has a variety of [third party receivers](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/master/receiver) that provide integration with a wide variety of metric sources.

Please note that not all metrics receivers available for the OpenTelemetry Collector have been tested by Lightstep Observability, and there may be bugs or unexpected issues in using these contributed receivers with Lightstep Observability metrics. File any issues with the appropriate OpenTelemetry community.
{: .callout}

## Requirements

* OpenTelemetry Collector Contrib v0.53.0+

## Prerequisites

You must have a Lightstep Observability [access token](/docs/create-and-manage-access-tokens) for the project to report metrics to.

## OpenSSL command to generate your private key and public certificate
```sh
openssl req -newkey rsa:2048 -nodes -keyout key.pem -x509 -days 365 -out certificate.pem
```

## Review the created certificate
```sh
openssl x509 -text -noout -in certificate.pem
```

## Running the Example

You can run this example with `docker-compose up` in this directory.

### Add Document Data to Solr Core
```sh
cd nodeapp && npm i && node index.js
```

## Configuration

Installation of the OpenTelemetry Collector varies, please refer to the [collector documentation](https://opentelemetry.io/docs/collector/) for more information.

The example configuration, used for this project shows using processors to add metrics with Lightstep Observability, add the following to your collector's configuration file:

``` yaml
# add the receiver configuration for your integration
receivers:
  jmx/solr:
    jar_path: /opt/opentelemetry-jmx-metrics.jar
    endpoint: solr:8983
    target_system: jvm,solr

processors:
  batch:

exporters:
  logging:
    loglevel: debug
  otlp:
    endpoint: ingest.lightstep.com:443
    headers: 
      "lightstep-access-token": "${LS_ACCESS_TOKEN}"

service:
  telemetry:
    logs:
      level: "debug"
  pipelines:
    metrics:
     receivers: [jmx/solr]
     processors: [batch]
     exporters: [logging,otlp]  

```
