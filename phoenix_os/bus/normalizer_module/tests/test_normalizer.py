import json
import time
import sys
import os

ROOT = os.path.dirname(os.path.dirname(__file__))
SCHEMA = os.path.join(ROOT, 'schema.json')
sys.path.insert(0, os.path.join(ROOT, 'src'))

from normalizer import normalize_event

def test_basic_normalization():
    evt = {'pid': 42, 'comm': 'tester', 'type': 'write', 'path': '/tmp/x', 'bytes': 1024}
    out = normalize_event(json.dumps(evt))
    o = json.loads(out)
    assert o['pid'] == 42
    assert o['type'] == 'WRITE'
    assert 'timestamp' in o

def test_schema_required_fields():
    evt = {'pid': 1, 'comm': 'c'}
    out = normalize_event(json.dumps(evt))
    o = json.loads(out)
    assert 'path' in o

if __name__ == '__main__':
    test_basic_normalization()
    test_schema_required_fields()
    print('OK')
