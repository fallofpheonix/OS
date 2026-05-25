# Phase 7 Build Gate: Reverse Engineering and Malware Analysis

## Static Analysis

- [ ] Analyze 3-4 binaries of increasing complexity.
- [ ] Identify file type and architecture.
- [ ] Extract strings and hardcoded data.
- [ ] Analyze imports and dependencies.
- [ ] Identify entry point and main function.
- [ ] Reconstruct function signatures.
- [ ] Identify key algorithms and crypto routines.
- [ ] Draw control flow graph for key functions.
- [ ] Estimate packing or obfuscation level.

## Decompilation

- [ ] Use Ghidra or equivalent decompiler.
- [ ] Reconstruct variable types and purposes.
- [ ] Identify input validation and bounds checks.
- [ ] Spot potential vulnerabilities.
- [ ] Explain key code paths.

## Dynamic Analysis

- [ ] Set breakpoints at entry point and key functions.
- [ ] Step through execution.
- [ ] Inspect and modify registers and memory.
- [ ] Trace syscalls.
- [ ] Monitor file and network activity.
- [ ] Record process tree and IPC behavior.

## Behavioral Malware Analysis

- [ ] Run sample only in controlled sandbox or VM.
- [ ] Document all system changes.
- [ ] Capture network connections.
- [ ] Identify C2 domains or IPs if present.
- [ ] Reconstruct attack chain.
- [ ] Extract IOCs.

## YARA Rules

- [ ] Write one YARA rule for a specific family or behavior.
- [ ] Include string patterns.
- [ ] Include hex patterns.
- [ ] Include file characteristics such as entropy, size, or sections.
- [ ] Test against positive samples.
- [ ] Test against benign samples.
- [ ] Document false positive rate.

Rule candidates:

- Packer or obfuscator signature.
- Cryptographic operation pattern.
- Code injection pattern.
- C2 communication artifact.

## Memory Forensics

- [ ] Identify process list and tree from memory dump.
- [ ] Extract suspicious processes.
- [ ] Dump process memory.
- [ ] Search strings and patterns in memory.
- [ ] Identify injected code or shellcode.
- [ ] Reconstruct command-line arguments.
- [ ] Recover decrypted strings if present.
- [ ] Use Volatility or equivalent for process/module/network analysis.

## Evasion

- [ ] Identify debugger checks.
- [ ] Identify emulator or VM checks.
- [ ] Patch anti-analysis checks.
- [ ] Bypass string encryption.
- [ ] Unpack or dump decompressed binary.
- [ ] Document bypass method.

## Malware Classification

- [ ] Determine malware type.
- [ ] Identify behavior and objectives.
- [ ] Map behavior to MITRE ATT&CK.
- [ ] Assess threat level and target.
- [ ] Create final report.

## Report

- [ ] Executive summary.
- [ ] Hashes and sample metadata.
- [ ] IOC table.
- [ ] Behavior analysis.
- [ ] Code analysis.
- [ ] Detection signatures.
- [ ] Remediation recommendations.
- [ ] Appendix with relevant disassembly or decompiled excerpts.

## Exit Criteria

- [ ] Reports promoted to `08_forensics/reports/`.
- [ ] Detection rules promoted to `07_security/yara/` or `07_security/detections/`.
- [ ] No live malware committed.
- [ ] Sample references are hash-only or external controlled-storage references.
