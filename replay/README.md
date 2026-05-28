# Replay Storage Architecture

This directory (`phoenix_os/replay/`) contains the storage tiers and related components for managing replay data within PhoenixOS. The architecture is designed to optimize for access speed and retention based on the criticality and age of replay data.

## Storage Tiers

### 1. `HOT/`
- **Purpose:** Stores replay data actively being generated or consumed. This tier is for live, in-memory, or very fast disk-backed replay streams.
- **Retention:** Live replay. Data in this tier is volatile and optimized for high-speed read/write access.
- **Use Case:** Current replay sessions, real-time analysis, immediate fault injection.

### 2. `WARM/`
- **Purpose:** Stores recent history of replay data. This tier serves as a buffer between live data and long-term archives.
- **Retention:** Recent history. Data is retained for a short to medium duration (e.g., hours to days) and is readily accessible.
- **Use Case:** Post-mortem debugging of recent events, short-term analytics, replay validation against current system state.

### 3. `COLD/`
- **Purpose:** Stores archived replay data for long-term retention and historical analysis.
- **Retention:** Archive. Data is retained for long durations (e.g., weeks, months, years) and may be stored on cost-effective, slower storage.
- **Use Case:** Compliance, long-term trend analysis, historical anomaly detection, deep forensic investigations.

## Auxiliary Storage Components

### 1. `index/`
- **Purpose:** Contains metadata and indices for navigating and querying replay data across all tiers. This includes mappings of replay IDs to storage locations, timestamps, and other relevant session details.
- **Content:** Index files, metadata databases, lookup tables.

### 2. `snapshots/`
- **Purpose:** Stores periodic snapshots of system state during replay. These snapshots enable efficient rollback and comparison operations without replaying from the very beginning.
- **Content:** Serialized system states, checkpoints.

### 3. `hashes/`
- **Purpose:** Stores cryptographic hashes (e.g., SHA-256) of replay data segments or system states. These hashes are crucial for verifying replay determinism and detecting divergence.
- **Content:** Hash chains, state hashes, input hashes.

## Integration
The replay flow utilizes these storage tiers:
```
input
   ↓
telemetry
   ↓
logical clock
   ↓
event ordering
   ↓
replay engine
   ↓
state hash (stored in hashes/)
   ↓
identity check
```
The output of the `state hash` step is stored in the `hashes/` directory.

---
**Note:** This documentation outlines the conceptual architecture. Specific implementation details (e.g., file formats, database choices) will be determined by the respective components interacting with these storage tiers.
