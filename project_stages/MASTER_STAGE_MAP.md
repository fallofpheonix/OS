# Master Stage Map

| Stage | Name | Class | Progression | Topic count | Doc count | Prerequisites | Scope |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Stage_00_Foundations | Foundations | FOUNDATIONAL | Learn | 12 | 9 | None | Programming, architecture, tools, research scaffolding, project scope |
| Stage_01_System_Internals | System Internals | FOUNDATIONAL | Learn -> Prototype | 12 | 6 | Stage_00_Foundations | CPU, memory, process, scheduler, boot, filesystem, userspace internals |
| Stage_02_Linux_and_Distro | Linux and Distro | IMPLEMENTATION | Prototype -> Integrate | 8 | 3 | Stage_00_Foundations, Stage_01_System_Internals | LFS, Arch/Kali base, package payloads, build reproducibility |
| Stage_03_Networking | Networking | FOUNDATIONAL | Learn -> Prototype | 10 | 2 | Stage_00_Foundations, Stage_01_System_Internals | TCP/IP, packet analysis, sockets, IDS-ready network telemetry |
| Stage_04_Security_Fundamentals | Security Fundamentals | FOUNDATIONAL | Learn -> Secure | 10 | 4 | Stage_00_Foundations, Stage_01_System_Internals | Low-level security, crypto, sandboxing, access control, hardening |
| Stage_05_Malware_and_RE | Malware and RE | IMPLEMENTATION | Prototype -> Secure | 8 | 2 | Stage_04_Security_Fundamentals, Stage_03_Networking | Static/dynamic malware analysis, YARA, sandboxing, reverse engineering |
| Stage_06_Forensics | Forensics | IMPLEMENTATION | Prototype -> Integrate | 8 | 2 | Stage_03_Networking, Stage_04_Security_Fundamentals, Stage_05_Malware_and_RE | DFIR acquisition, timelines, memory/disk evidence, case workflow |
| Stage_07_Threat_Intelligence | Threat Intelligence | IMPLEMENTATION | Integrate -> Secure | 7 | 1 | Stage_03_Networking, Stage_04_Security_Fundamentals, Stage_05_Malware_and_RE | IOC/rule/feed management, MITRE mapping, enrichment |
| Stage_08_Observability | Observability | IMPLEMENTATION | Observe | 8 | 1 | Stage_03_Networking, Stage_04_Security_Fundamentals | Logs, metrics, traces, OpenTelemetry, storage and dashboards |
| Stage_09_eBPF_and_Telemetry | eBPF and Telemetry | IMPLEMENTATION | Observe -> Extend | 8 | 1 | Stage_01_System_Internals, Stage_02_Linux_and_Distro, Stage_04_Security_Fundamentals, Stage_08_Observability | eBPF, probes, audit, runtime event schema, process graph |
| Stage_10_SOC_Stack | SOC Stack | IMPLEMENTATION | Integrate -> Observe | 9 | 4 | Stage_03_Networking, Stage_06_Forensics, Stage_07_Threat_Intelligence, Stage_08_Observability | SIEM, EDR/XDR, alert lifecycle, case management |
| Stage_11_AI_ML_Core | AI ML Core | FOUNDATIONAL | Learn -> Prototype | 8 | 2 | Stage_00_Foundations | ML fundamentals, feature engineering, datasets, evaluation |
| Stage_12_Security_AI | Security AI | IMPLEMENTATION | Prototype -> Integrate | 8 | 12 | Stage_10_SOC_Stack, Stage_11_AI_ML_Core, Stage_05_Malware_and_RE, Stage_06_Forensics | AI IDS, malware ML, SOC assistant, security RAG, adversarial ML |
| Stage_13_Containers_and_Cloud | Containers and Cloud | IMPLEMENTATION | Integrate -> Scale | 8 | 2 | Stage_02_Linux_and_Distro, Stage_03_Networking, Stage_04_Security_Fundamentals, Stage_08_Observability | Containers, Kubernetes, cloud security, service mesh, supply chain |
| Stage_14_Automation_and_SOAR | Automation and SOAR | IMPLEMENTATION | Automate | 6 | 0 | Stage_10_SOC_Stack, Stage_12_Security_AI | Playbooks, SOAR, response automation, workflow orchestration |
| Stage_15_Security_Distribution | Security Distribution | IMPLEMENTATION | Scale | 7 | 3 | Stage_02_Linux_and_Distro, Stage_10_SOC_Stack, Stage_13_Containers_and_Cloud, Stage_14_Automation_and_SOAR | Security distro packaging, payload policy, tools, release gates |
| Stage_16_Kernel_Extensions | Kernel Extensions | IMPLEMENTATION | OS Extension | 7 | 2 | Stage_09_eBPF_and_Telemetry, Stage_15_Security_Distribution | Kernel modules, LSM, drivers, custom telemetry hooks |
| Stage_17_Hybrid_OS | Hybrid OS | IMPLEMENTATION | OS Extension | 6 | 2 | Stage_15_Security_Distribution, Stage_16_Kernel_Extensions | LFS/Buildroot/custom kernel bridge, distro plus OS-owned layers |
| Stage_18_Custom_OS | Custom OS | IMPLEMENTATION | Custom Kernel | 9 | 2 | Stage_01_System_Internals, Stage_16_Kernel_Extensions, Stage_17_Hybrid_OS | Scratch kernel, boot, memory, scheduler, drivers, filesystem |
| Stage_19_Production_Platform | Production Platform | IMPLEMENTATION | Scale | 8 | 5 | Stage_12_Security_AI, Stage_14_Automation_and_SOAR, Stage_15_Security_Distribution, Stage_17_Hybrid_OS | Production Cybersecurity Platform, reliability, governance, release operations |
| Stage_20_Research_and_Future | Research and Future | RESEARCH_ONLY | Cyber AI OS | 7 | 5 | Stage_19_Production_Platform | Future research, prompts, thesis tracks, long-horizon Cyber AI OS |
