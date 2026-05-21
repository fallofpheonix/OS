import time
import json
from normalizer import normalize_event

def bench(n=100000):
    sample = json.dumps({'pid': 100, 'comm': 'bench', 'type': 'WRITE', 'path': '/var/log/test', 'bytes': 4096})
    start = time.time()
    for i in range(n):
        normalize_event(sample)
    end = time.time()
    total = end - start
    print('bench: calls=%d total_s=%.4f per_call_us=%.2f rate=%.0f ev/s' % (n, total, (total/n)*1e6, n/total if total>0 else 0))

if __name__ == '__main__':
    bench(20000)
