# RFC: Phoenix Nexus

## 1. Description
Phoenix Nexus provides L7 swarm coordination. It uses a gossip-style protocol to propagate "Security Disorder Index" (SDI) values and "Arbiter" policies across the network.

## 2. Specification
- **Protocol:** UDP-based Gossip.
- **State:** Conflict-free Replicated Data Types (CRDTs).
- **Consensus:** Quorum-based validation for critical containment actions.

## 3. Interface
```go
func PropagateState(sdi float64, nodeID string)
func RequestConsensus(action string) bool
```
