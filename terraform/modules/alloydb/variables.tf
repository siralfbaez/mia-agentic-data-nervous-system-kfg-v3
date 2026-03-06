variable "project_id" {
  type        = string
  description = "The GCP Project ID"
}

variable "region" {
  type        = string
  description = "The GCP region"
}

# This fixes the "Unsupported argument: cluster_id" error
variable "cluster_id" {
  type        = string
  description = "The ID for the AlloyDB cluster"
}

# This fixes the "Unsupported argument: vpc_id" error
variable "vpc_id" {
  type        = string
  description = "The VPC ID where AlloyDB will be peered"
}

variable "db_password" {
  type      = string
  sensitive = true
}