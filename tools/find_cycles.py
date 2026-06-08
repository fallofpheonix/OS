import sys

graph = {}

with open("dependency_raw.txt", "r") as f:
    for line in f:
        parts = line.split("[")
        if len(parts) < 2: continue
        pkg = parts[0].strip()
        imports_raw = parts[1].replace("]", "").strip()
        if not imports_raw: continue
        
        imports = imports_raw.split()
        internal_imports = [imp for imp in imports if "github.com/fallofpheonix/phoenix" in imp]
        graph[pkg] = internal_imports

def find_cycle(node, visited, path):
    visited.add(node)
    path.add(node)
    
    for neighbor in graph.get(node, []):
        if neighbor in path:
            return path
        if neighbor not in visited:
            res = find_cycle(neighbor, visited, path)
            if res: return res
            
    path.remove(node)
    return None

visited = set()
cycles = []

for node in graph:
    if node not in visited:
        cycle = find_cycle(node, visited, set())
        if cycle:
            cycles.append(cycle)

if cycles:
    print("# Circular Dependencies Detected")
    for c in cycles:
        print(" -> ".join(c))
else:
    print("# No Circular Dependencies Detected")
