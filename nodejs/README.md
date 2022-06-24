# js examples

## Environment variables

Export or add to a .env file

```bash
export LS_ACCESS_TOKEN=<lightstep access token>
```

optionally, set the lightstep metrics url

```bash
export LS_METRICS_URL=https://ingest.lightstep.com/metrics
```

## Start the client

```bash
docker-compose up
```

## Supported variables

| Name                     | Required | Default                              |
| ------------------------ | -------- | ------------------------------------ |
| LS_ACCESS_TOKEN          | yes      |
| LS_SERVICE_NAME | yes      |                                      |
| LS_METRICS_URL           | No       | https://ingest.lightstep.com/metrics |
