import ast
import os
from pathlib import Path
from collections import defaultdict
import json

class ImportGraphGenerator:
    def __init__(self, root_dir):
        self.root = Path(root_dir)
        self.imports = defaultdict(set) # repo -> {imported_modules}

    def analyze(self):
        print(f"[*] Analyzing imports in {self.root}")
        for root, _, files in os.walk(self.root):
            if any(p in root for p in ["__pycache__", ".venv", "venv", "site-packages", ".git"]):
                continue
            
            for file in files:
                if file.endswith(".py"):
                    path = Path(root) / file
                    self._process_file(path)

    def _process_file(self, path):
        try:
            tree = ast.parse(path.read_text(encoding="utf-8"))
        except:
            return

        rel_path = path.relative_to(self.root)
        source_repo = rel_path.parts[0] if rel_path.parts else "root"

        for node in ast.walk(tree):
            if isinstance(node, ast.Import):
                for alias in node.names:
                    self.imports[source_repo].add(alias.name)
            elif isinstance(node, ast.ImportFrom):
                if node.module:
                    self.imports[source_repo].add(node.module)

    def generate_json(self, output_path):
        # Convert sets to lists for JSON
        serializable = {k: sorted(list(v)) for k, v in self.imports.items()}
        with open(output_path, "w") as f:
            json.dump(serializable, f, indent=2)
        print(f"[+] Import graph saved to {output_path}")

if __name__ == "__main__":
    import sys
    root = sys.argv[1] if len(sys.argv) > 1 else "/Users/fallofpheonix/engineering"
    generator = ImportGraphGenerator(root)
    generator.analyze()
    generator.generate_json(Path(root) / "dependency_graph.json")
