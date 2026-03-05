# Create a dedicated Service Account for the AI Agent
resource "google_service_account" "ai_agent_sa" {
  account_id   = "ai-agent-sa"
  display_name = "MIA AI Agent Service Account"
}

# Bind Vertex AI User role to the SA
resource "google_project_iam_member" "vertex_ai_user" {
  project = var.project_id
  role    = "roles/aiplatform.user"
  member  = "serviceAccount:${google_service_account.ai_agent_sa.email}"
}

# Workload Identity Binding: Connects K8s Service Account to GCP Service Account
resource "google_service_account_iam_member" "workload_identity_user" {
  service_account_id = google_service_account.ai_agent_sa.name
  role               = "roles/iam.workloadIdentityUser"
  member             = "serviceAccount:${var.project_id}.svc.id.goog[ai-agent/ai-agent-sa]"
}
