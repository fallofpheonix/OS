#!/usr/bin/env bash
set -euo pipefail
ROOT_DIR=$(cd "$(dirname "$0")" && pwd)
OUT=${ROOT_DIR}/artifacts/replay
mkdir -p ${OUT}
python3 - <<'PY'
import json
from src.normalizer import normalize_event
data = json.load(open('replay/sample_events.json'))
out = []
for e in data:
    out.append(json.loads(normalize_event(json.dumps(e))))
open('${OUT}/replay.json','w').write(json.dumps(out, indent=2))
print('WROTE', '${OUT}/replay.json')
PY
