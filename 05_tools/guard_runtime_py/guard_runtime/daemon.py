import json
import os
from typing import Optional


class GuardDaemon:
    """Minimal prototype for Guard Daemon: validates messages and appends to ledger."""

    REQUIRED_KEYS = {"id", "time", "source", "confidence", "policy", "action"}

    def __init__(self, ledger_path: Optional[str] = None):
        self.ledger_path = ledger_path or os.environ.get("GUARD_LEDGER_PATH") or "guard_evidence.log"

    def validate_message(self, msg: dict) -> bool:
        if not isinstance(msg, dict):
            return False
        if not self.REQUIRED_KEYS.issubset(msg.keys()):
            return False
        # basic confidence check
        try:
            c = float(msg.get("confidence", 0.0))
            if c < 0.0 or c > 1.0:
                return False
        except Exception:
            return False
        # action must be a dict with type
        action = msg.get("action")
        if not isinstance(action, dict) or "type" not in action:
            return False
        return True

    def process_message(self, msg: dict) -> bool:
        """Validate and append the message to the ledger file as JSONL.

        Returns True on success, False on validation failure.
        """
        if not self.validate_message(msg):
            return False
        # append to ledger file
        entry = {
            "trace_hash": msg.get("id"),
            "sdi": None,
            "policy": msg.get("policy"),
            "action": msg.get("action"),
            "result": "submitted",
            "time": msg.get("time"),
            "confidence": msg.get("confidence"),
            "replay": False,
            "experiment": None,
        }
        os.makedirs(os.path.dirname(self.ledger_path), exist_ok=True) if os.path.dirname(self.ledger_path) else None
        with open(self.ledger_path, "a", encoding="utf-8") as f:
            f.write(json.dumps(entry, separators=(",", ":")) + "\n")
        return True
