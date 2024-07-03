<!-- action-docs-description -->
## Description

SDET composite build and run automation test
<!-- action-docs-description -->

<!-- action-docs-inputs -->
## Inputs

| parameter | description | required | default |
| --- | --- | --- | --- |
| gh_token | gh token | `true` |  |
| profile | profile | `false` |  |
| slack_webhook_url | slack webhook url | `false` |  |
| slack_webhook_debug_url | slack webhook debug url | `false` |  |
| testrail_url | testrail url | `false` |  |
| testrail_username | testrail username | `false` |  |
| testrail_password | testrail password | `false` |  |
| platform | platform | `false` |  |
| setup_pnpm | setup pnpm | `false` |  |
| setup_bun | setup bun | `false` |  |
| custom_command_packages | custom command packages | `false` |  |
| custom_command_run | custom command run | `false` |  |
| custom_command_create_test_run | custom command create test run | `false` |  |
| custom_command_setup_rsync | custom command setup rsync | `false` |  |
| custom_command_report_folder | custom command report folder | `false` |  |
| custom_command_insert_testrails | custom command insert testrails | `false` |  |
| custom_command_send_report_to_slack | custom command send report to slack | `false` |  |
<!-- action-docs-inputs -->

<!-- action-docs-runs -->
## Runs

This action is a `composite` action.
<!-- action-docs-runs -->
