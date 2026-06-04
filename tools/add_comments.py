import os

header_c_style = """/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
"""

header_py_style = """\"\"\"
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
\"\"\"
"""

header_sh_style = """# PHOENIX MATRIX SOVEREIGN ARCHITECTURE
# [STATUS]: 18-Repository Substrate Consolidated
# [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
# [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
# [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""

def prepend_to_file(filepath, header):
    try:
        with open(filepath, 'r') as f:
            content = f.read()
            if "PHOENIX MATRIX SOVEREIGN ARCHITECTURE" in content:
                return # Already added
        with open(filepath, 'w') as f:
            f.write(header + content)
    except Exception as e:
        print(f"Failed to write {filepath}: {e}")

def add_dir_notes(dirpath):
    notes_path = os.path.join(dirpath, "DIRECTORY_NOTES.md")
    if not os.path.exists(notes_path):
        try:
            with open(notes_path, 'w') as f:
                f.write(f"# {os.path.basename(dirpath)} Directory Notes\n\n"
                        "- **[FUTURE ENHANCEMENT]**: Expand test coverage and semantic validations within this component.\n"
                        "- **[POTENTIAL LOOPHOLE]**: Unhandled edge cases in domain translation.\n"
                        "- **[ERROR PRONE AREA]**: High complexity logic coupling.\n")
        except Exception:
            pass

domains = ["Phoenix.Nucleus", "Phoenix.Cognition", "Phoenix.Crucible", "Phoenix.Terminus", "Phoenix.UI", "Phoenix.Arbiter"]

for domain in domains:
    if not os.path.exists(domain):
        continue
    for root, dirs, files in os.walk(domain):
        if '.git' in root:
            continue
            
        add_dir_notes(root)
        
        for file in files:
            filepath = os.path.join(root, file)
            ext = os.path.splitext(file)[1].lower()
            
            if ext in ['.go', '.c', '.h', '.js', '.ts', '.java', '.cpp', '.cs', '.gd']:
                prepend_to_file(filepath, header_c_style)
            elif ext in ['.py']:
                prepend_to_file(filepath, header_py_style)
            elif ext in ['.sh', '.yml', '.yaml'] or file == 'Makefile':
                prepend_to_file(filepath, header_sh_style)

print("Annotation complete.")
