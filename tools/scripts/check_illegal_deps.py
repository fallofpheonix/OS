import os
import re

def check_illegal_dependencies():
    forbidden_rules = [
        ('09_telemetry', 'warden', 'Direct Telemetry -> Warden bypass'),
        ('06_ai', 'warden', 'Direct AI -> Warden control'),
        ('replay', 'containment', 'Replay bypassing Warden/Arbiter'),
        ('06_research', 'phoenix_os', 'Research code in Runtime'),
        ('15_archive', 'phoenix_os', 'Archive code in Runtime'),
        ('quantum', 'phoenix_os', 'Quantum code in Runtime')
    ]
    
    violations = []
    
    # Regex to find imports in Go files
    import_regex = re.compile(r'import\s+(?:\(\s*(.*?)\s*\)|"([^"]+)")', re.DOTALL)
    quoted_regex = re.compile(r'"([^"]+)"')

    for root, dirs, files in os.walk('.'):
        if '.git' in root or '.venv' in root or 'archive' in root: continue
        rel_root = os.path.relpath(root, '.')
        
        for f in files:
            if f.endswith('.go'):
                path = os.path.join(root, f)
                try:
                    with open(path, 'r', encoding='utf-8', errors='ignore') as f_obj:
                        content = f_obj.read()
                        
                        # Find all import blocks
                        imports = []
                        for match in import_regex.finditer(content):
                            if match.group(1): # Group 1 is ( ... )
                                for q_match in quoted_regex.finditer(match.group(1)):
                                    imports.append(q_match.group(1))
                            elif match.group(2): # Group 2 is single "..."
                                imports.append(match.group(2))

                        for imp in imports:
                            for source, target, reason in forbidden_rules:
                                if source in rel_root and target in imp:
                                    # Double check if target is just a substring of something allowed
                                    # e.g. "phoenix_os/common" contains "phoenix_os"
                                    if target == 'phoenix_os' and any(x in imp for x in ['common', 'bus', 'ledger', 'truth', 'state']):
                                         continue
                                    
                                    violations.append({
                                        'source': path,
                                        'target': imp,
                                        'allowed': 'NO',
                                        'reason': reason,
                                        'status': 'FAIL'
                                    })
                except:
                    pass

    with open('ILLEGAL_DEPENDENCY_MATRIX.md', 'w') as f:
        f.write("# Illegal Dependency Matrix\n\n")
        f.write("| source | target | allowed | reason | status |\n")
        f.write("|---|---|---|---|---|\n")
        if not violations:
            f.write("| Core | Core | YES | Baseline | PASS |\n")
        else:
            for v in sorted(violations, key=lambda x: x['source']):
                f.write(f"| {v['source']} | {v['target']} | {v['allowed']} | {v['reason']} | {v['status']} |\n")

if __name__ == '__main__':
    check_illegal_dependencies()
