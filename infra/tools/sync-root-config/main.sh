#!/bin/bash

set -euo pipefail

# Set app name and config path
APP_NAME=$(make get-app-name)
CONFIG_SOURCE="../.infra/secrets/${ENV}"
CONFIG_TARGET="infrastructure/kubernetes/apps/${APP_NAME}/${ENV}/common"

# Set branch target
if [[ "$ENV" == "prod" ]]; then
  BRANCH_TARGET="${APP_NAME}-image-updates"
else
  BRANCH_TARGET="main"
fi

# Clone infra repo
git clone --depth=1 --branch main "$INFRA_REPO_URL" tmp-root
cd tmp-root

# Set git config user
git config user.name "github-actions[bot]"
git config user.email "41898282+github-actions[bot]@users.noreply.github.com"

# Switch to target branch (use existing remote branch if available, otherwise branch off main)
if [[ "$BRANCH_TARGET" != "main" ]]; then
  if git ls-remote --exit-code --heads origin "$BRANCH_TARGET" >/dev/null 2>&1; then
    git fetch --depth=1 origin "$BRANCH_TARGET":"$BRANCH_TARGET"
    git checkout "$BRANCH_TARGET"
  else
    git checkout -b "$BRANCH_TARGET"
  fi
fi

# Sync to root repo
mkdir -p "$CONFIG_TARGET"
rsync -a --delete "$CONFIG_SOURCE"/ "$CONFIG_TARGET"/

# Commit & push to root repo
git add "$CONFIG_TARGET"
if git diff --cached --quiet; then
  echo "No config changes, skipping push"
else
  git commit -m "chore(${APP_NAME}): sync config ${ENV} from app repository @ ${GITHUB_SHA}"
  git push origin "$BRANCH_TARGET"
fi

# Remove temporary repo
cd ..
rm -rf tmp-root
