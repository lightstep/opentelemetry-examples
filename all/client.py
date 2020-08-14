#!/usr/bin/env python3
import time

from opentelemetry import trace
import requests

tracer = trace.get_tracer(__name__)


def send_requests():
    destinations = [
        "http://go-opentracing-server:8081/content/456",
        "http://go-otlp-server:8081/content/345",
        "http://go-launcher-server:8081/content/234",
        "http://py-lstrace-server:5000/redis/124",
        "http://py-collector-server:5000/sqlalchemy/123",
        "http://py-otlp-server:5000/pymongo/123",
        "http://py-launcher-server:5000/redis/123",
        "http://js-lstrace-server:8080/ping",
        "http://js-launcher-server:8080/ping",
        "http://java-specialagent-server:8083/content",
        "http://java-otlp-server:8083/content",
    ]
    with tracer.start_as_current_span("send_requests"):
        for url in destinations:
            try:
                res = requests.get(url)
                print(f"Request to {url}, got {len(res.content)} bytes")
            except Exception as e:
                print(f"Request to {url} failed {e}")


if __name__ == "__main__":
    while True:
        send_requests()
        time.sleep(2)
