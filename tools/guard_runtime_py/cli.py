"""Simple CLI for prototype guard runtime."""
import argparse
from guard_runtime.daemon import GuardDaemon
import json
import socket
from pathlib import Path


def main():
    p = argparse.ArgumentParser(prog="phoenix-guardctl")
    p.add_argument("command", choices=["status", "submit"], nargs="?", default="status")
    p.add_argument("--file", help="JSON file to submit as detection event")
    args = p.parse_args()
    daemon = GuardDaemon()
    if args.command == "status":
        print("Guard Daemon prototype — ledger:", daemon.ledger_path)
    elif args.command == "submit":
        if not args.file:
            print("Provide --file with a JSON message to submit")
            return
        with open(args.file, "r", encoding="utf-8") as f:
            msg = json.load(f)
        # if a socket path exists in env, use IPC; otherwise write locally
        sock_path = os.environ.get("GUARD_SOCKET_PATH")
        if sock_path and Path(sock_path).exists():
            try:
                with socket.socket(socket.AF_UNIX, socket.SOCK_STREAM) as s:
                    s.connect(sock_path)
                    s.send(json.dumps(msg).encode("utf-8"))
                    resp = s.recv(1024)
                    print(resp.decode("utf-8"))
                    return
            except Exception as e:
                print("IPC submit failed:", e)
        ok = daemon.process_message(msg)
        print("submitted" if ok else "validation failed")


if __name__ == "__main__":
    main()
