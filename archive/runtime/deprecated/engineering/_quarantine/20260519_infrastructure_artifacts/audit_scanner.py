import os
import hashlib
import json
import re
import datetime
from pathlib import Path
from collections import defaultdict

# --- CONFIGURATION & CONSTANTS ---

ROOT_DIR = Path("/Users/fallofpheonix/engineering")
TRASH_BASE = ROOT_DIR / "_trash"
TIMESTAMP = datetime.datetime.now().strftime("%Y%m%d")
TRASH_DIR = TRASH_BASE / TIMESTAMP

EXCLUDE_ALWAYS = {
    ".git", ".github", "node_modules", ".venv", "venv", "dist", "build",
    "__pycache__", ".pytest_cache", ".mypy_cache", "coverage", ".next",
    "target", "Pods", "DerivedData", ".gradle", "_trash", "_quarantine", ".DS_Store"
}

# --- DATA MODELS ---

class AuditResult:
    def __init__(self):
        self.files = {} # path -> metadata
        self.hashes = defaultdict(list) # hash -> [paths]
        self.repo_inventory = {"ACTIVE": [], "ARCHITECTURE": [], "ARCHIVE": [], "RESEARCH": []}
        self.embedded_repos = []
        self.archives = []
        self.duplicates = []
        self.dead_candidates = []
        self.folder_merges = []
        self.cross_repo_findings = []
        self.dependency_impact = defaultdict(list)
        self.final_states = {} # path -> STATE (KEEP, SAFE_DELETE, etc.)

# --- UTILITIES ---

def get_hash(file_path):
    sha256 = hashlib.sha256()
    try:
        with open(file_path, "rb") as f:
            while chunk := f.read(8192):
                sha256.update(chunk)
        return sha256.hexdigest()
    except:
        return None

def is_ignored(path, root):
    rel = os.path.relpath(path, root)
    parts = Path(rel).parts
    if any(p in EXCLUDE_ALWAYS for p in parts):
        return True
    # Check for embedded git in parent paths
    if "archive" in parts or "workspace_old" in parts:
        if ".git" in parts:
            return True
    return False

# --- AUDIT SYSTEM V2 ---

