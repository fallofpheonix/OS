# VPN Component Map

This document details the architectural analysis of each cloned VPN component.

---

## external/openvpn

```yaml
repo: external/openvpn
language: C/C++
purpose: Secure tunneling daemon, providing VPN connectivity.
entrypoints: openvpn (command-line executable)
ipc:
  - Command-line arguments
  - Configuration files (.ovpn)
  - Control channel/socket (potential, requires deeper dive)
auth:
  - X.509 Certificates
  - Pre-shared Keys
network:
  - UDP/TCP transport
  - TUN/TAP virtual network interfaces
config:
  - .ovpn configuration files
  - Command-line options
telemetry:
  - Logging (via syslog or file, configured through CLI/config)
state:
  - Connection status
  - Routing table modifications
  - Peer information
build:
  - Autotools (configure.ac, Makefile.am)
  - CMake (CMakeLists.txt)
runtime:
  - Daemon process (openvpn --daemon)
  - Command-line utility
dependencies:
  - OpenSSL (for cryptographic operations)
  - LZO (for compression)
  - TUN/TAP kernel driver/module
```

---

## external/protonvpn-go-vpn-lib

```yaml
repo: external/protonvpn-go-vpn-lib
language: Go (with bindings for Kotlin/Swift via Go Mobile)
purpose: Clients shared library for Wireguard, Ed25519 key tools, and state machine for Local Agents.
entrypoints:
  - ed25519 package (key generation/conversion)
  - localAgent package (state machine for local agent connections)
ipc:
  - Go package calls (native Go)
  - Native client interfaces (Kotlin/Swift callbacks via Go Mobile)
auth:
  - Client certificates (PEM)
  - Private keys (PEM)
  - Server root certificates
  - Ed25519 key pairs
network:
  - Wireguard VPN tunnel functionality (core purpose)
config:
  - Build config for gomobile (apple.json)
  - Runtime parameters for AgentConnection (certs, keys, server address)
telemetry:
  - NativeClient.log callback (debug logs)
  - NativeClient.onError callback (error messages)
state:
  - Local agent connection state machine (e.g., stateConnecting, stateHardJailed)
  - agent.status (last received status message)
build:
  - Go Modules (go.mod, go.sum)
  - gomobile (for Android .aar and Apple framework binding)
  - Gradle (build.gradle) for Android
  - Custom `make` script (Proton Go Mobile build script) for Apple
runtime:
  - Embedded library within client applications (Android, iOS, macOS)
dependencies:
  - go-srp (external Go library, expected to be cloned into project root)
  - Standard Go libraries
```

---
