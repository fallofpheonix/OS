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

## Artifact
An immutable, signed bundle of code, configuration, or data referenced by an Event in the Ledger, required for state reconstruction.
