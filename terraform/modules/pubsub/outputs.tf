output "topic_id" { value = google_pubsub_topic.signals.id }
output "subscription_id" { value = google_pubsub_subscription.worker_sub.id }
