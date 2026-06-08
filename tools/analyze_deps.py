import sys

layers = {
    "foundation": 1,
    "assurance": 2,
    "governance": 3,
    "cognition": 4,
    "platform": 5,
    "game": 6
}

def get_layer(path):
    for layer in layers:
        if f"phoenix/{layer}" in path:
            return layer
    return None

violations = []
deps = {}

with open("dependency_raw.txt", "r") as f:
    for line in f:
        parts = line.split("[")
        if len(parts) < 2: continue
        pkg = parts[0].strip()
        pkg_layer = get_layer(pkg)
        
        # Imports are space separated inside [ ]
        imports_raw = parts[1].replace("]", "").strip()
        if not imports_raw: continue
        
        imports = imports_raw.split()
        for imp in imports:
            if "github.com/fallofpheonix/phoenix" not in imp: continue
            imp_layer = get_layer(imp)
            
            if pkg_layer and imp_layer:
                if layers[imp_layer] > layers[pkg_layer]:
                    violations.append(f"VIOLATION: {pkg} ({pkg_layer}) imports {imp} ({imp_layer})")
            
            if pkg not in deps: deps[pkg] = []
            deps[pkg].append(imp)

print("# Dependency Violation Report")
for v in violations:
    print(v)

print("\n# Crucible Dependencies")
for pkg, pkg_deps in deps.items():
    if "crucible" in pkg:
        print(f"{pkg} ->")
        for d in pkg_deps:
            print(f"  {d}")

print("\n# Platform Dependencies")
for pkg, pkg_deps in deps.items():
    if "platform" in pkg and "crucible" not in pkg:
        print(f"{pkg} ->")
        for d in pkg_deps:
            print(f"  {d}")
