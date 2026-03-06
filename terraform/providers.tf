provider "google" {
  project = var.project_id
  region  = var.region
}

# Optional: Add the Beta provider if you're using cutting-edge Vertex AI features
provider "google-beta" {
  project = var.project_id
  region  = var.region
}
