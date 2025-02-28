# How to use

usage: |-
  ```yaml
    name: GHCR Auth
    on: [push]

    jobs:
      build:
        runs-on: ubuntu-latest
        permissions:
            contents: read    # required this to read package GCHR
            packages: write
        steps:
          - name: GHCR Auth
            uses: kitabisa/composite-actions/infra/tools/ghcr-credential@v2
  ```
