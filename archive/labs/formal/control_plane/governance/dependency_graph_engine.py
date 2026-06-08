"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
import yaml, os

def generate_dependency_graph():
    root = "/Users/fallofpheonix/engineering"
    registry_path = os.path.join(root, "control-plane/repo-registry/MASTER_REPO_INDEX.yaml")
    
    with open(registry_path, "r") as f:
        content = f.read().split("---\n")
        data = yaml.safe_load(content[1])
        
    output_path = os.path.join(root, "control-plane/dependency-map/DEPENDENCY_GRAPH.md")
    
    with open(output_path, "w") as f:
        f.write("# ECOSYSTEM DEPENDENCY GRAPH\n\n")
        f.write("```mermaid\n")
        f.write("graph TD\n")
        for entry in data:
            p = entry["project"]
            name = p["name"]
            deps = p.get("dependencies", [])
            for dep in deps:
                f.write(f"    {name} --> {dep}\n")
        f.write("```\n")
    print(f"Graph generated at {output_path}")

if __name__ == "__main__":
    generate_dependency_graph()
