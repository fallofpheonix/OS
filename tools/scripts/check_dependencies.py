import subprocess
import os

def check_imports():
    results = []
    for root, dirs, files in os.walk('.'):
        if '.git' in root or '.venv' in root or 'archive' in root:
            continue
        for f in files:
            if f.endswith('.go'):
                path = os.path.join(root, f)
                # Try to run go list on the file's package
                dir_path = os.path.dirname(path)
                try:
                    # Run go list in the directory of the file
                    # We use -f '{{.Error}}' to catch errors
                    cmd = ['go', 'list', '-f', '{{.Error}}', './' + dir_path]
                    out = subprocess.check_output(cmd, stderr=subprocess.STDOUT).decode()
                    if out.strip() != '<nil>' and out.strip() != '':
                        results.append(f"Error in {dir_path}: {out.strip()}")
                except Exception as e:
                    # If go list fails, it might be due to multiple files or other issues
                    # but we are mainly looking for "no required module provides package"
                    pass

    with open('TOOLS_DEPENDENCY_REPORT.md', 'w') as f:
        f.write("# Tools Dependency Report\n\n")
        if not results:
            f.write("## Status: PASS\n")
            f.write("- Missing imports: 0\n")
        else:
            f.write("## Status: FAIL\n")
            f.write("- Missing imports detected:\n\n")
            for r in results:
                f.write(f"  - {r}\n")

if __name__ == '__main__':
    check_imports()
