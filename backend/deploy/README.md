<!-- action-docs-description -->
## Description

Build and push docker image
<!-- action-docs-description -->

<!-- action-docs-inputs -->
## Inputs

| parameter | description | required | default |
| --- | --- | --- | --- |
| env | environment to deploy to | `true` | dev |
| gh_token | gh token | `true` |  |
| project_id | GCP project id | `true` |  |
| gcr_host | GCP container registry host | `true` |  |
| chartmuseum_host | ChartMuseum host | `true` |  |
| chartmuseum_user | ChartMuseum user | `true` |  |
| chartmuseum_pass | ChartMuseum password | `true` |  |
| gke_cluster_name | GKE cluster name | `true` |  |
| gke_cluster_zone | GKE cluster location zone | `true` |  |
| rancher_host | Rancher host | `true` |  |
| rancher_access_key | Rancher access key | `true` |  |
| rancher_secret_key | Rancher secret key | `true` |  |
| rancher_cluster_id | Rancher cluster id | `true` |  |
| working_directory | Set working directory | `false` | . |
| setup_helmfiles | Setup helmfiles | `false` | false |
<!-- action-docs-inputs -->

<!-- action-docs-outputs -->

<!-- action-docs-outputs -->

<!-- action-docs-runs -->
## Runs

This action is a `composite` action.
<!-- action-docs-runs -->
