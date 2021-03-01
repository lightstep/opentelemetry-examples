FROM python:3.8


ADD requirements.txt requirements.txt
RUN pip install -r requirements.txt

ADD test.py /app/test.py
ADD generate_config.py /app/generate_config.py
CMD ["opentelemetry-instrument", "pytest", "-vvvv", "/app/test.py"]
