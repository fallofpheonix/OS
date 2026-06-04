"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
import os, subprocess, json, time

def audit_ecosystem_sync():
    root = "/Users/fallofpheonix/engineering"
    state_dir = os.path.join(root, "control-plane/state")
    os.makedirs(state_dir, exist_ok=True)
    
    print(f"--- Ecosystem Git Sync Audit (Machine-Readable) ---")
    
    ecosystem_state = {
        "timestamp": time.strftime("%Y-%m-%dT%H:%M:%SZ", time.gmtime()),
        "root": root,
        "repositories": []
    }
    
    # Discovery via canonical git check
    repos = []
    for dirpath, dirnames, filenames in os.walk(root):
        if ".git" in dirnames:
            repos.append(dirpath)
            dirnames.remove(".git")
            
    for repo in sorted(repos):
        rel_path = os.path.relpath(repo, root)
        repo_state = {
            "name": os.path.basename(repo),
            "path": rel_path,
            "is_dirty": False,
            "sync_status": "UNKNOWN",
            "tracking_remote": False,
            "ahead_commits": 0,
            "behind_commits": 0,
            "violations": []
        }
        
        try:
            # 1. Check for uncommitted changes (Porcelain v2)
            status_raw = subprocess.check_output(
                ["git", "-C", repo, "status", "--porcelain=v2", "--branch"],
                stderr=subprocess.STDOUT
            ).decode()
            
            # Parse porcelain v2 output
            lines = status_raw.splitlines()
            for line in lines:
                if line.startswith("# branch.oid"): continue
                if line.startswith("# branch.head"): continue
                if line.startswith("# branch.upstream"):
                    repo_state["tracking_remote"] = True
                if line.startswith("# branch.ab"):
                    ab = line.split()
                    repo_state["ahead_commits"] = int(ab[2].replace("+", ""))
                    repo_state["behind_commits"] = int(ab[3].replace("-", ""))
                
                # If it doesn't start with '#', it's a file change
                if not line.startswith("#"):
                    repo_state["is_dirty"] = True

            # 2. Derive Sync Status
            if not repo_state["tracking_remote"]:
                repo_state["sync_status"] = "UNTRACKED"
                repo_state["violations"].append("No upstream remote configured")
            elif repo_state["ahead_commits"] > 0:
                repo_state["sync_status"] = "AHEAD"
            elif repo_state["behind_commits"] > 0:
                repo_state["sync_status"] = "BEHIND"
            else:
                repo_state["sync_status"] = "SYNCED"
                
            print(f"  [OK] {rel_path} -> {repo_state['sync_status']}")

        except Exception as e:
            repo_state["violations"].append(f"Audit failure: {str(e)}")
            print(f"  [FAIL] {rel_path}: {e}")

        ecosystem_state["repositories"].append(repo_state)

    # Emit Machine-Readable State
    state_file = os.path.join(state_dir, "git_sync_state.json")
    with open(state_file, "w") as f:
        json.dump(ecosystem_state, f, indent=2)
    
    print(f"\nAudit state emitted to {state_file}")

if __name__ == "__main__":
    audit_ecosystem_sync()
