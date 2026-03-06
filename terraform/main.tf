# Standardize the Terraform version across the Staff Engineering team
terraform {
  required_version = ">= 1.5.0"

  # We define the backend here but leave it partial 
  # so it can be filled by the CI/CD pipeline
  backend "gcs" {}
}
