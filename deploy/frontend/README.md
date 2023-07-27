<!-- action-docs-description -->
## Description

Frontend composite deployment
<!-- action-docs-description -->

<!-- action-docs-inputs -->
## Inputs

| parameter | description | required | default |
| --- | --- | --- | --- |
| project_id | GCP project id | `true` |  |
| credentials_json | GCP credentials services account | `true` |  |
| artifact_registry | Artifact registry host | `true` |  |
| gke_cluster_name | GKE cluster name | `true` |  |
| gke_cluster_zone | GKE cluster location zone | `true` |  |
| rancher_host | Rancher host | `true` |  |
| rancher_access_key | Rancher access key | `true` |  |
| rancher_secret_key | Rancher secret key | `true` |  |
| rancher_cluster_id | Rancher cluster id | `true` |  |
| working_directory | Set working directory | `false` | . |
| preview_url | Set preview url | `false` |  |
| prune | Run make prune | `false` | false |
<!-- action-docs-inputs -->

<!-- action-docs-outputs -->

<!-- action-docs-outputs -->

<!-- action-docs-runs -->
## Runs

This action is a `composite` action.
<!-- action-docs-runs -->
