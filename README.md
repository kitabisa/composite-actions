# Github Action Composite
This is repository of Composite GitHub Actions.

# Basic Usage
You can choose 2 option strategy.

Option 1: Single Job
See Example on Backend [Workflows](https://github.com/kitabisa/sangu/blob/0ceb5ee26f0550ab57201215a41eb79970b45b2f/.github/workflows/build-push-deploy-prod.yaml#L1-L42)
```
jobs:
  build-push-deploy:
    runs-on: ktbs-infra-k8s-runner

    steps:
      - name: Build
        uses: kitabisa/composite-actions/build/backend@v1
        with:
          gh_user: ${{ secrets.GH_USER }}
          gh_token: ${{ secrets.GH_TOKEN }}
          .....................
          <more input parameter>

      - name: Deploy
        uses: kitabisa/composite-actions/deploy/backend@v1
        with:
          env: ${{ env.ENV }}
          gh_token: ${{ secrets.GH_TOKEN }}
          project_id: ${{ secrets.GCP_PROJECT_ID_PROD }}
          .....................
          <more input parameter>
```


Option 2: Multiple Job
See Example on Frontend [Workflows](https://github.com/kitabisa/kanvas/blob/283050e921b632c9bc9dbcaab9e194f23cd9095a/.github/workflows/build-push-deploy-prod.yaml#L1-L56)
```
jobs:
  build-push:
    runs-on: ktbs-infra-k8s-runner
    steps:
      - name: Run build & push
        uses: kitabisa/composite-actions/build/frontend@main
        with:
          project_id: ${{ secrets.GCP_PROJECT_ID_PROD }}
          credentials_json: ${{ secrets.GCP_SA_KEY_PROD }}
          gcr_host: ${{ secrets.GCR_HOST }}
          .....................
          <more input parameter>

  deploy:
    runs-on: ktbs-infra-k8s-runner
    needs: build-push
    steps:
      - name: Run deploy & prune
        uses: kitabisa/composite-actions/deploy/frontend@main
        with:
          project_id: ${{ secrets.GCP_PROJECT_ID_PROD }}
          credentials_json: ${{ secrets.GCP_SA_KEY_PROD }}
          gcr_host: ${{ secrets.GCR_HOST }}
          .....................
          <more input parameter>
```
