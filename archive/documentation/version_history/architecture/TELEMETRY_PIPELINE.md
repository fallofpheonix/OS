# PhoenixOS Telemetry Pipeline

This document outlines the path of data from the Linux Kernel to the Phoenix Bus.

## Pipeline Architecture

1.  **Ingestion (L2):** eBPF probes capture syscall events in kernel space.
2.  **Buffering:** Events are passed to user space via a high-performance **eBPF Ring Buffer**.
3.  **Collection:** The `phoenix_os/kernel_agent` reads the ring buffer and serializes the raw events into the [Kernel Event Schema](./EVENT_SCHEMA.md).
4.  **Normalization (L3):** Raw syscalls are mapped to High-Level Events (e.g., `execve` -> `PROCESS_START`).
5.  **Distribution:** Normalized events are published to the **Phoenix Bus** for Trace (L4) and Warden (L5) consumption.

## Performance Requirements
- **Latency:** < 100μs from kernel trigger to user-space ingestion.
- **Throughput:** Must handle 10,000 events/sec without buffer overflow.
- **Integrity:** Sequential consistency must be maintained across core migrations.
