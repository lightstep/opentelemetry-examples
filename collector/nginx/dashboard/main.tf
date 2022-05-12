terraform {
  required_providers {
    lightstep = {
      source  = "lightstep/lightstep"
      version = "~> 1.60.2"
    }
  }
  required_version = ">= v1.0.11"
}

resource "lightstep_metric_dashboard" "otel_collector_nginxreceiver_dashboard" {
  project_name   = var.lightstep_project
  dashboard_name = "OpenTelemetry nginxreceiver Integration"



  chart {
    name = "nginx.connections_accepted"
    rank = "0"
    type = "timeseries"

    # These descriptive OTEL fields are provided for documentation purposes only:
    # - description: The total number of accepted client connections
    # - unit: connections

    query {
      query_name = "a"
      display    = "line"
      hidden     = false

      metric              = "nginx.connections_accepted"
      timeseries_operator = "rate"

      group_by {
        aggregation_method = "sum"
        keys               = []

      }

    }

  }

  chart {
    name = "nginx.connections_current"
    rank = "1"
    type = "timeseries"

    # These descriptive OTEL fields are provided for documentation purposes only:
    # - description: The current number of nginx connections by state
    # - unit: connections

    query {
      query_name = "a"
      display    = "line"
      hidden     = false

      metric              = "nginx.connections_current"
      timeseries_operator = "last"

      group_by {
        aggregation_method = "sum"
        keys               = ["state"]

      }

    }

  }

  chart {
    name = "nginx.connections_handled"
    rank = "2"
    type = "timeseries"

    # These descriptive OTEL fields are provided for documentation purposes only:
    # - description: The total number of handled connections. Generally, the parameter value is the same as nginx.connections_accepted unless some resource limits have been reached (for example, the worker_connections limit).
    # - unit: connections

    query {
      query_name = "a"
      display    = "line"
      hidden     = false

      metric              = "nginx.connections_handled"
      timeseries_operator = "rate"

      group_by {
        aggregation_method = "sum"
        keys               = []

      }

    }

  }

  chart {
    name = "nginx.requests"
    rank = "3"
    type = "timeseries"

    # These descriptive OTEL fields are provided for documentation purposes only:
    # - description: Total number of requests made to the server since it started
    # - unit: requests

    query {
      query_name = "a"
      display    = "line"
      hidden     = false

      metric              = "nginx.requests"
      timeseries_operator = "rate"

      group_by {
        aggregation_method = "sum"
        keys               = []

      }
    }

  }

  chart {
    name = "nginx.connections_current"
    rank = "4"
    type = "timeseries"

    # These descriptive OTEL fields are provided for documentation purposes only:
    # - description: The total number of current client connections aggregated by "instance.type"
    # - unit: connections
    query {
      query_name = "a"
      display    = "line"
      hidden     = false

      metric              = "nginx.connections_current"
      timeseries_operator = "count"

      group_by {
        aggregation_method = "sum"
        keys               = ["instance.type", ]
      }

    }

  }

}
