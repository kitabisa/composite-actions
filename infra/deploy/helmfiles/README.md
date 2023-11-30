## Description

Deploy Releases

## Inputs

| parameter | description | required | default |
| --- | --- | --- | --- |
| project_id | GCP project id | `true` |  |
| credentials_json | GCP credentials services account | `true` |  |
| gke_cluster_name | GKE cluster name | `true` |  |
| gke_cluster_zone | GKE cluster location zone | `true` |  |
| chartmuseum_host | ChartMuseum host | `true` |  |
| chartmuseum_user | ChartMuseum user | `true` |  |
| chartmuseum_pass | ChartMuseum password | `true` |  |


## Runs

This action is a `composite` action.
