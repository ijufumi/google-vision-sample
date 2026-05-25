output "bucket_name" {
  description = "Name of the created storage bucket"
  value       = google_storage_bucket.storage_bucket.name
}

output "bucket_url" {
  description = "Base URL of the created storage bucket"
  value       = google_storage_bucket.storage_bucket.url
}
