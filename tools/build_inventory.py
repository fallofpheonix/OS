import os
from datetime import datetime

DOCS_DIR = 'docs'
INVENTORY_FILE = os.path.join(DOCS_DIR, 'inventory', 'DOCUMENT_INVENTORY.md')

CLASSIFICATIONS = {
    'MASTER_': 'CANONICAL',
    'CONSTITUTION': 'CANONICAL',
    'GLOSSARY': 'CANONICAL',
    'SPECIFICATION': 'CANONICAL',
    'CHANGELOG': 'ACTIVE',
    'TECH_DEBT': 'ACTIVE',
    'SYSTEM_STATUS': 'CANONICAL',
    'README.md': 'REFERENCE',
    'GUIDE': 'REFERENCE',
    'AUDIT': 'ARCHIVED',
    'HISTORICAL': 'ARCHIVED',
    'ADR-': 'REFERENCE'
}

def get_classification(filename):
    for key, val in CLASSIFICATIONS.items():
        if key in filename.upper():
            return val
    return 'ACTIVE'

def build_inventory():
    inventory = [
        "# Document Inventory",
        "",
        "| File Path | Document Type | Owner Module | Classification | Status | Source of Truth |",
        "| :--- | :--- | :--- | :--- | :--- | :--- |"
    ]

    for root, dirs, files in os.walk(DOCS_DIR):
        for file in files:
            if file.endswith('.md'):
                path = os.path.join(root, file)
                rel_path = os.path.relpath(path, '.')
                doc_type = 'Technical'
                if 'GUIDE' in file.upper(): doc_type = 'Guide'
                if 'SPEC' in file.upper(): doc_type = 'Specification'
                if 'ROADMAP' in file.upper(): doc_type = 'Roadmap'
                
                owner = 'General'
                if 'architecture' in root: owner = 'Architecture'
                if 'governance' in root: owner = 'Governance'
                if 'game' in root: owner = 'Game'
                
                classification = get_classification(file)
                status = 'STABLE' if classification == 'CANONICAL' else 'ACTIVE'
                
                # Check for duplicates (filename only for now)
                # (Actual duplicate check will be in Phase 4)
                
                inventory.append(f"| {rel_path} | {doc_type} | {owner} | {classification} | {status} | Self |")

    # Also add root README
    inventory.append(f"| README.md | General | Project | CANONICAL | STABLE | Self |")

    with open(INVENTORY_FILE, 'w') as f:
        f.write('\n'.join(inventory))

if __name__ == "__main__":
    build_inventory()
