import math
import json
from collections import Counter

def calculate_shannon_entropy(data):
    """
    Calculates the Shannon entropy of a byte string.
    H(X) = -sum(P(x) * log2(P(x)))
    """
    if not data:
        return 0.0
    
    entropy = 0
    length = len(data)
    # Count frequency of each byte (0-255)
    counts = Counter(data)
    
    for count in counts.values():
        p_x = count / length
        entropy -= p_x * math.log2(p_x)
    
    return entropy

def process_telemetry_stream(input_line):
    """
    Processes a single line of JSON telemetry and enriches it with entropy.
    """
    try:
        event = json.loads(input_line)
        # In a real eBPF scenario, we'd have the raw buffer.
        # Here we simulate by generating a buffer based on 'bytes' or 'path'.
        # For demonstration, we'll assume the 'path' is the data for now, 
        # or use a placeholder string if bytes > 0.
        
        simulated_data = event.get('path', '').encode('utf-8')
        # If it's a WRITE, we simulate high entropy if bytes is large (placeholder logic)
        if event.get('type') == 'WRITE':
            # Simulated high entropy for demonstration if bytes > 1MB
            if event.get('bytes', 0) > 1024 * 1024:
                # Mock high entropy data
                simulated_data = b'\x00\xff\xde\xad\xbe\xef' * 100
            
        entropy = calculate_shannon_entropy(simulated_data)
        event['entropy'] = round(entropy, 4)
        
        # Tag as 'SUSPICIOUS' based on RFC-001C threshold
        if entropy > 7.5:
            event['alert_level'] = 'CRITICAL'
        elif entropy > 6.5:
            event['alert_level'] = 'WARNING'
            
        return json.dumps(event)
    except Exception as e:
        return json.dumps({"error": str(e)})

if __name__ == "__main__":
    # Test with sample data
    test_events = [
        {"pid": 1234, "comm": "log_daemon", "type": "WRITE", "path": "/var/log/syslog", "bytes": 1024},
        {"pid": 5678, "comm": "ransom_ware", "type": "WRITE", "path": "/home/user/photo.jpg.enc", "bytes": 2048576}
    ]
    
    print("Testing Entropy Engine Logic...")
    for event in test_events:
        enriched = process_telemetry_stream(json.dumps(event))
        print(enriched)
