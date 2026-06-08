# Knowledge Transfer: Phoenix.Terminus/Core

## Beginner Explanation
This is the main entry point and the "Body" of the PhoenixOS operating system. It contains the message bus (the nervous system), the monitor (the eyes), and the recovery loops (the immune system).

## Intermediate Explanation
The core files in `phoenix_os/` provide the foundational infrastructure:
- `bus.go`: High-performance, cryptographically secure event routing.
- `monitor.go`: Statistical event analysis.
- `trace.go`: Forensic persistence.
- `recovery.go`: Automated health checks.

## Expert Explanation
The core architecture is built for **High-Assurance Distributed Telemetry**.
- **Bus:** Implements causal hashing and Lamport clocks for strict event ordering. Use ed25519 for message integrity.
- **Telemetry Chain:** Kernel -> Bus -> Monitor -> AI -> Warden -> Kernel.
- **Persistence:** Trace storage is decoupled from the hot-path via bounded queues to ensure forensic logging doesn't block kernel execution.
- **Self-Healing:** Recovery loops operate out-of-band to monitor for "Watchdog" timeouts in the AI reasoning loop.
