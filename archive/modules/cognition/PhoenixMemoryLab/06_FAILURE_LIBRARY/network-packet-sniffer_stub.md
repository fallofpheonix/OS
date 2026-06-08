# Failure: Scapy Sudo PermissionError

## Date
2026-05-12

## Project
[[05_PROJECTS/ACTIVE/network-packet-sniffer]]

## Environment
MacOS Sandbox Python 3.13

## Symptom
Running the Python script locally immediately exits with `PermissionError` when calling `sniff()`.

## Timeline
Month 3 network tier implementation. Built the `packet-sniffer` module. Attempted to run it with `python3 main.py`. Instantly crashed.

## Root Cause
Scapy attempts to open raw sockets on the BPF (Berkeley Packet Filter) interface to intercept raw network traffic. macOS restricts raw socket creation to the `root` user only.

## Fix
Executed the script with `sudo python3 main.py` or configured the deployment environment to run the sniffer daemon with elevated privileges.

## Why It Was Hard To Find
The exception stack trace points deep inside Scapy's OS abstraction layer rather than explicitly stating "Run with Sudo".

## Prevention
Added a runtime check inside the `packet-sniffer` module to detect `os.geteuid() == 0` and throw a descriptive human-readable error before invoking Scapy.

## What I Should Have Caught Earlier
Network tooling fundamentally operates at OSI Layer 2/3, which OS security models naturally protect.

## Pattern This Belongs To
[[Permission Errors]]

## Related Concepts
- [[OSI Model]]
- [[Raw Sockets]]

## Related Failures
- [[iptables-firewall_stub]]
