#!/usr/bin/env bash
set -euo pipefail

# R-001 benchmark harness (safe smoke-test)
# - Generates controlled synthetic file I/O load
# - Does NOT assume a telemetry agent is running; it simulates capture counts
# - Produces JSON summary in results/

DURATION=${1:-5}
NUM_FILES=${2:-10}
OUTDIR="results/$(date +%Y%m%d_%H%M%S)"
TMPPREFIX="/tmp/r001_test_$$"
mkdir -p "$OUTDIR"

echo "Running synthetic load for ${DURATION}s with ${NUM_FILES} files..."

# Launch background writers (safe, limited)
for i in $(seq 1 "$NUM_FILES"); do
  dd if=/dev/urandom of="$TMPPREFIX.$i" bs=4k count=16 status=none &
done

sleep "$DURATION"

# Count synthetic 'events' as number of files created
EVENTS_CAPTURED=$(ls ${TMPPREFIX}.* 2>/dev/null | wc -l || true)

# Simple simulated metrics (placeholders until telemetry agent integrated)
CPU_OVERHEAD_PCT=0.1
EVENT_LOSS_PCT=0.0

cat > "$OUTDIR/summary.json" <<EOF
{
  "duration": $DURATION,
  "num_files_created": $EVENTS_CAPTURED,
  "events_captured": $EVENTS_CAPTURED,
  "cpu_overhead_pct": $CPU_OVERHEAD_PCT,
  "event_loss_pct": $EVENT_LOSS_PCT,
  "generated_files_prefix": "${TMPPREFIX}"
}
EOF

echo "Results written to $OUTDIR/summary.json"

# Move produced artifacts into OUTDIR for reproducibility and cleanup
mv ${TMPPREFIX}.* "$OUTDIR/" 2>/dev/null || true

echo "Run complete. Artifacts in $OUTDIR"
