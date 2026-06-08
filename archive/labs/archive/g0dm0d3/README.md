# G0DM0D3 — LLM Oracle for PhoenixOS

> Multi-model chat interface and ULTRAPLINIAN evaluation engine for cognitive queries.

## Overview

G0DM0D3 is the LLM Oracle that provides strategic directives to the PhoenixOS AI orchestrator. It implements the ULTRAPLINIAN multi-model evaluation engine, GODMODE CLASSIC parallel racing, and Parseltongue red-teaming capabilities.

**All Oracle responses must be validated before triggering actuation.**

## Features

- **50+ Models** via OpenRouter (Claude, GPT-5, Gemini, Grok, etc.)
- **GODMODE CLASSIC** — 5 battle-tested prompt + model combos racing in parallel
- **ULTRAPLINIAN** — Multi-model evaluation across 5 tiers (10-55 models)
- **Parseltongue** — Input perturbation engine with 33 techniques
- **AutoTune** — Context-adaptive sampling parameters
- **STM Modules** — Real-time output normalization
- **Privacy-First** — API key in browser only, no PII tracking

## Project Structure

```
G0DM0D3/
├── index.html          # Complete single-file application
├── api/                # Optional API server (Node.js/Express)
│   ├── server.ts       # Express server
│   └── routes/         # API routes
├── HF/                 # Hugging Face integration
├── research/           # Evaluation scripts
├── paper/              # Research paper
├── public/             # Static assets
├── API.md              # API documentation
├── PAPER.md            # Research paper
├── SECURITY.md         # Security policy
└── TERMS.md            # Terms of service
```

## Quick Start

### Self-Host
```bash
# Clone
git clone https://github.com/fallofpheonix/G0DM0D3.git
cd G0DM0D3

# Open in browser
open index.html

# Or serve locally
python3 -m http.server 8000
```

### API Server
```bash
cd api
npm install
npm start
```

## PhoenixOS Integration

The Oracle is queried by `PhoenixMind/intelligence/nexus_bridge.go`:
- Endpoint: `http://127.0.0.1:7860/v1/chat/completions`
- Model: `ultraplinian/standard`
- Format: OpenAI-compatible
- Response: JSON directive with command, confidence, reasoning, graph_proof

## Security

- API key stored in browser localStorage only
- No server-side data storage
- Opt-out telemetry
- AGPL-3.0 license

## License

AGPL-3.0 — Forever free, irrevocably open.
