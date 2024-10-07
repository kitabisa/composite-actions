<!-- action-docs-description -->
## Description

This action is responsible for building and pushing a Docker image. It allows for customization of the build process through various input parameters, enabling flexibility in how the Docker image is constructed and tested. The action supports both default and custom commands, making it adaptable to different project requirements.
<!-- action-docs-description -->

<!-- action-docs-inputs -->
## Inputs

| parameter | description | required | default |
| --- | --- | --- | --- |
| working_directory | Set the working directory for the action. | `false` | `.` |
| gh_user | GitHub username required for authentication. | `true` | N/A |
| gh_token | GitHub token required for authentication. | `true` | N/A |
| go_version | Specify the Go version to use. | `false` | `~1.22` |
| custom_command_flag | A flag to determine if a custom command should be executed. When set to `true`, the `custom_command` will be executed instead of the default behavior. | `false` | `false` |
| custom_command | The custom command to execute if `custom_command_flag` is set to `true`. This can be used to override the default build or test commands. | `false` | `echo no command` |
| unit_test_command | The command used to run unit tests. This command should cover all necessary packages and generate a coverage report. The default command is tailored for Go projects. | `false` | `go test ./internal/... -coverpkg=./... -coverprofile=coverage.out -covermode=atomic` |
| hash_files | Files to hash for cache key generation. | `false` | `**` |
| default_unit_test | Flag to determine if default unit tests should be run. | `false` | `true` |
| setup_go_cache | Flag to set up Go cache. | `false` | `false` |
| setup_action_cache | Flag to set up action cache. | `false` | `true` |
| cache_key_suffix | Optional cache key suffix. | `false` | N/A |
| diff_file | Path to the diff file. | `false` | `pr.diff` |
| coverage_file | Path to the coverage file. | `false` | `coverage.out` |
<!-- action-docs-inputs -->

<!-- action-docs-outputs -->
## Outputs

Currently, this action does not produce any outputs. Future versions may include outputs such as the Docker image ID or a URL to the pushed image.
<!-- action-docs-outputs -->

<!-- action-docs-runs -->
## Runs

This action is a `composite` action, meaning it is composed of multiple steps that are executed in sequence. It leverages the flexibility of composite actions to allow for complex workflows involving Docker image creation and testing. The action can be customized to fit specific needs by using the provided input parameters.
<!-- action-docs-runs -->

## Prerequisites

- Ensure Docker is installed and running on the host machine where this action is executed.
- The necessary permissions to push images to the Docker registry must be configured.
- Ensure Go is installed if using Go-specific commands.

## Examples

To use this action with a custom command:

```yaml
- name: Build and Push Docker Image
  uses: ./path/to/action
  with:
    custom_command_flag: true
    custom_command: "your-custom-command"


