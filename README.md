# Pheonix OS (formerly SentinelOS)

![Pheonix Logo](assets/pheonix/logo_simple.svg)

Pheonix is a security-focused operating system project (formerly SentinelOS) built from scratch or as a custom Linux derived from Arch/Kali.

Project symbols (SVGs): `assets/pheonix/logo_simple.svg`, `assets/pheonix/logo_fire.svg`, `assets/pheonix/logo_shield.svg` — raster fallbacks: `assets/pheonix/raster/*.png`

This repository contains the minimal documentation, source layout, and build planning files for:

- Bare-metal OS development.
- Linux From Scratch style custom Linux.
- Arch-based custom ISO work.
- Kali-based security-focused live image work.
- Cyber-defense OS architecture with AI/ML support.

## Goals

- Boot a 64-bit x86 system in QEMU.
- Provide a minimal shell and basic filesystem.
- Wire together toolchain, kernel, bootloader, root filesystem, and init.
- Keep Arch/Kali-derived paths documented separately from scratch kernel work.
- Integrate IDS, IPS, malware detection, forensics, sandboxing, and AI-assisted defense.

## Status

- [ ] Boot sector loads in QEMU.
- [ ] C kernel starts and prints to screen.
- [ ] Simple shell or initramfs exists.
- [ ] Linux-derived root filesystem builds.
- [ ] ISO image boots.

## Prerequisites

Bare-metal path:

- `gcc` or cross `i686-elf-gcc` / `x86_64-elf-gcc`.
- `nasm`.
- `ld`.
- `grub-mkrescue` or `xorriso`.
- `qemu-system-x86_64`.

Linux-derived path:

- `make`.
- `patch`.
- `binutils`.
- `glibc` or cross libc build.
- Linux kernel source.
- BusyBox.
- Distro-specific image tools.

## Build And Run

### From-Scratch OS

```sh
make clean
make
make run
```

### Linux-Based Distro

Read [docs/01-lfs-build.md](docs/01-lfs-build.md).

## Path Selection

| Path | Use When | Primary Cost | Primary Risk |
|---|---|---:|---|
| From scratch | Need kernel/runtime control, research OS, embedded target | Very high | Hardware, drivers, toolchain, filesystem |
| Arch-based | Need general-purpose Linux with custom packages and defaults | Medium | Release drift, package conflicts |
| Kali-based | Need offensive/security tooling preloaded | Medium | Legal scope, unsafe defaults, package bloat |

## Repository Structure

```text
.
├── boot/
├── config/
├── docs/
│   └── decisions/
├── drivers/
├── iso/
├── kernel/
├── manifests/
├── rootfs/
├── scripts/
├── security/
│   ├── ids/
│   ├── ips/
│   ├── malware/
│   ├── sandbox/
│   ├── forensics/
│   ├── honeypots/
│   ├── threatintel/
│   └── zero_trust/
├── ai/
│   ├── models/
│   ├── anomaly_detection/
│   ├── malware_classification/
│   ├── threat_prediction/
│   ├── llm_assistant/
│   ├── log_analysis/
│   └── behavior_engine/
├── src/
│   ├── boot/
│   └── kernel/
├── tools/
└── userspace/
```

## Document Index

