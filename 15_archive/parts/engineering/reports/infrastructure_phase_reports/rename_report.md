# Repository Identity Normalization Report

## Overview
This report provides a comprehensive plan for normalizing the naming of repositories within the Astraeus ecosystem to align with the `domain-subsystem` standard.

## Safe Keeps (NO RENAME)
The following repositories are protected and will remain unchanged to preserve primary identity and stable core references:
- **fallofpheonix**: Primary profile identity.
- **forge-agent**: Primary execution agent.
- **control-plane**: System orchestration layer.
- **infrastructure**: Foundation logic.
- **astraeus-core**: Framework kernel.

## Recommended Renames
The following repositories are proposed for renaming to improve clarity and consistency:
- **ecosystem_os** → **ecosystem-os** (Mandatory: Standardize hyphenation)
- **runtime** → **runtime-core** (Review Required: Distinguish from other runtimes)
- **brain** → **brain-core** (Review Required: Align with cognitive core identity)
- **TrustLab** → **trust-research** (Normalize experimental name)

## Scientific Stack Names
Alignment for the P1-P7 layers:
- **physics** → **physics-runtime**
- **mathematics** → **mathematics-engine**
- **simulation** → **simulation-runtime**
- **reinforcement-learning** → **rl-infrastructure**
- **memory-engine** → **memory-system**
- **neuromorphic-computing** → **neuromorphic-research**
- **agent-swarm** → **swarm-runtime**

## Future Naming Rules
1. Pattern: `domain-subsystem`.
2. Suffixes allowed: `-core`, `-engine`, `-runtime`, `-system`, `-lab`, `-research`, `-sim`, `-agent`, `-os`, `-gateway`, `-graph`, `-registry`.
3. Forbidden: `temp`, `test`, `final`, `new`, abbreviations.

## Legacy Aliases
All renamed repositories will maintain an alias system to ensure that existing links, local paths, and remote URLs remain functional during the transition.

---
Report generated on 2026-05-19.
