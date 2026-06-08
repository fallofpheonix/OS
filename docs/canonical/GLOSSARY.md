---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Glossary of Terms

## Authority
The explicit, cryptographically verifiable right to perform a specific action within a bounded scope. Authority is never implicit.

## Identity
A cryptographic entity (keypair or signed certificate) that can hold and delegate Authority.

## Ledger
The append-only, immutable sequence of all authoritative Events. It is the sole source of truth for the state of PhoenixOS.

## Truth
The canonical state of the system at any given Logical Time, as derived deterministically by replaying the Ledger. 

## Replay
The process of deterministically reconstructing the system state from the Genesis block by sequentially applying the Events recorded in the Ledger through the replay contract.

## Recovery
The process of fully resurrecting a destroyed or corrupted node using only its preserved Ledger and verified Artifacts.

## Containment
The enforcement of strict operational boundaries around a process or module to prevent unauthorized mutation of system state or environment.

## Contract
The authoritative public interface for a bounded context. Contracts define the only permitted source of truth for cross-boundary behavior.

## Federation
A network of independent PhoenixOS nodes that mutually verify each other's adherence to the Constitution via Proof Exchange.

## Causal Frontier (Heads)
The set of cryptographic hashes representing the "leaf" nodes of the Evidence Merkle DAG. Resuming replay from a snapshot requires the Causal Frontier to correctly link future events to history.

## Evidence Merkle DAG
The structural implementation of the Ledger where events are cryptographically linked to multiple parents, allowing for branching and concurrent execution traces while maintaining absolute causal integrity.

## Fixed-Point Arithmetic
A numeric representation using integer scaling (e.g., 10^6) and truncation rounding. Used to eliminate non-determinism inherent in floating-point operations across different hardware architectures.

## Sovereign Payload
 A 43-byte versioned binary format used for Warden forensic records. Designed for O(1) parsing, zero-allocation serialization, and bit-perfect cross-platform hashing.

## Write-Ahead Enforcement
A consistency protocol where system state mutations are only applied to memory after the corresponding intent has been successfully persisted and hashed into the Ledger.
