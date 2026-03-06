module "vpc" {
  source     = "../../modules/vpc-network"
  region     = var.region
  project_id = var.project_id
}

module "gke" {
  source       = "../../modules/gke-cluster"
  project_id   = var.project_id
  cluster_name = "mia-agentic-data-nervous-system-prod"
  vpc_id       = module.vpc.vpc_id
  subnet_id    = module.vpc.subnet_id
}

module "alloydb" {
  source      = "../../modules/alloydb"
  cluster_id  = "kfg-v3-state-prod"
  vpc_id      = module.vpc.vpc_id
  db_password = var.db_password
}
