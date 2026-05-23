import json
import os
import re
from pathlib import Path
from collections import defaultdict

# --- CONFIGURATION ---
ROOT_DIR = Path("/Users/fallofpheonix/engineering")
AUDIT_FILES = [
    "repo_inventory.json", "duplicate_report.md", "dead_files_report.md",
    "delete_candidates.json", "archive_inventory.md", "embedded_repo_report.md",
    "dependency_impact.md", "hash_index.json", "safe_delete_script.sh",
    "restore_script.sh", "final_matrix.md"
]

class Validator:
    def __init__(self, root):
        self.root = root
        self.candidates = []
        self.hashes = {}
        self.embedded = []
        self.archives = []
        self.metadata = {} # path -> meta
        self.validation_log = []
        
    def load_data(self):
        print("[*] Loading Audit Data...")
        with open(self.root / "delete_candidates.json", "r") as f:
            self.candidates = json.load(f)
        with open(self.root / "hash_index.json", "r") as f:
            self.hashes = json.load(f)
        with open(self.root / "repo_inventory.json", "r") as f:
            self.inventory = json.load(f)
        
        # Load embedded repos from report
        with open(self.root / "embedded_repo_report.md", "r") as f:
            self.embedded = [line.strip("- ").strip() for line in f if line.startswith("- ")]
            
        # Load archives from report
        with open(self.root / "archive_inventory.md", "r") as f:
            self.archives = [line.strip("- ").strip() for line in f if line.startswith("- ")]

    def phase_a_consistency(self):
        print("[*] Phase A: Audit Consistency Check")
        rejected = []
        validated = []
        
        # Protective patterns
        protective_dirs = ["brain", "governance", "docs", "runtime", "control-plane", "workspace/active"]
        
        for c in self.candidates:
            path = c["path"]
            rel = c["rel"]
            reason = "VALID"
            
            # Check for archive presence
            if any(path.startswith(a) for a in self.archives):
                reason = "Archive Conflict"
            # Check for .git
            elif "/.git" in path:
                reason = "Git Internal"
            # Check for governance/active
            elif any(p in path for p in protective_dirs):
                reason = "Architectural Protection"
            
            if reason != "VALID":
                rejected.append({"path": path, "reason": reason})
            else:
                validated.append(c)
                
        with open(self.root / "audit_validation.md", "w") as f:
            f.write("# Audit Validation Report\n\n")
            f.write(f"Validated Candidates: {len(validated)}\n")
            f.write(f"Rejected Candidates: {len(rejected)}\n\n")
            f.write("## Rejected Details\n")
            for r in rejected:
                f.write(f"- {r['path']} ({r['reason']})\n")
        
        self.validated_candidates = validated
        self.rejected_candidates = rejected

    def phase_b_review(self):
        print("[*] Phase B: Safe Delete Review")
        review = defaultdict(list)
        
        for c in self.validated_candidates:
            path = c["path"]
            rel = c["rel"]
            
            cat = "UNKNOWN"
            if "research" in path: cat = "RESEARCH_DELETE"
            elif "_tmp" in path or ".tmp" in path: cat = "TEMP_DELETE"
            elif c.get("confidence", 0) >= 90: cat = "DUPLICATE_DELETE"
            
            review[cat].append(c)
            
        with open(self.root / "safe_delete_review.md", "w") as f:
            f.write("# Safe Delete Review\n\n")
            for cat, items in review.items():
                f.write(f"## {cat}\n")
                for i in items:
                    f.write(f"- ID: {i.get('hash', 'N/A')[:8]}\n")
                    f.write(f"  Path: {i['path']}\n")
                    f.write(f"  Reason: {i.get('reason', 'N/A')}\n")
                    f.write(f"  Conf: {i.get('confidence', 0)}\n\n")

    def phase_c_research(self):
        print("[*] Phase C: Research Cleanup Plan")
        with open(self.root / "research_cleanup_plan.md", "w") as f:
            f.write("# Research Cleanup Plan\n\n")
            for c in self.validated_candidates:
                if "research" in c["path"]:
                    f.write(f"- Action: QUARANTINE | {c['path']}\n")

    def phase_d_hotspots(self):
        print("[*] Phase D: Duplicate Hotspot Analysis")
        hotspots = ["LAMP", "ParticleStimulator", "LifeTrack", "AutoEIT-STS"]
        with open(self.root / "duplicate_hotspots.md", "w") as f:
            f.write("# Duplicate Hotspot Analysis\n\n")
            for h in hotspots:
                f.write(f"## Hotspot: {h}\n")
                matches = [c for c in self.validated_candidates if h in c["path"]]
                f.write(f"Candidates: {len(matches)}\n")
                for m in matches[:10]: f.write(f"- {m['rel']}\n")
                if len(matches) > 10: f.write("- ...\n")

    def phase_e_repos(self):
        print("[*] Phase E: Repository Cleanup Plan")
        with open(self.root / "repo_cleanup_plan.md", "w") as f:
            f.write("# Repository Cleanup Plan\n\n")
            for r in self.embedded:
                # Heuristic classification
                status = "ARCHIVE_REPO"
                if "workspace/active" in r: status = "ACTIVE_REPO"
                f.write(f"- {status} | {r}\n")

    def phase_f_simulation(self):
        print("[*] Phase F: Quarantine Simulation")
        timestamp = datetime.datetime.now().strftime("%Y%m%d")
        dest_base = f"engineering/_quarantine_sim/{timestamp}/"
        
        quarantine_manifest = []
        restore_manifest = []
        
        with open(self.root / "move_preview.md", "w") as f:
            f.write("# Quarantine Move Preview\n\n")
            f.write(f"Target Destination: {dest_base}\n\n")
            for c in self.validated_candidates:
                dest = f"{dest_base}{c['rel']}"
                quarantine_manifest.append({"src": c["path"], "dest": dest})
                restore_manifest.append({"src": dest, "dest": c["path"]})
                f.write(f"mv '{c['path']}' '{dest}'\n")
                
        with open(self.root / "quarantine_manifest.json", "w") as f:
            json.dump(quarantine_manifest, f, indent=2)
        with open(self.root / "restore_manifest.json", "w") as f:
            json.dump(restore_manifest, f, indent=2)

    def summary(self):
        print("\n[*] Validation Summary:")
        print(f"Validated Safe Deletes: {len(self.validated_candidates)}")
        print(f"Rejected Deletes: {len(self.rejected_candidates)}")
        
        state = "READY_FOR_QUARANTINE" if len(self.validated_candidates) > 0 else "REQUIRES_MANUAL_REVIEW"
        print(f"\nFINAL STATE: {state}")

if __name__ == "__main__":
    import datetime
    v = Validator(ROOT_DIR)
    v.load_data()
    v.phase_a_consistency()
    v.phase_b_review()
    v.phase_c_research()
    v.phase_d_hotspots()
    v.phase_e_repos()
    v.phase_f_simulation()
    v.summary()
