# Lightstep Examples

This repo contains example client/server applications using different mechanism for sending data to Lightstep. The following examples are configured in the docker-compose file:

| name           | description |
| -------------- | ----------- |
| go-opentracing | client/server example instrumented via lightstep-tracer-go |
| go-otlp        | client/server example instrumented via OpenTelemetry and the OTLP exporter |
| go-launcher    | client/server example instrumented via OpenTelemetry and the Launcher for configuration |
| py-lstrace     | client/server example instrumented via ls-trace-py |
| py-collector   | client/server example instrumented via OpenTelemetry and the OTLP exporter combined with the OpenTelemetry Collector |
| py-otlp        | client/server example instrumented via OpenTelemetry and the OTLP exporter |
| py-launcher    | client/server example instrumented via OpenTelemetry and the Launcher for configuration |
| js-lstrace     | client/server example instrumented via ls-trace-js |
| java           | client/server example instrumented via special agent |
| java-otlp      | client/server example instrumented via OpenTelemetry and the OTLP exporter |

## Getting started

```bash
git clone https://github.com/lightstep/opentelemetry-examples && cd opentelemetrys-examples
cp example.env .env
# edit .env file with your access token
docker-compose up
```
