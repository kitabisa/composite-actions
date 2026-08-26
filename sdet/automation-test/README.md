<!-- action-docs-description -->

## Description

SDET composite setup and run automation test

<!-- action-docs-description -->

<!-- action-docs-inputs -->

## Inputs

| parameter                                     | description                                   | required | default |
| --------------------------------------------- | --------------------------------------------- | -------- | ------- |
| gh_token                                      | gh token                                      | `true`   |         |
| profile                                       | profile                                       | `false`  |         |
| platform                                      | platform                                      | `false`  |         |
| setup_pnpm                                    | setup pnpm                                    | `false`  |         |
| setup_bun                                     | setup bun                                     | `false`  |         |
| bun_version                                   | bun version                                   | `false`  | latest  |
| setup_gh_cli                                  | setup github cli                              | `false`  |         |
| custom_command_packages                       | custom command packages                       | `false`  |         |
| custom_command_setup                          | custom command setup                          | `true`   |         |
| custom_command_run                            | custom command run                            | `false`  |         |
| create_test_run                               | create test run                               | `false`  |         |
| custom_command_post_run                       | custom command post run                       | `false`  |         |
| custom_command_report_folder                  | custom command report folder                  | `false`  |         |
| custom_command_report_subfolder               | custom command report subfolder               | `false`  |         |
| insert_testiny_flag                           | insert testiny                                | `false`  |         |
| deploy_gh_pages                               | dispatch deploy to gh pages                   | `false`  |         |
| custom_command_send_report_to_slack_flag      | custom command send report to slack flag      | `false`  |         |
| custom_command_send_report_to_slack           | custom command send report to slack           | `false`  |         |
| custom_command_send_report_to_dashboard_flag  | custom command send report to dashboard flag  | `false`  |         |
| custom_command_send_report_to_dashboard       | custom command send report to dashboard       | `false`  |         |
| custom_command_send_report_to_analyze_flag    | custom command send report to analyze flag    | `false`  |         |
| custom_command_send_report_to_analyze         | custom command send report to analyze         | `false`  |         |
| custom_command_send_result_to_pr_release_flag | custom command send result to pr release flag | `false`  |         |
| custom_command_send_result_to_pr_release      | custom command send result to pr release      | `false`  |         |
| app_id                                        | GitHub App ID for generating token to create verified commits | `false`  |         |
| app_private_key                               | GitHub App private key for generating token to create verified commits | `false`  |         |

<!-- action-docs-inputs -->

<!-- action-docs-runs -->

## Runs

This action is a `composite` action.

<!-- action-docs-runs -->

## Triggering kresek report analysis

Use `custom_command_send_report_to_analyze_flag` / `custom_command_send_report_to_analyze` to
trigger kresek's report analysis (e.g. `npx test-kit upload-report ...`). This runs as its own
"send report to analyze" step, with the same execution semantics as the dashboard step
(`if: always()`, `continue-on-error: true`), so it fires regardless of test outcome and never
fails the job on its own:

```yaml
- uses: kitabisa/composite-actions/sdet/automation-test@v2
  with:
    # ...other inputs
    custom_command_send_report_to_analyze_flag: "true"
    custom_command_send_report_to_analyze: |
      npx test-kit upload-report --run-id "${{ github.run_id }}"
```

**Decision:** add a dedicated input pair rather than reusing
`custom_command_send_report_to_dashboard`. Multiple consuming repos need to trigger the kresek
upload independently of the dashboard step, so folding it into the dashboard command would have
forced those repos to either duplicate the dashboard step's command elsewhere or couple two
unrelated concerns behind one flag. A dedicated pair keeps the two steps independently
enable-able and follows the existing `custom_command_send_report_to_*` flag+command convention.
