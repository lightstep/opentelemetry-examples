FROM maven:3-eclipse-temurin-11 AS build

RUN apt-get update
RUN apt-get install -y curl
RUN update-ca-certificates -f

WORKDIR /usr/src/app
RUN curl -o opentracing-specialagent-1.7.0.jar https://repo1.maven.org/maven2/io/opentracing/contrib/specialagent/opentracing-specialagent/1.7.0/opentracing-specialagent-1.7.0.jar
COPY src ./src
COPY pom.xml pom.xml
RUN mvn -f /usr/src/app/pom.xml clean package

FROM ibmjava:8-jre 

COPY --from=build /usr/src/app/opentracing-specialagent-1.7.0.jar /app/
COPY --from=build /usr/src/app/target/client-1.0-SNAPSHOT.jar /app/

ENTRYPOINT java -javaagent:/app/opentracing-specialagent-1.7.0.jar  \
        -Dsa.tracer=lightstep \
        -Dls.componentName=$LS_SERVICE_NAME \
        -Dls.accessToken=$LS_ACCESS_TOKEN \
        -Dls.collectorHost=$LS_COLLECTOR_HOST \
        -Dls.metricsUrl=$LS_METRICS_URL \
        -Dls.propagator=b3 \
        -cp /app/client-1.0-SNAPSHOT.jar \
        com.lightstep.examples.client.App
