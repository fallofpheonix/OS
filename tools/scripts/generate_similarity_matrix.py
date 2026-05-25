import os
import re

def get_tokens(text):
    text = re.sub(r'[^a-zA-Z0-9\s]', '', text.lower())
    return set(text.split())

def calculate_similarity(setA, setB):
    if not setA or not setB: return 0
    intersection = len(setA.intersection(setB))
    union = len(setA.union(setB))
    return intersection / union

def main():
    root = '.'
    md_files = []
    for dirpath, dirnames, filenames in os.walk(root):
        if any(ex in dirpath for ex in ['.git', '.venv', 'node_modules', 'archive']):
            continue
        for f in filenames:
            if f.endswith('.md'):
                md_files.append(os.path.join(dirpath, f))

    file_tokens = {}
    for path in md_files:
        try:
            with open(path, 'r', encoding='utf-8', errors='ignore') as f:
                file_tokens[path] = get_tokens(f.read())
        except:
            pass

    matrix = []
    processed = set()
    for pathA in md_files:
        for pathB in md_files:
            if pathA == pathB: continue
            pair = tuple(sorted((pathA, pathB)))
            if pair in processed: continue
            processed.add(pair)
            
            sim = calculate_similarity(file_tokens.get(pathA, set()), file_tokens.get(pathB, set()))
            if sim > 0.4: # Threshold for similarity
                matrix.append({
                    'fileA': os.path.relpath(pathA, root),
                    'fileB': os.path.relpath(pathB, root),
                    'sim': sim,
                    'merge': 'Yes' if sim > 0.8 else 'Maybe'
                })

    with open('DOC_SIMILARITY_MATRIX.md', 'w') as f:
        f.write("# Document Similarity Matrix\n\n")
        f.write("| fileA | fileB | semantic_similarity | merge_candidate |\n")
        f.write("|---|---|---|---|\n")
        for m in sorted(matrix, key=lambda x: x['sim'], reverse=True):
            f.write(f"| {m['fileA']} | {m['fileB']} | {m['sim']:.2f} | {m['merge']} |\n")

if __name__ == '__main__':
    main()
