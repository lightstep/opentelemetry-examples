# Importing InfluxDB LP Format to Lightstep with Telegraf

## Export data

The official docs for the [influx_inspect export](https://docs.influxdata.com/influxdb/v1.8/tools/influx_inspect/#export) command explain how to extract your data from InfluxDB. 

For testing the migration to Lightstep I obtained Influx line protocol formatted data from the [InfluxDB sample data repo](https://github.com/influxdata/influxdb2-sample-data). 

In this walkthrough I'll use the [air-sensor-data](https://github.com/influxdata/influxdb2-sample-data/tree/master/air-sensor-data), that I fetched like this:

```bash
wget -q -O - https://raw.githubusercontent.com/influxdata/influxdb2-sample-data/master/air-sensor-data/air-sensor-data.lp > data/in/air-sensor-data.lp
```

## Configure Telegraf

### Configure Telegraf: Consume Influx Line Protocol 

In the Telegraf configuration you need to configure `directory_monitor` plugin. There are multiple input plugins that process files. I'm using this input plugin as opposed to others that work for the filesystem, because this will only process the file once.

```
[[inputs.directory_monitor]]
  files = "data/in"
  finished_directory = "data/done"
  data_format = "influx"
```

### Configure Telegraf: Output OTLP

You can use Telegraf's OpenTelemetry output plugin to send OTLP over gRPC to Lightstep with configuration similar to this.

```
[[outputs.opentelemetry]]
  service_address = "ingest.lightstep.com:443"
  insecure_skip_verify = true

  [outputs.opentelemetry.headers]
    lightstep-access-token = "$LS_ACCESS_TOKEN"
```
## Edit the Sample Data

This step is solely to make our dataset easier to see in Lightstep. Unlike the other steps, we won't do anything similar in our real workflows.

The sample data is likely to be older than anything that will show up in your Lightstep account. I used the six most significant digits of [Unix Timestamp](https://www.unixtimestamp.com/) and replaced the first 6 digits I found in timestamps of the sample data. For example, it's September 1, 2022 and the first 6 digits of the current Unix timestamp are 166204. The timestamps in the sample data are 166198, so I replaced 166198 with 166204. It's only worth the trouble if you need
to see how the data appears in Lightstep Observability.

## Process the Data to OTLP Conventions 

The rename plugin can change attributes. To make our data more like the conventions in OpenTelemetry we'll transform the name first to report this data in a particular namespace. 

```bash
[[processors.rename]]
  [[processors.rename.replace]]
    measurement = "airSensors"
    dest = "sensors.air"
```

That takes care of using `sensors.air` in place of the `airSensors` measurement name from the data, but the measurement name and fields are still separated by an `_` like `air.sensors_co` and we'd like a dot separation. For this we can use templates, which are a Telegraf mini-language that describe how to map a dot delimited string to and from outputs. 

```bash
separator = "."
templates = [
    "*.*.field"
]
```

## Run Telegraph

If you have Telegraph installed locally then with this configuration you can simply run the command in this example directory using the flag to indicate where to get the config file:

```bash
telegraph --config telegraf/telegraf.conf
```

Or with Docker use `docker run` as follows:

```bash
docker run --rm -v $(pwd)/telegraf:/telegraf -e LS_ACCESS_TOKEN={$LS_ACCESS_TOKEN} telegraf --config /telegraf/telegraf.conf 
```

## View the Results in Lightstep

After running Telegraf in this example we find in the Lightstep app the following metrics `airSensors_co`, `airSensors_humidity`, `airSensors_temperature`.

