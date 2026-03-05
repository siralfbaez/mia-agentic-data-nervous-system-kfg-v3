resource "google_kms_key_ring" "keyring" {
  name     = var.key_ring_name
  location = var.region
}

resource "google_kms_crypto_key" "data_key" {
  name     = "mia-data-encryption-key"
  key_ring = google_kms_key_ring.keyring.id
  rotation_period = "7776000s" # 90 days rotation for NIST compliance
}

output "key_id" { value = google_kms_crypto_key.data_key.id }
