# Runtime Risk Matrix

## Severity: CRITICAL
| Risk | Description | Impact |
| --- | --- | --- |
| Python Version Drift | Root env is 3.14 (Experimental), Astraeus-core is 3.12. | Potential runtime incompatibilities and bytecode issues. |

## Severity: HIGH
| Risk | Description | Impact |
| --- | --- | --- |
| Duplicate Packages | `chromadb`, `pydantic`, `fastapi` installed in multiple envs. | Increased disk usage, inconsistent behavior across envs. |
| Shadow Runtimes | Repos like `workspace` claim to be `VENV` but lack envs. | Broken automation or deployment scripts. |

## Severity: MEDIUM
| Risk | Description | Impact |
| --- | --- | --- |
| Archive Sprawl | Dozens of `requirements.txt` and `pyproject.toml` in `archive/`. | Clutters search results and complicates dependency analysis. |
| Unmanaged Root Env | Root `.venv` exists without a clear `pyproject.toml` or `requirements.txt`. | "Zombie" environment; source of truth is unclear. |

## Severity: LOW
| Risk | Description | Impact |
| --- | --- | --- |
| False Positive Venvs | Mypy typeshed folders detected as venvs. | Noise in analysis tools. |
