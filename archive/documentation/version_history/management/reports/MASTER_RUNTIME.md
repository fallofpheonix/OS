# MASTER RUNTIME

## Status: UNIFIED

### Source: 02_docs/01_architecture/00-overview.md

# Project Overview

## Project Name

PhoenixOS

## Type

Custom operating system or Linux-derived distribution.

## Development Models

### Scratch Build

- Bootloader integration.
- Kernel.
- Drivers.
- Userspace.
- Filesystem.
- Package system.

### Linux-Derived Build

Base options:

- Arch Linux.
- Kali Linux.
- Debian.
- Buildroot.
- Linux From Scratch.

## Initial Target

| Field | Value |
|---|---|
| Architecture | `x86_64` |
| Boot mode | UEFI |
| Kernel type | Hybrid initially |
| Primary languages | C, C++, optional Rust |
| First artifact | Bootable ISO |
| First runtime | Shell |

## Goal

Build an educational and experimental OS with modular architecture and a path from Linux-derived prototype to custom kernel work.



### Source: RUNTIME_GRAPH.md

# Runtime Graph

```mermaid
graph TD
    agents --> src
    ai --> arbiter
    ai --> bus
    ai --> monitor
    ai --> serialization
    ai --> src
    ai --> tcs
    ai --> trace
    ai --> warden
    arbiter --> bus
    arbiter --> monitor
    arbiter --> warden
    boot --> src
    concurrency --> resource
    contracts --> state
    control --> types
    decision --> arbiter
    decision --> bus
    decision --> replay
    decision --> warden
    detector --> events
    detector --> process_lineage
    forensics --> types
    game --> src
    game --> types
    game --> warden
    gov --> ai
    gov --> arbiter
    gov --> bus
    gov --> common
    gov --> guard
    gov --> ledger
    gov --> monitor
    gov --> tcs
    gov --> trace
    gov --> warden
    graph --> types
    guard --> bus
    kernel --> bus
    kernel --> phoenix.kernel.ebpf
    kernel --> phoenix.normalizer
    main --> agents
    main --> ai
    main --> arbiter
    main --> boot
    main --> bus
    main --> control
    main --> detector
    main --> entropy_engine
    main --> events
    main --> forensics
    main --> game
    main --> graph
    main --> guard
    main --> kernel
    main --> logical_clock
    main --> monitor
    main --> normalizer
    main --> phoenix.bus
    main --> phoenix.sys.boot
    main --> phoenix_trace.db
    main --> physics
    main --> process_graphs
    main --> process_lineage
    main --> resource
    main --> sandbox
    main --> src
    main --> stackelberg
    main --> tcs
    main --> telemetry
    main --> trace
    main --> types
    main --> warden
    monitor --> bus
    monitor --> kalman
    monitor --> phoenix.monitor
    normalizer --> bus
    normalizer --> clock
    physics --> disorder
    physics --> monitor
    physics --> types
    replay --> bus
    rollback --> containment
    rollback --> file
    rollback --> network
    router --> phoenix-mind
    security --> src
    snapshot --> truth
    telemetry --> phoenix-dev-node
    telemetry --> types
    tooling --> replay
    tooling --> src
    trace --> bus
    truth --> phoenix-contracts
    verifier --> 
				isInCore := strings.HasPrefix(relPath, 
    verifier -->  as anything outside the agents module or phoenix_os core (sentinel, bus, monitor, warden, arbiter).
			if impPath == 
    verifier --> verifier can import Kernel Agent
			if impPath == 
    warden --> bus
    warden --> phoenix-contracts
```


### Source: WORKING_MODEL.md

# Working Model

## Implementation Reality

### Advisor
Status: **EMPTY**

### Arbiter
Status: **EMPTY**

### Containment
Status: **ACTIVE**

### Distributed
Status: **EMPTY**

### Memory
Status: **EMPTY**

### Metrics
Status: **EMPTY**

### Quantum
Status: **ACTIVE**

### Recovery
Status: **ACTIVE**

### Replay
Status: **EMPTY**

### Runtime
Status: **ACTIVE**

### Truth
Status: **ACTIVE**

### Validation
Status: **EMPTY**

### Warden
Status: **PARTIAL**



