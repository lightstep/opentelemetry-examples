# Run synthethic checks OTel Collector's HTTP Check receiver

The OpenTelemetry Collector [HTTP Check receiver](httpcheckreceiver) connects to a configured endpoint via HTTP to validate that the endpoint is responding with a status code 200. The examples in this repo show how to configure an HTTP endpoint and the Collector to send metrics to Lightstep Observability.

## Requirements

* OpenTelemetry Collector Contrib v0.61.0+
* Docker Compose

## Prerequisites

You must have a Lightstep Observability [access token][ls-docs-access-token] for the project to report metrics to.

## Running the Example

**Set LS_ACCESS_TOKEN as an environment variable**

The `docker-compose.yml` assumes your access token has been set as an environment variable named `LS_ACCESS_TOKEN`. Set the environment variable using the method of your choosing, for example:

```
export LS_ACCESS_TOKEN=<YOUR-TOKEN>
```

Run Docker compose

```
docker-compose up
```

[httpcheckreceiver]: https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/httpcheckreceiver
[ls-docs-access-token]: https://docs.lightstep.com/docs/create-and-manage-access-tokens
[ls-docs-dashboards]: https://docs.lightstep.com/docs/create-and-manage-dashboards