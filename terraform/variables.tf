variable "project_id" {
  type        = string
  description = "The GCP Project ID where the Nervous System is hosted"
}

variable "region" {
  type        = string
  default     = "us-central1"
  description = "Primary region for high-performance AI compute"
}

variable "db_password" {
  type        = string
  sensitive   = true
  description = "Master password for the AlloyDB KFG state store"
}
