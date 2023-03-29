# Monitoring CouchDB with the OpenTelemetry Collector and Lightstep

This demo shows an example of how to use OpenTelemetry Collector to collect metrics from CouchDB and send them to Lightstep for monitoring.

## Prerequisites
- Docker
- Docker Compose

## Getting started
1. Clone this repository and navigate to the `couchdb` directory:
   ```
   git clone https://github.com/lightstep/opentelemetry-examples.git
   cd opentelemetry-examples/collector/couchdb
   ```

2. Export your access token in the environment:
   You can get an access token for your project at app.lightstep.com.
   ```
   export LS_ACCESS_TOKEN=<your-lightstep-access-token>
   ```

3. Start the demo environment:
   To run a CouchDB cluster you will run...
   ```
   docker-compose -f docker-compose.cluster.yml up -d
   ```
   To run a single instance example run...
   ```
   docker-compose up -d
   ```

4. Wait for a few minutes for the metrics to be collected and sent to Lightstep.

5. Open the Lightstep web app and navigate to the Service Dashboard for the CouchDB service. 
You should see metrics such as request rate, response time, and throughput.

7. To stop the demo, run:
   ```
   docker-compose down
   ```

## Configuration
The `collector.yml` file configures the OpenTelemetry Collector to collect metrics from CouchDB and export them to Lightstep. The relevant parts of the configuration file are:

- `receivers`: The `couchdb` receiver is configured to scrape metrics from the CouchDB instance running at `http://couchdb:5984`. The `username` and `password` fields are used to authenticate to the CouchDB instance. The `collection_interval` field specifies how often the metrics are collected.

- `exporters`: The `otlp` exporter is used to send the collected metrics to Lightstep. The `endpoint` field specifies the Lightstep endpoint to send the metrics to, and the `headers` field is used to authenticate to Lightstep using your access token.

- `processors`: The `batch` processor is used to batch the collected metrics before exporting them.

- `service`: The `metrics` pipeline is used to collect CouchDB metrics using the `couchdb` receiver, process them using the `batch` processor, and export them using the `otlp` exporter.

## Environment variables
There are no environment variables required for this demo.

## Troubleshooting
- If you are not seeing any metrics in Lightstep, check that the CouchDB instance is running and accessible at `http://couchdb:5984`. Also, check that the OpenTelemetry Collector is running and configured correctly.

- If you are still having issues, refer to the OpenTelemetry and Lightstep documentation for more information.
