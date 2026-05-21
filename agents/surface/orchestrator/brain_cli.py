"""CLI runner for the MinimalBrain prototype.

Usage:
  python agents/surface/orchestrator/brain_cli.py --signal telemetry-spike --score 0.92 --evidence 3
"""
from __future__ import annotations

import argparse
import json

from agents.surface.orchestrator.minimal_brain import BrainSample, build_minimal_brain


def main():
    p = argparse.ArgumentParser(description="Minimal Brain CLI")
    p.add_argument("--signal", required=True)
    p.add_argument("--score", type=float, required=True)
    p.add_argument("--evidence", type=int, default=0)
    p.add_argument("--mode", default="shadow")
    args = p.parse_args()

    brain = build_minimal_brain()
    sample = BrainSample(signal=args.signal, score=args.score, evidence_count=args.evidence, mode=args.mode)
    dec = brain.evaluate(sample)
    out = {
        "action": dec.action,
        "risk": dec.risk,
        "explanation": dec.explanation,
        "trace": list(dec.trace),
    }
    print(json.dumps(out, indent=2))


if __name__ == "__main__":
    main()