| File | Purpose |
|---|---|
| [docs/guides/01-os-building-guide-overview.md](docs/guides/01-os-building-guide-overview.md) | Comprehensive overview of scratch vs Linux-based OS building |
| [docs/guides/02-from-scratch-os-development.md](docs/guides/02-from-scratch-os-development.md) | From-scratch OS development quick-start |
| [docs/guides/03-linux-based-os-customization.md](docs/guides/03-linux-based-os-customization.md) | Arch, Kali, Debian, kernel, ISO, and deployment guide |
| [docs/guides/04-os-technical-specifications-template.md](docs/guides/04-os-technical-specifications-template.md) | Professional OS technical specification template |
| [docs/00-overview.md](docs/00-overview.md) | Project type, targets, development models |
| [docs/specs/SAD.md](docs/specs/SAD.md) | System Architecture Specification |
| [docs/specs/toolchain-guide.md](docs/specs/toolchain-guide.md) | Build environment and cross-toolchain guide |
| [docs/specs/boot-init-sequence.md](docs/specs/boot-init-sequence.md) | Boot and initialization sequence |
| [docs/specs/package-payload-ecosystem.md](docs/specs/package-payload-ecosystem.md) | Linux-path package and payload policy |
| [docs/specs/bridge-lfs.md](docs/specs/bridge-lfs.md) | LFS bridge approach |
| [docs/specs/ai-ml-security-layer.md](docs/specs/ai-ml-security-layer.md) | AI/ML cybersecurity support layer |
| [docs/ai-cyber/01-ai-cyber-overview.md](docs/ai-cyber/01-ai-cyber-overview.md) | AI+Cyber project overview and study plan |
| [docs/ai-cyber/02-threat-model.md](docs/ai-cyber/02-threat-model.md) | AI security threat model |
| [docs/ai-cyber/03-architecture.md](docs/ai-cyber/03-architecture.md) | AI security data pipeline architecture |
| [docs/ai-cyber/04-evaluation.md](docs/ai-cyber/04-evaluation.md) | Metrics, datasets, and adversarial tests |
| [docs/cybersecurity/01-cybersecurity-with-aiml-overview.md](docs/cybersecurity/01-cybersecurity-with-aiml-overview.md) | Open-source AI/ML cybersecurity overview |
| [docs/cybersecurity/02-open-source-tools-and-stack.md](docs/cybersecurity/02-open-source-tools-and-stack.md) | Security tool stack and deployment tiers |
| [docs/cybersecurity/03-aiml-threat-detection-analysis.md](docs/cybersecurity/03-aiml-threat-detection-analysis.md) | ML threat detection algorithms and patterns |
| [docs/cybersecurity/04-practical-implementation-guide.md](docs/cybersecurity/04-practical-implementation-guide.md) | Step-by-step security infrastructure setup |
| [docs/cybersecurity/05-github-repos-and-resources.md](docs/cybersecurity/05-github-repos-and-resources.md) | Open-source repositories and resources |
| [docs/cybersecurity/06-advanced-topics-best-practices.md](docs/cybersecurity/06-advanced-topics-best-practices.md) | Model drift, XAI, adversarial ML, scale, governance |
| [docs/research/01-cyber-ai-os-research-map.md](docs/research/01-cyber-ai-os-research-map.md) | Complete research map for Cyber AI OS development |
| [docs/research/02-research-modules.md](docs/research/02-research-modules.md) | Arrow-by-arrow research modules and concrete artifacts |
| [docs/research/03-os-engineering-curriculum.md](docs/research/03-os-engineering-curriculum.md) | OS-engineering curriculum mapped to kernel projects |
| [docs/research/04-low-level-security-labs.md](docs/research/04-low-level-security-labs.md) | Low-level security lab map for mitigations and exploit analysis |
| [docs/research/05-networking-packet-analysis-labs.md](docs/research/05-networking-packet-analysis-labs.md) | Networking and packet-analysis lab map |
| [docs/research/06-threat-detection-stack-labs.md](docs/research/06-threat-detection-stack-labs.md) | IDS, EDR, SIEM, SOAR, hunting, and attack-graph labs |
| [docs/research/07-malware-research-labs.md](docs/research/07-malware-research-labs.md) | Malware triage, reverse engineering, sandboxing, and YARA labs |
| [docs/research/08-dfir-labs.md](docs/research/08-dfir-labs.md) | Digital forensics acquisition, timeline, memory, and IR labs |
| [docs/research/09-ai-ml-research-topics.md](docs/research/09-ai-ml-research-topics.md) | AI/ML research topics, frameworks, and project ideas |
| [docs/research/10-security-ai-research-topics.md](docs/research/10-security-ai-research-topics.md) | Security-AI, SOC copilot, RAG, agents, and adversarial ML topics |
| [docs/research/11-kernel-telemetry-layer.md](docs/research/11-kernel-telemetry-layer.md) | eBPF, probes, syscall streams, process graphs, Falco/Tracee, and telemetry budgets |
| [docs/research/12-containers-cloud-security.md](docs/research/12-containers-cloud-security.md) | Containers, Kubernetes, SBOM, runtime security, and cloud policy research |
| [docs/research/13-observability-stack.md](docs/research/13-observability-stack.md) | Metrics, logs, traces, OTel, Prometheus, Grafana, Loki, SLOs, and dashboards |
| [docs/research/14-containers-cloud-ms-research-roadmap.md](docs/research/14-containers-cloud-ms-research-roadmap.md) | MSc-style containers and cloud security research roadmap |
| [docs/prompts/research-alignment-master-prompt.md](docs/prompts/research-alignment-master-prompt.md) | Reusable master prompt for research alignment and execution planning |
| [docs/prompts/staged-restructuring-master-prompt.md](docs/prompts/staged-restructuring-master-prompt.md) | Reusable prompt for staging the full project into execution folders |
| [docs/01-os-from-scratch.md](docs/01-os-from-scratch.md) | Minimal bare-metal OS outline |
| [docs/01-lfs-build.md](docs/01-lfs-build.md) | LFS-style custom Linux build plan |
| [docs/01-goals-and-scope.md](docs/01-goals-and-scope.md) | Goals, non-goals, acceptance criteria |
| [docs/02-architecture.md](docs/02-architecture.md) | System architecture, invariants, base paths |
| [docs/02-quickstart-arch.md](docs/02-quickstart-arch.md) | Arch host and custom ISO quickstart |
| [docs/02-quickstart-kali.md](docs/02-quickstart-kali.md) | Kali host and security image quickstart |
| [docs/03-build-system.md](docs/03-build-system.md) | Toolchain, build targets, outputs |
| [docs/04-boot-process.md](docs/04-boot-process.md) | Firmware, bootloader, kernel handoff |
| [docs/05-kernel-design.md](docs/05-kernel-design.md) | Kernel subsystems and failure handling |
| [docs/06-userspace.md](docs/06-userspace.md) | Init, shell, utilities, execution model |
| [docs/07-filesystem.md](docs/07-filesystem.md) | Filesystem phases and constraints |
| [docs/08-drivers.md](docs/08-drivers.md) | Driver order, model, hardware scope |
| [docs/09-security.md](docs/09-security.md) | Access control, AI/ML support layer, unsafe defaults |
| [docs/10-networking.md](docs/10-networking.md) | Network stack and Linux-derived network checks |
| [docs/11-packaging.md](docs/11-packaging.md) | Arch, Kali/Debian, and scratch packaging |
| [docs/12-testing.md](docs/12-testing.md) | Build, boot, runtime, and release tests |
| [docs/13-roadmap.md](docs/13-roadmap.md) | Phased delivery plan |
| [docs/14-cybersecurity-vision.md](docs/14-cybersecurity-vision.md) | SentinelOS cybersecurity vision |
| [docs/15-ai-ml-architecture.md](docs/15-ai-ml-architecture.md) | AI/ML defense architecture |
| [docs/16-malware-detection.md](docs/16-malware-detection.md) | Static, dynamic, and AI malware detection |
| [docs/17-threat-intelligence.md](docs/17-threat-intelligence.md) | IOC, CVE, ATT&CK, YARA, Sigma design |
| [docs/18-ai-assistant.md](docs/18-ai-assistant.md) | Offline AI security assistant |
| [docs/19-forensics.md](docs/19-forensics.md) | Digital forensics modules |
| [docs/20-security-modules.md](docs/20-security-modules.md) | IDS, IPS, EDR, XDR, SIEM, SOAR modules |
| [docs/21-red-team-tools.md](docs/21-red-team-tools.md) | Controlled red-team tool environment |
| [docs/decisions/ADR-001-base-system.md](docs/decisions/ADR-001-base-system.md) | Base-system architecture decision |

## Recommended First Target

Use Arch-based first unless the goal is explicitly kernel development.

Rationale:

- Fastest route to a bootable custom OS image.
- Full Linux userspace and package ecosystem.
- Cleaner base than Kali for non-security distributions.
- Easier to reproduce than ad hoc from-scratch systems.

Use Kali-based only when the system purpose is security testing and tooling. Use from-scratch only when Linux is not an acceptable kernel/runtime base.

## Practical Development Order

```text
Arch/Linux base
  -> security tools integration
  -> custom security modules
  -> sandbox / monitoring / eBPF
  -> AI anomaly detection and malware classifier
  -> custom distro
  -> kernel modifications
  -> partial custom OS
  -> full security OS
```

Starting directly from a custom kernel for this cybersecurity scope is inefficient. Begin with Arch/LFS, build the cyber stack, then progressively replace components.
