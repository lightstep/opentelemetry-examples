# Monitoring Elasticsearch with Lightstep

## About this Configuration

This is intended to be the simplest practical example for collecting Elasticsearch metrics with the OTEL Collector, sending data to Lightstep. While Elasticsearch is usually run in multinode clusters and often as part of the "ELK stack" with Logstash and Kibana, this example focuses on illustrating Collector configuration for one node for clarity.

## Running the Example

The file `.env` contains important configuration values which are also referenced in other files such as `collector.yml` and the docker compose configuration. Be sure to set the values here appropriately for your situation.

Make sure that you have set the environment variable `LS_ACCESS_TOKEN`. Then you can run `make up`. If you don't already have certificates then it will automatically run `make setup` for you to generate certificates and the Elastic keystore. See the Makefile for additional make rules.

## Additional Resources

You can find more details about configuration in OTEL Collector's [elasticsearchreceiver docs](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/elasticsearchreceiver). 

After you run `docker compose -f docker-compose.setup.yml`, you can run the example with `docker compose up`. For convenience there's also a make rule you can run with `make up`. This will detect whether setup has been run and will run it if needed.

## License

Some parts of this configuration are derived from the [elastdocker](https://github.com/sherifabdlnaby/elastdocker/) model configuration. It is incorporated and the work provided under a commercially permissive MIT license.

