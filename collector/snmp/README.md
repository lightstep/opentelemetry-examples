---
# Ingest metrics using the SNMP integration

The OTel Collector has a variety of [third party receivers](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/master/receiver) that provide integration with a wide variety of metric sources.

{: .callout}

## Prerequisites

You must have a Lightstep Observability [access token](/docs/create-and-manage-access-tokens) for the project to report metrics to.

## Running the Example

You can run this example with `docker-compose up` in this directory.

## Configuration

Installation of the OpenTelemetry Collector varies, please refer to the [collector documentation](https://opentelemetry.io/docs/collector/) for more information.

SNMP allows to access diffirent metrics by using unique OIDs (Object Identifiers) and sysObjectIDs(System Object Identifiers).

The example collector's configuration, used for this project shows using processors to add metrics with Lightstep Observability:

``` yaml
receivers:
  snmp:
    collection_interval: 60s
    endpoint: udp://snmpd:161
    version: v3
    security_level: auth_priv
    user: collector_user
    auth_type: "MD5"
    auth_password: password
    privacy_type: "DES"
    privacy_password: priv_password

    resource_attributes:
      resource_attr.name.1:
        indexed_value_prefix: probe

    metrics:
      snmp_cpu_user:
        unit: "By"
        gauge:
          value_type: int
        column_oids:
          - oid: "1.3.6.1.4.1.2021.11.9"
            resource_attributes:
              - resource_attr.name.1

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
    metrics/snmp:
      receivers: [snmp]
      processors: [batch]
      exporters: [logging, otlp/public]

```
