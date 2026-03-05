resource "google_alloydb_cluster" "main" {
  cluster_id = var.cluster_id
  location   = var.region
  network    = var.vpc_id

  initial_user {
    password = var.db_password
    user     = "admin"
  }
}

resource "google_alloydb_instance" "primary" {
  cluster       = google_alloydb_cluster.main.name
  instance_id   = "${var.cluster_id}-primary"
  instance_type = "PRIMARY"

  machine_config {
    cpu_count = 4 # High performance for stream-to-db mapping
  }
}

output "cluster_ip" {
  value = google_alloydb_cluster.main.ip_address
}
