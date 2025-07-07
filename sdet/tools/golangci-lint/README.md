<!-- action-docs-description -->
## Description

Run golangci-lint with optional caching and Go setup
<!-- action-docs-description -->

<!-- action-docs-inputs -->
## Inputs

| parameter | description | required | default |
| --- | --- | --- | --- |
| gh_user | gh user | `true` |  |
| gh_token | gh token | `true` |  |
| setup_go | Setup Go | `false` | `true` |
| go_version | go version | `false` | `^1.21` |
| setup_go_cache | go version | `false` | `false` |
| setup_action_cache | setup cache | `false` | `true` |
| cache_key_suffix | Additional suffix for cache key | `false` |  |
| hash_files | hashFiles location | `false` | `**` |
| run_golangci_lint | Run golangci-lint | `false` | `true` |
<!-- action-docs-inputs -->

<!-- action-docs-runs -->
## Runs

This action is a `composite` action.
<!-- action-docs-runs -->
