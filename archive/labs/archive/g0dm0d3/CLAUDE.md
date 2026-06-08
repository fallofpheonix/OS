# G0DM0D3 — LLM Oracle

## Agent Skills
### Issue Tracker
GitHub issue tracker. See `docs/agents/issue-tracker.md`.

### Triage Labels
Default triage label vocabulary. See `docs/agents/triage-labels.md`.

### Domain Docs
Multi-context layout. See `docs/agents/domain.md`.

## Build & Test
```bash
# Single-file app — no build required
open index.html

# API server
cd api && npm install && npm start
```

## Architecture
G0DM0D3 is the LLM Oracle that provides strategic directives to the PhoenixOS AI orchestrator. It implements ULTRAPLINIAN multi-model evaluation and GODMODE CLASSIC parallel racing.

## Key Components
- **index.html** — Complete single-file application
- **api/** — Optional API server (Node.js/Express)
- **HF/** — Hugging Face integration
- **research/** — Evaluation scripts

## Invariants
- API key stored in browser only
- No server-side data storage
- All responses must be validated before actuation
- Privacy-first design
