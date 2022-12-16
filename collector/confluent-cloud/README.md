---
# Ingest metrics from Confluent Cloud

The OpenTelemetry Collector has a variety of [third party receivers](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/master/receiver) that provide integration with a wide variety of metric sources.

Please note that not all metrics receivers available for the OpenTelemetry Collector have been tested by Lightstep Observability, and there may be bugs or unexpected issues in using these contributed receivers with Lightstep Observability metrics. File any issues with the appropriate OpenTelemetry community.
{: .callout}

## Requirements

* OpenTelemetry Collector Contrib v0.61.0+

## Prerequisites

You must have a Lightstep Observability [access token](/docs/create-and-manage-access-tokens) for the project to report metrics to.

## Running the Example

You can run this example with `docker-compose up` in this directory. 

## Configuration

Installation of the OpenTelemetry Collector varies, please refer to the [collector documentation](https://opentelemetry.io/docs/collector/) for more information.

The example configuration, used for this project shows using processors to add metrics with Lightstep Observability, add the following to your collector's configuration file.

The following environment variables are required:

#### [Cloud API](https://confluent.cloud/settings/api-keys) key
* CONFLUENT_API_ID
* CONFLUENT_API_SECRET

#### Kafka Client API key
* CLUSTER_API_KEY
* CLUSTER_API_SECRET

#### Lightstep Access Tokenj
* LS_ACCESS_TOKEN

#### Confluent Cloud Cluster Details
* CLUSTER_ID
* CLUSTER_BOOTSTRAP_SERVER

Below are some redacted example values for connecting to a Confluent Cloud cluster:

```
  # Available from https://confluent.cloud/settings/api-keys
  export CONFLUENT_API_ID=M4...J47
  export CONFLUENT_API_SECRET=h34W...mL

  # Available from Cluster Overview > API Keys
  export CLUSTER_API_KEY=BZ..I
  export CLUSTER_API_SECRET=/bff+2JJKEy..2Z

  # Available from the Cluster Settings page in Confluent Cloud
  export CLUSTER_BOOTSTRAP_SERVER=pkc-8ozv2.us-west4.gcp.confluent.cloud:9092
  export CLUSTER_ID=lkc-4ddff62
```

``` yaml
# add the receiver configuration for your integration
receivers:
  kafkametrics:
    brokers:
      - "${CLUSTER_BOOTSTRAP_SERVER}"
    protocol_version: 2.0.0
    scrapers:
      - brokers
      - topics
      - consumers
    auth:
      sasl:
        username: "${CLUSTER_API_KEY}"
        password: "${CLUSTER_API_SECRET}"
        mechanism: PLAIN
      tls:
        insecure_skip_verify: false
    collection_interval: 30s


  prometheus:
    config:
      scrape_configs:
        - job_name: "confluent"
          scrape_interval: 60s # Do not go any lower than this or you'll hit rate limits
          static_configs:
            - targets: ["api.telemetry.confluent.cloud"]
          scheme: https
          basic_auth:
            username: "${CONFLUENT_API_ID}"
            password: "${CONFLUENT_API_SECRET}"
          metrics_path: /v2/metrics/cloud/export
          params:
            "resource.kafka.id":
              - ${CLUSTER_ID}

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
    metrics/confluence:
      receivers: [prometheus, kafkametrics]
      processors: [batch]
      exporters: [logging, otlp/public]
```
