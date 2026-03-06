output "db_cluster_ip" {
  description = "The private IP address of the AlloyDB primary instance"
  # Reference the INSTANCE, not the CLUSTER
  value = google_alloydb_instance.primary.ip_address
}