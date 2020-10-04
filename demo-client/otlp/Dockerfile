FROM python:3.8
RUN pip install opentelemetry-launcher requests yaml
RUN opentelemetry-bootstrap -a install

ADD client.py /app/client.py
CMD ["opentelemetry-instrument", "/app/client.py"]
