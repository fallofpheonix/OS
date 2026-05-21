"""Minimal brain bootstrap for SentinelOS prototyping.

This keeps the first runnable slice intentionally small: score a telemetry
sample, choose a coarse action, and emit a trace that can be debugged.
"""
from __future__ import annotations

from dataclasses import dataclass


@dataclass(frozen=True)
class BrainSample:
    """Input sample for the minimal brain setup."""

    signal: str
    score: float
    evidence_count: int = 0
    mode: str = "shadow"


@dataclass(frozen=True)
class BrainDecision:
    """Decision produced by the minimal brain setup."""

    action: str
    risk: str
    explanation: str
    trace: tuple[str, ...]


class MinimalBrain:
    """Tiny deterministic brain used for the first runnable prototype."""

    def __init__(self, investigate_threshold: float = 0.4, escalate_threshold: float = 0.8):
        self.investigate_threshold = investigate_threshold
        self.escalate_threshold = escalate_threshold

    def evaluate(self, sample: BrainSample) -> BrainDecision:
        if sample.score >= self.escalate_threshold:
            action = "escalate"
            risk = "high"
        elif sample.score >= self.investigate_threshold:
            action = "investigate"
            risk = "medium"
        else:
            action = "observe"
            risk = "low"

        explanation = (
            f"signal={sample.signal} score={sample.score:.2f} "
            f"evidence={sample.evidence_count} mode={sample.mode}"
        )
        trace = (
            f"thresholds investigate={self.investigate_threshold:.2f} escalate={self.escalate_threshold:.2f}",
            f"sample signal={sample.signal} score={sample.score:.2f} evidence={sample.evidence_count} mode={sample.mode}",
            f"decision action={action} risk={risk}",
        )
        return BrainDecision(action=action, risk=risk, explanation=explanation, trace=trace)

    def debug_trace(self, sample: BrainSample) -> str:
        """Render a compact debug trace for manual inspection."""

        decision = self.evaluate(sample)
        return "\n".join(decision.trace)


def build_minimal_brain() -> MinimalBrain:
    """Factory for the default minimal brain configuration."""

    return MinimalBrain()