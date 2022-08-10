---
# Ingest Docker metrics using OTEL Collector's `dockerstats` or `prometheus` receivers

This example illustrates how you can ingest Docker metrics using the [OTEL collector's `dockerstats` receiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/dockerstatsreceiver#readme) or alternatively with the [`prometheus` receiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/prometheusreceiver#readme).

> WARNING: The `dockerstats` receiver is only supported on Linux.

## Requirements

* OpenTelemetry Collector Contrib v0.51.0+

## Prerequisites

You must have a Lightstep Observability [access token](/docs/create-and-manage-access-tokens) for the project you want to report metrics. This example is configured to recognize this in the environment variable `LS_ACCESS_TOKEN`.

## Configuration Requirements

A container that will use the Docker socket needs read/write permission on the host. And that's how we use the Docker API. So you can either run the container collecting metrics as a user with this permission (usually `root`, a.k.a. `0`) or under the group with this permission (usually `docker`). It's considered bad practice to run as the root user, but you should know that the docker group has equivalent privileges to root. There's more information about configuration and relevant security issues in the docs on [post installation steps for Linux](https://docs.docker.com/engine/install/linux-postinstall/) and [Docker daemon attack surface](https://docs.docker.com/engine/security/#docker-daemon-attack-surface).

You set the user and group in Docker Compose with the `user` key. The value is a string containing the id integer for the user or a colon separated user and group pair - `uid:gid`.

To run as the docker group, we can pick any integer other than 0 for the uid and set the group ID to that of the docker group. You can see the `gid` with `getent group docker`. On my Debian 11 machine it's 997. 

You also have to tell the Docker daemon to expose the metrics at a particular location. To do this you add something like the following to your `/etc/docker/daemon.json`. 

```
{
        "metrics-addr": "127.0.0.1:9100",
}
```

You may need to restart the Docker daemon after you change the configuration. If you use a systemd distribution then it's `sudo systemctl restart docker.service`.

## Running the example

You can just run `docker compose up -d --build` after you confirm your configuration.
