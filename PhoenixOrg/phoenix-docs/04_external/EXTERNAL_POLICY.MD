# External Repositories

## 1. Repository Inventory

| Name | Purpose | License | Version | Fork Status | Integration Level |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **Linux Kernel** | Core OS platform | GPL-2.0 | 6.x | Upstream | L2 dependency |
| **BCC / bpftrace** | eBPF toolchain | Apache-2.0 | Latest | Integrated | L2/L3 utility |
| **Protobuf** | Serialization | BSD-3-Clause | 3.x | Standard | Cross-layer comms |
| **gRPC** | RPC Framework | Apache-2.0 | Latest | Standard | L4-L7 communication |
| **Hashicorp Raft** | Consensus library | MPL-2.0 | Latest | Integrated | L7 (Nexus) |
| **go-sqlite3** | Graph lineage storage | MIT | Latest | Standard | L4 (Trace) database |

## 2. Fork Management
- **Security Patches:** All forks must be audited and patched within 48 hours of upstream CVE disclosure.
- **Divergence Policy:** Minimize divergence from upstream to ensure maintainability.
