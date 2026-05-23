import json
import os
import hmac
import hashlib
import socket
import threading
from typing import Optional
from pathlib import Path


class GuardDaemon:
    """Minimal prototype for Guard Daemon: validates messages and appends to ledger."""

    REQUIRED_KEYS = {"id", "time", "source", "confidence", "policy", "action"}

    def __init__(self, ledger_path: Optional[str] = None):
        self.ledger_path = ledger_path or os.environ.get("GUARD_LEDGER_PATH") or "guard_evidence.log"
        # signing key for evidence; default is development key
        self.signing_key = os.environ.get("GUARD_SIGNING_KEY", "dev-guard-key").encode("utf-8")

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
        # sign the entry using HMAC-SHA256 and append signature field
        sig = hmac.new(self.signing_key, json.dumps(entry, separators=(",", ":")).encode("utf-8"), hashlib.sha256).hexdigest()
        entry["signature"] = sig
        with open(self.ledger_path, "a", encoding="utf-8") as f:
            f.write(json.dumps(entry, separators=(",", ":")) + "\n")
        return True


class GuardIPCServer:
    """Simple Unix-domain socket server to receive JSON messages and hand them to GuardDaemon."""

    def __init__(self, socket_path: str, ledger_path: Optional[str] = None):
        self.socket_path = socket_path
        self.daemon = GuardDaemon(ledger_path)
        self._sock = None
        self._thread = None
        self._stop = threading.Event()

    def start(self):
        # ensure parent dir exists
        sockfile = Path(self.socket_path)
        if sockfile.exists():
            sockfile.unlink()
        parent = sockfile.parent
        parent.mkdir(parents=True, exist_ok=True)
        self._sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
        self._sock.bind(self.socket_path)
        self._sock.listen(1)
        self._thread = threading.Thread(target=self._serve_loop, daemon=True)
        self._thread.start()

    def _serve_loop(self):
        while not self._stop.is_set():
            try:
                conn, _ = self._sock.accept()
            except OSError:
                break
            with conn:
                data = conn.recv(65536)
                if not data:
                    continue
                try:
                    msg = json.loads(data.decode("utf-8"))
                except Exception:
                    conn.send(b"ERROR: invalid json")
                    continue
                ok = self.daemon.process_message(msg)
                conn.send(b"OK" if ok else b"INVALID")

    def stop(self):
        self._stop.set()
        try:
            if self._sock:
                self._sock.close()
        except Exception:
            pass
        # cleanup socket file
        try:
            Path(self.socket_path).unlink()
        except Exception:
            pass
