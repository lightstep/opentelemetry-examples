# Monitoring Elasticsearch

## About this Configuration

Elasticsearch is frequently run as part of the ELK stack with Logstash and Kibana. For clarity we limited this example to a single Elasticsearch node to cover just the information that you need to integrate with Lightstep.

## Setup

Make sure that you have set the environment variable `LS_ACCESS_TOKEN`. Then to prepare the environment you can run `make setup` which will create certificates. However, it isn't necessary to run `make setup` before running `make up`, because the `up` rule will check for the directory called `secrets` and generate certificates if it doesn't exist.

The file `.env` contains key configuration values which are also referenced in other files such as `collector.yml` and the docker compose configuration. Be sure to set the values here appropriately for your situation.

Other useful `make` rules are provided such as `down` and `prune` for cleaning up the Docker host after running the example. Consult the Makefile for more details.

## Running this Example

After you run `docker compose -f docker-compose.setup.yml`, you can run the example with `docker compose up`. For convenience there's also a make rule you can run with `make up`. This will detect whether setup has been run and will run it if needed.

Edit the .env file to adjust variables for your configuration.

## About the Configuration

The base `docker-commpose.yml` file includes the Elasticsearch node. The file `docker-compose.override.yml` includes the OTEL Collector. And `docker-compose.setup.yml` includes services that setup the requisite Elastic keystore and certificates.

Note that the file receiver of the OTEL receiver is also configured for this example to simplify inspection of the output.

## License

Some parts of this configuration are derived from the [elastdocker](https://github.com/sherifabdlnaby/elastdocker/) model configuration. It is used and provided under a commercially permissive MIT license.
