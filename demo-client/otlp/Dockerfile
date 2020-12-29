FROM python:3.8
RUN pip install opentelemetry-launcher requests pyyaml protobuf==3.12.4
RUN opentelemetry-bootstrap -a install

ADD client.py /app/client.py
CMD ["opentelemetry-instrument", "/app/client.py"]
