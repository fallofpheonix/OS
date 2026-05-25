import subprocess

labels = [
    ("docs-cleanup", "Cleanup of redundant or outdated documentation", "ededed"),
    ("docs-merge", "Merging of duplicate documentation files", "ededed"),
    ("research-review", "Review of research modules or notes", "ededed"),
    ("duplicate-doc", "Highly similar documentation files detected", "ededed"),
    ("runtime-validation", "Validation of runtime behavior and axioms", "ededed"),
    ("test-missing", "Missing test coverage for critical paths", "ededed"),
    ("architecture-conflict", "Architectural misalignment or violations", "ededed"),
    ("quantum-review", "Review of quantum-related research/logic", "ededed"),
    ("archive-candidate", "File or folder recommended for archival", "ededed"),
    ("security-blocker", "Security issues blocking further progress", "d73a4a"),
    ("replay-validation", "Validation of deterministic replayability", "ededed"),
    ("determinism", "Determinism-related issues or improvements", "ededed"),
    ("technical-debt", "Technical debt needing resolution", "ededed"),
    ("future-phase", "Tasks deferred to future maturity phases", "ededed"),
    ("integration", "Tasks related to external repository integration", "0052cc"),
    ("experiments", "Experimental research and prototypes", "ededed"),
    ("F0", "Phase 0 stabilization tasks", "0052cc")
]

for name, desc, color in labels:
    cmd = [
        "gh", "label", "create", name,
        "--description", desc,
        "--color", color
    ]
    try:
        subprocess.run(cmd, check=True, capture_output=True, text=True)
        print(f"Created label: {name}")
    except subprocess.CalledProcessError as e:
        if "already exists" in e.stderr:
            print(f"Label already exists: {name}")
        else:
            print(f"Failed to create label {name}: {e.stderr}")
