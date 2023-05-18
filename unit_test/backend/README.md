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
| cache | enable cache for go modules | `false` | true |
| cache_dependency_path | go sum location | `false` | go.sum |
| go_version | go version | `false` | ^1.13.1 |
| need_mockery | need mockery | `false` | false |
| custom_command_flag | custom command flag | `false` | false |
| custom_command | custom command | `false` | echo no command |
| unit_test_command | unit test command | `false` | go test ./internal/... -coverpkg=./... -coverprofile=coverage.out -covermode=atomic |
<!-- action-docs-inputs -->

<!-- action-docs-outputs -->

<!-- action-docs-outputs -->

<!-- action-docs-runs -->
## Runs

This action is a `composite` action.
<!-- action-docs-runs -->
