<!-- action-docs-description -->
## Description

Frontend composite deployment
<!-- action-docs-description -->

<!-- action-docs-inputs -->
## Inputs

| parameter | description | required | default |
| --- | --- | --- | --- |
| install | Run make install | `false` | true |
| config | Run make config | `false` | true |
| build | Run make build | `false` | true |
| package | Run make package | `false` | true |
| cache | Run action cache | `false` | true |
| project_id | GCP project id | `true` |  |
| credentials_json | GCP credentials services account | `true` |  |
| gcr_host | GCP container registry host | `true` |  |
| setup_pnpm | Setup pnpm and cache modules | `false` | false |
| setup_yarn | Setup yarn and cache modules | `false` | false |
| chartmuseum_host | ChartMuseum host | `true` |  |
| chartmuseum_user | ChartMuseum user | `true` |  |
| chartmuseum_pass | ChartMuseum password | `true` |  |
| working_directory | Set working directory | `false` | . |
| pnpm_version | Setup pnpm version | `false` | latest |
| using_nextjs | Setup next.js cache modules | `false` | false |
| using_cdn | Setup cdn for static assets | `false` | false |
| delete_oldest_cdn | Setup delete oldest cdn | `false` | true |
| cdn_aws_s3_bucket | CDN aws s3 bucket | `false` |  |
| cdn_aws_access_key_id | CDN aws access key id | `false` |  |
| cdn_aws_secret_access_key | CDN aws secret access key | `false` |  |
| cdn_aws_region | CDN aws region | `false` |  |
| cdn_source_dir | CDN source dir | `false` |  |
| cdn_dest_dir | CDN dest dir | `false` |  |
| cdn_keep_latest_version | CDN keep latest version length | `false` | 2 |
<!-- action-docs-inputs -->

<!-- action-docs-outputs -->
## Outputs

| parameter | description |
| --- | --- |
| build-time | Define build time |
<!-- action-docs-outputs -->

<!-- action-docs-runs -->
## Runs

This action is a `composite` action.
<!-- action-docs-runs -->
