# OpenTelemetry Metrics Example (Vanilla Setup)

This example shows how to configure OpenTelemetry JS to export metrics to Lightstep without any additional (non-OTel) dependencies.

## Installation

```
npm i
```

## Run the Application

- Export your access token as LS_ACCESS_TOKEN

```
export LS_ACCESS_TOKEN=<YOUR_TOKEN>
```

- Optionally, export a prefix for your metrics. By default the metrics will be prefixed `demo.`

```
export METRICS_PREFIX=foo
```

- Run the example

```
npm run start
```
