#!/usr/bin/env python
#
# example code to test opentelemetry
#
# usage:
#   LS_ACCESS_TOKEN=${SECRET_TOKEN} \
#   python server.py \

from opentelemetry import trace
from opentelemetry.instrumentation.flask import FlaskInstrumentor
from opentelemetry.instrumentation.requests import RequestsInstrumentor
import random
import string
from flask import Flask, request
import requests
from common import get_tracer
import uuid
from opentelemetry.trace.propagation.tracecontext import TraceContextTextMapPropagator

import redis
from pymongo import MongoClient
from sqlalchemy import Column, ForeignKey, Integer, String, create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import relationship


# Init tracer
tracer = get_tracer()

# Init autoinstrumentation with Flask
app = Flask(__name__)
FlaskInstrumentor().instrument_app(app)
RequestsInstrumentor().instrument()

Base = declarative_base()

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

@app.route("/rolldice")
def roll_dice():
    return str(do_roll())

def get_header_from_flask_request(request, key):
    # print(f"Request headers {request.headers}")
    return request.headers.get_all(key)

# @tracer.start_as_current_span("do_roll")
def do_roll():
    traceparent = get_header_from_flask_request(request, "traceparent")
    carrier = {"traceparent": traceparent[0]}
    print(f"Carrier value: {carrier}")
    
    ctx = TraceContextTextMapPropagator().extract(carrier)
    
    with tracer.start_as_current_span("do_roll", context=ctx) as current_span:
        
        res = random.randint(1, 6)
        # current_span = trace.get_current_span()
            
        # request.headers["carrier"] = carrier
        # print(f"Header value: {request.headers}")
        # print(f"Carrier: {carrier}")
        current_span.set_attribute("roll.value", res)
        current_span.set_attribute("operation.name", "Saying hello!")
        current_span.set_attribute("operation.other-stuff", [1, 2, 3])
        current_span.add_event("Suuuuuppp")    
        print(f"Returning {res}")
        return res


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8081, debug=True, use_reloader=False)