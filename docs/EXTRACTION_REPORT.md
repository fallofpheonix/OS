This file will contain details about extracted assets. It will be populated during the extraction phase.

---

### Extracted from pheonixos/archive/old_repos/PhoenixOrg/phoenix-core

*   **Source:** `archive/old_repos/PhoenixOrg/phoenix-core/main.go`
    *   **Destination:** `core/phoenix_os/main.go`
    *   **Description:** Main entry point for the phoenix-core runtime.
*   **Source:** `archive/old_repos/PhoenixOrg/phoenix-core/go.mod`
    *   **Destination:** `core/phoenix_os/go.mod`
    *   **Description:** Go module definition for phoenix-core.
*   **Source:** `archive/old_repos/PhoenixOrg/phoenix-core/INVARIANTS.md`
    *   **Destination:** `docs/INVARIANTS.md`
    *   **Description:** Formal mathematical and systems invariants for deterministic replayability.
*   **Source:** `archive/old_repos/PhoenixOrg/phoenix-core/security/` (full directory)
    *   **Destination:** `security/`
    *   **Description:** Comprehensive security module, including control, detections, EDR, IDS, IPS, SIEM, SOAR, XDR, and YARA components.
*   **Source:** `archive/old_repos/PhoenixOrg/phoenix-core/observability/`
    *   **Destination:** None
    *   **Description:** Subdirectory found to be empty (only contains `.gitkeep`), no functional assets extracted.

---

### Cloned External VPN Components

*   **Source:** `https://github.com/openvpn/openvpn.git`
    *   **Destination:** `external/openvpn`
    *   **Description:** Official OpenVPN repository.
*   **Source:** `https://github.com/ProtonVPN/go-vpn-lib.git`
    *   **Destination:** `external/protonvpn-go-vpn-lib`
    *   **Description:** ProtonVPN's Go VPN library.
*   **Source:** `https://github.com/ProtonVPN/python-proton-vpn-api-core.git`
    *   **Destination:** `external/protonvpn-python-api-core`
    *   **Description:** ProtonVPN's Python API core logic.
*   **Source:** `https://github.com/ProtonVPN/proton-vpn-daemon.git`
    *   **Destination:** `external/protonvpn-daemon`
    *   **Description:** ProtonVPN's system daemon for Linux VPN connections.
*   **Source:** `https://github.com/ProtonVPN/local-agent-rs.git`
    *   **Destination:** `external/protonvpn-local-agent-rs`
    *   **Description:** ProtonVPN's local agent written in Rust.