import json
import socket
from guard_runtime.daemon import GuardDaemon


def test_process_message_success(tmp_path):
    ledger = tmp_path / "ledger.log"
    daemon = GuardDaemon(str(ledger))
    msg = {
        "id": "uuid-1234",
        "time": "2026-05-21T12:00:00Z",
        "source": "detector-1",
        "confidence": 0.92,
        "policy": "policy-xyz",
        "action": {"type": "isolate", "params": {"pid": 1234}},
    }
    ok = daemon.process_message(msg)
    assert ok is True
    # ledger should contain one JSON line with the trace_hash
    content = ledger.read_text(encoding="utf-8").strip().splitlines()
    assert len(content) == 1
    entry = json.loads(content[0])
    assert entry["trace_hash"] == "uuid-1234"


def test_process_message_validation_fail(tmp_path):
    ledger = tmp_path / "ledger.log"
    daemon = GuardDaemon(str(ledger))
    msg = {"bad": "message"}
    ok = daemon.process_message(msg)
    assert ok is False
    assert not ledger.exists()


def test_ipc_server_accepts_and_signs(tmp_path):
    # Use a short path under /tmp to avoid AF_UNIX path length limits
    short_name = tmp_path.name[:8]
    sock = f"/tmp/guard_{short_name}.sock"
    ledger = str(tmp_path / "ledger.log")
    server = None
    try:
        from guard_runtime.daemon import GuardIPCServer

        server = GuardIPCServer(sock, ledger_path=ledger)
        server.start()
        # send a valid message
        msg = {
            "id": "uuid-999",
            "time": "2026-05-21T12:00:00Z",
            "source": "detector-ipc",
            "confidence": 0.5,
            "policy": "policy-ipc",
            "action": {"type": "throttle", "params": {"rate": 10}},
        }
        with socket.socket(socket.AF_UNIX, socket.SOCK_STREAM) as c:
            c.connect(sock)
            c.send(json.dumps(msg).encode("utf-8"))
            resp = c.recv(1024)
            assert resp == b"OK"
        # ledger should exist and contain signature
        lines = (tmp_path / "ledger.log").read_text(encoding="utf-8").strip().splitlines()
        assert len(lines) == 1
        entry = json.loads(lines[0])
        assert entry["trace_hash"] == "uuid-999"
        assert "signature" in entry and len(entry["signature"]) > 10
    finally:
        if server:
            server.stop()
