import os
import shutil

def documentation_collapse():
    sacred_8 = {
        'SYSTEM_IDENTITY.md': '02_docs/00_governance',
        'ROADMAP.md': '02_docs/00_governance',
        'PHASE_LOCK.md': '02_docs/00_governance',
        'STATE_MODEL.md': '02_docs/01_architecture',
        'TRUTH_MODEL.md': '02_docs/01_architecture',
        'REPLAY_SPEC.md': '02_docs/01_architecture',
        'DECISION_MODEL.md': '02_docs/02_runtime',
        'VALIDATION_RULES.md': '02_docs/03_validation'
    }
    
    archive_dir = '02_docs/09_archive'
    os.makedirs(archive_dir, exist_ok=True)
    
    # Iterate through all files in 02_docs (except the ones we just created)
    for root, dirs, files in os.walk('02_docs', topdown=False):
        # Skip the target directories themselves during walk
        if any(root.startswith(os.path.join('02_docs', d)) for d in ['00', '01', '02', '03', '04', '05', '06', '07', '08', '09']):
            continue
            
        for f in files:
            if f.endswith('.md'):
                path = os.path.join(root, f)
                if f in sacred_8:
                    target = sacred_8[f]
                    os.makedirs(target, exist_ok=True)
                    dest = os.path.join(target, f)
                    if not os.path.exists(dest):
                        shutil.move(path, dest)
                        print(f"Kept sacred doc: {path} -> {dest}")
                    else:
                        # If conflict, archive it
                        dest_archive = os.path.join(archive_dir, f + "_" + str(hash(path))[:4])
                        shutil.move(path, dest_archive)
                        print(f"Archived duplicate sacred: {path} -> {dest_archive}")
                else:
                    # Move to archive
                    dest = os.path.join(archive_dir, f)
                    if os.path.exists(dest):
                        dest = dest + "_" + str(hash(path))[:4]
                    shutil.move(path, dest)
                    print(f"Archived: {path} -> {dest}")

if __name__ == '__main__':
    documentation_collapse()
