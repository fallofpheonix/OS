import json
import time

def dump_event(evt):
    print('RAW EVENT:', json.dumps(evt))

def timeline_note(note):
    print('TIMELINE:', time.time(), note)
