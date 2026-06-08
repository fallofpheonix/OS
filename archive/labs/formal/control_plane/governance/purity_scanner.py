"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
import os, sys, json, time

# Policy Tiers
SEVERITY_CRITICAL = "CRITICAL"
SEVERITY_ERROR = "ERROR"
SEVERITY_WARNING = "WARNING"
SEVERITY_INFO = "INFO"

# Context-Sensitive Policy Zones
ZONES = {
    "brain": {
        "path": "/Users/fallofpheonix/engineering/brain",
        "forbidden_dirs": {"node_modules": SEVERITY_CRITICAL, ".venv": SEVERITY_CRITICAL, "target": SEVERITY_ERROR},
        "forbidden_exts": {".exe": SEVERITY_CRITICAL, ".bin": SEVERITY_CRITICAL, ".so": SEVERITY_CRITICAL},
    },
    "workspace": {
        "path": "/Users/fallofpheonix/engineering/workspace",
        "forbidden_dirs": {"node_modules": SEVERITY_INFO},
    },
    "infrastructure": {
        "path": "/Users/fallofpheonix/engineering/infrastructure",
        "forbidden_dirs": {"node_modules": SEVERITY_ERROR},
    }
}

def scan_zones():
    root = "/Users/fallofpheonix/engineering"
    state_dir = os.path.join(root, "control-plane/state")
    os.makedirs(state_dir, exist_ok=True)
    
    print("--- Ecosystem Governance Audit: Machine-Readable Scanning ---")
    
    audit_results = {
        "timestamp": time.strftime("%Y-%m-%dT%H:%M:%SZ", time.gmtime()),
        "violations": []
    }
    
    for zone_name, config in ZONES.items():
        root_path = config["path"]
        if not os.path.exists(root_path): continue
            
        for root_dir, dirs, files in os.walk(root_path):
            for d in dirs:
                if d in config.get("forbidden_dirs", {}):
                    severity = config["forbidden_dirs"][d]
                    audit_results["violations"].append({
                        "zone": zone_name,
                        "type": "DIRECTORY",
                        "artifact": d,
                        "path": os.path.join(root_dir, d),
                        "severity": severity
                    })
            
            for f in files:
                _, ext = os.path.splitext(f)
                if ext in config.get("forbidden_exts", {}):
                    severity = config["forbidden_exts"][ext]
                    audit_results["violations"].append({
                        "zone": zone_name,
                        "type": "FILE",
                        "artifact": ext,
                        "path": os.path.join(root_dir, f),
                        "severity": severity
                    })

    # Emit Machine-Readable State
    state_file = os.path.join(state_dir, "governance_state.json")
    with open(state_file, "w") as f:
        json.dump(audit_results, f, indent=2)
    
    print(f"Governance state emitted to {state_file}")
    
    critical_errors = [v for v in audit_results["violations"] if v["severity"] in [SEVERITY_CRITICAL, SEVERITY_ERROR]]
    if critical_errors:
        print(f"\n!!! AUDIT FAILED: {len(critical_errors)} critical violations found. !!!")
        sys.exit(1)
        
    print("\nAudit PASSED.")

if __name__ == "__main__":
    scan_zones()
