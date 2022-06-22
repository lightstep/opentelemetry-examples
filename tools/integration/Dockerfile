FROM python:3-alpine3.15


ADD requirements.txt requirements.txt
RUN pip install -r requirements.txt

ADD test.py /app/test.py
ADD generate_config.py /app/generate_config.py
CMD ["opentelemetry-instrument", "pytest", "/app/test.py"]
