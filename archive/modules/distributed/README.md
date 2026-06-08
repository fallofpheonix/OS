# PhoenixDistributed — Distributed Coordination Layer

> Proof-of-Authority consensus, peer discovery, and replicated ledger for the PhoenixOS ecosystem.

## Overview

PhoenixDistributed provides the distributed coordination layer for PhoenixOS. It implements PoA (Proof of Authority) consensus, UDP multicast peer discovery, and a consensus-backed distributed ledger for multi-node operation.

**All high-risk state transitions require cluster quorum. No single node can act alone.**

## Repository Structure

```
PhoenixDistributed/
├── consensus/           # PoA consensus engine
│   └── poa.go           # Quorum voting with Ed25519 signatures
├── discovery/           # Peer discovery
│   ├── discovery.go     # PeerDiscovery interface
│   ├── beacon.go        # UDP multicast beacon transport
│   └── registry.go      # Peer registry with reputation
├── identity/            # Node identity management
│   ├── node.go          # AuthorityCertificate, NodeRegistry
│   └── node_test.go     # Registration tests
├── ledger/              # Distributed ledger
│   ├── ledger.go        # ConsensusLedger interface
│   ├── consensus.go     # ConsensusEngine
│   ├── distributed_ledger.go  # Consensus-backed implementation
│   └── stub.go          # Development stub
├── docs/                # Design and benchmark documents
│   ├── design.md
│   └── benchmarks.md
├── go.mod
└── go.sum
```

## Core Principles

1. **2/3 Quorum**: High-risk transitions require ≥66.7% cluster weight
2. **Ed25519 Signatures**: All votes are cryptographically signed
3. **Reputation Weighting**: Node authority is reputation-based
4. **Deterministic Consensus**: Same inputs produce same outcomes

## Build & Test

```bash
# Run all tests
go test ./...

# Run with race detector
go test -race ./...
```

## Dependencies

- **Depends on**: PhoenixCore (ledger, contracts)
- **Depended by**: PhoenixMind (orchestrator), PhoenixOS (integration)

## Invariants

- Quorum requires ≥2/3 of total cluster weight
- All votes must be cryptographically verified
- No single node can authorize high-risk transitions
- Consensus is deterministic given the same input

## License

PhoenixDistributed is part of the PhoenixOS ecosystem.
