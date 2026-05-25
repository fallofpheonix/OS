"""Sandbox policy definitions."""

from __future__ import annotations

from dataclasses import dataclass


@dataclass(frozen=True)
class DockerPolicy:
    network: str = "none"
    memory: str = "512m"
    cpus: str = "1"
    read_only: bool = True
    timeout_s: int = 30
