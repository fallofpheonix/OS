#!/usr/bin/env bash
set -euo pipefail
ROOT_DIR=$(cd "$(dirname "$0")" && pwd)
OUT=${ROOT_DIR}/artifacts/validate
mkdir -p ${OUT}
python3 - <<'PY'
import json, sys
replay = json.load(open('artifacts/replay/replay.json'))
ok = True
for r in replay:
    # simple validation: required fields exist
    for f in ('timestamp','pid','comm','type','path'):
        if f not in r:
            print('MISSING', f, 'in', r)
            ok = False
open('${OUT}/validation.json','w').write(json.dumps({'pass': ok}))
print('VALIDATION_PASS' if ok else 'VALIDATION_FAIL')
PY
