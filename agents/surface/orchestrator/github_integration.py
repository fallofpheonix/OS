"""GitHub integration helpers for the Surface Orchestrator.

This module provides functions to create GitHub issues from workflow tasks.

Requirements: set `GITHUB_TOKEN` in the environment before running.
"""
from __future__ import annotations

import os
import re
import subprocess
from typing import Dict, List, Optional

import requests

GITHUB_API = "https://api.github.com"


def _get_repo_from_git() -> Optional[str]:
    try:
        out = subprocess.check_output(["git", "remote", "get-url", "origin"]).decode().strip()
        # supports ssh and https urls
        m = re.search(r"github.com[:/](.+?)(?:\.git)?$", out)
        if m:
            return m.group(1)
    except Exception:
        return None


def create_issue(owner_repo: str, title: str, body: str, labels: Optional[List[str]] = None) -> Dict:
    token = os.environ.get("GITHUB_TOKEN")
    if not token:
        raise RuntimeError("GITHUB_TOKEN not set in environment")
    url = f"{GITHUB_API}/repos/{owner_repo}/issues"
    headers = {"Authorization": f"token {token}", "Accept": "application/vnd.github.v3+json"}
    payload = {"title": title, "body": body}
    if labels:
        payload["labels"] = labels
    resp = requests.post(url, json=payload, headers=headers, timeout=15)
    resp.raise_for_status()
    return resp.json()


def create_issues_for_workflow(service, wf_id: str, state_filter: Optional[List[str]] = None, labels: Optional[List[str]] = None) -> List[Dict]:
    """Create GitHub issues for tasks in a workflow matching `state_filter`.

    `service` should be an `OrchestratorService` instance.
    """
    owner_repo = _get_repo_from_git()
    if not owner_repo:
        raise RuntimeError("could not determine GitHub repo from git remote; set remote 'origin'")

    wf = service.orch.create_workflow(wf_id)
    created = []
    for t in wf.tasks.values():
        if state_filter and t.state.value not in state_filter:
            continue
        title = f"[{wf_id}] {t.title}"
        body_lines = [f"Task ID: {t.id}", f"State: {t.state.value}", f"Type: {t.task_type}", "", "Metadata:"]
        for k, v in (t.metadata or {}).items():
            body_lines.append(f"- {k}: {v}")
        body = "\n".join(body_lines)
        issue = create_issue(owner_repo, title, body, labels=labels)
        created.append(issue)
    return created


def create_pull_request(owner_repo: str, title: str, head: str, base: str = "main", body: Optional[str] = None) -> Dict:
    """Create a pull request using the GitHub API. Requires GITHUB_TOKEN in env."""
    token = os.environ.get("GITHUB_TOKEN")
    if not token:
        raise RuntimeError("GITHUB_TOKEN not set in environment")
    url = f"{GITHUB_API}/repos/{owner_repo}/pulls"
    headers = {"Authorization": f"token {token}", "Accept": "application/vnd.github.v3+json"}
    payload = {"title": title, "head": head, "base": base}
    if body:
        payload["body"] = body
    resp = requests.post(url, json=payload, headers=headers, timeout=15)
    resp.raise_for_status()
    return resp.json()
