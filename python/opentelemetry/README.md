# Python OTel Auto-Instrumentation README

## Setup

```bash
python3 -m venv .
source ./bin/activate

# Installs OTel libraries
pip install -r requirements.txt
opentelemetry-bootstrap -a install
```

## Send data to Lightstep via OTel Collector

> Note: This setup assumes that you have an OTel Collector running at `localhost:4317`. To run an OTel Collector instance locally, check out docs [here](../../../collector/vanilla/readme.md)

```bash
# Enable Flask debugging
export FLASK_DEBUG=1

# gRPC debug flags
export GRPC_VERBOSITY=debug
export GRPC_TRACE=http,call_error,connectivity_state

# Run Python app with auto-instrumentation
opentelemetry-instrument \
    --traces_exporter console,otlp \
    --metrics_exporter console \
    --service_name test-py-auto-collector-server \
    python server.py
```

To send over HTTP, replace `otlp` with `otlp_proto_http` in the `--traces_exporter` line.

# Send data to Lightstep direct from app (OTLP)

```bash
# Enable Flask debugging
export FLASK_DEBUG=1

# gRPC debug flags
export GRPC_VERBOSITY=debug
export GRPC_TRACE=http,call_error,connectivity_state

export OTEL_EXPORTER_OTLP_TRACES_HEADERS="lightstep-access-token=<LS_ACCESS_TOKEN>"

# Run Python app with auto-instrumentation
opentelemetry-instrument \
    --traces_exporter console,otlp_proto_grpc \
    --metrics_exporter console \
    --service_name test-py-auto-otlp-server \
    --exporter_otlp_traces_endpoint "ingest.lightstep.com:443" \
    python server.py
```

Be sure to replace `<LS_ACCESS_TOKEN>` with your own [Lightstep Access Toekn](https://docs.lightstep.com/docs/create-and-manage-access-tokens).

To use HTTP instead of gRPC:

```bash
opentelemetry-instrument \
    --traces_exporter console,otlp_proto_http \
    --metrics_exporter console,otlp_proto_http \
    --service_name test-py-auto-otlp-server \
    --exporter_otlp_traces_endpoint "https://ingest.lightstep.com/traces/otlp/v0.9" \
    python server.py
```

## Run the Client

In a separate terminal window:

```bash
source ./bin/activate

export OTEL_EXPORTER_OTLP_TRACES_HEADERS="lightstep-access-token=<LS_ACCESS_TOKEN>"

opentelemetry-instrument \
    --traces_exporter console,otlp \
    --metrics_exporter console,otlp \
    --service_name test-py-auto-client \
    --exporter_otlp_endpoint "ingest.lightstep.com:443" \
    python client.py test
```

Where `test` is the parameter being passed to `client.py`.

Be sure to replace `<LS_ACCESS_TOKEN>` with your own [Lightstep Access Toekn](https://docs.lightstep.com/docs/create-and-manage-access-tokens).

## References

More info on `opentelemetry-instrument` [here](https://github.com/open-telemetry/opentelemetry-python-contrib/tree/main/opentelemetry-instrumentation).