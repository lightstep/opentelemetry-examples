from opentelemetry import trace
from opentelemetry.instrumentation.flask import FlaskInstrumentor
from opentelemetry.instrumentation.requests import RequestsInstrumentor
from random import randint
from flask import Flask, request
from common import get_tracer
import uuid

# Init tracer
tracer = get_tracer()

# Init autoinstrumentation with Flask
app = Flask(__name__)
FlaskInstrumentor().instrument_app(app)
RequestsInstrumentor().instrument()

@app.route("/ping")
def handle_ping():
    return str(handle_ping())

@app.route("/rolldice")
def roll_dice():
    return str(do_roll())

@tracer.start_as_current_span("do_work")
def do_roll():
    res = randint(1, 6)
    current_span = trace.get_current_span()
    current_span.set_attribute("roll.value", res)
    current_span.set_attribute("operation.name", "Saying hello!")
    current_span.set_attribute("operation.other-stuff", [1, 2, 3])
    return res

@tracer.start_as_current_span("handle_ping")
def handle_ping():
    res = uuid.uuid4().hex
    current_span = trace.get_current_span()
    current_span.set_attribute("library.language", "python"),
    current_span.set_attribute("library.version", "v1.7.0"),
    current_span.set_status(1)
    
    current_span.add_event("Suuuuuppp")
    
    print(f"Returning {res}")
    return res


if __name__ == "__main__":
    app.run(host="0.0.0.0")
