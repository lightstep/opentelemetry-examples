---
# Ingest metrics using the External DNS integration

The OTel Collector has a variety of [third party receivers](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/master/receiver) that provide integration with a wide variety of metric sources.

Please note that not all metrics receivers available for the OpenTelemetry Collector have been tested by Lightstep Observability, and there may be bugs or unexpected issues in using these contributed receivers with Lightstep Observability metrics. File any issues with the appropriate OpenTelemetry community.
{: .callout}

## External DNS supports

ExternalDNS allows to configure external services, like [AWS Route53](https://github.com/kubernetes-sigs/external-dns/blob/master/docs/tutorials/aws.md), [Cloudflare DNS](https://github.com/kubernetes-sigs/external-dns/blob/master/docs/tutorials/cloudflare.md) and other ones. However for this example we provide configuration for CoreDNS.

## Prerequisites for local installation with CoreDNS

You must have a Lightstep Observability [access token](/docs/create-and-manage-access-tokens) for the project to report metrics to.

#### minikube

You can use any approach to managing your cluster, but the Makefile builds a cluster in `minikube`.

#### helm

We use helm charts to install some apps in this example.

## Running the Example

You can run this example with `make all` in this directory.
After tests just need to run `make delete-cluster`.

## Steps

### 1. Create a cluster

First you'll need to create a cluster by a method of your choice. `minikube start` works well for local development on Linux.

### 2. Required libraries

External DNS requires ETCD, DNS manager, for this example we picked CoreDNS. We add repos for coreDNS and otel-collector.

### 3. Installation

#### a. Install ETCD

ETCD is used to manage data inside the cluster, will be required for CoreDNS.

#### b. Install CoreDNS

Installation CoreDNS. It should be configured to point to the ETCD endpoint.

```sh
helm install my-coredns coredns/coredns -f values-coredns.yaml
```

#### c. Install ExternalDNS

ExternalDNS configs also has to be configured to point to the ETCD client.

#### d. Install Nginx ingress

For this example Nginx Ingress was chosen, but other ingress controllers may be used as well.

#### e. Collector installation

Installing the Collector is straightforward. Assuming you have already added the chart repo, you can install it with the helm chart like this.

```sh
helm install my-collector open-telemetry/opentelemetry-collector -f values-collector.yaml
```

## Configuration

Installation of the OpenTelemetry Collector varies, please refer to the [collector documentation](https://opentelemetry.io/docs/collector/) for more information.

Detailed description of available [External DNS metrics](https://github.com/kubernetes-sigs/external-dns/blob/e2b86a114612cb334145c5cff3876495b67b8988/docs/faq.md#what-metrics-can-i-get-from-externaldns-and-what-do-they-mean).

Collector Prometheus receiver has to be pointer to External DNS metrics endpoint, which is exposed by default at :7979 port.

The example configuration to add metrics with Lightstep Observability, add the following to your collector's configuration file:

``` yaml
# add the receiver configuration for your integration
receivers:
  prometheus:
    config:
      scrape_configs:
        - job_name: otel-external-dns
          static_configs:
            - targets: [external-dns:7979]

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
      receivers: [prometheus]
      processors: [batch]
      exporters: [logging, otlp/public]
```

