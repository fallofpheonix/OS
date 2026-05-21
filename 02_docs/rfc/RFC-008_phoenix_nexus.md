# RFC-008: Simulation Telemetry Agent

## Status
Approved

## 1. Purpose
This RFC specifies the macOS-native Simulation Telemetry Agent. Because production eBPF telemetry is Linux-specific and cannot run directly on macOS development environments, the Phoenix Nexus is required to mock system calls, process lineage trees, and file activities. This provides a high-throughput, schema-conformant telemetry baseline to validate the Phoenix Bus normalizers, correlation logic, and local SOC dashboard on the macOS host.

## 2. Architecture & Data Flow

```mermaid
graph TD
    Sim[Simulation Engine] --> Scenario[Scenario Generator]
    Scenario --> Gen[Event Encoder (JSON)]
    Gen -- "Unix Domain Socket / TCP" --> Bus[Phoenix Bus Ingestion]
```

---

## 3. Simulation Scenarios

The simulator supports two modes of execution:

### 3.1 Normal Baseline Scenario
Generates standard background system noise:
*   Periodic process spawns (e.g., shell commands, background daemons).
*   Mock system health checks writing to `/var/log/syslog`.
*   Web server handling network requests (spawning worker threads, accepting mock TCP states).

### 3.2 Ransomware Threat Scenario
Simulates a multi-stage ransomware attack:
1.  **Stage 1 (Persistence/Privilege):** A system daemon (e.g., `cron`) spawns an unexpected shell process (e.g. `/bin/bash` with suspicious parameters).
2.  **Stage 2 (Discovery):** The shell runs script files to index filesystem paths.
3.  **Stage 3 (Encryption):** Spawns an encryption utility (e.g. `/usr/bin/gpg` or a custom `encryptor`). The encryptor opens multiple documents and writes high-entropy buffers to simulate encryption.
4.  **Stage 4 (Impact):** The original shell deletes backup volumes or files (`vfs_unlink`).

---

## 4. Transmission & IPC
*   **Socket Path:** Default is a project-local Unix Domain Socket (UDS) located at `./14_experiments/telemetry_replay/sentinel.sock` or fallback TCP port `127.0.0.1:9092`.
*   **Protocol:** Framed JSON over UDS. Each event is serialized to JSON and terminated with a newline character (`\n`).

---

## 5. Performance Targets & Validation Gates
*   **Throughput Target:** Able to emit up to **100,000 events/second** to load-test the Phoenix Bus.
*   **Resource Budget:** CPU usage under **3%** on macOS when generating standard baseline noise (100 events/second).
*   **Event Integrity:** 0% corruption rate on generated payloads.
