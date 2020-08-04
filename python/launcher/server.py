#!/usr/bin/env python
import os
import random
import string

import redis
from pymongo import MongoClient
from sqlalchemy import Column, ForeignKey, Integer, String, create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import relationship

from opentelemetry import trace
from opentelemetry.launcher import configure_opentelemetry


configure_opentelemetry()
tracer = trace.get_tracer(__name__)
from flask import Flask

app = Flask(__name__)

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
    app.run(host="0.0.0.0")
