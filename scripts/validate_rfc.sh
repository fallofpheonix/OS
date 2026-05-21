#!/bin/bash
# scripts/validate_rfc.sh

# Target branch (usually main)
TARGET_BRANCH=${1:-main}

echo "Comparing against $TARGET_BRANCH..."

# Check if we are in a git repo
if ! git rev-parse --is-inside-work-tree > /dev/null 2>&1; then
  echo "Not a git repository."
  exit 0
fi

# Fetch target branch if not present
git fetch origin $TARGET_BRANCH --quiet

# Get changed files
CHANGED_FILES=$(git diff --name-only origin/$TARGET_BRANCH)

RESEARCH_CHANGES=$(echo "$CHANGED_FILES" | grep "^01_research/")
RFC_CHANGES=$(echo "$CHANGED_FILES" | grep "^02_docs/rfc/")

if [ -n "$RESEARCH_CHANGES" ] && [ -z "$RFC_CHANGES" ]; then
    echo "Error: Changes detected in 01_research/ but no corresponding RFC changes in 02_docs/rfc/."
    echo "Research changes:"
    echo "$RESEARCH_CHANGES"
    exit 1
fi

if [ -n "$RESEARCH_CHANGES" ]; then
    echo "Research changes accompanied by RFC changes."
else
    echo "No research changes detected."
fi

exit 0
