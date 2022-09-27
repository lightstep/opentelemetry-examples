#!/usr/bin/env python
#
# example code to test opentelemetry
#
# usage:
#  OTEL_EXPORTER_OTLP_TRACES_HEADERS="lightstep-access-token=<LS_ACCESS_TOKEN>"
#   opentelemetry-instrument \
#       --traces_exporter console,otlp_proto_grpc \
#       --metrics_exporter console \
#       --service_name test-py-auto-otlp-server \
#       --exporter_otlp_traces_endpoint "ingest.lightstep.com:443" \
#       python server.py
#
# See README.md for more details.

from typing import Iterable
import time

import random
import string
import flask

import redis
from pymongo import MongoClient
from sqlalchemy import Column, ForeignKey, Integer, String, create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import relationship

from opentelemetry import trace

# Metrics stuff
from opentelemetry.exporter.otlp.proto.grpc.metric_exporter import (
    OTLPMetricExporter,
)
from opentelemetry.metrics import (
    CallbackOptions,
    Observation,
    get_meter_provider,
    set_meter_provider,
)
from opentelemetry.sdk.metrics import MeterProvider
from opentelemetry.sdk.metrics.export import PeriodicExportingMetricReader

exporter = OTLPMetricExporter(insecure=True)
# reader = PeriodicExportingMetricReader(exporter)
# provider = MeterProvider(metric_readers=[reader])
# set_meter_provider(provider)

meter = get_meter_provider().get_meter(__name__)


# from common import get_tracer
tracer = trace.get_tracer_provider().get_tracer(__name__)

# tracer = get_tracer()

app = flask.Flask(__name__)

Base = declarative_base()

### ------- Metrics stuff
def observable_counter_func(options: CallbackOptions) -> Iterable[Observation]:
    yield Observation(1, {})


def observable_up_down_counter_func(
    options: CallbackOptions,
) -> Iterable[Observation]:
    yield Observation(-10, {})


def observable_gauge_func(options: CallbackOptions) -> Iterable[Observation]:
    yield Observation(9, {})


counter = meter.create_counter(
    name="requests_counter",
    description="number of requests",
    unit="1"
)
# counter.add(1)

# Async Counter
observable_counter = meter.create_observable_counter(
    "observable_counter",
    [observable_counter_func],
)

# UpDownCounter
updown_counter = meter.create_up_down_counter("updown_counter")
updown_counter.add(1)
updown_counter.add(-5)

# Async UpDownCounter
observable_updown_counter = meter.create_observable_up_down_counter(
    "observable_updown_counter", [observable_up_down_counter_func]
)

# Histogram
histogram = meter.create_histogram(
    name="request_size_bytes",
    description="size of requests",
    unit="byte"
)    
# histogram = meter.create_histogram("histogram")
# histogram.record(99.9)

# Async Gauge
gauge = meter.create_observable_gauge("gauge", [observable_gauge_func])

staging_attributes = {"environment": "staging"}


def set_metrics(length):

    print("Setting metrics")
    
    counter.add(random.randint(0, 25), staging_attributes)
    histogram.record(length, staging_attributes)
    updown_counter.add(random.randint(-5, 25), staging_attributes)

### ------- END Metrics stuff

class Person(Base):
    __tablename__ = "person"
    # Here we define columns for the table person
    # Notice that each column is also a normal Python instance attribute.
    id = Column(Integer, primary_key=True)
    name = Column(String(250), nullable=False)


class Address(Base):
    __tablename__ = "address"
    # Here we define columns for the table address.
    # Notice that each column is also a normal Python instance attribute.
    id = Column(Integer, primary_key=True)
    street_name = Column(String(250))
    street_number = Column(String(250))
    post_code = Column(String(250), nullable=False)
    person_id = Column(Integer, ForeignKey("person.id"))
    person = relationship(Person)


def _random_string(length):
    """Generate a random string of fixed length """
    letters = string.ascii_lowercase
    return "".join(random.choice(letters) for i in range(int(length)))


@app.route("/ping")
def ping():
    length = random.randint(1, 1024)
    set_metrics(length)
    # redis_integration(length)
    # pymongo_integration(length)
    # sqlalchemy_integration(length)
    return _random_string(length)


@app.route("/redis/<length>")
def redis_integration(length):
    with tracer.start_as_current_span("server redis operation"):
        r = redis.Redis(host="redis", port=6379)
        r.mset({"length": _random_string(length)})
        return str(r.get("length"))


@app.route("/pymongo/<length>")
def pymongo_integration(length):
    with tracer.start_as_current_span("server pymongo operation"):
        client = MongoClient("mongo", 27017, serverSelectionTimeoutMS=2000)
        db = client["opentelemetry-tests"]
        collection = db["tests"]
        collection.find_one()
        return _random_string(length)


@app.route("/sqlalchemy/<length>")
def sqlalchemy_integration(length):
    with tracer.start_as_current_span("server sqlalchemy operation"):
        # Create an engine that stores data in the local directory's
        # sqlalchemy_example.db file.
        engine = create_engine("sqlite:///sqlalchemy_example.db")

        # Create all tables in the engine. This is equivalent to "Create Table"
        # statements in raw SQL.
        Base.metadata.create_all(engine)
        return str(_random_string(length))


if __name__ == "__main__":
    # set_metrics()
    # Counter
    
    app.run(host="0.0.0.0", port=8081, debug=True, use_reloader=False)
