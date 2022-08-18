# How to run the Collector locally

1. Edit the [Collector config YAML](collector.yml)

    Replace `${LIGHTSTEP_ACCESS_TOKEN}` with your own [Lightstep Access Token](https://docs.lightstep.com/docs/create-and-manage-access-tokens)

2. Run the Collector's Docker container instance

    Ensure that you are in the repo root folder (`opentelemetry-examples`), then run:

    ```bash
    cd collector/vanilla
    docker run -it --rm -p 4317:4317 -p 4318:4318 \
        -v $(pwd)/collector.yml:/otel-config.yaml \
        --name otelcol otel/opentelemetry-collector-contrib:0.53.0  \
        "/otelcol-contrib" \
        "--config=otel-config.yaml"
    ```
