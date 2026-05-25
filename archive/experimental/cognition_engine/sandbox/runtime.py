"""Docker sandbox runtime."""

from __future__ import annotations

import subprocess
from pathlib import Path
from typing import Any

from .policies import DockerPolicy


class DockerSandbox:
    def __init__(self, policy: DockerPolicy | None = None):
        self.policy = policy or DockerPolicy()

    def run_python(self, code: str, workdir: Path) -> dict[str, Any]:
        workdir.mkdir(parents=True, exist_ok=True)
        script = workdir / "snippet.py"
        script.write_text(code, encoding="utf-8")
        cmd = [
            "docker",
            "run",
            "--rm",
            "--network",
            self.policy.network,
            "--memory",
            self.policy.memory,
            "--cpus",
            self.policy.cpus,
        ]
        if self.policy.read_only:
            cmd.append("--read-only")
        cmd += [
            "-v",
            f"{workdir.resolve()}:/work:rw",
            "-w",
            "/work",
            "python:3.12-slim",
            "python",
            "snippet.py",
        ]
        try:
            result = subprocess.run(
                cmd,
                capture_output=True,
                text=True,
                timeout=self.policy.timeout_s,
                check=False,
            )
            return {
                "ok": result.returncode == 0,
                "return_code": result.returncode,
                "stdout": result.stdout,
                "stderr": result.stderr,
            }
        except FileNotFoundError:
            return {"ok": False, "return_code": 127, "stdout": "", "stderr": "docker not found"}
        except subprocess.TimeoutExpired:
            return {"ok": False, "return_code": 124, "stdout": "", "stderr": "sandbox timeout"}
