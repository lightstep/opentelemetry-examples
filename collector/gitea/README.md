# Monitor Gitea with the OpenTelemetry Collector

## Overview

Gitea is a user-friendly self-hosted Git service. Proper monitoring is crucial for the reliability and efficiency of any Gitea instance. With Gitea's metrics exposure capability and the OpenTelemetry Collector, these metrics can be easily forwarded to Lightstep for in-depth analysis and visualization. This README guides you through the process of setting up the OpenTelemetry Collector to funnel Gitea's metrics into Lightstep.

## Prerequisites

* Docker
* Docker Compose
* A Lightstep Observability account
* Lightstep Observability [access token][ls-docs-access-token]

## How to set it up

1. **Export your Lightstep access token**:
    ```bash
    export LS_ACCESS_TOKEN=<YOUR_LIGHTSTEP_TOKEN>
    ```
2. **Run the docker compose example to spin up Gitea and the OpenTelemetry Collector**:
    ```bash
    docker-compose up -d
    ```
3. **Access Gitea's web interface**: Visit http://localhost:3000.
4. **Monitor Gitea Metrics in Lightstep**: After setting things up, Gitea metrics should start populating in your Lightstep dashboard.
5. **Shutting down the monitoring setup**:
    ```bash
    docker-compose down
    ```
 