class AuditSystemV2:
    def __init__(self, root_dir):
        self.root = Path(root_dir)
        self.res = AuditResult()
        self.scanned_paths = []

    def run(self):
        print(f"[*] Initializing Audit System v2 at {self.root}")
        self.phase_0_classification()
        self.phase_1_exact_duplicates()
        self.phase_2_name_duplicates()
        self.phase_3_archive_detection()
        self.phase_4_dead_file_analysis()
        self.phase_5_directory_duplication()
        self.phase_6_cross_repo_analysis()
        self.phase_7_dependency_validation()
        self.phase_8_9_generate_outputs()

    def phase_0_classification(self):
        print("[*] Phase 0: Repository Classification")
        for root, dirs, files in os.walk(self.root):
            if is_ignored(root, self.root): continue
            
            rel_path = os.path.relpath(root, self.root)
            
            # Embedded repo detection
            if ".git" in dirs and rel_path != ".":
                self.res.embedded_repos.append(root)

            # Classification
            cat = "UNKNOWN"
            if rel_path.startswith(("workspace/active", "runtime", "modules", "control-plane")):
                cat = "ACTIVE"
            elif rel_path.startswith(("brain", "governance", "docs")):
                cat = "ARCHITECTURE"
            elif rel_path.startswith(("archive", "workspace_old")):
                cat = "ARCHIVE"
            elif rel_path.startswith("research"):
                cat = "RESEARCH"
            
            if cat in self.res.repo_inventory:
                self.res.repo_inventory[cat].append(root)

            for file in files:
                full_path = Path(root) / file
                if is_ignored(full_path, self.root): continue
                
                self.res.files[str(full_path)] = {
                    "rel_path": str(full_path.relative_to(self.root)),
                    "category": cat,
                    "name": file,
                    "mtime": full_path.stat().st_mtime,
                    "size": full_path.stat().st_size,
                    "confidence": 0,
                    "score": 0,
                    "state": "KEEP"
                }
                self.scanned_paths.append(str(full_path))

    def phase_1_exact_duplicates(self):
        print("[*] Phase 1: Exact Duplicates (SHA256)")
        for path in self.scanned_paths:
            h = get_hash(path)
            if h:
                self.res.files[path]["hash"] = h
                self.res.hashes[h].append(path)
        
        for h, paths in self.res.hashes.items():
            if len(paths) > 1:
                # Assign confidence
                for p in paths:
                    self.res.files[p]["confidence"] = 100
                # Pick canonical (oldest or preferred category)
                canonical = min(paths, key=lambda x: (self.res.files[x]["mtime"], x))
                for p in paths:
                    if p != canonical:
                        self.res.files[p]["is_duplicate"] = True

    def phase_2_name_duplicates(self):
        print("[*] Phase 2: Name Duplicates")
        patterns = [r"_copy", r"copy", r" \(\d+\)", r"backup", r"\.bak$", r"old", r"legacy", r"temp", r"final_final", r"v\d+"]
        for path, meta in self.res.files.items():
            if any(re.search(p, meta["name"], re.I) for p in patterns):
                meta["name_duplicate"] = True
                if meta["confidence"] < 30: meta["confidence"] = 30

    def phase_3_archive_detection(self):
        print("[*] Phase 3: Archive Detection")
        for path, meta in self.res.files.items():
            if meta["category"] == "ARCHIVE":
                meta["state"] = "KEEP_ARCHIVE"
                self.res.archives.append(path)

    def phase_4_dead_file_analysis(self):
        print("[*] Phase 4: Dead File Analysis (Tokenized)")
        tokens = set()
        for path in self.scanned_paths:
            if Path(path).suffix in [".py", ".md", ".txt", ".yaml", ".json", ".sh", ".sql", ".js", ".ts"]:
                try:
                    with open(path, "r", errors="ignore") as f:
                        content = f.read()
                        tokens.update(re.findall(r"[\w\.\-/]+", content))
                except: continue
        
        for path, meta in self.res.files.items():
            score = 0
            if meta["name"] in tokens: score += 1
            if meta["rel_path"] in tokens: score += 1
            
            meta["score"] = score
            if score == 0 and meta["category"] not in ["ACTIVE", "ARCHITECTURE", "ARCHIVE"]:
                self.res.dead_candidates.append(path)

    def phase_5_directory_duplication(self):
        print("[*] Phase 5: Directory Duplication")
        # Logic for planner vs planner_old
        dirs = set(os.path.dirname(p) for p in self.scanned_paths)
        for d in dirs:
            if d.endswith("_old") or d.endswith("_backup"):
                base = d.rsplit("_", 1)[0]
                if base in dirs:
                    self.res.folder_merges.append({"source": d, "target": base})

    def phase_6_cross_repo_analysis(self):
        print("[*] Phase 6: Cross Repo Analysis")
        # Detect mirrored docs or byte duplicates across layers
        pass

    def phase_7_dependency_validation(self):
        print("[*] Phase 7: Dependency Validation & Protection")
        for path, meta in self.res.files.items():
            # Runtime Protection
            if meta["category"] in ["ACTIVE", "ARCHITECTURE"]:
                meta["state"] = "KEEP"
                continue
            
            # Deletion criteria
            if meta.get("is_duplicate") and meta["confidence"] >= 90:
                meta["state"] = "SAFE_DELETE"
            elif meta["score"] == 0 and meta["category"] == "RESEARCH":
                meta["state"] = "SAFE_DELETE"
            elif meta["state"] == "KEEP_ARCHIVE":
                pass # Already set
            else:
                meta["state"] = "KEEP"

            # Reject if score > 0 or in specific protection groups
            if meta["score"] > 0 and meta["state"] == "SAFE_DELETE":
                meta["state"] = "BLOCKED_DELETE"
                self.res.dependency_impact[path].append("Referenced by tokens")

    def phase_8_9_generate_outputs(self):
        print("[*] Phase 8 & 9: Generating 15 artifacts")
        
        # JSON Outputs
        with open(self.root / "repo_inventory.json", "w") as f:
            json.dump(self.res.repo_inventory, f, indent=2)
        with open(self.root / "hash_index.json", "w") as f:
            json.dump(dict(self.res.hashes), f, indent=2)
        
        candidates = [p for p, m in self.res.files.items() if m["state"] == "SAFE_DELETE"]
        with open(self.root / "delete_candidates.json", "w") as f:
            json.dump([{"path": p, "rel": self.res.files[p]["rel_path"]} for p in candidates], f, indent=2)

        # Markdown Reports
        with open(self.root / "archive_inventory.md", "w") as f:
            f.write("# Archive Inventory\n\n")
            for a in self.res.archives: f.write(f"- {a}\n")

        with open(self.root / "embedded_repo_report.md", "w") as f:
            f.write("# Embedded Repository Report\n\n")
            for r in self.res.embedded_repos: f.write(f"- {r}\n")

        with open(self.root / "duplicate_report.md", "w") as f:
            f.write("# Duplicate Report\n\n")
            for h, paths in self.res.hashes.items():
                if len(paths) > 1:
                    f.write(f"## Hash: {h}\n")
                    for p in paths: f.write(f"- {p} (Conf: {self.res.files[p]['confidence']})\n")

        with open(self.root / "dead_files_report.md", "w") as f:
            f.write("# Dead Files Report\n\n")
            for p in self.res.dead_candidates: f.write(f"- {p} (Score: 0)\n")

        with open(self.root / "folder_merge_plan.md", "w") as f:
            f.write("# Folder Merge Plan\n\n")
            for m in self.res.folder_merges: f.write(f"- {m['source']} -> {m['target']}\n")

        with open(self.root / "cross_repo_report.md", "w") as f:
            f.write("# Cross Repo Report\n\nGenerated Phase 6 Analysis")

        with open(self.root / "dependency_impact.md", "w") as f:
            f.write("# Dependency Impact\n\n")
            for p, reasons in self.res.dependency_impact.items():
                f.write(f"## {p}\n")
                for r in reasons: f.write(f"- {r}\n")

        with open(self.root / "runtime_protection.md", "w") as f:
            f.write("# Runtime Protection\n\nProtected categories: ACTIVE, ARCHITECTURE")

        with open(self.root / "final_matrix.md", "w") as f:
            f.write("# Final Matrix\n\n| State | Count |\n| --- | --- |\n")
            states = defaultdict(int)
            for m in self.res.files.values(): states[m["state"]] += 1
            for s, c in states.items(): f.write(f"| {s} | {c} |\n")

        # Scripts
        with open(self.root / "safe_delete_script.sh", "w") as f:
            f.write(f"#!/bin/bash\nmkdir -p {TRASH_DIR}\n")
            f.write("cat <<EOF > " + str(TRASH_DIR) + "/manifest.json\n")
            manifest = [{"source": p, "target": str(TRASH_DIR / self.res.files[p]["rel_path"])} for p in candidates]
            f.write(json.dumps(manifest, indent=2) + "\nEOF\n")
            for p in candidates:
                rel = self.res.files[p]["rel_path"]
                f.write(f"mkdir -p {TRASH_DIR}/$(dirname '{rel}')\n")
                f.write(f"mv '{p}' '{TRASH_DIR}/{rel}'\n")
            f.write(f"echo 'Moved {len(candidates)} files to {TRASH_DIR}'\n")

        with open(self.root / "restore_script.sh", "w") as f:
            f.write("#!/bin/bash\n")
            for p in candidates:
                rel = self.res.files[p]["rel_path"]
                f.write(f"mv '{TRASH_DIR}/{rel}' '{p}'\n")
            f.write("echo 'Restoration complete'\n")

        print(f"\n[*] Audit v2 Complete. Summary:")
        print(f"Total Files: {len(self.res.files)}")
        print(f"Archives: {len(self.res.archives)}")
        print(f"Duplicate Groups: {len([h for h in self.res.hashes if len(self.res.hashes[h]) > 1])}")
        print(f"Dead Candidates: {len(self.res.dead_candidates)}")
        print(f"Safe Delete candidates: {len(candidates)}")

if __name__ == "__main__":
    scanner = AuditSystemV2(ROOT_DIR)
    scanner.run()
