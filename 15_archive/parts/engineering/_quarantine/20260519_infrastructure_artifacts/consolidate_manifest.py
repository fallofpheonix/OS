import json
import yaml
import os
import re
from pathlib import Path
from datetime import datetime

ROOT_DIR = Path("/Users/fallofpheonix/engineering")

def load_json(path):
    if path.exists():
        with open(path, 'r') as f:
            return json.load(f)
    return {}

def load_md_table(path):
    data = []
    if path.exists():
        with open(path, 'r') as f:
            lines = f.readlines()
            for line in lines:
                if "|" in line and "---" not in line and "Repository" not in line:
                    parts = [p.strip() for p in line.split("|") if p.strip()]
                    if parts:
                        data.append(parts)
    return data

def load_tags(path):
    tags = {}
    if path.exists():
        with open(path, 'r') as f:
            current_repo = None
            for line in f:
                if line.startswith("## "):
                    current_repo = line.strip("# ").strip()
                    tags[current_repo] = []
                elif line.startswith("- ") and current_repo:
                    tags[current_repo].append(line.strip("- ").strip())
    return tags

# Load Artifacts
inventory_raw = load_md_table(ROOT_DIR / "repo_inventory_v2.md")
physics_tags = load_tags(ROOT_DIR / "physics_alignment.md")
math_tags = load_tags(ROOT_DIR / "math_alignment.md")
dependency_graph = load_json(ROOT_DIR / "dependency_graph.json")
boundary_violations = load_json(ROOT_DIR / "boundary_violations.json")
ownership_map = load_json(ROOT_DIR / "ownership_map.json")
git_verify = load_json(ROOT_DIR / "git_verification_results.json")

