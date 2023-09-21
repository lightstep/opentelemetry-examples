FROM debian:11.7

RUN apt update && apt -y install wget
RUN wget https://github.com/open-telemetry/opentelemetry-collector-releases/releases/download/v0.77.0/otelcol-contrib_0.77.0_linux_amd64.tar.gz
RUN tar -xf otelcol-contrib_0.77.0_linux_amd64.tar.gz -C /

ENTRYPOINT ["/otelcol-contrib"]
