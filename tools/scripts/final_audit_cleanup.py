import os
import shutil

def main():
    # 1. Quantum Evaluation (Audit Prompt Logic)
    quantum_root = '06_research/quantum_os'
    accepted = '06_research/quantum/accepted'
    experimental = '06_research/quantum/experimental'
    rejected = '06_research/quantum/rejected'
    
    remove_log = []
    
    if os.path.exists(quantum_root):
        for dirpath, dirnames, filenames in os.walk(quantum_root):
            for f in filenames:
                path = os.path.join(dirpath, f)
                content_lower = ""
                try:
                    with open(path, 'r', errors='ignore') as f_obj:
                        content_lower = f_obj.read().lower()
                except:
                    continue
                
                # Rule 2: Reject Noise
                reject_keywords = ['consciousness', 'agi', 'tutorial', 'intro', 'generic']
                if any(k in content_lower for k in reject_keywords):
                    dest = os.path.join(rejected, f)
                    if not os.path.exists(dest):
                        shutil.copy2(path, dest)
                    remove_log.append(f"Rejected Noise: {path}")
                    continue
                
                # Rule 3: Keep only if improves core
                core_improve = ['kalman', 'bayesian', 'monte carlo', 'search optimization', 'constraint solving', 'simulation', 'risk estimation']
                if any(ci in content_lower for ci in core_improve):
                    dest = os.path.join(accepted, f)
                    if not os.path.exists(dest):
                        shutil.copy2(path, dest)
                else:
                    dest = os.path.join(experimental, f)
                    if not os.path.exists(dest):
                        shutil.copy2(path, dest)

    # 2. Doc Cleanup & Restoration (Cleanup Prompt Logic)
    # Target directories
    gov = '02_docs/00_governance'
    arch = '02_docs/01_architecture'
    rnt = '02_docs/02_runtime'
    val = '02_docs/03_validation'
    
    # Restructure specific files
    mappings = {
        'F0_MASTER.md': os.path.join(gov, 'F0_MASTER.md'),
        'MASTER_ROADMAP.md': os.path.join(gov, 'ROADMAP.md'),
        'F0_EXIT_STATUS.md': os.path.join(gov, 'F0_EXIT_STATUS.md'),
        'REPLAY_IDENTITY_REPORT.md': os.path.join(val, 'REPLAY_IDENTITY_REPORT.md'),
        'DETERMINISM_REPORT.md': os.path.join(gov, 'DETERMINISM_REPORT.md'),
        'TOPOLOGY_AUDIT.md': os.path.join(arch, 'TOPOLOGY.md'),
        'RUNTIME_GRAPH.md': os.path.join(arch, 'EVENT_FLOW.md'),
        'WORKING_MODEL.md': os.path.join(arch, 'CONTAINMENT_MODEL.md')
    }
    
    for src, dst in mappings.items():
        if os.path.exists(src):
            shutil.copy2(src, dst)
            print(f"Mapped {src} -> {dst}")

    # 3. Final Logs
    with open('02_docs/00_governance/REMOVE_LOG.md', 'w') as f:
        f.write("# Remove Log\n\n")
        for log in remove_log:
            f.write(f"- {log}\n")

    print("Audit and Cleanup complete.")

if __name__ == '__main__':
    main()
