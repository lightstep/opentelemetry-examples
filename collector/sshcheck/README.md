# Run synthethic checks OTel Collector's SSH Check receiver

The OpenTelemetry Collector [SSH Check receiver](sshcheckreceiver) connects to a configured endpoint via SSH to validate that the client can connect. The examples in this repo show how to configure an SSH endpoint and the Collector to send metrics to Lightstep Observability.

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

[sshcheckreceiver]: https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/sshcheckreceiver
[ls-docs-access-token]: https://docs.lightstep.com/docs/create-and-manage-access-tokens
[ls-docs-dashboards]: https://docs.lightstep.com/docs/create-and-manage-dashboards
