<!-- action-docs-description -->
## Description

Destroy release deployment
<!-- action-docs-description -->

<!-- action-docs-inputs -->
## Inputs

| parameter | description | required | default |
| --- | --- | --- | --- |
| working_directory | Set working directory | `false` | . |
| credentials_json | GCP credentials services account | `true` |  |
| gke_cluster_name | GKE cluster name | `true` |  |
| gke_cluster_zone | GKE cluster location zone | `true` |  |
| gcr_host | GCP container registry host | `true` |  |
| project_id | GCP project id | `true` |  |
| prune_all | make prune-all | `false` | false |
| using_cdn | Setup cdn for static assets | `false` | false |
| delete_oldest_cdn | Setup delete oldest cdn | `false` | true |
| cdn_aws_s3_bucket | CDN aws s3 bucket | `false` |  |
| cdn_aws_access_key_id | CDN aws access key id | `false` |  |
| cdn_aws_secret_access_key | CDN aws secret access key | `false` |  |
| cdn_aws_region | CDN aws region | `false` |  |
<!-- action-docs-inputs -->

<!-- action-docs-outputs -->

<!-- action-docs-outputs -->

<!-- action-docs-runs -->
## Runs

This action is a `composite` action.
<!-- action-docs-runs -->
