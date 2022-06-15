FROM curlimages/curl:7.82.0 as curler
ARG JMX_JAR_VERSION=v1.14.0
USER root
RUN curl -L \
    --output /opentelemetry-jmx-metrics.jar \
    "https://github.com/open-telemetry/opentelemetry-java-contrib/releases/download/${JMX_JAR_VERSION}/opentelemetry-jmx-metrics.jar"

RUN curl -L -s \
    "https://github.com/open-telemetry/opentelemetry-collector-releases/releases/download/v0.53.0/otelcol-contrib_0.53.0_linux_amd64.tar.gz" | \
    tar -xvz -C /

FROM ibmjava:8-jre
WORKDIR /

COPY --from=curler /opentelemetry-jmx-metrics.jar /opt/opentelemetry-jmx-metrics.jar
COPY --from=curler /otelcol-contrib /otelcol-contrib

ENTRYPOINT [ "/otelcol-contrib" ]
CMD ["--config", "/etc/otel/config.yaml"]