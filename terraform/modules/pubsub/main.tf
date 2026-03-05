resource "google_pubsub_topic" "signals" {
  name = "mia-signals-v3"
}

resource "google_pubsub_subscription" "worker_sub" {
  name  = "mia-worker-subscription"
  topic = google_pubsub_topic.signals.name
  ack_deadline_seconds = 20
}

output "topic_name" { value = google_pubsub_topic.signals.name }
