resource "google_container_cluster" "primary" {
  name     = var.cluster_name
  location = var.region

  # Enabling Autopilot for a "just-in-time" consumption experience
  # This aligns with the Confluent "Serverless Flink" philosophy
  enable_autopilot = true

  # Networking: Private nodes for NIST 800-53 compliance
  network    = var.vpc_id
  subnetwork = var.subnet_id

  ip_allocation_policy {
    cluster_secondary_range_name  = "pods"
    services_secondary_range_name = "services"
  }

  private_cluster_config {
    enable_private_nodes    = true
    enable_private_endpoint = false # Keep master public for easier dev access, or true for full locked-down
    master_ipv4_cidr_block  = "172.16.0.0/28"
  }

  # Workload Identity: The "Synapse" for secure GCP API access
  workload_identity_config {
    workload_pool = "${var.project_id}.svc.id.goog"
  }

  release_channel {
    channel = "REGULAR"
  }
}

# IAM Binding for the AI Agent's K8s Service Account to use Vertex AI
resource "google_project_iam_member" "vertex_ai_user" {
  project = var.project_id
  role    = "roles/aiplatform.user"
  member  = "serviceAccount:${var.project_id}.svc.id.goog[ai-agent/ai-agent-sa]"
}
