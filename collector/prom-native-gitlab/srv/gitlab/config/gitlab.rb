# Enable HTTP for the Prometheus endpoint
prometheus['enable'] = true
prometheus['listen_address'] = 'localhost:9090'
prometheus['ssl_enabled'] = false

# Allow Prometheus to access the metrics endpoint
gitlab_workhorse['prometheus_listen_addr'] = "localhost:9229"

external_url 'http://gitlab.example.com'