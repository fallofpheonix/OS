import os
import json
import re
from pathlib import Path
from collections import defaultdict

# --- CONFIGURATION & CONSTANTS ---
ROOT_DIR = Path("/Users/fallofpheonix/engineering")
GITHUB_SYNC_ARCHIVE = ROOT_DIR / "archive" / "github_synced"

EXCLUSIONS = {
    "graphify-out/cache",
    "site-packages",
    "torch",
    "jsonschema",
    ".git",
    "node_modules",
    "__pycache__",
    ".venv",
    "venv",
    "artifacts",
    "backups",
}

# Ownership Categories (Phase L0)
KEEP_ACTIVE = {
    "astraeus-core", "brain", "control-plane", "forge-agent", "modules", 
    "infrastructure", "aegis-auth", "ledger-core", "repo-analyzer"
}
REMOVE_CANDIDATES = {
    "my-portfolio", "fallofpheonix", "audio_transcription", "agentskill", 
    "ArtExtract", "AutoTRandHD", "AI-PFI", "AutoMation-Engine"
}
MANUAL_REVIEW = {
    "LifeTrack", "AutoEIT-STS", "ParticleStimulator", "Noesis", 
    "healingstone", "TerraHerb", "TrustLab", "LAMP"
}

# Repos for Physics/Math Alignment
PHYSICS_TARGETS = {
    "AutoEIT-STS",
    "ParticleStimulator",
    "TerraHerb",
    "AI4MH",
    "healingstone",
    "AutoTRandHD"
}

# Infrastructure Repos
INFRA_REPOS = {
    "aegis-auth",
    "smart-api-limiter",
    "ledger-core",
    "infrastructure",
    "control-plane",
    "truenotes",
    "repo-analyzer"
}

# Classification & Maturity Rules
CORE_SYSTEMS = {
    "astraeus-core", "brain", "control-plane", "forge-agent", "modules", 
    "repo-analyzer", "infrastructure", "aegis-auth", "smart-api-limiter", "ledger-core"
}
RESEARCH_SYSTEMS = {
    "AI4MH", "AI-PFI", "AutoEIT-STS", "ParticleStimulator", "Noesis", 
    "TerraHerb", "TrustLab", "ArtExtract", "AutoTRandHD", "AutoMation-Engine"
}
PRODUCTS = {"LifeTrack", "truenotes", "cognitron-game", "healingstone", "UDIE", "sira"}
TOOLS = {"agentskill", "audio_transcription", "autoeit-suite"}
PERSONAL = {"portfolio", "idea"}

# --- DATA MODELS ---

class EcosystemReport:
    def __init__(self):
        self.repos = {} # name -> metadata
        self.relations = []
        self.physics_map = {}
        self.math_map = {}
        self.provenance = {}
        self.merges = []
        self.ownership = {}

# --- UTILITIES ---

def is_excluded(path):
    parts = Path(path).parts
    for ex in EXCLUSIONS:
        if ex in parts or any(re.search(rf"{re.escape(ex)}", str(p)) for p in parts):
            return True
    return False

# --- ANALYSIS ENGINE ---

