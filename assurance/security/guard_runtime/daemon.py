import json
import socket
import threading
import os
import hashlib
import hmac

class GuardDaemon:
    def __init__(self, ledger_path: str):
        self.ledger_path = ledger_path

    def process_message(self, msg: dict) -> bool:
        # Validate minimal required fields
        required_fields = ["id", "time", "source", "confidence", "policy", "action"]
        if not all(field in msg for field in required_fields):
            return False

        # In a real system, we would use proper cryptographic keys.
        # For this prototype, we simulate a signature.
        secret_key = b"PHOENIX_MATRIX_INSECURE_DEFAULT_KEY"
        msg_str = json.dumps(msg, sort_keys=True)
        signature = hmac.new(secret_key, msg_str.encode('utf-8'), hashlib.sha256).hexdigest()

        entry = {
            "trace_hash": msg["id"],
            "payload": msg,
            "signature": signature
        }

        # Append to ledger
        with open(self.ledger_path, 'a') as f:
            f.write(json.dumps(entry) + '\n')
            
        return True


class GuardIPCServer:
    def __init__(self, sock_path: str, ledger_path: str):
        self.sock_path = sock_path
        self.daemon = GuardDaemon(ledger_path)
        self.server = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
        self.running = False
        self.thread = None

    def start(self):
        if os.path.exists(self.sock_path):
            os.remove(self.sock_path)
            
        self.server.bind(self.sock_path)
        self.server.listen(5)
        self.running = True
        self.thread = threading.Thread(target=self._run)
        self.thread.start()

    def _run(self):
        self.server.settimeout(1.0)
        while self.running:
            try:
                conn, _ = self.server.accept()
                self._handle_connection(conn)
            except socket.timeout:
                pass
            except Exception as e:
                if self.running:
                    print(f"IPC Server Error: {e}")

    def _handle_connection(self, conn):
        try:
            data = conn.recv(4096)
            if data:
                try:
                    msg = json.loads(data.decode('utf-8'))
                    success = self.daemon.process_message(msg)
                    if success:
                        conn.send(b"OK")
                    else:
                        conn.send(b"ERR: Invalid Format")
                except json.JSONDecodeError:
                    conn.send(b"ERR: Invalid JSON")
        finally:
            conn.close()

    def stop(self):
        self.running = False
        if self.thread:
            self.thread.join()
        self.server.close()
        if os.path.exists(self.sock_path):
            os.remove(self.sock_path)
