variable "project" {
  description = "GCP project ID"
  type        = string
}

variable "credentials_file" {
  description = "Path to the GCP service account key JSON file"
  type        = string
}

variable "location" {
  description = "Location of the storage bucket"
  type        = string
  default     = "us-west1"
}

variable "bucket_name" {
  description = "Name of the storage bucket"
  type        = string
  default     = "sample-bucket"
}
