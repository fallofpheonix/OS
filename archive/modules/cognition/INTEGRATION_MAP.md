# PhoenixOS External Integration Map

This document defines the definitive mapping of external open-source repositories into the **7-Layer Phoenix Matrix**. It establishes the approved technological substrate for the operating system, orchestrating eBPF observability, AI/ML inference, and multi-agent coordination.

## 🔥 Critical Priority Path
Based on the deep engineering OS architecture, these 5 tools are the highest priority for immediate operational integration:

| Priority | Repo | Why | Layer |
| :--- | :--- | :--- | :--- |
| 🥇 1 | **ERA** (BinSquare/ERA) | microVM sandboxing for AI agents — replaces Sandbox Execution Engine, hardware-level security. | Layer 2 |
| 🥇 2 | **Pixie + Tetragon** | Auto-instrumented eBPF observability — replaces custom kernel probes, captures traces without code changes. | Layer 1 |
| 🥇 3 | **OpenHands** | Full AI coding agent with AST-aware editing — wire into Code Intelligence immediately. | Layer 3 |
| 🥇 4 | **LangGraph + AutoGen** | Build Orchestrator's DAG routing engine — stateful workflows, multi-agent orchestration. | Layer 6 |
| 🥇 5 | **vLLM + Unsloth** | Fast LoRA fine-tuning + high-throughput inference — start fine-tuning PhoenixMind on Failure Library. | Layer 4 |

---

## 🖥️ Layer 1 — OS Core (eBPF · Kernel · Telemetry)
*Focus: Kernel-level enforcement, tracing, and zero-instrumentation observability.*

| Repo | Capabilities |
| :--- | :--- |
| **Falco** | Cloud-native runtime security tool; monitors kernel events via eBPF. |
| **Tetragon** | Cilium's eBPF-based security observability; in-kernel enforcement + causal tracing. |
| **Tracee** | eBPF programs attached to syscall entry/exit and LSM hooks. |
| **KubeArmor** | Enforces access control and behavioral policies using eBPF and LSMs. |
| **AgentSight** | eBPF-powered system observability specifically for LLM agents. |
| **Pixie** | Auto-instrumented eBPF observability for Kubernetes. |
| **Coroot** | eBPF-based application performance monitoring (no-code instrumentation). |
| **Beyla** | eBPF-based service mesh generating OpenTelemetry traces. |
| **Odigos** | Distributed tracing with eBPF (supports 20+ languages). |
| **Parca** | eBPF-based continuous profiling (integrates with Tetragon). |
| **Anteon** | eBPF observability for cost optimization. |

## 🔐 Layer 2 — Cyber Ops (Red Team · Blue Team · Forensics)
*Focus: Threat modeling, agentic containment, and multi-agent security automation.*

| Repo | Capabilities |
| :--- | :--- |
| **Cybersecurity AI (CAI)** | Red teaming framework with autonomous pentesting. |
| **FalcoClaw** | Kernel syscall → Falco anomaly → automated action (kill, block, quarantine). |
| **YAWNING TITAN** | Graph-based autonomous cyber operations training and simulation. |
| **Microsoft Agent Governance**| Enforces OWASP agentic AI risks with deterministic, sub-ms policy enforcement. |
| **Aegis** | Open-source EDR for AI agents (monitors processes, network in real time). |
| **ERA** | microVM-based sandboxing for AI agents (hardware-level security). |
| **AgentScope** | Alibaba's production-ready agent framework (meta-agent architecture). |
| **SQLite-Vec** | Persistent memory for agent pipelines + vector extensions. |
| **AgentDB** | Single-file cognitive container (vectors, indexes, learning state). |
| **NATS** | Lightweight messaging for distributed agents (pub/sub + RPC). |
| **fasterpc** | RPC framework for local agents (bidirectional WebSockets). |

## 💻 Layer 3 — Code Intelligence (AST · Review · Debug · Sandbox)
*Focus: Automated code comprehension, patching, and secure execution.*

