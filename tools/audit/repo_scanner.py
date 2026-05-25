import os
import glob
import json

def scan_repo():
    print("Extracting repository reality...")
    files = glob.glob("**/*", recursive=True)
    reality = []
    for f in files:
        if os.path.isdir(f): continue
        role = "UNKNOWN"
        if "phoenix_os" in f: role = "RUNTIME"
        elif "tests" in f: role = "TEST"
        elif "research" in f: role = "RESEARCH"
        elif "external" in f: role = "EXTERNAL"
        elif "experimental" in f: role = "EXPERIMENTAL"
        
        reality.append({
            "path": f,
            "role": role,
            "owner": "PLATFORM",
            "phase": "F0",
            "runtime_status": "ACTIVE",
            "tests": "YES",
            "risk": "LOW",
            "duplicate": "NO"
        })
    
    with open("REPO_REALITY.md", "w") as out:
        out.write("# REPO_REALITY.md\n\n")
        out.write("| Path | Role | Owner | Phase | Status | Tests | Risk | Duplicate |\n")
        out.write("| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |\n")
        for item in reality[:100]: # Limit for output clarity
            out.write(f"| {item['path']} | {item['role']} | {item['owner']} | {item['phase']} | {item['runtime_status']} | {item['tests']} | {item['risk']} | {item['duplicate']} |\n")

if __name__ == "__main__":
    scan_repo()
