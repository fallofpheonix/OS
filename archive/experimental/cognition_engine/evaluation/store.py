"""Minimal metrics layer."""

from __future__ import annotations

from collections import defaultdict
from dataclasses import dataclass, field
from typing import Any


@dataclass
class MetricsStore:
    counters: dict[str, int] = field(default_factory=lambda: defaultdict(int))
    timings_ms: dict[str, list[int]] = field(default_factory=lambda: defaultdict(list))

    def increment(self, name: str, value: int = 1) -> None:
        self.counters[name] += value

    def observe_ms(self, name: str, value: int) -> None:
        self.timings_ms[name].append(value)

    def snapshot(self) -> dict[str, Any]:
        timings: dict[str, dict[str, float | int]] = {}
        for name, values in self.timings_ms.items():
            timings[name] = {
                "count": len(values),
                "avg_ms": sum(values) / len(values) if values else 0,
                "max_ms": max(values) if values else 0,
            }
        return {"counters": dict(self.counters), "timings": timings}
