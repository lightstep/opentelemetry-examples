output "dashboard_url" {
  value       = "https://app.lightstep.com/${var.lightstep_project}/dashboard/${lightstep_metric_dashboard.otel_collector_nginxreceiver_dashboard.id}"
  description = "OpenTelemetry Nginx Dashboard URL"
}
