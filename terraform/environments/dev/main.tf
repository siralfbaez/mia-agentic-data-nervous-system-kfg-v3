module "vpc" {
  source     = "../../modules/vpc-network"
  project_id = var.project_id
  region     = var.region
}

module "gke" {
  source       = "../../modules/gke-cluster"
  project_id   = var.project_id
  cluster_name = "mia-kfg-v3-dev"
  vpc_id       = module.vpc.vpc_id
  subnet_id    = module.vpc.subnet_id
}
