FROM python:3.8


ADD requirements.txt requirements.txt
RUN pip install -r requirements.txt
RUN opentelemetry-bootstrap -a install
RUN pip install protobuf~=3.12.2

ADD test.py /app/test.py
CMD ["opentelemetry-instrument", "pytest", "/app/test.py"]
