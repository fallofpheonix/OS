# GitHub Repository Normalization Report

## Overview
This report outlines the plan for normalizing GitHub repository names within the Astraeus organization.

## Protected Repositories (MANUAL_APPROVAL_REQUIRED)
- **fallofpheonix**: PRIMARY_PROFILE. No rename planned.
- **forge-agent**: CORE. No rename planned.
- **control-plane**: CORE. No rename planned.
- **astraeus-core**: CORE. No rename planned.
- **infrastructure**: CORE. No rename planned.

## Review Required
- **brain** → **brain-core**
- **runtime** → **runtime-core**
- **ecosystem_os** → **ecosystem-os**

## Scientific Stack Target Names
- **physics** → **physics-runtime**
- **mathematics** → **mathematics-engine**
- **simulation** → **simulation-runtime**
- **reinforcement-learning** → **rl-infrastructure**
- **memory-engine** → **memory-system**
- **neuromorphic-computing** → **neuromorphic-research**
- **agent-swarm** → **swarm-runtime**

## Rename Batches

### Batch P0 (Scientific Core)
- physics → physics-runtime
- mathematics → mathematics-engine
- simulation → simulation-runtime

### Batch P1 (Scientific Extended)
- reinforcement-learning → rl-infrastructure
- memory-engine → memory-system
- neuromorphic-computing → neuromorphic-research
- agent-swarm → swarm-runtime

### Batch P2 (Core Normalization)
- ecosystem_os → ecosystem-os
- runtime → runtime-core
- brain → brain-core

## Redirect & Impact Analysis
- **Redirects**: GitHub provides automatic redirects for renamed repositories. Existing `git clone` and `git fetch` commands will continue to work via the old URL, but updating to the new URL is recommended for long-term stability.
- **Clones**: Local Git remotes will need to be updated to the new URLs to avoid reliance on redirects.
- **Installs**: Any automation relying on hardcoded repository names or URLs (e.g., `uv add git+...`) will need to be updated.
- **Cache**: Local caches indexed by repo name will need to be re-indexed or moved.

## Risk Assessment
- **Risk Level**: LOW to MEDIUM.
- **Mitigation**: Automated redirects and a comprehensive reference update plan.

---
Report generated on 2026-05-19.
