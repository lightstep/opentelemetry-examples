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
    --service_name test-py-auto-collector-server \
    python server.py
```


# Send data to Lightstep direct from app (OTLP)

```bash
# Enable Flask debugging
export FLASK_DEBUG=1

# gRPC debug flags
export GRPC_VERBOSITY=debug
export GRPC_TRACE=http,call_error,connectivity_state

export OTEL_EXPORTER_OTLP_TRACES_HEADERS="<LS_ACCESS_TOKEN>"

# Run Python app with auto-instrumentation
opentelemetry-instrument \
    --traces_exporter console,otlp \
    --service_name test-py-auto-otlp \
    --exporter_otlp_endpoint "ingest.lightstep.com:443" \
    python server.py
```

Be sure to replace `<LS_ACCESS_TOKEN>` with your own [Lightstep Access Toekn](https://docs.lightstep.com/docs/create-and-manage-access-tokens)
