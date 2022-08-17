ui = true
api_addr  = "http://127.0.0.1:8200"

storage "consul" {
  address = "consul-server:8500"
  path    = "vault/"
}

telemetry {
  disable_hostname = true
  prometheus_retention_time = "12h"
}