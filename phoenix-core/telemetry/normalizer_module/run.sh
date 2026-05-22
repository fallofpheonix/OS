#!/usr/bin/env bash
set -euo pipefail
ROOT_DIR=$(cd "$(dirname "$0")" && pwd)
echo "Running normalizer with sample replay"
cat ${ROOT_DIR}/replay/sample_events.json | python3 -c "import sys,json; from src.normalizer import normalize_event; data=json.load(sys.stdin); print('\n'.join([normalize_event(json.dumps(x)) for x in data]))"
