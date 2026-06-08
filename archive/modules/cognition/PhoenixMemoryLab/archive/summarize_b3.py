"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[LABELS]: tech-debt, utility
"""

"""
FILE HEADER:
Purpose: Summarizes markdown files by filtering content based on structural rules (headers, code blocks, bullet points, etc.).
Subsystem: Memory Lab / Documentation Utility
Dependencies: os, sys
Dependents: Manual use, potential CI scripts.
Security Considerations: Performs in-place overwrite; creates .bak backup to prevent data loss.
Performance Considerations: O(N) complexity where N is the number of lines in the file.
Risk Score: 3/10 (due to destructive overwrite)
Complexity Score: 3/10
"""
"""Summarize a markdown text file by filtering content based on structural rules.

Reads the file specified by FILENAME, keeps code blocks, headers, all-caps
lines, bullet points, colon-terminated lines, and short sentences (<=5 words
ending with period). Writes the filtered content back to the same file.

WARNING: Destructive in-place overwrite. Creates a .bak backup.
"""

import os
import sys

FILENAME = "b3"


def summarize(filepath):
    """
    Function: summarize
    Purpose: Summarize the given file by keeping structurally significant lines.
    Responsibilities: Read file, filter lines based on rules, return original and filtered counts.
    Inputs: filepath (str) - Path to the file to summarize.
    Outputs: Tuple (original_count, filtered_lines)
    Complexity: O(N) where N is the number of lines.
    """
    with open(filepath, "r", encoding="utf-8") as f:
        lines = f.readlines()

    output = []
    in_code_block = False

    for line in lines:
        clean_line = line.strip()
        if not clean_line:
            continue

        if clean_line.startswith("```"):
            in_code_block = not in_code_block
            output.append(clean_line)
            continue

        if in_code_block:
            output.append(clean_line)
            continue

        if (clean_line.startswith("#") or
                clean_line.isupper() or
                clean_line.startswith("-") or
                clean_line.startswith("*") or
                clean_line.endswith(":")):
            output.append(clean_line)
            continue

        words = clean_line.split()
        if len(words) <= 5 and clean_line.endswith("."):
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

    original_count, output = summarize(filepath)

    with open(filepath, "w", encoding="utf-8") as f:
        for line in output:
            f.write(line + "\n")

    print(f"Reduced from {original_count} to {len(output)} lines.")
    print(f"Backup saved to: {backup_path}")


if __name__ == "__main__":
    main()
