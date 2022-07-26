# How to run the Collector locally

From the repo root:

```bash
cd collector/vanilla
docker run -it --rm -p 4317:4317 -p 4318:4318 \
    -v (pwd)/collector.yaml:/otel-config.yaml \
    --name otelcol otel/opentelemetry-collector-contrib:0.53.0  \
    "/otelcol-contrib" \
    "--config=otel-config.yaml"
```
