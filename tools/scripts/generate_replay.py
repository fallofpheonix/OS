#!/usr/bin/env python3
"""Generate a large replay dataset of synthetic telemetry events."""
import json
import random
from datetime import datetime

OUT = '09_telemetry/replay_events_large.jsonl'

def gen_event(i):
    pid = 1000 + (i % 50000)
    evt = {
        'timestamp': datetime.utcnow().isoformat() + 'Z',
        'event_id': f'e{i:012d}',
        'category': 'filesystem' if random.random() > 0.4 else 'process',
        'event_type': 'write' if random.random() > 0.4 else 'execve',
        'host_id': 'macos-dev-host',
        'pid': pid,
        'ppid': pid - 1,
        'uid': 501,
        'gid': 20,
        'comm': 'gpg' if random.random() > 0.9 else 'sh',
        'exe_path': '/usr/bin/gpg' if random.random() > 0.9 else '/bin/sh',
        'payload': {
            'file_path': f'/tmp/f{i}.bin',
            'bytes_transferred': 4096 if random.random() > 0.9 else 512,
            'entropy_score': 7.85 if random.random() > 0.95 else 3.2
        }
    }
    return evt

def main():
    N = 200000
    with open(OUT, 'w') as f:
        for i in range(N):
            e = gen_event(i)
            f.write(json.dumps(e) + '\n')
    print('WROTE', OUT)

if __name__ == '__main__':
    main()
