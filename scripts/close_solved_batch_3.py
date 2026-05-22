import subprocess

# Solving issues related to hashing, endianness, and checkpoints
solved_ids = [
    972, 973, # Endianness and Cross-platform hashing (BigEndian/SHA256 used)
    1001, 1000 # Checkpoint hashing and verifier (Checkpoint method added)
]

for issue_id in solved_ids:
    subprocess.run(["gh", "issue", "close", str(issue_id), "--comment", "Resolved by hardening serialization with BigEndian and implementing deterministic Checkpoint hashing in the Ledger."])
