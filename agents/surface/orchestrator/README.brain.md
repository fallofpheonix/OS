Minimal Brain Prototype
======================

This tiny prototype exposes the deterministic `MinimalBrain` for quick experimentation.

CLI
---
Run locally:

```bash
PYTHONPATH=$(pwd) python agents/surface/orchestrator/brain_cli.py --signal telemetry-spike --score 0.92 --evidence 3
```

API
---
Start the FastAPI app and POST to `/brain/evaluate` (the service is small and safe):

```bash
PYTHONPATH=$(pwd) uvicorn agents.surface.orchestrator.api:app --reload --port 8001
curl -s -X POST "http://localhost:8001/brain/evaluate" -H "Content-Type: application/json" -d '{"signal":"telemetry-spike","score":0.92,"evidence_count":3}' | jq
```

Notes
-----
- The brain is intentionally deterministic and shadow-mode by default.
- Use this as a harness to prototype RAG, memory, and policy gates incrementally.
