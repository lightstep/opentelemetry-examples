#!/bin/bash

export OTEL_EXPORTER_OTLP_ENDPOINT="ingest.lightstep.com:443"
export OTEL_EXPORTER_OTLP_HEADERS="lightstep-access-token=<your_token_here>"
