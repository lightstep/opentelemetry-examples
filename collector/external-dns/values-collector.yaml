mode: "deployment"

config:
  exporters:
    logging:
      loglevel: debug
    otlp/public:
      endpoint: ingest.lightstep.com:443
      headers:
        - "lightstep-access-token": "LS_ACCESS_TOKEN"
  processors:
    batch: {}
    memory_limiter: null
  receivers:
    prometheus:
      config:
        scrape_configs:
          - job_name: opentelemetry-collector
            scrape_interval: 10s
            static_configs:
              - targets:
                  - external-dns.kube-system.svc.cluster.local:7979
  service:
    telemetry:
      metrics:
        address: 0.0.0.0:8888
    pipelines:
      logs:
        exporters:
          - logging
          - otlp/public
        processors:
          - memory_limiter
          - batch
        receivers:
          - otlp
      metrics:
        exporters:
          - logging
          - otlp/public
        processors:
          - memory_limiter
          - batch
        receivers:
          - prometheus

image:
  repository: otel/opentelemetry-collector-contrib
  pullPolicy: IfNotPresent
  tag: "0.59.0"
