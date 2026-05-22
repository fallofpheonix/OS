import json
import time
from jsonschema import validate, ValidationError
import math

SCHEMA_PATH = 'schema.json'

def load_schema():
    with open(SCHEMA_PATH, 'r') as f:
        return json.load(f)

SCHEMA = None
try:
    SCHEMA = load_schema()
except Exception:
    SCHEMA = None

def calculate_shannon_entropy_bytes(data_bytes: bytes) -> float:
    if not data_bytes:
        return 0.0
    counts = {}
    for b in data_bytes:
        counts[b] = counts.get(b, 0) + 1
    entropy = 0.0
    L = len(data_bytes)
    for c in counts.values():
        p = c / L
        entropy -= p * math.log2(p)
    return entropy

def normalize_event(raw_json: str) -> str:
    """Normalize a raw event JSON string into canonical schema JSON string."""
    try:
        evt = json.loads(raw_json)
    except Exception:
        raise ValueError('invalid json')

    # Basic shaping and defaults
    now = time.time()
    out = {
        'timestamp': float(evt.get('timestamp', now)),
        'pid': int(evt.get('pid', 0)),
        'comm': str(evt.get('comm', ''))[:64],
        'type': str(evt.get('type', 'OTHER')).upper(),
        'path': str(evt.get('path', ''))[:1024],
        'bytes': int(evt.get('bytes', 0)),
        'entropy': None,
        'alert_level': None,
    }

    # Entropy heuristic: if bytes present and >0, simulate entropy on path bytes
    try:
        if out['bytes'] > 0:
            data = out['path'].encode('utf-8')
            out['entropy'] = round(calculate_shannon_entropy_bytes(data), 4)
            if out['entropy'] > 7.5:
                out['alert_level'] = 'CRITICAL'
            elif out['entropy'] > 6.5:
                out['alert_level'] = 'WARNING'
    except Exception:
        out['entropy'] = None

    # Schema validation if available
    if SCHEMA:
        try:
            validate(instance=out, schema=SCHEMA)
        except ValidationError as e:
            raise ValueError('schema validation failed: %s' % e)

    return json.dumps(out)

if __name__ == '__main__':
    # simple run: read lines from stdin and write normalized lines to stdout
    import sys
    for line in sys.stdin:
        try:
            print(normalize_event(line.strip()))
        except Exception as e:
            print(json.dumps({'error': str(e)}))
