#!/bin/bash

REGEX="^(feat|fix|docs|refactor|perf|test|build|ci|chore)\/[A-Za-z0-9._-]+$|release-please--.*"

if [[ -z "${BRANCH}" ]]; then
    BRANCH=$(git rev-parse --abbrev-ref HEAD)
else
    BRANCH="${BRANCH}"
fi

echo $BRANCH

if ! [[ $BRANCH =~ $REGEX ]]; then
  echo "Your commit was rejected due to branching name"
  echo "Please rename your branch with '(feat|fix|docs|refactor|perf|test|build|ci|chore)/lowercase-name' syntax"
  echo "Or you can use better-branch tools"
  exit 1
fi
