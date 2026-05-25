import os

def main():
    # 1. Update RESEARCH_INDEX.md with refined rules
    # REJECT IF: no runtime relevance, no validation path, no benchmark, no integration target
    research_items = []
    if os.path.exists('RESEARCH_INDEX.md'):
        with open('RESEARCH_INDEX.md', 'r') as f:
            lines = f.readlines()[2:] # Skip header
            for line in lines:
                parts = line.split('|')
                if len(parts) >= 4:
                    path = parts[1].strip()
                    status = parts[2].strip()
                    reason = parts[3].strip()
                    
                    # Refine rejection
                    if 'agi' in reason.lower() or 'consciousness' in reason.lower():
                        status = 'REJECTED'
                        reason = "No runtime relevance / speculative"
                    
                    research_items.append({'path': path, 'status': status, 'reason': reason})

    with open('RESEARCH_INDEX.md', 'w') as f:
        f.write("# Research Index (Refined)\n\n")
        f.write("| Path | Status | Reason |\n")
        f.write("|---|---|---|\n")
        for item in research_items:
            f.write(f"| {item['path']} | {item['status']} | {item['reason']} |\n")

    # 2. Generate Master Documents (Merges)
    masters = {
        'F0_MASTER.md': ['02_docs/00_governance/OPERATIONAL_ROADMAP.md', '02_docs/00_governance/F0_EXIT_CHECKLIST.md', 'F0_EXIT_STATUS.md'],
        'MASTER_ROADMAP.md': ['00_program_management/roadmap/main_roadmap.md', '00_program_management/roadmap/TASK_BREAKDOWN.md', 'UPDATED_IMPLEMENTATION_PLAN.md'],
        'MASTER_VALIDATION.md': ['02_docs/00_governance/meta/TESTING_ROADMAP.md', 'VALIDATION_COVERAGE.md', 'FINAL_AUDIT.md'],
        'MASTER_RUNTIME.md': ['02_docs/01_architecture/00-overview.md', 'RUNTIME_GRAPH.md', 'WORKING_MODEL.md'],
        'MASTER_RESEARCH.md': ['RESEARCH_INDEX.md', '06_research/README.md']
    }

    for master, sources in masters.items():
        with open(master, 'w') as f_out:
            f_out.write(f"# {master.replace('.md', '').replace('_', ' ')}\n\n")
            f_out.write("## Status: UNIFIED\n\n")
            for src in sources:
                if os.path.exists(src):
                    f_out.write(f"### Source: {src}\n\n")
                    try:
                        with open(src, 'r') as f_in:
                            content = f_in.read()
                            # Strip frontmatter or redundant titles if needed
                            f_out.write(content + "\n\n")
                    except:
                        f_out.write(f"*Error reading {src}*\n\n")
                else:
                    f_out.write(f"### Source: {src} (NOT FOUND)\n\n")

    # 3. Final maturity decision
    # Evidence based on P2 test run
    with open('FINAL_DECISION.md', 'w') as f:
        f.write("# Final Decision & Project Maturity\n\n")
        f.write("- **Current Maturity**: Runtime Security Research Platform (Stage A Hardening)\n")
        f.write("- **Working Runtime %**: UNKNOWN (Needs measured line coverage)\n")
        f.write("- **Research %**: UNKNOWN (Needs evidence of integration targets)\n")
        f.write("- **Dead Docs %**: UNKNOWN (Needs measured semantic overlap across all layers)\n")
        f.write("- **Userspace Replay**: PARTIAL VERIFIED (100% precision on large traces)\n")
        f.write("- **Global Determinism**: UNVERIFIED\n")
        f.write("- **F0 Status**: PARTIAL (Core substrate verified, security tests compile but need deeper logic verification)\n")
        f.write("- **F1 Status**: LOCKED (Blockers resolved: build errors. Remaining: unverified kernel determinism)\n\n")
        f.write("## Final Action\n")
        f.write("F1 remains LOCKED until global determinism is verified under kernel jitter simulation.\n")

if __name__ == '__main__':
    main()
