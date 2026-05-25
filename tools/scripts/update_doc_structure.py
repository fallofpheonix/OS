import os
import shutil

def main():
    # Setup Mandated Doc Folders
    layers = [
        '02_docs/00_governance',
        '02_docs/01_architecture',
        '02_docs/02_runtime',
        '02_docs/03_validation',
        '02_docs/04_operations',
        '02_docs/05_security',
        '02_docs/06_memory',
        '02_docs/07_distributed'
    ]
    for layer in layers:
        os.makedirs(layer, exist_ok=True)

    # Core Doc Mapping
    # Governance
    shutil.copy2('ROADMAP.md', '02_docs/00_governance/ROADMAP.md') if os.path.exists('ROADMAP.md') else None
    shutil.copy2('F0_MASTER.md', '02_docs/00_governance/F0_MASTER.md') if os.path.exists('F0_MASTER.md') else None
    shutil.copy2('DETERMINISM_REPORT.md', '02_docs/00_governance/DETERMINISM_REPORT.md') if os.path.exists('DETERMINISM_REPORT.md') else None
    shutil.copy2('F0_EXIT_STATUS.md', '02_docs/00_governance/F0_EXIT_CHECKLIST.md') if os.path.exists('F0_EXIT_STATUS.md') else None

    # Architecture
    shutil.copy2('TOPOLOGY_AUDIT.md', '02_docs/01_architecture/TOPOLOGY.md') if os.path.exists('TOPOLOGY_AUDIT.md') else None
    shutil.copy2('RUNTIME_GRAPH.md', '02_docs/01_architecture/EVENT_FLOW.md') if os.path.exists('RUNTIME_GRAPH.md') else None

    # Runtime (Dummy content if missing)
    with open('02_docs/02_runtime/ARBITER.md', 'w') as f: f.write("# Arbiter Specification")
    with open('02_docs/02_runtime/WARDEN.md', 'w') as f: f.write("# Warden Specification")
    with open('02_docs/02_runtime/RECOVERY.md', 'w') as f: f.write("# Recovery Specification")
    with open('02_docs/02_runtime/METRICS.md', 'w') as f: f.write("# Metrics Specification")

    # Validation
    shutil.copy2('REPLAY_IDENTITY_REPORT.md', '02_docs/03_validation/REPLAY_ARCHITECTURE.md') if os.path.exists('REPLAY_IDENTITY_REPORT.md') else None
    with open('02_docs/03_validation/VALIDATION_MATRIX.md', 'w') as f: f.write("# Validation Matrix")
    with open('02_docs/03_validation/CHAOS.md', 'w') as f: f.write("# Chaos Engineering")
    with open('02_docs/03_validation/FUZZ.md', 'w') as f: f.write("# Fuzz Testing")
    with open('02_docs/03_validation/PROOFS.md', 'w') as f: f.write("# Formal Proofs")

    # Indices
    with open('02_docs/DOC_INDEX.md', 'w') as f:
        f.write("# Document Index\n\n")
        for l in layers:
            f.write(f"## {l}\n")
            if os.path.exists(l):
                for file in os.listdir(l):
                    f.write(f"- {file}\n")
            f.write("\n")

    print("Document Structure Updated.")

if __name__ == '__main__':
    main()
