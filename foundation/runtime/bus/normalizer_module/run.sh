# PHOENIX MATRIX SOVEREIGN ARCHITECTURE
# [STATUS]: 18-Repository Substrate Consolidated
# [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
# [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
# [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
#!/usr/bin/env bash
set -euo pipefail
ROOT_DIR=$(cd "$(dirname "$0")" && pwd)
echo "Running normalizer with sample replay"
cat ${ROOT_DIR}/replay/sample_events.json | python3 -c "import sys,json; from src.normalizer import normalize_event; data=json.load(sys.stdin); print('\n'.join([normalize_event(json.dumps(x)) for x in data]))"
