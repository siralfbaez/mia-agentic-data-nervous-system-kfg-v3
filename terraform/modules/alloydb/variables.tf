variable "cluster_id"  { type = string }
variable "region"      { type = string }
variable "vpc_id"      { type = string }
variable "db_password" { type = string; sensitive = true }
