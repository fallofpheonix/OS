"""Test the Astraeus CLI."""

from __future__ import annotations

import subprocess
import sys


def test_cli_help():
    """Verify that the CLI can be invoked and shows help."""
    result = subprocess.run(
        [sys.executable, "-m", "cli.main", "--help"],
        capture_output=True,
        text=True,
        check=True
    )
    assert "Astraeus Execution Substrate CLI" in result.stdout
    assert "Available commands" in result.stdout
    assert "run" in result.stdout
    assert "replay" in result.stdout
    assert "verify" in result.stdout
    assert "audit" in result.stdout


def test_cli_verify():
    """Verify the 'verify' command executes (even if it finds no journal)."""
    # This might return exit code 0 or 1 depending on whether data/journal.jsonl exists.
    # We just want to see it run and output something about verification.
    result = subprocess.run(
        [sys.executable, "-m", "cli.main", "verify"],
        capture_output=True,
        text=True
    )
    assert "Verifying Substrate Integrity" in result.stdout


if __name__ == "__main__":
    test_cli_help()
    test_cli_verify()
    print("CLI tests passed!")
