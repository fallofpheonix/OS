import os
import re

def main():
    root = '.'
    go_files = []
    for dirpath, dirnames, filenames in os.walk(root):
        # Include all files to catch illegal imports from research/archive
        for f in filenames:
            if f.endswith('.go'):
                go_files.append(os.path.join(dirpath, f))

    illegal_paths = []
    # (From, To) pairs that are illegal
    rules = [
        ('telemetry', 'warden'),
        ('ai', 'warden'),
        ('advisor', 'containment'),
        ('memory', 'containment'),
        ('replay', 'containment'),
        ('telemetry', 'recovery'),
        ('quantum', 'runtime'),
        ('research', 'core'),
        ('archive', 'runtime')
    ]

    dependencies = []
    for path in go_files:
        rel_path = os.path.relpath(path, root)
        try:
            with open(path, 'r', encoding='utf-8', errors='ignore') as f:
                content = f.read()
            
            # Simple keyword matching for source/target in path and imports
            pkg_match = re.search(r'^package\s+([a-zA-Z0-9_]+)', content, re.M)
            pkg = pkg_match.group(1) if pkg_match else "unknown"
            
            imports = re.findall(r'"([^"]+)"', content)
            
            source_tags = []
            if '09_telemetry' in rel_path or 'telemetry' in pkg: source_tags.append('telemetry')
            if '06_ai' in rel_path or 'ai' in pkg: source_tags.append('ai')
            if '06_research' in rel_path: source_tags.append('research')
            if '15_archive' in rel_path: source_tags.append('archive')
            if 'quantum' in rel_path.lower(): source_tags.append('quantum')
            if 'advisor' in rel_path.lower(): source_tags.append('advisor')
            if 'memory' in rel_path.lower(): source_tags.append('memory')
            if 'replay' in rel_path.lower(): source_tags.append('replay')

            for imp in imports:
                target_tags = []
                if 'warden' in imp: target_tags.append('warden')
                if 'containment' in imp: target_tags.append('containment')
                if 'recovery' in imp: target_tags.append('recovery')
                if 'runtime' in imp: target_tags.append('runtime')
                if 'phoenix_os' in imp and not any(x in imp for x in ['common', 'bus', 'ledger']): target_tags.append('core')

                for st in source_tags:
                    for tt in target_tags:
                        for s_rule, t_rule in rules:
                            if st == s_rule and tt == t_rule:
                                illegal_paths.append({
                                    'file': rel_path,
                                    'source': st,
                                    'target': tt,
                                    'import': imp
                                })
        except:
            pass

    with open('ILLEGAL_DEPENDENCY_MATRIX.md', 'w') as f:
        f.write("# Illegal Dependency Matrix\n\n")
        f.write("| File | Source Role | Illegal Target | Actual Import |\n")
        f.write("|---|---|---|---|\n")
        if not illegal_paths:
            f.write("| None | N/A | N/A | N/A |\n")
        else:
            for p in illegal_paths:
                f.write(f"| {p['file']} | {p['source']} | {p['target']} | {p['import']} |\n")

if __name__ == '__main__':
    main()
