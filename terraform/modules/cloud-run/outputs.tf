output "service_url" {
  description = "The URL of the deployed Cloud Run service"
  value       = google_cloud_run_v2_service.default.uri
}

output "service_id" {
  description = "The unique identifier for the Cloud Run service"
  value       = google_cloud_run_v2_service.default.id
}
