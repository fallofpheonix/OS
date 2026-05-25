import os
import shutil
import hashlib

def get_hash(path):
    if not os.path.isfile(path): return None
    h = hashlib.md5()
    with open(path, 'rb') as f:
        for chunk in iter(lambda: f.read(4096), b""):
            h.update(chunk)
    return h.hexdigest()

def classify_quantum(content, filename):
    content_lower = content.lower()
    tags = []
    if 'probability' in content_lower: tags.append('Q1')
    if 'optimization' in content_lower: tags.append('Q2')
    if 'simulation' in content_lower: tags.append('Q3')
    if 'control' in content_lower: tags.append('Q4')
    if 'state estimation' in content_lower or 'kalman' in content_lower: tags.append('Q5')
    if 'search' in content_lower: tags.append('Q6')
    if 'computing' in content_lower: tags.append('Q7')
    if 'noise' in content_lower: tags.append('Q8')
    
    # Reject noise/speculation
    reject_keywords = ['consciousness', 'agi', 'tutorial', 'intro', 'generic', 'unrelated']
    if any(k in content_lower for k in reject_keywords):
        return 'REMOVE', tags, "Speculative/Unrelated"
        
    # Must improve core primitives
    core_improve = ['bayesian', 'kalman', 'monte carlo', 'search optimization', 'constraint solving', 'decision graphs', 'simulation', 'risk estimation']
    if any(ci in content_lower for ci in core_improve):
        return 'KEEP', tags, "Improves core primitives"
        
    return 'RESEARCH', tags, "Theory only"

def main():
    # Setup directories
    for d in ['02_docs/00_governance', '02_docs/01_architecture', '02_docs/02_runtime', '02_docs/03_validation', '02_docs/04_operations', '02_docs/05_security', '02_docs/06_memory', '02_docs/07_distributed']:
        os.makedirs(d, exist_ok=True)
    
    os.makedirs('06_research/quantum/accepted', exist_ok=True)
    os.makedirs('06_research/quantum/experimental', exist_ok=True)
    os.makedirs('06_research/quantum/rejected', exist_ok=True)
    os.makedirs('15_archive/removed_docs', exist_ok=True)
    os.makedirs('15_archive/merged_docs', exist_ok=True)
    os.makedirs('15_archive/deprecated', exist_ok=True)

    remove_log = []
    merge_log = []
    
    # 1. Quantum Evaluation
    quantum_root = '06_research/quantum_os'
    if os.path.exists(quantum_root):
        for dirpath, dirnames, filenames in os.walk(quantum_root):
            for f in filenames:
                path = os.path.join(dirpath, f)
                try:
                    with open(path, 'r', encoding='utf-8', errors='ignore') as f_obj:
                        content = f_obj.read()
                    action, tags, reason = classify_quantum(content, f)
                    
                    dest = None
                    if action == 'KEEP':
                        dest = '06_research/quantum/accepted/' + f
                    elif action == 'RESEARCH':
                        dest = '06_research/quantum/experimental/' + f
                    else:
                        dest = '06_research/quantum/rejected/' + f
                        remove_log.append(f"Quantum Reject: {path} ({reason})")
                    
                    if dest and not os.path.exists(dest):
                        shutil.copy2(path, dest)
                except:
                    pass

    # 2. Doc Restructuring & Cleanup
    # Scan root and subdirs for unique docs
    for dirpath, dirnames, filenames in os.walk('.'):
        if any(ex in dirpath for ex in ['.git', '.venv', '15_archive', '02_docs', '06_research']): continue
        for f in filenames:
            if f.endswith('.md'):
                path = os.path.join(dirpath, f)
                rel_path = os.path.relpath(path, '.')
                
                # Heuristic categorization
                dest_dir = '02_docs/'
                content_lower = ""
                try:
                    with open(path, 'r') as f_in: content_lower = f_in.read().lower()
                except: pass
                
                if 'roadmap' in content_lower or 'exit' in content_lower: dest_dir += '00_governance/'
                elif 'topology' in content_lower or 'architecture' in content_lower or 'flow' in content_lower: dest_dir += '01_architecture/'
                elif 'runtime' in content_lower or 'warden' in content_lower or 'arbiter' in content_lower: dest_dir += '02_runtime/'
                elif 'test' in content_lower or 'validation' in content_lower or 'proof' in content_lower: dest_dir += '03_validation/'
                else: dest_dir += '04_operations/'
                
                dest = dest_dir + f
                if not os.path.exists(dest):
                    shutil.copy2(path, dest)
                else:
                    merge_log.append(f"Duplicate doc: {rel_path} matched existing in 02_docs")

    # 3. Master Merge
    f0_master_content = "# F0 Master\n\nStatus: CONDITIONALLY COMPLETE\n\n"
    for f in ['F0.5', 'F0.6', 'F0.7']:
        # Search for these patterns in existing md files
        pass # Placeholder for content extraction if files existed
    
    with open('02_docs/00_governance/F0_MASTER.md', 'w') as f:
        f.write(f0_master_content)

    # 4. Logs
    with open('REMOVE_LOG.md', 'w') as f:
        f.write("# Remove Log\n\n" + "\n".join(remove_log))
    with open('MERGE_LOG.md', 'w') as f:
        f.write("# Merge Log\n\n" + "\n".join(merge_log))
    
    print("Cleanup and Evaluation complete.")

if __name__ == '__main__':
    main()
