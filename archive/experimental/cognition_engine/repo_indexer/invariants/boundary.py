import json
from pathlib import Path

class BoundaryInvariantChecker:
    def __init__(self, dependency_graph_path):
        self.graph_path = Path(dependency_graph_path)
        self.violations = []

    def load_graph(self):
        with open(self.graph_path, "r") as f:
            return json.load(f)

    def check(self):
        graph = self.load_graph()
        print("[*] Checking architectural boundaries...")
        
        # Example rule: runtime should not depend on research directly
        # Example rule: cognition should not depend on products
        
        rules = [
            {"source": "runtime", "forbidden": ["research", "products"]},
            {"source": "brain", "forbidden": ["products"]},
            {"source": "control-plane", "forbidden": ["research"]}
        ]

        for repo, imported in graph.items():
            for rule in rules:
                if rule["source"] in repo.lower():
                    for forbidden in rule["forbidden"]:
                        for imp in imported:
                            if forbidden in imp.lower():
                                self.violations.append({
                                    "repo": repo,
                                    "imported": imp,
                                    "rule": f"Layer {rule['source']} cannot import from {forbidden}"
                                })

    def generate_report(self, output_path):
        with open(output_path, "w") as f:
            json.dump(self.violations, f, indent=2)
        print(f"[+] Boundary violations report saved to {output_path}")
        print(f"[*] Found {len(self.violations)} violations.")

if __name__ == "__main__":
    import sys
    graph_path = sys.argv[1] if len(sys.argv) > 1 else "/Users/fallofpheonix/engineering/dependency_graph.json"
    checker = BoundaryInvariantChecker(graph_path)
    checker.check()
    checker.generate_report(Path(graph_path).parent / "boundary_violations.json")
