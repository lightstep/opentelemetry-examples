FROM python:3-slim-buster
RUN pip install opentelemetry-launcher requests pyyaml protobuf==3.20.1
RUN opentelemetry-bootstrap -a install

ADD client.py /app/client.py
CMD ["opentelemetry-instrument", "/app/client.py"]
