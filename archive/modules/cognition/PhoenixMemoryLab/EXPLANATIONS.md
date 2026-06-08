# EXPLANATIONS.md - Memory Lab Utilities

## summarize_b3.py
### Beginner (What it does)
Takes a long text file and makes it shorter by keeping the most important parts like headers, code examples, and lists.
### Intermediate (How it interacts)
It's a standalone tool that reads a file, creates a backup, and then rewrites the file with only structurally significant lines.
### Expert (Architectural role)
A documentation processing utility used to compress large telemetry or archival logs into human-readable summaries while preserving technical context (code blocks, headers).

## summarize_b3_v2.py
### Beginner (What it does)
A "super-summarizer" that only keeps the absolute bare essentials: headers and code.
### Intermediate (How it interacts)
Similar to version 1, but with much stricter rules for what content to keep.
### Expert (Architectural role)
High-compression summary utility for quick architectural overviews of large data sets or logs.
