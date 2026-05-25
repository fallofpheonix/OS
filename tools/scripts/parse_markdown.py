import csv

with open('ecosystem/inventory/portfolio_inventory.csv', 'r') as f:
    reader = csv.DictReader(f)
    repos = list(reader)

with open('ecosystem/inventory/repo_status.md', 'w') as f:
    f.write('# Portfolio Repository Status\n\n')
    f.write('| Repository | Domain | Language | Updated | Archived | Initial Status |\n')
    f.write('|---|---|---|---|---|---|\n')
    for r in repos:
        # Simple heuristic for initial status
        status = 'ACTIVE'
        if r['is_archived'] == 'true' or r['domain'] == 'Archive / Legacy':
            status = 'ARCHIVE'
        elif r['domain'] == 'Science / Physics / Simulation':
            status = 'RESEARCH'
        elif r['domain'] == 'Products / UX' or r['domain'] == 'Health / Bio':
            status = 'PRODUCT'
        elif r['domain'] == 'Runtime Core' or r['domain'] == 'Security' or r['domain'] == 'Infrastructure':
            status = 'INTEGRATION'
        
        # Override some based on user prompt examples
        if r['repo_name'] == 'brain':
            status = 'EXPERIMENTAL'
        elif r['repo_name'] == 'physics':
            status = 'RESEARCH'
        elif r['repo_name'] == 'legacy':
            status = 'ARCHIVE'
        elif r['repo_name'] == 'my-portfolio':
            status = 'PRODUCT'
        
        f.write(f"| {r['repo_name']} | {r['domain']} | {r['language']} | {r['updated_at'][:10]} | {r['is_archived']} | {status} |\n")

with open('ecosystem/inventory/dependency_scan.md', 'w') as f:
    f.write('# Portfolio Dependency Scan Summary\n\n')
    f.write('*Note: This is a pre-clone metadata scan based on GitHub language stats.*\n\n')
    langs = {}
    for r in repos:
        l = r['language']
        langs[l] = langs.get(l, 0) + 1
    
    for l, c in langs.items():
        f.write(f"- **{l}**: {c} repositories\n")
