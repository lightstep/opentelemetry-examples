# notes

collector command

```bash
docker run -it --rm \
     -p 4317:4317 \
     -p 4318:4318 \
     -v $(pwd)/collector.yaml:/otel-config.yaml \
     --network="otel-collector-demo" \
     -h otel-collector \
     --name otel-collector \
     otel/opentelemetry-collector-contrib:0.63.0  \
     "/otelcol-contrib" \
     "--config=otel-config.yaml"
```