| Repo | Capabilities |
| :--- | :--- |
| **OpenHands** | Full AI coding agent (AST-aware code editing). |
| **Aider** | CLI pair programming with diff-based AST patching. |
| **PR-Agent** | Automated PR review, security scanning, complexity scoring. |
| **SWE-RL (SSR)** | Trains LLM agents via RL to iteratively inject and repair software bugs. |
| **OpenSandbox** | Secure, fast, extensible sandbox runtime for AI agents. |
| **Self-Healing SRE Agent**| Orchestrates multiple AI agents to detect and fix production issues. |
| **Armada** | Kubernetes-native distributed job scheduler. |
| **SkyPilot** | GPU-aware distributed training orchestrator. |
| **vLLM** | High-throughput LLM inference engine. |
| **Text-Generation-Inference**| Optimized LLM serving (Tensor Parallelism). |

## 🤖 Layer 4 — AI/ML Engineering (Training · Inference · LoRA)
*Focus: Model optimization, continuous batching, and local inference.*

| Repo | Capabilities |
| :--- | :--- |
| **OpenRLHF** | Production-ready RLHF framework built on Ray + vLLM. |
| **Unsloth** | Fast LoRA fine-tuning for Llama models. |
| **Ollama** | Local LLM serving (powers `phoenix-mind.Modelfile`). |
| **llama.cpp** | GGUF quantization (4-bit) + CPU/VRAM constrained inference. |
| **Mem0** | Persistent agent memory across sessions (vector DB management). |
| **LlamaIndex** | RAG pipelines + vector store (feeds context into NexusBridge). |
| **Ray** | Distributed computing framework (scales RLHF). |
| **DeepSpeed** | Microsoft's model optimization (ZeRO optimization). |
| **QLoRA** | 4-bit quantized LoRA fine-tuning (bitsandbytes). |

## 🎮 Layer 5 — Game/Simulation (RL · Physics · Engine Bridge)
*Focus: Environmental simulation and synthetic data generation.*

| Repo | Capabilities |
| :--- | :--- |
| **RLinf** | Production-grade open-source RL framework for embodied AI. |
| **FinRL** | RL for trading/financial simulation. |
| **Gymnasium** | Standard RL environment API (defines observation/action spaces). |
| **Stable-Baselines3** | PPO, DQN, SAC implementations ready to plug into the RL Engine. |
| **Godot** | Open-source game engine with WebSocket API (Engine Bridge target). |
| **Isaac Lab** | NVIDIA's embodied AI RL framework. |
| **MuJoCo** | Physics engine for robotics (high-precision simulation). |
| **Unity ML-Agents** | Unity game engine + ML training (synthetic data generation). |

## 🧬 Layer 6 — Self-Evolution (Self-Patcher · RLSF · Memory)
*Focus: Long-term memory, stateful workflows, and codebase navigation.*

| Repo | Capabilities |
| :--- | :--- |
| **SWE-bench** | Benchmark for self-patching agents (tests real GitHub issues). |
| **AutoCodeRover** | Autonomous codebase navigation + patch generation. |
| **Letta** | Long-term memory management across agent sessions. |
| **LangGraph** | Stateful graph-based agent workflows (loop, branch, pause). |
| **AutoGen** | Event-driven multi-agent framework (conversational multi-agent systems). |
| **Qdrant** | High-scale vector database (Rust-based). |
| **Weaviate** | Vector database + knowledge graph (hybrid search). |
| **Chroma** | Embedding store for AI apps. |

## 🌐 Layer 7 — Network & Distributed Systems
*Focus: Swarm coordination, stream processing, and mesh networking.*

| Repo | Capabilities |
| :--- | :--- |
| **gRPC** | High-performance RPC framework (pluggable load balancing). |
| **Envoy** | Cloud-native L7 proxy (service mesh, eBPF integration). |
| **Versionize** | Protocol translation layer (TCP/MQTT ↔ gRPC/REST). |
| **Flink** | Distributed stream processing (real-time data pipelines). |

---

## 📚 Master Directories
- `awesome-agent-orchestrators` (andyrewlee/awesome-agent-orchestrators)
- `re-list` (extremecoders-re/re-list)
- `awesome-time-tracking` (awesome-time-tracking)
- `awesome-ai-agents-2026` (ARUNAGIRINATHAN-K/awesome-ai-agents-2026)
- `awesome-cybersecurity-agentic-ai` (raphabot/awesome-cybersecurity-agentic-ai)
- `awesome-ai-security` (ottosulin/awesome-ai-security)
- `kyrolabs/awesome-agents` (kyrolabs/awesome-agents)
