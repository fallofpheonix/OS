---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Phoenix Security — Attack Surface Analysis

> Last verified: 2026-06-04

This document outlines the entry points, trust boundaries, and protocols exposing the Phoenix OS runtime to potential compromises.

## 1. System Call Interface
- **Exposure**: Processes running in user namespaces call Linux syscalls.
- **Risk**: Kernel exploits bypassing sandbox boundaries.
- **Control**: eBPF probes monitor syscall execution and block unauthorized attempts.

## 2. Event Ingestion API
- **Exposure**: Incoming events over network sockets or IPC.
- **Risk**: Malformed JSON payloads causing resource exhaustion during parsing.
- **Control**: Enforce strict JSON length limits and schema validation inside `events/` before processing.

## 3. Storage Subsystem
- **Exposure**: Database and ledger log files on disk.
- **Risk**: Modifying ledger files directly.
- **Control**: Cryptographic integrity checks verify parent-hashes recursively on startup.

## 4. Telemetry Normalizer Socket
- **Exposure**: UNIX domain socket transmitting event streams.
- **Risk**: Injection of fake audit metrics.
- **Control**: Secure file permissions limit socket access to root processes only.
