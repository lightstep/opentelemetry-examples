# Python OTel Manual Instrumentation README

## Setup

```bash
python3 -m venv .
source ./bin/activate

# Installs OTel libraries
pip install -r requirements.txt
```

## Send data to Lightstep direct from app

```bash
export FLASK_DEBUG=1

export GRPC_VERBOSITY=debug
export GRPC_TRACE=http,call_error,connectivity_state

export LS_ACCESS_TOKEN="<LS_ACCESS_TOKEN>"

python server.py
```

Be sure to replace `<LS_ACCESS_TOKEN>` with your own [Lightstep Access Toekn](https://docs.lightstep.com/docs/create-and-manage-access-tokens).

## Run the Client

In a separate terminal window:

```bash
curl http://localhost:8082/rolldice
```