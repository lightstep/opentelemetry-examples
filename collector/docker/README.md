---
# Ingest Docker metrics using OTEL Collector's Prometheus receiver

This example illustrates how you can ingest Docker metrics using the OTEL collector's Prometheus receiver. Docker exposes a Prometheus compatible metrics endpoint.

## Requirements

* OpenTelemetry Collector Contrib v0.51.0+

## Prerequisites

You must have a Lightstep Observability [access token](/docs/create-and-manage-access-tokens) for the project you want to report metrics. This example is configured to recognize this in the environment variable `LS_ACCESS_TOKEN`.

## Configuration Requirements

A container that will use the Docker socket needs read/write permission on the host. And that's how we use the Docker API. So you can either run the container collecting metrics as a user with this permission (usually `root`, a.k.a. `0`) or under the group with this permission (usually `docker`). It's considered bad practice to run as the root user, but you should know that the docker group has equivalent privileges to root. There's more information about configuration and relevant security issues in the docs on [post installation steps for Linux](https://docs.docker.com/engine/install/linux-postinstall/) and [Docker daemon attack surface](https://docs.docker.com/engine/security/#docker-daemon-attack-surface).

You set the user and group in Docker Compose with the `user` key. The value is a string containing the id integer for the user or a colon separated user and group pair - `uid:gid`.

To run as the docker group, we can pick any integer other than 0 for the uid and set the group ID to that of the docker group. You can see the `gid` with `getent group docker`. On my Debian 11 machine it's 997. 

## Running the Example

> WARNING: This example mounts the Docker socket in a container which creates certain security risks. You can learn more about how to manage this risk [Docker's documentation on protecting access](https://docs.docker.com/engine/security/protect-access/) to the daemon socket for real world suggestions.

The command to run this example simply is `docker compose up -d --build`. 

