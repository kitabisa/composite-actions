# Github Action Composite
This is main of Composite GitHub Actions.
- build
- deploy
- rollback
- destroy (dev release)

# Basic Usage
You can choose 2 option strategy.

Option 1: Single Job
```
jobs:
  build-push-deploy:
    runs-on: k8s-runner

    steps:
      - name: Build
        uses: kitabisa/composite-actions/backend/build@v2
        with:
          gh_user: ${{ secrets.GH_USER }}
          gh_token: ${{ secrets.GH_TOKEN }}
          .....................
          <more input parameter>

      - name: Deploy
        uses: kitabisa/composite-actions/backend/deploy@v2
        with:
          env: ${{ env.ENV }}
          gh_token: ${{ secrets.GH_TOKEN }}
          project_id: ${{ secrets.GCP_PROJECT_ID_PROD }}
          .....................
          <more input parameter>
```


Option 2: Multiple Job
```
jobs:
  build-push:
    runs-on: k8s-runner
    steps:
      - name: Run build & push
        uses: kitabisa/composite-actions/frontend/build@v2
        with:
          project_id: ${{ secrets.GCP_PROJECT_ID_PROD }}
          credentials_json: ${{ secrets.GCP_SA_KEY_PROD }}
          gcr_host: ${{ secrets.GCR_HOST }}
          .....................
          <more input parameter>

  deploy:
    runs-on: k8s-runner
    needs: build-push
    steps:
      - name: Run deploy & prune
        uses: kitabisa/composite-actions/frontend/deploy@v2
        with:
          project_id: ${{ secrets.GCP_PROJECT_ID_PROD }}
          credentials_json: ${{ secrets.GCP_SA_KEY_PROD }}
          gcr_host: ${{ secrets.GCR_HOST }}
          .....................
          <more input parameter>
```
