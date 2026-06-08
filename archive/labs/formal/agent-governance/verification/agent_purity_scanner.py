"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
import os, sys, json, hashlib

PROTECTED_ZONES = ["control-plane/core", "infrastructure/security", "control-plane/agent-governance"]
MANIFEST_PATH = "/Users/fallofpheonix/engineering/control-plane/agent-governance/invariants/mutation_manifest.json"

def verify_invariants():
    root = "/Users/fallofpheonix/engineering"
    if not os.path.exists(MANIFEST_PATH):
        print("No mutation manifest found. Validation skipped.")
        return []

    with open(MANIFEST_PATH, "r") as f:
        manifest = json.load(f)

    violations = []
    for rel_path, expected_hash in manifest["hashes"].items():
        abs_path = os.path.join(root, rel_path)
        if not os.path.exists(abs_path):
            violations.append({"type": "missing_protected_file", "path": rel_path})
            continue
            
        with open(abs_path, "rb") as f:
            current_hash = hashlib.sha256(f.read()).hexdigest()
            if current_hash != expected_hash:
                violations.append({"type": "unauthorized_mutation", "path": rel_path})
                
    return violations


def verify_authenticity(manifest_data):
    key_path = "/Users/fallofpheonix/engineering/control-plane/agent-governance/invariants/authority.key"
    if not os.path.exists(key_path): return False
    
    with open(key_path, "rb") as f:
        key = f.read()
    
    provided_sig = manifest_data.get("signature")
    data_to_verify = {k: v for k, v in manifest_data.items() if k not in ["signature", "authenticity"]}
    
    serialized = json.dumps(data_to_verify, sort_keys=True).encode("utf-8")
    expected_sig = hmac.new(key, serialized, hashlib.sha256).hexdigest()
    
    return provided_sig == expected_sig

if __name__ == "__main__":
    print("--- AI Agent Governance Scan: Authenticity Verification ---")
    with open(MANIFEST_PATH, "r") as f:
        manifest_data = json.load(f)
    
    if not verify_authenticity(manifest_data):
        print("!!! AUTHENTICITY FAILURE: Manifest signature invalid or missing !!!")
        sys.exit(1)
    violations = verify_invariants()
    if violations:
        print(f"!!! MUTATION VIOLATION: {len(violations)} issues found !!!")
        for v in violations:
            print(f"  - {v['type']}: {v['path']}")
        sys.exit(1)
    print("Agent Governance Audit: PASSED (Immutable Invariants Verified)")
