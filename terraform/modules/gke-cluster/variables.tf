variable "project_id" {
  type        = string
  description = "The GCP Project ID"
}

variable "cluster_name" {
  type    = string
  default = "mia-kfg-v3-cluster"
}

variable "region" {
  type    = string
  default = "us-central1"
}

variable "vpc_id" {
  type        = string
  description = "The VPC network to host the cluster"
}

variable "subnet_id" {
  type        = string
  description = "The subnetwork to host the cluster"
}
