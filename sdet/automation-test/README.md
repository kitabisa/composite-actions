<!-- action-docs-description -->
## Description

SDET composite setup and run automation test
<!-- action-docs-description -->

<!-- action-docs-inputs -->
## Inputs

| parameter | description | required | default |
| --- | --- | --- | --- |
| gh_token | gh token | `true` |  |
| profile | profile | `false` |  |
| platform | platform | `false` |  |
| github_run_number | github run number | `false` |  |
| setup_pnpm | setup pnpm | `false` |  |
| setup_bun | setup bun | `false` |  |
| setup_gh_cli | setup github cli | `false` |  |
| custom_command_packages | custom command packages | `false` |  |
| custom_command_setup | custom command setup | `true` |  |
| restore_cache_summary | restore cache summary | `false` |  |
| custom_command_run | custom command run | `false` |  |
| create_test_run | create test run | `false` |  |
| custom_command_report_folder | custom command report folder | `false` |  |
| custom_command_report_subfolder | custom command report subfolder | `false` |  |
| insert_testrails | insert testrails | `false` |  |
| deploy_gh_pages | dispatch deploy to gh pages | `false` |  |
| custom_command_send_report_to_slack | custom command send report to slack | `false` |  |
| custom_command_send_report_to_dashboard | custom command send report to dashboard | `false` |  |
| cache_summary | cache summary | `false` |  |
<!-- action-docs-inputs -->

<!-- action-docs-runs -->
## Runs

This action is a `composite` action.
<!-- action-docs-runs -->