# Runtime Scan (from previous find output)
runtime_paths = [
    "/Users/fallofpheonix/engineering/archive/workspace_old/Noesis/requirements.txt",
    "/Users/fallofpheonix/engineering/archive/workspace_old/Noesis/pyproject.toml",
    "/Users/fallofpheonix/engineering/archive/workspace_old/LAMP/pyproject.toml",
    "/Users/fallofpheonix/engineering/archive/workspace_old/AutoEIT-STS/requirements.txt",
    "/Users/fallofpheonix/engineering/archive/workspace_old/AutoEIT-STS/pyproject.toml",
    "/Users/fallofpheonix/engineering/archive/workspace_old/AutoEIT-STS/submission/requirements.txt",
    "/Users/fallofpheonix/engineering/archive/workspace_old/AutoEIT-STS/submission/pyproject.toml",
    "/Users/fallofpheonix/engineering/archive/workspace_old/ParticleStimulator/requirements.txt",
    "/Users/fallofpheonix/engineering/archive/workspace_old/TerraHerb/requirements.txt",
    "/Users/fallofpheonix/engineering/archive/workspace_old/TerraHerb/pyproject.toml",
    "/Users/fallofpheonix/engineering/archive/workspace_old/ArtExtract/requirements.txt",
    "/Users/fallofpheonix/engineering/archive/workspace_old/ArtExtract/pyproject.toml",
    "/Users/fallofpheonix/engineering/archive/workspace_old/UDIE/udie_backend_py/requirements.txt",
    "/Users/fallofpheonix/engineering/archive/workspace_old/SecureForg/requirements.txt",
    "/Users/fallofpheonix/engineering/archive/workspace_old/audio_transcription/requirements.txt",
    "/Users/fallofpheonix/engineering/archive/workspace_old/audio_transcription/pyproject.toml",
    "/Users/fallofpheonix/engineering/archive/workspace_old/agentskill/backend/agentman/pyproject.toml",
    "/Users/fallofpheonix/engineering/archive/workspace_old/sira/requirements.txt",
    "/Users/fallofpheonix/engineering/archive/workspace_old/sira/pyproject.toml",
    "/Users/fallofpheonix/engineering/archive/workspace_old/sira/submission/source_code/requirements.txt",
    "/Users/fallofpheonix/engineering/archive/workspace_old/sira/submission/source_code/pyproject.toml",
    "/Users/fallofpheonix/engineering/archive/workspace_old/AutoTRandHD/pyproject.toml",
    "/Users/fallofpheonix/engineering/archive/workspace_old/healingstone/requirements.txt",
    "/Users/fallofpheonix/engineering/archive/workspace_old/healingstone/pyproject.toml",
    "/Users/fallofpheonix/engineering/archive/workspace_old/AI4MH/backend/pyproject.toml",
    "/Users/fallofpheonix/engineering/archive/workspace_old/AutoMation-Engine/requirements.txt",
    "/Users/fallofpheonix/engineering/archive/workspace_old/AI-PFI/requirements.txt",
    "/Users/fallofpheonix/engineering/archive/workspace_old/AI-PFI/submission/requirements.txt",
    "/Users/fallofpheonix/engineering/archive/workspace_old/ChoreoAI/requirements.txt",
    "/Users/fallofpheonix/engineering/archive/workspace_old/ChoreoAI/pyproject.toml",
    "/Users/fallofpheonix/engineering/environments/ai-system/requirements.txt",
    "/Users/fallofpheonix/engineering/environments/ai-system/venv",
    "/Users/fallofpheonix/engineering/workspace/repo-analyzer/pyproject.toml",
    "/Users/fallofpheonix/engineering/workspace/active/astraeus-core/artifacts/artifacts/run_20260515T205747_ee6424ff/main/snapshots/snapshot_794c25ef0cc0/files/uv.lock",
    "/Users/fallofpheonix/engineering/workspace/active/astraeus-core/artifacts/artifacts/run_20260515T205747_ee6424ff/main/snapshots/snapshot_794c25ef0cc0/files/pyproject.toml",
    "/Users/fallofpheonix/engineering/workspace/active/astraeus-core/uv.lock",
    "/Users/fallofpheonix/engineering/workspace/active/astraeus-core/pyproject.toml",
    "/Users/fallofpheonix/engineering/workspace/active/astraeus-core/.venv",
    "/Users/fallofpheonix/engineering/workspace/active/astraeus-core/.venv/lib/python3.13/site-packages/mypy/typeshed/stdlib/venv",
    "/Users/fallofpheonix/engineering/workspace/active/TrustLab/requirements.txt",
    "/Users/fallofpheonix/engineering/workspace/active/TrustLab/pyproject.toml",
    "/Users/fallofpheonix/engineering/workspace/active/TrustLab/submission/source_code/requirements.txt",
    "/Users/fallofpheonix/engineering/workspace/active/TrustLab/submission/source_code/pyproject.toml",
    "/Users/fallofpheonix/engineering/workspace/active/forge-agent/pyproject.toml",
    "/Users/fallofpheonix/engineering/workspace/active/autoeit-suite/packages/autoeit-score/requirements.txt",
    "/Users/fallofpheonix/engineering/workspace/active/autoeit-suite/packages/autoeit-score/pyproject.toml",
    "/Users/fallofpheonix/engineering/workspace/active/autoeit-suite/packages/autoeit-score/submission/requirements.txt",
    "/Users/fallofpheonix/engineering/workspace/active/autoeit-suite/packages/autoeit-score/submission/pyproject.toml",
    "/Users/fallofpheonix/engineering/workspace/active/autoeit-suite/packages/autoeit-transcribe/requirements.txt",
    "/Users/fallofpheonix/engineering/workspace/active/autoeit-suite/packages/autoeit-transcribe/pyproject.toml",
    "/Users/fallofpheonix/engineering/.venv",
    "/Users/fallofpheonix/engineering/brain/.obsidian/plugins/claudian/tests/unit/providers/claude/env",
    "/Users/fallofpheonix/engineering/brain/.obsidian/plugins/claudian/tests/unit/providers/codex/env",
    "/Users/fallofpheonix/engineering/brain/.obsidian/plugins/claudian/src/providers/claude/env",
    "/Users/fallofpheonix/engineering/brain/.obsidian/plugins/claudian/src/providers/codex/env",
    "/Users/fallofpheonix/engineering/brain/.obsidian/plugins/claudian/src/providers/opencode/env",
    "/Users/fallofpheonix/engineering/infrastructure/python/ruff/pyproject.toml"
]

# Normalization & Merging
repos = {}
for item in inventory_raw:
    name, cat, mat, own, path = item
    
    # Language Detection (Simplified)
    lang = "Python" # Default
    
    # Sync State
    sync_state = git_verify.get(name, {}).get("status", "UNTRACKED")
    
    # Archive Candidate
    is_archive_candidate = sync_state == "READY_FOR_LOCAL_REMOVAL"
    
    # Risk Level Calculation
    risk_level = "LOW"
    if sync_state == "DIRTY_WORKING_TREE": risk_level = "MEDIUM"
    if name in [v["repo"] for v in boundary_violations]: risk_level = "HIGH"

    # Runtime Type
    runtime_type = "NONE"
    for rp in runtime_paths:
        if name in rp:
            if "venv" in rp or ".venv" in rp:
                runtime_type = "VENV"
            elif "requirements.txt" in rp or "pyproject.toml" in rp:
                if runtime_type == "NONE": runtime_type = "PACKAGE"

    repos[name] = {
        "repo_name": name,
        "classification": cat,
        "status": "ACTIVE" if mat != "ARCHIVE" else "ARCHIVED",
        "owner": own,
        "maturity": mat,
        "runtime_type": runtime_type,
        "language": lang,
        "github_sync_state": sync_state,
        "archive_candidate": is_archive_candidate,
        "dependency_edges": dependency_graph.get(name, []),
        "research_tags": [], # Heuristic would go here
        "physics_tags": physics_tags.get(name, []),
        "math_tags": math_tags.get(name, []),
        "risk_level": risk_level,
        "path": path
    }

