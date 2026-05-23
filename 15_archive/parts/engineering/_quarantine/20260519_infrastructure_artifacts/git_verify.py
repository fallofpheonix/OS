import os
import subprocess
from pathlib import Path
import json

# --- CONFIGURATION ---
ROOT_DIR = Path("/Users/fallofpheonix/engineering")
INVENTORY_FILE = ROOT_DIR / "repo_inventory_v2.md"

def get_git_status(repo_path):
    try:
        # Check if it's a git repo
        if not (repo_path / ".git").exists():
            return "NOT_A_GIT_REPO", None
        
        # 1. git status
        status = subprocess.check_output(["git", "status", "--porcelain"], cwd=repo_path).decode().strip()
        if status:
            return "DIRTY_WORKING_TREE", status
        
        # 2. git remote -v
        remotes = subprocess.check_output(["git", "remote", "-v"], cwd=repo_path).decode().strip()
        if not remotes:
            return "NO_REMOTE", None
        
        # 3. git fetch origin (Safe read)
        subprocess.check_call(["git", "fetch", "origin", "--quiet"], cwd=repo_path)
        
        # 4. git branch -vv (Check sync status)
        branches = subprocess.check_output(["git", "branch", "-vv"], cwd=repo_path).decode().strip()
        if "[ahead " in branches or "[behind " in branches:
            return "OUT_OF_SYNC", branches
            
        return "READY_FOR_LOCAL_REMOVAL", remotes
    except subprocess.CalledProcessError as e:
        return "GIT_ERROR", str(e)

def run_verification():
    print("[*] Starting Phase L1: Git Verification")
    candidates = []
    
    # Read inventory to find GITHUB_ONLY_REMOVE_CANDIDATE
    with open(INVENTORY_FILE, "r") as f:
        for line in f:
            if "GITHUB_ONLY_REMOVE_CANDIDATE" in line:
                parts = line.split("|")
                name = parts[1].strip()
                path = Path(parts[5].strip())
                candidates.append((name, path))
                
    results = {}
    for name, path in candidates:
        print(f"[*] Verifying {name}...")
        status, details = get_git_status(path)
        results[name] = {"path": str(path), "status": status, "details": details}
        
    with open(ROOT_DIR / "git_verification_results.json", "w") as f:
        json.dump(results, f, indent=2)
        
    print(f"\n[*] Verification Complete. Results saved to git_verification_results.json")
    for name, data in results.items():
        print(f"- {name}: {data['status']}")

if __name__ == "__main__":
    run_verification()
