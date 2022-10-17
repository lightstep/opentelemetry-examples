
# Kubernetes kubelet, etcd, and API Server Tracing

Demo of experimental distributed tracing features in Kubernetes 1.25+.

## Instructions

Start an OpenTelemetry collector to ingest traces:

```
$ export LS_ACCESS_TOKEN=<your-lightstep-access-token>
$ docker-compose up
```

Update *.yaml and *.toml config files to point to the *external* IP address of the collector started above (localhost won't work).

Start minikube (> v1.26.1) with experimental tracing feature gates:

```
$ ./run-minikube.sh
```
