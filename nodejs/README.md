# js examples

## Environment variables

Export or add to a .env file

```bash
export LS_ACCESS_TOKEN=<lightstep access token>
export DD_TRACE_AGENT_URL=https://ingest.lightstep.com
```

optionally, set the lightstep metrics url

```bash
export LS_METRICS_URL=https://ingest.staging.lightstep.com/metrics
```

## Start the client

```bash
docker-compose up
```

## Supported variables

| Name                     | Required | Default                              |
| ------------------------ | -------- | ------------------------------------ |
| LS_ACCESS_TOKEN          | yes      |
| LIGHTSTEP_COMPONENT_NAME | yes      |                                      |
| DD_TRACE_AGENT_URL       | yes      |
| LS_METRICS_URL           | No       | https://ingest.lightstep.com/metrics |
