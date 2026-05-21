# PhoenixOS Layer 1: External Agents

Layer 1 agents interact with the outside world. They are the "sensory and creative" organs of the OS. They NEVER touch the kernel directly; instead, they populate the research, intelligence, and experimental layers that the internal OS services consume.

## Agents

### 1. Research Agent (`agents/external/research/`)
- **Purpose:** Discovery and theorizing.
- **Owns:** Math models, physics notes, game theory papers.
- **Outputs:** `01_research/`, `02_docs/`.

### 2. Threat Intel Agent (`agents/external/intel/`)
- **Purpose:** External threat awareness.
- **Owns:** MITRE TTPs, malware feeds, IOCs.
- **Outputs:** `07_security/feeds/`.

### 3. Benchmark Agent (`agents/external/benchmark/`)
- **Purpose:** Empirical validation.
- **Owns:** Performance metrics, replay datasets, regression scoring.
- **Outputs:** `14_experiments/`.

### 4. Build Agent (`agents/external/build/`)
- **Purpose:** System synthesis.
- **Owns:** CI/CD, QEMU, VM artifacts.
- **Outputs:** `build/`, `reports/`.
