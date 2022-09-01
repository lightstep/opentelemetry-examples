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

## Run Telegraph

If you have Telegraph installed locally then with this configuration you can simply run the command in this example directory using the flag to indicate where to get the config file:

```bash
telegraph --config telegraf/telegraf.conf
```

Or with Docker use `docker run` as follows:

```bash
docker run --rm -v $(pwd)/telegraf:/telegraf -e LS_ACCESS_TOKEN={$LS_ACCESS_TOKEN} telegraf --config telegraf/telegraf.conf
```

## View the Results in Lightstep

After running Telegraf in this example we find in the Lightstep app the following metrics `airSensors_co`, `airSensors_humidity`, `airSensors_temperature`.

