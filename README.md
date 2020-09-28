# Lightstep OpenTelemetry Examples

This repo contains example code and resources for working with OpenTelemetry.


## With docker-compose

The following client/server applications using different mechanism for sending data to Lightstep. The following examples are configured in the `docker-compose.yml` file:

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
| java-launcher  | client/server example instrumented via OpenTelemetry and the Launcher for configuration |
| js-ot-shim     | client/server example instrumented via OpenTelemetry and JS Launcher with OpenTracing |

### running examples

```bash
git clone https://github.com/lightstep/opentelemetry-examples && cd opentelemetrys-examples

# copy the example environment variable file
# and update the access token
cp example.env .env
sed -i '' 's/<ACCESS TOKEN>/YOUR TOKEN HERE/' .env
cp example-collector-config.yaml ./collector/collector-config.yaml
sed -i '' 's/<ACCESS TOKEN>/YOUR TOKEN HERE/' ./collector/collector-config.yaml

docker-compose up
```
