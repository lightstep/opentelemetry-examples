# Monitor ArangoDB with OpenTelemetry Collector and Lightstep

This example shows how to monitor ArangoDB using OpenTelemetry Collector with receivers for Prometheus metrics. The collected metrics are sent to Lightstep and can also be logged for debugging purposes.

## Collector Configuration

The configuration file for the collector is `collector.yml`. This section will explain the relevant parts of the configuration file.

### Receivers

This example uses the Prometheus receiver to ingest the OpenMetrics format published by ArangoDB. The configuration includes the following settings:

- `job_name`: The name of the job, set to `arangodb`.
- `scrape_interval`: The interval at which metrics are scraped, set to `3s`.
- `metrics_path`: The path to the ArangoDB metrics endpoint, set to `/_admin/metrics/v2`.
- `scheme`: The URL scheme to use when connecting to ArangoDB, set to `http`.
- `tls_config`: The TLS configuration, with `insecure_skip_verify` set to `true`.
- `static_configs`: The static configurations for the target, set to `arangodb:8529`.

### Exporters

Two exporters are defined in this configuration:

1. `logging`: This exporter logs the collected metrics with a log level of `debug`. This can be useful for debugging purposes.
2. `otlp`: This exporter sends the metrics to Lightstep with the following settings:
   - `endpoint`: The Lightstep endpoint, set to `ingest.lightstep.com:443`.
   - `headers`: A header with the Lightstep access token, set to `${LS_ACCESS_TOKEN}`. This requires an environment variable to be set with the actual access token.

### Processors

A batch processor is used to process the metrics before exporting them.

### Service

The service is configured to use the defined receivers, processors, and exporters for the metrics pipeline.

## Running the Demo

To run the demo, you will need Docker Compose installed on your machine. You will also need to set the environment variable `LS_ACCESS_TOKEN` to your Lightstep access token.

1. Clone the `lightstep/collector` repository and navigate to the `arangodb` folder.
2. Run `docker-compose up -d` to start the ArangoDB service and the OpenTelemetry Collector.
3. To view the metrics collected by the OpenTelemetry Collector, check the logs using `docker-compose logs collector`.
4. To view the metrics in Lightstep, navigate to the Metrics dashboard in your Lightstep account.

You can choose a particular version of ArangoDB by setting APP_VERSION. For example:

    ```
    export APP_VERSION=3.10 # ArangoDB released October 4th, 2022
    ```

## Teardown

To stop and remove the services, run `docker-compose down` in the `arangodb` folder.
