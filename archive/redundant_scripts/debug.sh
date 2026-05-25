#!/usr/bin/env bash
set -euo pipefail
ROOT_DIR=$(cd "$(dirname "$0")" && pwd)
python3 - <<'PY'
import json,sys
from debug.debug_helper import dump_event, timeline_note
from src.normalizer import normalize_event
data = json.load(open('replay/sample_events.json'))
for e in data:
    dump_event(e)
    timeline_note('before normalize')
    print(normalize_event(json.dumps(e)))
    timeline_note('after normalize')
PY
