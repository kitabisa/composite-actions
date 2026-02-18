<!-- action-docs-description -->
## Description

PR Opened composite action for validating PR branch naming convention, checking JIRA/GitHub tickets in PR body and title, and auto-updating PR body and title with ticket information.

Features:
- Validates branch name against configurable regex pattern
- Checks for allowed branch prefixes (feat/, fix/, docs/, etc.)
- Detects JIRA or GitHub issue tickets from branch name
- Auto-updates PR body with ticket link (JIRA or GitHub)
- Auto-updates PR title with ticket ID
- Comments on PR if branch name is incorrect
<!-- action-docs-description -->

<!-- action-docs-inputs -->
## Inputs

| parameter | description | required | default |
| --- | --- | --- | --- |
| branch_regex | Regex pattern for branch name validation | `false` | `^([a-z]+)[\/]+(([a-z\-]+)-)?([0-9]{1,5})(\.[a-z]+\|(-[a-z0-9\-]+)?)$` |
| allowed_prefixes | Comma-separated list of allowed branch prefixes | `false` | `feat/,fix/,docs/,chore/,ci/,test/,refactor/` |
| ignore_branches | Branches to ignore from validation | `false` | `master,main` |
| jira_base_url | Base URL for JIRA ticket links | `false` | `https://kitabisa.atlassian.net/browse` |
| auto_update_body | Automatically update PR body with ticket link | `false` | `true` |
| auto_update_title | Automatically update PR title with ticket ID | `false` | `true` |
| gh_token | GitHub token for API operations | `true` | - |
| check_branch_name | Enable branch naming convention validation | `false` | `true` |
| github_issue_keyword | Keyword for GitHub issue linking (e.g., Resolve, Fix, Close) | `false` | `Resolve` |
<!-- action-docs-inputs -->

<!-- action-docs-outputs -->
## Outputs

| output | description |
| --- | --- |
| branch_name | The current branch name |
| ticket_type | Type of ticket detected (JIRA or GITHUB) |
| ticket_id | The detected ticket ID |
| body_updated | Whether PR body was updated |
| title_updated | Whether PR title was updated |
<!-- action-docs-outputs -->

<!-- action-docs-runs -->
## Runs

This action is a `composite` action.
<!-- action-docs-runs -->

## Usage

### Basic Usage

```yaml
name: PR Opened

on:
  pull_request:
    types: [opened, reopened, ready_for_review, synchronize]

jobs:
  pr-check:
    runs-on: ktbs-small-runner
    if: github.event.pull_request.draft == false && github.actor != 'renovate[bot]'

    steps:
      - name: PR Check
        uses: kitabisa/composite-actions/sdet/tools/pr-check@main
        with:
          gh_token: ${{ secrets.GH_TOKEN }}
```

### Custom Configuration

```yaml
name: PR Opened

on:
  pull_request:
    types: [opened, reopened, ready_for_review, synchronize]

jobs:
  pr-check:
    runs-on: ktbs-small-runner
    if: github.event.pull_request.draft == false && github.actor != 'renovate[bot]'

    steps:
      - name: PR Check
        uses: kitabisa/composite-actions/sdet/tools/pr-check@main
        with:
          gh_token: ${{ secrets.GH_TOKEN }}
          branch_regex: '^([a-z]+)[\/]+(([a-z\-]+)-)?([0-9]{1,5})(\.[a-z]+|(-[a-z0-9\-]+)?)$'
          allowed_prefixes: 'feat/,fix/,docs/,chore/,ci/,test/,refactor/,hotfix/'
          ignore_branches: 'master,main,develop'
          jira_base_url: 'https://your-org.atlassian.net/browse'
          auto_update_body: 'true'
          auto_update_title: 'true'
          check_branch_name: 'true'
          github_issue_keyword: 'Fix'
```

### Disable Auto-Updates

```yaml
- name: PR Check (validation only)
  uses: kitabisa/composite-actions/sdet/tools/pr-check@main
  with:
    gh_token: ${{ secrets.GH_TOKEN }}
    github_issue_keyword: 'Close'
    auto_update_body: 'false'
    auto_update_title: 'false'
```

## Branch Naming Convention

The action expects branch names to follow this pattern:

- `feat/SDET-123-description` (JIRA ticket)
- `fix/123-description` (GitHub issue)

### Supported Prefixes (default)
- `feat/` - New features
- `fix/` - Bug fixes
- `docs/` - Documentation changes
- `chore/` - Maintenance tasks
- `ci/` - CI/CD changes
- `test/` - Test additions/modifications
- `refactor/` - Code refactoring

## What Happens

1. **Branch Validation**: Validates the branch name against the configured regex and allowed prefixes
2. **Ticket Detection**: Extracts JIRA or GitHub issue ticket from branch name
3. **PR Body Check**: Checks if PR body already contains the ticket link
4. **PR Title Check**: Checks if PR title already contains the ticket ID
5. **Auto-Update Body**: Adds ticket section to PR body if missing
6. **Auto-Update Title**: Appends ticket ID to PR title if missing

### PR Body Format

For JIRA tickets:
```markdown
## :tickets: JIRA Ticket
[SDET-123](https://kitabisa.atlassian.net/browse/SDET-123)
```

For GitHub issues (example with `github_issue_keyword: 'Resolve'`):
```markdown
## :tickets: Github Ticket
Resolve #123
```

> Note: Set `github_issue_keyword` to your preferred keyword. Supported: `Close`, `Closes`, `Closed`, `Fix`, `Fixes`, `Fixed`, `Resolve`, `Resolves`, `Resolved`. These will auto-close the linked issue when PR is merged.
