import os
import re

def fix_package_conflicts(dir_path):
    # Find all go files in the directory
    files = [f for f in os.listdir(dir_path) if f.endswith('.go')]
    if not files: return
    
    # Check package names
    pkg_counts = {}
    file_pkgs = {}
    for f in files:
        path = os.path.join(dir_path, f)
        with open(path, 'r') as f_obj:
            content = f_obj.read()
            match = re.search(r'^package\s+([a-zA-Z0-9_]+)', content, re.M)
            if match:
                pkg = match.group(1)
                pkg_counts[pkg] = pkg_counts.get(pkg, 0) + 1
                file_pkgs[f] = pkg

    if len(pkg_counts) > 1:
        # Multiple packages found. 
        # Usually, tests should be in 'tests' or same package.
        # Let's align them to the one that is NOT 'tests' if possible, or 'tests' if majority.
        target_pkg = max(pkg_counts, key=pkg_counts.get)
        if target_pkg == 'tests' and len(pkg_counts) > 1:
             # Prefer the actual module name if it's there
             for p in pkg_counts:
                 if p != 'tests':
                     target_pkg = p
                     break
        
        print(f"Fixing package conflicts in {dir_path} -> aligning to {target_pkg}")
        for f, pkg in file_pkgs.items():
            if pkg != target_pkg:
                path = os.path.join(dir_path, f)
                with open(path, 'r') as f_obj:
                    content = f_obj.read()
                new_content = re.sub(r'^package\s+[a-zA-Z0-9_]+', f'package {target_pkg}', content, 1, re.M)
                with open(path, 'w') as f_obj:
                    f_obj.write(new_content)

if __name__ == '__main__':
    dirs_to_check = [
        'tests/server/arbiter_server',
        'tests/server/recovery_server',
        'tests/server/replay_server',
        'tests/server/telemetry_server',
        'tests/server/truth_server',
        'tests/server/warden_server',
        'tests/validation',
        'tests/security'
    ]
    for d in dirs_to_check:
        if os.path.exists(d):
            fix_package_conflicts(d)
