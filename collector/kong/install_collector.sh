#!/bin/bash

helm install my-collector open-telemetry/opentelemetry-collector -f values-collector.yaml  -f - <<EOF
config:
  exporters:
    otlp/public:
      headers:
        - "lightstep-access-token": "${LS_ACCESS_TOKEN}"
EOF
