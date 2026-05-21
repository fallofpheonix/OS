# Stage 05: Reverse Engineering and Malware Analysis

## Purpose

Build capability to inspect binaries without source code, analyze malware behavior, write detection signatures, and produce defensible technical reports.

## Scope

- Binary formats: ELF, PE, Mach-O.
- Static analysis: disassembly, decompilation, strings, imports, metadata.
- Dynamic analysis: debuggers, syscall tracing, behavioral analysis, sandboxing.
- Packing, obfuscation, and anti-analysis.
- Memory forensics.
- x86/x64 instruction and architecture analysis.
- Malware classification and ATT&CK mapping.
- YARA and network/signature-based detection.

## Classification

- Type: `SECURITY_RESEARCH`
- Status: `RESEARCH_ONLY`
- Difficulty: expert
- Estimated duration: 8-10 weeks
- Upstream prerequisites:
  - Stage 01 System Internals
  - Stage 04 Security
- Downstream blockers:
  - Stage 06 Forensics
  - Stage 07 Threat Intelligence
  - Stage 12 Security AI

## Research Modules

| Module | Path |
|---|---|
| Phase 7 Research Plan | `phase_07_reverse_engineering_malware.md` |
| Phase 7 Build Gate | `build_gate.md` |

## Internal Dependency Order

```text
Binary formats
-> Static analysis
-> Debugging and dynamic tracing
-> Packing and obfuscation
-> Memory forensics
-> YARA and detection signatures
-> Malware classification and reporting
```

## Gate

Do not store live malware samples in this repository. Store only hashes, IOCs, reports, rules, and sanitized metadata.
