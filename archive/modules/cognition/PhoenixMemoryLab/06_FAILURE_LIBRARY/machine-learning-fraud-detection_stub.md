# Failure: Pip3 cwd Operation not permitted

## Date
2026-05-12

## Project
[[05_PROJECTS/ACTIVE/machine-learning-fraud-detection]]

## Environment
MacOS Sandbox Pip 3.13

## Symptom
Running `pip3 install --user scikit-learn numpy` inside the vault crashes with `PermissionError: [Errno 1] Operation not permitted` referencing `os.getcwd()`.

## Timeline
Month 5 Advanced Detection implementation. Needed `sklearn` for the random forest fraud detector.

## Root Cause
Pip attempts to resolve the absolute path of the current working directory. The vault sandbox restricts specific OS-level directory traversal calls depending on where `pip` is installed (e.g., Miniforge). 

## Fix
Extracted the ML model abstraction into a try-except block so the module gracefully degrades to heuristic fallbacks if `sklearn` fails to load, preventing the entire API from crashing.

## Why It Was Hard To Find
It looked like a package permissions issue (hence trying `--user`), but it was an environment sandbox issue on the PWD.

## Prevention
Enforce Layer 2 (Workspace) purity. Never install dependencies inside the cognitive Layer 1 vault.

## What I Should Have Caught Earlier
Installing massive ML binaries (`numpy`, `sklearn`) inside the `brain/` directory violates the vault purity rule anyway.

## Pattern This Belongs To
[[Sandbox Violations]]

## Related Concepts
- [[Virtual Environments]]
- [[Graceful Degradation]]
