#!/usr/bin/env python3
import json
from pathlib import Path

res_dir = Path('14_experiments/R-001-ebpf-bench/results')
# find latest run
runs = sorted([p for p in res_dir.iterdir() if p.is_dir()])
if not runs:
    print('no runs')
    raise SystemExit(1)
latest = runs[-1]
summary = json.loads((latest / 'summary.json').read_text())
raw = json.loads((latest / 'raw_events.json').read_text()) if (latest / 'raw_events.json').exists() else []
report = {
    'run_id': latest.name,
    'summary': summary,
    'raw_event_count': len(raw)
}
(latest / 'report.md').write_text('# R-001 Benchmark Report\n\n' + json.dumps(report, indent=2))
print('report written to', latest / 'report.md')
