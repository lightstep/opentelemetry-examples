# go examples

## Export your access token
```bash
export SECRET_TOKEN=<lightstep access token>
```

## Start the server
```bash
LS_ACCESS_TOKEN=${SECRET_TOKEN} \
LS_SERVICE_NAME=demo-client-go \
LS_SERVICE_VERSION=0.1.8 \
go run server.go
```

## Start the client
```bash
LS_ACCESS_TOKEN=${SECRET_TOKEN} \
LS_SERVICE_NAME=demo-client-go \
LS_SERVICE_VERSION=0.1.8 \
go run client.go
```

## Supported variables


| Name | Required | Default |
| ---- | -------- | ------- |
|LS_ACCESS_TOKEN| yes|
|LS_SERVICE_NAME|yes|
|LS_SERVICE_VERSION|yes|
|LIGHTSTEP_HOST| No | ingest.lightstep.com|
|LIGHTSTEP_PORT| No | 443 |
|LIGHTSTEP_SECURE| No | 1 |