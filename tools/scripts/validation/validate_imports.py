import os
import re
import sys

# Architectural boundaries
BLOCKED_MODULES = ["phoenix_os/ai", "phoenix_os/memory", "phoenix_os/nexus"]
ALLOWED_IMPORTS = ["phoenix_os/contracts", "phoenix_os/truth", "phoenix_os/state", "phoenix_os/warden", "phoenix_os/arbiter", "phoenix_os/monitor"]

def validate_imports(file_path):
    with open(file_path, 'r') as f:
        content = f.read()
        
    imports = re.findall(r'import \(\n(.*?)\n\)', content, re.DOTALL)
    if not imports:
        imports = re.findall(r'import "(.*?)"', content)
    else:
        # Split multiline imports
        imports = [i.strip().strip('"') for i in imports[0].split('\n') if i.strip()]

    violations = []
    for imp in imports:
        for blocked in BLOCKED_MODULES:
            if imp.startswith(blocked):
                violations.append(f"BLOCKED: {imp}")
        
    return violations

def main():
    root_dir = "phoenix_os"
    all_violations = {}
    
    for root, dirs, files in os.walk(root_dir):
        for file in files:
            if file.endswith(".go"):
                file_path = os.path.join(root, file)
                violations = validate_imports(file_path)
                if violations:
                    all_violations[file_path] = violations

    if all_violations:
        print("IMPORT FREEZE VIOLATIONS DETECTED:")
        for file, viols in all_violations.items():
            print(f"  {file}:")
            for v in viols:
                print(f"    - {v}")
        sys.exit(1)
    else:
        print("Import check passed. No blocked module dependencies found.")

if __name__ == "__main__":
    main()
