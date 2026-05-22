"""Lightweight GH CLI helpers for issue and label management.

Uses the local `gh` CLI for operations so we don't need to manage tokens here.
"""
from __future__ import annotations

import json
import subprocess
from typing import List, Optional


def run_cmd(args: List[str]) -> str:
    out = subprocess.check_output(args)
    return out.decode().strip()


def ensure_label(repo: str, name: str, color: str = "ededed", description: str = "") -> None:
    # create label if missing
    try:
        run_cmd(["gh", "label", "view", name, "--repo", repo, "--json", "name"])
    except subprocess.CalledProcessError:
        args = ["gh", "label", "create", name, "--color", color, "--description", description, "--repo", repo]
        run_cmd(args)


def add_labels(repo: str, issue_number: int, labels: List[str]) -> None:
    for lbl in labels:
        run_cmd(["gh", "issue", "edit", str(issue_number), "--add-label", lbl, "--repo", repo])


def remove_labels(repo: str, issue_number: int, labels: List[str]) -> None:
    for lbl in labels:
        run_cmd(["gh", "issue", "edit", str(issue_number), "--remove-label", lbl, "--repo", repo])


def create_issue_cli(repo: str, title: str, body: str, labels: Optional[List[str]] = None) -> int:
    args = ["gh", "issue", "create", "--title", title, "--body", body, "--repo", repo, "--json", "number"]
    if labels:
        for label in labels:
            args.extend(["--label", label])
    out = run_cmd(args)
    try:
        obj = json.loads(out)
        return int(obj["number"])
    except Exception:
        # fallback: try to parse number from output
        return int(out.strip().split()[-1])
