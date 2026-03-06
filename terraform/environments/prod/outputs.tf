output "alloydb_address" {
  description = "The internal IP for the AI Agent's state store"
  value       = module.alloydb.db_cluster_ip
}

output "gke_endpoint" {
  description = "The Kubernetes Control Plane address"
  value       = module.gke.cluster_endpoint
}