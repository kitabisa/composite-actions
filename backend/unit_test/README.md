<!-- action-docs-description -->
## Description

Build and push docker image
<!-- action-docs-description -->

<!-- action-docs-inputs -->
## Inputs

| parameter | description                                                                   | required | default |
| --- |-------------------------------------------------------------------------------| --- | --- |
| custom_command_flag | custom command flag to determine whether use default unit test command or not | `false` | false |
| custom_command | command before running the unit test                                          | `false` | echo no command |
| unit_test_command | actual unit test command                                                      | `false` | go test ./internal/... -coverpkg=./... -coverprofile=coverage.out -covermode=atomic |
<!-- action-docs-inputs -->

<!-- action-docs-outputs -->

<!-- action-docs-outputs -->

<!-- action-docs-runs -->
## Runs

This action is a `composite` action.
<!-- action-docs-runs -->

## Example from Caller
```
uses: kitabisa/composite-actions/backend/unit_test@44cea52512bc5b9c2de21a34401140c479644c82
with:
  gh_user: ${{ secrets.GH_USER }}
  gh_token: ${{ secrets.GH_TOKEN }}
  custom_command_flag: true
  cache_key_suffix: "kulonuwun"
  default_unit_test: false
  custom_command: "cp ./params/.env.example ./params/.env && cp ./params/firebase.json.sample ./params/firebase.json"
  unit_test_command: "./... -json -coverpkg=./... -coverprofile=coverage.out -covermode=atomic"
```