class EcosystemAnalyzer:
    def __init__(self, root_dir):
        self.root = Path(root_dir)
        self.report = EcosystemReport()

    def run(self):
        print(f"[*] Starting Ecosystem Analysis at {self.root}")
        self.phase_0_inventory()
        self.phase_1_relations()
        self.phase_2_provenance()
        self.phase_4_physics_math_alignment()
        self.phase_5_merge_analysis()
        self.generate_reports()

    def phase_0_inventory(self):
        print("[*] Phase 0: Global Repository Inventory & Scoring")
        for root, dirs, files in os.walk(self.root):
            if is_excluded(root):
                continue
            
            if ".git" in dirs or Path(root).parent == self.root:
                name = Path(root).name
                if name == "engineering" or name == "shared": continue
                
                # Ownership Classification (Phase L0)
                own_cat = "UNKNOWN"
                if name in KEEP_ACTIVE: own_cat = "KEEP_ACTIVE"
                elif name in REMOVE_CANDIDATES: own_cat = "GITHUB_ONLY_REMOVE_CANDIDATE"
                elif name in MANUAL_REVIEW: own_cat = "MANUAL_REVIEW"
                
                # Classification
                category = "UNKNOWN"
                if name in CORE_SYSTEMS: category = "CORE"
                elif name in RESEARCH_SYSTEMS: category = "RESEARCH"
                elif name in PRODUCTS: category = "PRODUCT"
                elif name in TOOLS: category = "TOOL"
                elif name in PERSONAL: category = "PERSONAL"
                
                if category == "UNKNOWN":
                    rel_path = os.path.relpath(root, self.root)
                    if "archive" in rel_path or "workspace_old" in rel_path: category = "RESEARCH"
                    elif "workspace" in rel_path: category = "PRODUCT"

                # Maturity
                maturity = "PROTOTYPE"
                if category == "PRODUCT": maturity = "PRODUCT"
                elif category == "CORE": maturity = "ACTIVE"
                
                rel_path = os.path.relpath(root, self.root)
                if "archive" in rel_path or "workspace_old" in rel_path:
                    maturity = "ARCHIVE"
                    category = "RESEARCH"
                elif name == "LAMP":
                    maturity = "ARCHIVE"
                
                self.report.repos[name] = {
                    "path": root,
                    "category": category,
                    "maturity": maturity,
                    "ownership": own_cat,
                    "files": len(files)
                }

    def phase_1_relations(self):
        print("[*] Phase 1: Relation Graph")
        infra_markers = {"docker-compose.yml", "Dockerfile", "Jenkinsfile", ".gitlab-ci.yml"}
        auth_markers = {"auth", "login", "jwt", "oauth"}
        for name, meta in self.report.repos.items():
            repo_path = Path(meta["path"])
            try:
                files = os.listdir(repo_path)
            except: continue
            if any(m in files for m in infra_markers):
                self.report.relations.append((name, "INFRA", "Infrastructure present"))
            if any(auth_markers.intersection(set(re.findall(r"\w+", f.lower()))) for f in files):
                self.report.relations.append((name, "AUTH", "Potential auth logic detected"))

    def phase_2_provenance(self):
        print("[*] Phase 2: Provenance Detection")
        for name, meta in self.report.repos.items():
            prov = "OWNED"
            rel_path = os.path.relpath(meta["path"], self.root)
            if "fork" in rel_path: prov = "FORK"
            elif "archive" in rel_path or "workspace_old" in rel_path: prov = "ARCHIVE"
            self.report.provenance[name] = prov

    def phase_4_physics_math_alignment(self):
        print("[*] Phase 4: Physics & Math Alignment")
        mappings = {
            "AutoEIT-STS": {
                "physics": ["electromagnetics", "tomography"],
                "math": ["inverse problems", "optimization", "PDEs"]
            },
            "ParticleStimulator": {
                "physics": ["particle dynamics", "motion systems"],
                "math": ["numerical methods", "ODE/PDE"]
            }
        }
        for name in PHYSICS_TARGETS:
            if name in mappings:
                self.report.physics_map[name] = mappings[name]["physics"]
                self.report.math_map[name] = mappings[name]["math"]

    def phase_5_merge_analysis(self):
        print("[*] Phase 5: Merge Analysis")
        self.report.merges = [
            ("autoeit-suite", "AutoEIT-STS", "MERGE_REVIEW"),
            ("forge-agent", "astraeus-core", "KEEP_SEPARATE"),
            ("infrastructure", "control-plane", "KEEP_SEPARATE"),
            ("repo-analyzer", "brain", "SHARED_INTERFACE")
        ]

    def generate_reports(self):
        print("[*] Generating Reports...")
        with open(self.root / "repo_inventory_v2.md", "w") as f:
            f.write("# Global Repository Inventory v2\n\n")
            f.write("| Repository | Category | Maturity | Ownership | Path |\n")
            f.write("| --- | --- | --- | --- | --- |\n")
            for name, meta in sorted(self.report.repos.items()):
                f.write(f"| {name} | {meta['category']} | {meta['maturity']} | {meta['ownership']} | {meta['path']} |\n")
        with open(self.root / "repo_relation_graph.md", "w") as f:
            f.write("# Repository Relation Graph\n\n")
            for src, rel, desc in self.report.relations:
                f.write(f"- {src} -> [{rel}] {desc}\n")
        with open(self.root / "foreign_influence_report.md", "w") as f:
            f.write("# Foreign Influence & Provenance Report\n\n")
            f.write("| Repository | Provenance | Action |\n")
            f.write("| --- | --- | --- |\n")
            for name, prov in self.report.provenance.items():
                f.write(f"| {name} | {prov} | KEEP |\n")
        with open(self.root / "physics_alignment.md", "w") as f:
            f.write("# Physics Alignment Mapping\n\n")
            for name, phys in self.report.physics_map.items():
                f.write(f"## {name}\n- " + "\n- ".join(phys) + "\n\n")
        with open(self.root / "math_alignment.md", "w") as f:
            f.write("# Math Alignment Mapping\n\n")
            for name, math in self.report.math_map.items():
                f.write(f"## {name}\n- " + "\n- ".join(math) + "\n\n")
        with open(self.root / "ENGINEERING_MANIFEST.md", "w") as f:
            f.write("# Engineering Ecosystem Manifest\n\n")
            f.write("## Core Architecture (Astraeus Model)\n")
            f.write("- **Cognition**: brain\n- **Execution**: workspace\n- **Components**: modules\n\n")
            f.write("## Meta Layout\n")
            f.write("- governance/\n- architecture/\n- research/\n- shared/\n- archives/\n")
        with open(self.root / "merge_matrix.md", "w") as f:
            f.write("# Merge Analysis Matrix\n\n")
            f.write("| Source | Target | Recommendation |\n")
            f.write("| --- | --- | --- |\n")
            for src, tgt, rec in self.report.merges:
                f.write(f"| {src} | {tgt} | {rec} |\n")
        with open(self.root / "ecosystem_history.md", "w") as f:
            f.write("# Ecosystem History & Timeline\n\n")
            f.write("## Origin\nUnified Engineering Ecosystem under Astraeus Paradigm.\n\n")
            f.write("## Major Shifts\n- Transition to Physics-First Reasoning (Active)\n- Mathematics-First Design Alignment (Active)\n")
        with open(self.root / "future_architecture.md", "w") as f:
            f.write("# Future System Architecture Model\n\n")
            f.write("```text\nengineering/\n├── cognition\n├── physics\n├── mathematics\n├── simulation\n├── agents\n├── runtime\n├── governance\n├── research\n├── products\n└── archives\n```\n")

    def validate(self):
        print("[*] Starting Phase R0: Report Validation")
        reports = ["repo_inventory_v2.md", "repo_relation_graph.md", "foreign_influence_report.md", "ENGINEERING_MANIFEST.md", "physics_alignment.md", "math_alignment.md", "merge_matrix.md", "ecosystem_history.md", "future_architecture.md"]
        missing = [r for r in reports if not (self.root / r).exists()]
        if missing:
            print(f"[!] MISSING REPORTS: {missing}")
            return False
        print("[+] All reports exist. Checking content consistency...")
        with open(self.root / "foreign_influence_report.md", "r") as f:
            content = f.read()
            valid_prov = {"OWNED", "FORK", "ARCHIVE", "EXTERNAL", "VENDOR", "GENERATED"}
            lines = content.splitlines()
            found_prov = set()
            for line in lines:
                match = re.search(r"\| [\w\-]+ \| (\w+) \|", line)
                if match:
                    prov = match.group(1)
                    if prov != "Provenance":
                        found_prov.add(prov)
            if not found_prov.issubset(valid_prov):
                print(f"[!] ERROR: Invalid provenance found: {found_prov - valid_prov}")
                return False
        with open(self.root / "repo_inventory_v2.md", "r") as f:
            for line in f:
                if "archive/workspace_old" in line and "ARCHIVE" not in line:
                    print(f"[!] ERROR: Archival repo not isolated correctly: {line.strip()}")
                    return False
        print("[+] Report Validation PASSED.")
        return True

if __name__ == "__main__":
    import sys
    analyzer = EcosystemAnalyzer(ROOT_DIR)
    if "--validate" in sys.argv:
        analyzer.phase_0_inventory()
        analyzer.phase_2_provenance()
        analyzer.validate()
    else:
        analyzer.run()
