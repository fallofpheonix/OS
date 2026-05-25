#!/bin/bash
set -euo pipefail

repo_dir="/Users/fallofpheonix/os"
cd "$repo_dir"

if [[ -f .venv/bin/activate ]]; then
  # shellcheck disable=SC1091
  source .venv/bin/activate
fi

export PYTHONPATH="$repo_dir"
python agents/surface/orchestrator/run_demo.py