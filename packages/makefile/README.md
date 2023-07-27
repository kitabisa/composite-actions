<!-- action-docs-description -->
## Description

Makefile build push deploy
<!-- action-docs-description -->

<!-- action-docs-inputs -->
## Inputs

| parameter | description | required | default |
| --- | --- | --- | --- |
| project_id | GCP project id | `false` |  |
| artifact_registry_host | Artifact registry host | `true` |  |
| install | make install | `false` | false |
| config | make config | `false` | false |
| build | make build | `false` | false |
| package | make package | `false` | false |
| deploy | make deploy | `false` | false |
| rollback | make rollback | `false` | false |
| rollback_release | rollback release | `false` |  |
| rollback_revision | rollback revision | `false` |  |
| prune | make prune | `false` | false |
| prune_all | make prune-all | `false` | false |
| destroy | make destroy | `false` | false |
| working_directory | Set working directory | `false` | . |
| using_cdn | Setup cdn for static assets | `false` | false |
| cdn_keep_latest_version | CDN keep latest version length | `false` | 2 |
| setup_helmfiles | Setup helmfiles | `false` | true |
<!-- action-docs-inputs -->

<!-- action-docs-outputs -->
## Outputs

| parameter | description |
| --- | --- |
| build-time | Define build time |
| cdn-upload-path | Define cdn upload path |
| cdn-delete-path | Define cdn delete path |
<!-- action-docs-outputs -->

<!-- action-docs-runs -->
## Runs

This action is a `composite` action.
<!-- action-docs-runs -->
