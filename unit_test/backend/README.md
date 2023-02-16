# How to use this composite action
Example of usage:
```
name: Test Coverage
on:
  pull_request:
    branches:
      - master

jobs:
  test-coverage:
    if: ${{ !github.event.pull_request.draft }}
    runs-on: ktbs-infra-k8s-runner
    steps:
      - name: Test Coverage
        uses: kitabisa/composite-actions/unit_test/backend@878932e5ea13a1db81014c8b27635d21c69f2b96
        with:
          gh_user: ${{ secrets.GH_USER }}
          gh_token: ${{ secrets.GH_TOKEN }}
          cache: true
          cache_dependency_path: go.sum
          go_version: 1.19
          custom_command_flag: true
          custom_command: "cp ./params/.env.sample ./params/.env"
          need_mockery: true

```

Parameters that you can configure
- custom_command_flag: by default is false, set to true if you have custom command for unit test
- custom_command: custom command that you want to run for executing unit test. For example, our unit test need to use `.env` so we have to copy `.env.sample`
- need_mockery: by default is false, set to true if you need mockery for unit test (e.g you want to re-generate mock before running UT)
