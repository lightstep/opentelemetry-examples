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
    --service_name test-py-auto-launcher-server \
    --exporter_otlp_traces_endpoint "0.0.0.0:4317" \
    --exporter_otlp_traces_insecure true \
    python server.py
```

# Send data to Lightstep direct from app (OTLP)

```bash
# Enable Flask debugging
export FLASK_DEBUG=1

# gRPC debug flags
export GRPC_VERBOSITY=debug
export GRPC_TRACE=http,call_error,connectivity_state

export LS_ACCESS_TOKEN="<LS_ACCESS_TOKEN>"

# Run Python app with auto-instrumentation
opentelemetry-instrument \
    --service_name test-py-auto-launcher-server \
    python server.py
```

Be sure to replace `<LS_ACCESS_TOKEN>` with your own [Lightstep Access Toekn](https://docs.lightstep.com/docs/create-and-manage-access-tokens).

## Run the Client

In a separate terminal window:

```bash
source ./bin/activate
export LS_ACCESS_TOKEN="<LS_ACCESS_TOKEN>"

# Run Python app with auto-instrumentation
opentelemetry-instrument \
    --service_name test-py-auto-launcher-client \
    python client.py test
```

Where `test` is the parameter being passed to `client.py`.

Be sure to replace `<LS_ACCESS_TOKEN>` with your own [Lightstep Access Toekn](https://docs.lightstep.com/docs/create-and-manage-access-tokens).


## References

More info on `opentelemetry-instrument` [here](https://github.com/open-telemetry/opentelemetry-python-contrib/tree/main/opentelemetry-instrumentation).

More info on the Python Launcher [here](https://github.com/lightstep/otel-launcher-python).