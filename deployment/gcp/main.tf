terraform {
  required_version = ">= 1.5"

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 7.33"
    }
  }
}

provider "google" {
  project     = var.project
  credentials = file(var.credentials_file)
}

resource "google_storage_bucket" "storage_bucket" {
  project       = var.project
  name          = var.bucket_name
  location      = var.location
  force_destroy = true

  cors {
    method          = ["*"]
    origin          = ["*"]
    max_age_seconds = 3600
  }
}
