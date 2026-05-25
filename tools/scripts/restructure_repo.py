import os
import shutil

def main():
    # Mandated Doc Structure
    doc_layers = [
        '02_docs/00_governance',
        '02_docs/01_architecture',
        '02_docs/02_runtime',
        '02_docs/03_validation',
        '02_docs/04_operations',
        '02_docs/05_security',
        '02_docs/06_memory',
        '02_docs/07_distributed'
    ]
    for d in doc_layers:
        os.makedirs(d, exist_ok=True)
    
    # Mandated Research Structure
    research_layers = [
        '06_research/quantum/accepted',
        '06_research/quantum/experimental',
        '06_research/quantum/rejected',
        '06_research/control_theory',
        '06_research/replay_math',
        '06_research/decision_systems',
        '06_research/simulation',
        '06_research/formal_methods',
        '06_research/optimization'
    ]
    for d in research_layers:
        os.makedirs(d, exist_ok=True)

    # 1. Update ROADMAP.md
    roadmap_content = """# PhoenixOS Master Roadmap

## Status: F0 CONDITIONALLY COMPLETE / F1 PREPARED

### Phases
- **F0 Foundation**: COMPLETE (Conditionally - Pending Proofs)
- **F1 Runtime Hardening**: IN PROGRESS
- **F2 Formal Proof**: PLANNED
- **F3 OS Primitives**: PLANNED
- **F4 Memory**: PLANNED
- **F5 Advisor**: PLANNED
- **F6 Distributed**: PLANNED
- **F7 Cognitive Layer**: PLANNED

### Parallel Research Tracks
- **R1 Replay Math**
- **R2 Control Theory**
- **R3 Optimization**
- **R4 Simulation**
- **R5 Formal Systems**

### Current Priorities
1. Replay identity close (100% precision achieved)
2. Cleanup docs (Active)
3. Merge research (Active)
4. Build F1 runtime (Locked)
5. Evaluate quantum usefulness (Active)
6. Reject noise (Active)
"""
    with open('02_docs/00_governance/ROADMAP.md', 'w') as f:
        f.write(roadmap_content)

    # 2. Update DOC_INDEX.md
    with open('02_docs/DOC_INDEX.md', 'w') as f:
        f.write("# Document Index\n\n")
        for layer in doc_layers:
            f.write(f"## {layer}\n")
            if os.path.exists(layer):
                files = os.listdir(layer)
                for file in files:
                    f.write(f"- {file}\n")
            f.write("\n")

    # 3. Update RESEARCH_INDEX.md
    with open('06_research/RESEARCH_INDEX.md', 'w') as f:
        f.write("# Research Index\n\n")
        f.write("Keep only: Kalman, Bayesian, Monte Carlo, MCTS, Constraint solvers, Control theory, TLA+, Replay mathematics.\n\n")
        for layer in research_layers:
            f.write(f"## {layer}\n")
            if os.path.exists(layer):
                files = os.listdir(layer)
                for file in files:
                    f.write(f"- {file}\n")
            f.write("\n")

    print("Restructuring complete.")

if __name__ == '__main__':
    main()