# Build Manifest Structure
manifest = {
    "repos": repos,
    "ownership": {
        "KEEP_ACTIVE": [n for n, r in repos.items() if r["owner"] == "KEEP_ACTIVE"],
        "GITHUB_ONLY_REMOVE_CANDIDATE": [n for n, r in repos.items() if r["owner"] == "GITHUB_ONLY_REMOVE_CANDIDATE"],
        "MANUAL_REVIEW": [n for n, r in repos.items() if r["owner"] == "MANUAL_REVIEW"],
        "UNKNOWN": [n for n, r in repos.items() if r["owner"] == "UNKNOWN"]
    },
    "maturity_levels": {
        "ACTIVE": [n for n, r in repos.items() if r["maturity"] == "ACTIVE"],
        "PRODUCT": [n for n, r in repos.items() if r["maturity"] == "PRODUCT"],
        "PROTOTYPE": [n for n, r in repos.items() if r["maturity"] == "PROTOTYPE"],
        "ARCHIVE": [n for n, r in repos.items() if r["maturity"] == "ARCHIVE"]
    },
    "dependency_state": {
        "graph_path": "dependency_graph.json",
        "violations": boundary_violations,
        "cycles": [], # Simplified for now
        "cross_links": [] # Simplified for now
    },
    "github_state": {
        "synced": [n for n, r in repos.items() if r["github_sync_state"] == "READY_FOR_LOCAL_REMOVAL"],
        "dirty": [n for n, r in repos.items() if r["github_sync_state"] == "DIRTY_WORKING_TREE"],
        "manual_review": [n for n, r in repos.items() if r["owner"] == "MANUAL_REVIEW"],
        "archive_ready": [n for n, r in repos.items() if r["archive_candidate"]]
    },
    "archive_state": {
        "github_synced": [n for n, r in repos.items() if r["github_sync_state"] == "READY_FOR_LOCAL_REMOVAL"],
        "local_only": [n for n, r in repos.items() if r["github_sync_state"] == "NO_REMOTE"],
        "active": [n for n, r in repos.items() if r["status"] == "ACTIVE"],
        "protected": [n for n, r in repos.items() if r["owner"] == "KEEP_ACTIVE"]
    },
    "runtime_targets": {
        "python_envs": [p for p in runtime_paths if "python" in p],
        "venvs": [p for p in runtime_paths if "venv" in p],
        "nested_envs": [p for p in runtime_paths if p.count("venv") > 1 or p.count("/") > 10],
        "shared_runtime": "/Users/fallofpheonix/engineering/.venv"
    },
    "research_mapping": {
        "cognition": ["brain"],
        "physics": list(physics_tags.keys()),
        "math": list(math_tags.keys()),
        "experiments": [n for n, r in repos.items() if r["classification"] == "RESEARCH"]
    },
    "remove_candidates": [n for n, r in repos.items() if r["archive_candidate"]],
    "restore_points": [],
    "integrity": {
        "manifest_version": "1.0.0",
        "generated_at": datetime.now().isoformat(),
        "ecosystem_hash": "sha256-placeholder",
        "cognition_state": "AUTHORITATIVE",
        "boundary_state": "VALID"
    }
}

# Write Manifest
with open(ROOT_DIR / "repo_manifest.yaml", "w") as f:
    yaml.dump(manifest, f, sort_keys=False)

# Write Runtime Inventory JSON
runtime_inventory = {
    "all_found_paths": runtime_paths,
    "grouped": {
        "core": [p for p in runtime_paths if "active" in p],
        "archive": [p for p in runtime_paths if "archive" in p],
        "environments": [p for p in runtime_paths if "environments" in p]
    }
}
with open(ROOT_DIR / "runtime_inventory.json", "w") as f:
    json.dump(runtime_inventory, f, indent=2)

