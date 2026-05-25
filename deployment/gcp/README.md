# GCP deployment (Terraform)

Provisions the Google Cloud resources required by the application. Currently this
is a single Cloud Storage bucket (with permissive CORS) used to store uploaded
images/PDFs.

> Migrated from CDK for Terraform (cdktf) to plain Terraform HCL, since cdktf is
> no longer actively developed.

## Prerequisites

- [Terraform](https://developer.hashicorp.com/terraform/install) >= 1.5
- A GCP service account key JSON file with permission to manage Cloud Storage

## Configuration

Copy the example tfvars file and fill in your values:

```bash
cp terraform.tfvars.example terraform.tfvars
```

| Variable           | Description                              | Default       |
| ------------------ | ---------------------------------------- | ------------- |
| `credentials_file` | Path to the service account key JSON     | _(required)_  |
| `project`          | GCP project ID                           | _(required)_  |
| `location`         | Bucket location                          | `us-west1`    |
| `bucket_name`      | Bucket name                              | `sample-bucket` |

## Usage

```bash
terraform init     # download the provider
terraform plan     # preview changes
terraform apply    # create / update resources
terraform destroy  # tear everything down
```

## Migrating existing state from cdktf

The old cdktf workspace stored state in `terraform.gcp.tfstate` with the bucket
at resource address `google_storage_bucket.StorageBucket`. With plain Terraform
the address is `google_storage_bucket.storage_bucket`, so a fresh `apply` would
try to create a bucket that already exists.

If the bucket is already deployed, import it into the new state instead of
recreating it:

```bash
terraform init
terraform import google_storage_bucket.storage_bucket <bucket_name>
terraform plan   # should report "No changes"
```

Once import succeeds you can delete the obsolete cdktf state files
(`terraform.gcp.tfstate*`).
