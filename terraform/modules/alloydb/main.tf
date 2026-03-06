resource "google_alloydb_cluster" "default" {
  cluster_id = var.cluster_id # Using the variable here
  location   = var.region
  project    = var.project_id
  network_config {
    network = var.vpc_id # Using the variable here
  }
}

resource "google_alloydb_instance" "primary" {
  cluster       = google_alloydb_cluster.default.name
  instance_id   = "${var.cluster_id}-primary"
  instance_type = "PRIMARY"

  # Staff Tip: Use a small tier for Dev/Prod unless high-load is expected
  machine_config {
    cpu_count = 2
  }
}