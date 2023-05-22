<!-- action-docs-description -->
## Description

Build and push docker image
<!-- action-docs-description -->

<!-- action-docs-inputs -->
## Inputs

| parameter | description | required | default |
| --- | --- | --- | --- |
| gh_user | gh user | `true` |  |
| gh_token | gh token | `true` |  |
| cache_dependency_path | go sum location | `false` | go.sum |
| go_version | go version | `false` | ^1.13.1 |
| build | Run make build | `false` | true |
| build_push_image | build and push image for deployment | `false` | false |
| build_push_image_swagger | build and push image for swagger | `false` | false |
| build_push_image_mockoon | build and push image for mockoon | `false` |  |
| credentials_json | GCP credentials services account | `false` |  |
| project_id | GCP project id | `false` |  |
| gcr_host | GCP container registry host | `false` |  |
| swagger_script_path | swagger script path | `false` |  |
| openapi_input_file | openapi input file | `false` |  |
| openapi_output_file | openapi output file | `false` |  |
| working_directory | Set working directory | `false` | . |
<!-- action-docs-inputs -->

<!-- action-docs-outputs -->

<!-- action-docs-outputs -->

<!-- action-docs-runs -->
## Runs

This action is a `composite` action.
<!-- action-docs-runs -->
