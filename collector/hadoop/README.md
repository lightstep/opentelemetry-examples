
---
# Ingest metrics using the Hadoop integration

The OTEL Collector has a variety of [third party receivers](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/master/receiver) that provide integration with a wide variety of metric sources.

Please note that not all metrics receivers available for the OpenTelemetry Collector have been tested by Lightstep Observability, and there may be bugs or unexpected issues in using these contributed receivers with Lightstep Observability metrics. File any issues with the appropriate OpenTelemetry community.
{: .callout}

## Requirements

* OpenTelemetry Collector Contrib v0.53.0+

## Prerequisites

You must have a Lightstep Observability [access token](/docs/create-and-manage-access-tokens) for the project to report metrics to.

## To enable JMX in Hadoop

* [hadoop-env.sh](/collector/hadoop/conf/hadoop-env.sh)
```sh
export HDFS_NAMENODE_OPTS="-Dcom.sun.management.jmxremote=true -Dcom.sun.management.jmxremote.authenticate=false -Dcom.sun.management.jmxremote.ssl=false -Dcom.sun.management.jmxremote.port=8004 $HDFS_NAMENODE_OPTS"
export HDFS_DATANODE_OPTS="-Dcom.sun.management.jmxremote=true -Dcom.sun.management.jmxremote.authenticate=false -Dcom.sun.management.jmxremote.ssl=false -Dcom.sun.management.jmxremote.port=8006 $HDFS_DATANODE_OPTS"
```

* [yarn-env.sh](/collector/hadoop/conf/yarn-env.sh)
```sh
export YARN_RESOURCEMANAGER_OPTS="-Dcom.sun.management.jmxremote=true -Dcom.sun.management.jmxremote.authenticate=false -Dcom.sun.management.jmxremote.ssl=false -Dcom.sun.management.jmxremote.port=8002 $YARN_RESOURCEMANAGER_OPTS"
export YARN_NODEMANAGER_OPTS="-Dcom.sun.management.jmxremote=true -Dcom.sun.management.jmxremote.authenticate=false -Dcom.sun.management.jmxremote.ssl=false -Dcom.sun.management.jmxremote.port=8002 $YARN_NODEMANAGER_OPTS"
```


## Running the Example

You can run this example with `docker-compose up` in this directory. You'll want to view this in Lightstep with a dashboard. 

#### Hadoop Docker

### Quick Start

To deploy an example HDFS cluster, run to pull and start:
``` sh
  docker compose up
```

To stop and remove, run:
``` sh
  docker compose down
```

Run example wordcount job:
``` sh
  make wordcount
```

To clean/remove HDFS container and wordcount job:
``` sh
  make remove_all_images
```

`docker-compose` creates a docker network that can be found by running `docker network list`, e.g. `hadoop_integrations`.

Run `docker network inspect` on the network (e.g. `hadoop_integrations`) to find the IP the hadoop interfaces are published on.

## Configure Environment Variables

The configuration parameters can be specified in the hadoop.env file or as environmental variables for specific services (e.g. namenode, datanode etc.):
```
  CORE_CONF_fs_defaultFS=hdfs://namenode:8020
```

CORE_CONF corresponds to core-site.xml. fs_defaultFS=hdfs://namenode:8020 will be transformed into:
```
  <property><name>fs.defaultFS</name><value>hdfs://namenode:8020</value></property>
```
To define dash inside a configuration parameter, use triple underscore, such as YARN_CONF_yarn_log___aggregation___enable=true (yarn-site.xml):
```
  <property><name>yarn.log-aggregation-enable</name><value>true</value></property>
```

The available configurations are:
* /etc/hadoop/core-site.xml CORE_CONF
* /etc/hadoop/hdfs-site.xml HDFS_CONF
* /etc/hadoop/yarn-site.xml YARN_CONF
* /etc/hadoop/httpfs-site.xml HTTPFS_CONF
* /etc/hadoop/kms-site.xml KMS_CONF
* /etc/hadoop/mapred-site.xml  MAPRED_CONF

If you need to extend some other configuration file, refer to base/entrypoint.sh bash script.


## Configuration

Installation of the OpenTelemetry Collector varies, please refer to the [collector documentation](https://opentelemetry.io/docs/collector/) for more information.

The example configuration, used for this project shows using processors to add metrics with Lightstep Observability, add the following to your collector's configuration file:

``` yaml
# add the receiver configuration for your integration
receivers:
  jmx/hadoop:
    jar_path: /opt/opentelemetry-jmx-metrics.jar
    endpoint: namenode:8004
    target_system: jvm,hadoop

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
     receivers: [jmx/hadoop]
     processors: [batch]
     exporters: [logging,otlp]  
```


