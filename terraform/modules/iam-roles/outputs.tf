output "ai_agent_sa_email" {
  description = "The email of the Google service account for the AI Agent"
  value       = google_service_account.ai_agent_sa.email
}

output "ai_agent_sa_name" {
  description = "The fully qualified name of the AI Agent service account"
  value       = google_service_account.ai_agent_sa.name
}
