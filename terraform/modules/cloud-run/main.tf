resource "google_cloud_run_v2_service" "default" {
  name     = var.service_name
  location = var.region

  template {
    containers {
      image = var.image_url
    }
  }
}

output "service_url" { value = google_cloud_run_v2_service.default.uri }
