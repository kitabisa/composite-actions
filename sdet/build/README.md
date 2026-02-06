<!-- action-docs-description -->

## Description

SDET composite build

<!-- action-docs-description -->

<!-- action-docs-inputs -->

## Inputs

| parameter               | description               | required | default |
| ----------------------- | ------------------------- | -------- | ------- |
| gh_token                | github token for checkout | `false`  |         |
| checkout                | checkout repository       | `false`  | true    |
| fetch_depth             | git fetch depth           | `false`  | 1       |
| setup_pnpm              | setup pnpm                | `false`  | false   |
| setup_gh_cli            | setup gh cli              | `false`  | false   |
| setup_bun               | setup bun                 | `false`  | false   |
| bun_version             | bun version               | `false`  | latest  |
| custom_command_packages | custom command packages   | `false`  |         |
| custom_command_setup    | custom command setup      | `false`  |         |

<!-- action-docs-inputs -->

<!-- action-docs-runs -->

## Runs

This action is a `composite` action.

<!-- action-docs-runs -->
