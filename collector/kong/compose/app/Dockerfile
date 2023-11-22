#COPY ./httpd-info.conf /usr/local/apache2/conf/extra/httpd-info.conf
FROM python:3-alpine3.15

RUN apk add build-base

RUN mkdir /app
WORKDIR /app
ADD requirements.txt .
RUN pip install -r requirements.txt

ADD *.py ./
CMD ["python", "-u", "/app/client.py"]