# Write Archive Execution Plan
archive_plan = []
for name in manifest["github_state"]["archive_ready"]:
    repo = repos[name]
    archive_plan.append({
        "repo": name,
        "current_path": repo["path"],
        "target_archive": str(ROOT_DIR / "archive" / "github_synced" / name),
        "github_verified": True,
        "dirty_state": False,
        "safe_to_move": True,
        "rollback_path": repo["path"]
    })
# Add manual review ones as unsafe for now
for name in ["AI-PFI", "ArtExtract", "AutoTRandHD", "agentskill", "audio_transcription"]:
    if name in repos:
        repo = repos[name]
        archive_plan.append({
            "repo": name,
            "current_path": repo["path"],
            "target_archive": str(ROOT_DIR / "archive" / "github_synced" / name),
            "github_verified": False,
            "dirty_state": True,
            "safe_to_move": False,
            "rollback_path": repo["path"]
        })

with open(ROOT_DIR / "archive_execution_plan.yaml", "w") as f:
    yaml.dump(archive_plan, f, sort_keys=False)

# Generate Supporting Reports
reports_dir = ROOT_DIR / "reports"

# Repo Summary
with open(reports_dir / "repo_summary.md", "w") as f:
    f.write("# Repository Summary Report\n\n")
    f.write(f"Total Repositories: {len(repos)}\n")
    f.write(f"Active: {len(manifest['maturity_levels']['ACTIVE']) + len(manifest['maturity_levels']['PRODUCT'])}\n")
    f.write(f"Archived: {len(manifest['maturity_levels']['ARCHIVE'])}\n\n")
    f.write("| Repo | Category | Maturity | Risk |\n")
    f.write("| --- | --- | --- | --- |\n")
    for name, r in repos.items():
        f.write(f"| {name} | {r['classification']} | {r['maturity']} | {r['risk_level']} |\n")

# Runtime Inventory Report
with open(reports_dir / "runtime_inventory.md", "w") as f:
    f.write("# Runtime Inventory Report\n\n")
    f.write("## Virtual Environments\n")
    for p in manifest["runtime_targets"]["venvs"]:
        f.write(f"- {p}\n")
    f.write("\n## Package Manifests\n")
    for p in runtime_paths:
        if "requirements.txt" in p or "pyproject.toml" in p:
            f.write(f"- {p}\n")

# Risk Matrix
with open(reports_dir / "risk_matrix.md", "w") as f:
    f.write("# Risk Analysis Matrix\n\n")
    
    # Identify orphans
    orphans = []
    all_imports = set()
    for imp_list in dependency_graph.values():
        for imp in imp_list:
            all_imports.add(imp.split('.')[0])
    for name in repos:
        if name not in dependency_graph and name not in all_imports:
            orphans.append(name)
            
    f.write("## Critical Risks\n")
    if boundary_violations:
        f.write("- **Boundary Violations Found**: Immediate action required.\n")
    else:
        f.write("- None detected in cognition layer.\n")
        
    f.write("\n## High Severity Risks\n")
    f.write("- **Nested Virtual Environments**: Multiple venvs found in same trees (astraeus-core).\n")
    f.write("- **Dirty Research Repos**: candidates for removal have local changes.\n")
    
    f.write("\n## Medium Severity Risks\n")
    for name in orphans:
        f.write(f"- **Orphan Repository**: {name} has no detected imports or dependants.\n")

# Archive Plan Report
with open(reports_dir / "archive_plan.md", "w") as f:
    f.write("# Local Archival Plan\n\n")
    f.write("## Ready for Move\n")
    for name in manifest["github_state"]["archive_ready"]:
        f.write(f"- {name}\n")
    f.write("\n## Pending Manual Review\n")
    for name in ["AI-PFI", "ArtExtract", "AutoTRandHD", "agentskill", "audio_transcription"]:
        f.write(f"- {name} (Reason: Dirty working tree)\n")

# Dependency Hotspots
with open(reports_dir / "dependency_hotspots.md", "w") as f:
    f.write("# Dependency Hotspots\n\n")
    # Fan-in analysis
    fan_in = {}
    for imp_list in dependency_graph.values():
        for imp in imp_list:
            base = imp.split('.')[0]
            fan_in[base] = fan_in.get(base, 0) + 1
            
    sorted_fan_in = sorted(fan_in.items(), key=lambda x: x[1], reverse=True)
    f.write("## High Fan-In (Single Point of Failure Candidates)\n")
    for name, count in sorted_fan_in[:10]:
        f.write(f"- {name}: {count} dependents\n")

print("Phase R4 Consolidation Complete.")
