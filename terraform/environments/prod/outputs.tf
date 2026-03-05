output "gke_endpoint" { value = module.gke.cluster_endpoint }
output "db_ip"       { value = module.alloydb.db_cluster_ip }
