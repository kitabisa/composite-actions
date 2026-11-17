## Description

Upgrade Go dependencies that have known vulnerabilities to versions that fix them

## Inputs

| name | description | required | default |
| --- | --- | --- | --- |
| `args` | <p>Run with arguments (flags and patterns)</p> | `false` | `""` |
| `working-directory` | <p>Directory of the module to upgrade</p> | `false` | `.` |
| `create-pr` | <p>Commit the upgrades and open or update a pull request with govulnbump's report</p> | `false` | `false` |
| `token` | <p>GitHub token for create-pr, with <code>contents: write</code> and <code>pull-requests: write</code> permissions</p> | `false` | `${{ github.token }}` |


## Outputs

| name | description |
| --- | --- |
| `exit-code` | <p>govulnbump exit code: 0 if no vulnerabilities remain, 1 on error, 2 if the command line is invalid, 3 if vulnerabilities remain</p> |


## Runs

This action is a `composite` action.

## Usage

```yaml
- uses: kitabisa/composite-actions/security/govulnbump@main
  with:
    args:
    # Run with arguments (flags and patterns)
    #
    # Required: false
    # Default: ""

    working-directory:
    # Directory of the module to upgrade
    #
    # Required: false
    # Default: .

    create-pr:
    # Commit the upgrades and open or update a pull request with govulnbump's report
    #
    # Required: false
    # Default: false

    token:
    # GitHub token for create-pr, with `contents: write` and `pull-requests: write` permissions
    #
    # Required: false
    # Default: ${{ github.token }}
```
