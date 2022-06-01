FROM maven:3-eclipse-temurin-11 AS build

RUN apt-get update
RUN apt-get install -y curl
RUN update-ca-certificates -f

WORKDIR /usr/src/app
RUN curl -o opentelemetry-javaagent.jar https://github.com/open-telemetry/opentelemetry-java-instrumentation/releases/latest/download/opentelemetry-javaagent.jar
COPY src ./src
COPY pom.xml pom.xml
COPY client ./client
RUN mvn -f /usr/src/app/pom.xml clean package

ENTRYPOINT mvn package exec:exec
