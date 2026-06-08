"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[LABELS]: tech-debt, utility
"""

"""
FILE HEADER:
Purpose: Aggressively summarizes markdown files by keeping only headers, all-caps lines, and code blocks.
Subsystem: Memory Lab / Documentation Utility
Dependencies: os, sys
Dependents: Manual use.
Security Considerations: Performs in-place overwrite; creates .bak backup.
Performance Considerations: O(N) complexity.
Risk Score: 3/10 (due to destructive overwrite)
Complexity Score: 2/10
"""
"""Aggressively summarize a markdown text file.

More restrictive than summarize_b3.py — only keeps headers,
all-caps lines, and code blocks. Drops bullet points,
colon-terminated lines, and short sentences.

WARNING: Destructive in-place overwrite. Creates a .bak backup.
"""

import os
import sys

FILENAME = "b3"


def summarize_aggressive(filepath):
    """
    Function: summarize_aggressive
    Purpose: Summarize the given file with aggressive filtering.
    Responsibilities: Read file, keep only headers, all-caps, and code blocks.
    Inputs: filepath (str) - Path to the file to summarize.
    Outputs: Tuple (original_count, filtered_lines)
    Complexity: O(N)
    """
    with open(filepath, "r", encoding="utf-8") as f:
        lines = f.readlines()

    output = []
    in_code_block = False

    for line in lines:
        clean_line = line.strip()

        if clean_line.startswith("```"):
            in_code_block = not in_code_block
            output.append(clean_line)
            continue

        if in_code_block:
            output.append(clean_line)
            continue

        if clean_line.startswith("#") or clean_line.isupper():
            output.append(clean_line)
            continue

    return len(lines), output


def main():
    if len(sys.argv) > 1:
        filepath = sys.argv[1]
    else:
        filepath = FILENAME

    if not os.path.isfile(filepath):
        print(f"Error: File '{filepath}' not found.", file=sys.stderr)
        sys.exit(1)

    backup_path = filepath + ".bak"
    try:
        with open(filepath, "r", encoding="utf-8") as f:
            with open(backup_path, "w", encoding="utf-8") as bak:
                bak.write(f.read())
    except OSError as e:
        print(f"Warning: Could not create backup: {e}", file=sys.stderr)

    original_count, output = summarize_aggressive(filepath)

    with open(filepath, "w", encoding="utf-8") as f:
        for line in output:
            f.write(line + "\n")

    print(f"Reduced from {original_count} to {len(output)} lines.")
    print(f"Backup saved to: {backup_path}")


if __name__ == "__main__":
    main()
