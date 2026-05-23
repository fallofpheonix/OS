# Environment Management

Purpose
-------
Prevents drift by specifying pinned dependencies, environment isolation, and reproducibility guarantees.

Structure
---------
- Dependency Pinning: lockfiles and hashes
- Python Versions: supported interpreters
- Model Versions: model checksums and provenance
- Environment Isolation: containers and virtualenvs
- Virtual Environments: recommended layouts and activation
- Reproducibility: validation and CI checks

Next steps
----------
- Create `bootstrap.sh` and CI checks that verify environment reproducibility.
