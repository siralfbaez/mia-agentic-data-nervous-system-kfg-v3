output "key_id" {
  description = "The self-link of the KMS crypto key"
  value       = google_kms_crypto_key.data_key.id
}

output "key_ring_id" {
  description = "The ID of the KMS key ring"
  value       = google_kms_key_ring.keyring.id
}
