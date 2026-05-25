#!/usr/bin/env python3
import os
import csv

ROOT = os.path.abspath(os.path.join(os.path.dirname(__file__), '..'))
OUT = os.path.join(ROOT, '14_experiments', 'file_action_matrix.csv')

PATTERNS = [
    ('/01_research/stage_32', 'Theoretical island', 'N/A', 'DELETE', 'Unconnected theory', 'Low', 1, ''),
    ('/06_ai/rag', 'AI RAG', 'Frozen AI advisory', 'FREEZE', 'Unvalidated AI before telemetry', 'Medium', 4, '06_ai/rag'),
    ('/10_kernel', 'Kernel prototypes', 'Deferred kernel work', 'DEFER', 'Kernel changes blocked until userspace validation', 'High', 2, '10_kernel'),
    ('/09_telemetry', 'Telemetry', 'Core telemetry', 'KEEP', 'Telemetry foundation', 'Low', 10, '09_telemetry'),
    ('/07_security', 'Security', 'Security engines', 'KEEP', 'Core security modules', 'Low', 9, '07_security'),
    ('/05_tools', 'Tools', 'Tooling', 'KEEP', 'Required tooling & benchmarks', 'Low', 7, '05_tools'),
    ('/14_experiments', 'Experiments', 'Standardize experiments', 'MOVE', 'Standardize experiment layout under R###', 'Low', 6, '14_experiments'),
    ('/02_docs/rfc/RFC-001B', 'RFC fragment', 'Merge into RFC-001', 'MERGE', 'Consolidate telemetry math into RFC-001', 'Low', 6, '02_docs/rfc/RFC-001_telemetry_schema.md'),
    ('/02_docs/rfc/RFC-001C', 'RFC fragment', 'Merge into RFC-001', 'MERGE', 'Consolidate entropy logic into RFC-001', 'Low', 6, '02_docs/rfc/RFC-001_telemetry_schema.md'),
]

def decide(path):
    for pat, current, future, action, reason, risk, score, target in PATTERNS:
        if pat in path.replace('\\', '/'):
            return (current, future, action, reason, risk, score, target)
    # Default heuristic
    if '/02_docs' in path:
        return ('Docs', 'Docs', 'KEEP', 'Documentation', 'Low', 5, '02_docs')
    if '/11_distribution' in path or '/images' in path:
        return ('Distribution', 'Distribution', 'KEEP', 'Packaging', 'Low', 4, '11_distribution')
    return ('Misc', 'Misc', 'KEEP', 'Default keep', 'Low', 3, '')

def walk_and_write():
    rows = []
    for dirpath, dirnames, filenames in os.walk(ROOT):
        # skip hidden
        rel_dir = os.path.relpath(dirpath, ROOT)
        if rel_dir.startswith('.'):
            continue
        for f in filenames:
            if f.startswith('.'):
                continue
            full = os.path.join(dirpath, f)
            rel = os.path.relpath(full, ROOT)
            current, future, action, reason, risk, score, target = decide('/' + rel)
            rows.append({
                'file_name': rel.replace(os.sep, '/'),
                'current_role': current,
                'future_role': future,
                'action': action,
                'reason': reason,
                'risk': risk,
                'dependency_score': score,
                'integration_target': target,
            })

    os.makedirs(os.path.dirname(OUT), exist_ok=True)
    with open(OUT, 'w', newline='') as csvfile:
        writer = csv.DictWriter(csvfile, fieldnames=['file_name','current_role','future_role','action','reason','risk','dependency_score','integration_target'])
        writer.writeheader()
        for r in sorted(rows, key=lambda x: (-x['dependency_score'], x['file_name'])):
            writer.writerow(r)
    print('Wrote', OUT)

if __name__ == '__main__':
    walk_and_write()
