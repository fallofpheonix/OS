#!/usr/bin/env python3
import time
import json
import os

ROOT = os.path.abspath(os.path.join(os.path.dirname(__file__), '..'))
ENTROPY = os.path.join(ROOT, '14_experiments', 'ebpf_lab', 'file_activity_capture', 'entropy_engine.py')

def run_entropy_sim(n=10000):
    """Run the Python entropy engine on synthetic events and measure per-event latency."""
    import importlib.util
    spec = importlib.util.spec_from_file_location('entropy_engine', ENTROPY)
    mod = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(mod)

    # generate synthetic events
    sample = {'pid': 9999, 'comm': 'sim_writer', 'type': 'WRITE', 'path': '/tmp/f.bin', 'bytes': 1024*4}
    start = time.time()
    for i in range(n):
        sample['bytes'] = 4096
        sample['path'] = '/tmp/file_%d' % i
        _ = mod.process_telemetry_stream(json.dumps(sample))
    end = time.time()
    total = end - start
    print('Entropy sim: events=%d total_s=%.4f per_event_us=%.2f' % (n, total, (total/n)*1e6))

def run_r001_sim(event_rate=100000, duration_s=2):
    """Simulate R001 by generating events at target rate and measuring loop throughput."""
    total_events = event_rate * duration_s
    start = time.time()
    count = 0
    for i in range(total_events):
        # simulate work
        _ = {'pid': i%5000, 'comm':'sim', 'type':'WRITE', 'path':'/tmp/f%d'%i, 'bytes': 512}
        count += 1
    end = time.time()
    total = end - start
    print('R001 sim: generated=%d total_s=%.4f rate=%.0f ev/s' % (count, total, count/total if total>0 else 0))

def run_r035_sim(n=100000):
    """Simulate normalizer micro-bench by parsing JSON and normalizing a few fields."""
    def normalize(raw_json):
        e = json.loads(raw_json)
        out = {
            'ts': e.get('timestamp', time.time()),
            'pid': int(e.get('pid',0)),
            'comm': e.get('comm','').strip(),
            'type': e.get('type','').upper(),
            'path': e.get('path','')[:256]
        }
        return out

    # generate raw lines
    raws = []
    for i in range(1000):
        raws.append(json.dumps({'pid': i, 'comm': 'proc%d'%i, 'type':'WRITE', 'path':'/tmp/f%d'%i, 'timestamp': time.time()}))

    start = time.time()
    for i in range(n):
        raw = raws[i % len(raws)]
        _ = normalize(raw)
    end = time.time()
    total = end - start
    print('R035 sim: calls=%d total_s=%.4f per_call_us=%.2f rate=%.0f ev/s' % (n, total, (total/n)*1e6, n/total if total>0 else 0))

if __name__ == '__main__':
    print('Running lightweight simulated benchmarks (non-root)')
    run_entropy_sim(20000)
    run_r001_sim(event_rate=50000, duration_s=1)
    run_r035_sim(n=200000)
