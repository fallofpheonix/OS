---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Phoenix Security — Threat Model

> Last verified: 2026-06-04

This document defines the threat model for the Phoenix OS runtime based on the STRIDE framework.

## Threat Classification

### 1. Spoofing Identity
- **Threat**: Attackers inject mock events posing as trusted authorities.
- **Mitigation**: Every event inside the ledger envelope is signed with the authority key. The `Serializer` interface and `Ledger` verify signatures on append.

### 2. Tampering with Data
- **Threat**: Direct filesystem modification of the ledger log or memory variables.
- **Mitigation**: The ledger computes a linear Parent-Hash chain. Tampering breaks the verification hashes, halting consensus.

### 3. Repudiation
- **Threat**: Malicious nodes deny performing state changes.
- **Mitigation**: Event envelopes log cryptographic validation hashes linking the state transition before and after values directly to node signatures.

### 4. Information Disclosure
- **Threat**: Leakage of cognitive memory state to user space.
- **Mitigation**: Containment level sandbox isolates mounts, blocking reading memory map dumps.

### 5. Denial of Service
- **Threat**: User code exhausting process handles or memory (Forkbomb / Memory exhaustion).
- **Mitigation**: The `ResourceAllocator` bounds ledger storage. The security warden actively terminates high-disorder processes.

### 6. Elevation of Privilege
- **Threat**: Compromised process breaking sandbox barriers.
- **Mitigation**: eBPF monitors syscall namespaces and denies unrecognized system calls.
