# Wiring Hermes-Agent + Obsidian + PhoenixOS

This guide details how to integrate your local **Hermes Agent** (cloned under `parts/hermes-agent/`) with your Obsidian knowledge base (**phoenix_brain/**) and the core **PhoenixOS** runtime.

---

## 1. Architectural Map

```
                     ┌──────────────────┐
                     │  phoenix_brain/  │
                     │  (Obsidian Vault)│
                     └────────┬─────────┘
                              │ (Reads Specs / Context)
                              ▼
  ┌──────────────┐   ┌──────────────────┐   ┌──────────────────┐
  │  PhoenixOS   ├──>│   Hermes Agent   ├──>│ Local LLM Server │
  │  (Replay/Go) │   │  (Python CLI)    │   │ (Jan AI / Ollama)│
  └──────────────┘   └──────────────────┘   └──────────────────┘
    (State Logs &        (Synthesizes
    Ledger Hashes)        Policies)
```

---

## 2. Setting Up the Agent & Inference Engine

### A. Run the Local Inference Engine (Jan AI)
1. Install and run **Jan AI** (source code cloned under [parts/jan](file:///Users/fallofpheonix/os/parts/jan)). Alternatively, download the client application from [jan.ai](https://jan.ai).
2. Start the local server inside Jan (defaults to standard OpenAI-compatible API endpoint: `http://localhost:1337/v1`).
3. Load a model (e.g., `Llama-3-8B-Instruct` or `Hermes-2-Pro-Llama-3-8B`).

### B. Configure Hermes Agent
1. Navigate to the Hermes Agent directory:
   ```bash
   cd parts/hermes-agent/
   ```
2. Copy the configuration template:
   ```bash
   cp cli-config.yaml.example config.yaml
   ```
3. Edit `config.yaml` to point to your local Jan inference endpoint:
   ```yaml
   model:
     provider: openai
     base_url: http://localhost:1337/v1
     name: hermes-2-pro-llama-3-8b
     api_key: local-key
   ```
4. Install python dependencies (recommended using `uv`):
   ```bash
   uv pip install -e .
   ```

---

## 3. Integrating Obsidian (phoenix_brain) as Long-Term Memory

Hermes Agent reads files directly from the local workspace. By mounting `phoenix_brain/` as a reference context folder, you can prompt Hermes with full awareness of your system design and previous decisions.

### A. Context Mounting
Configure your agent startup command or routine file to index the Markdown notes:
```bash
python run_agent.py --knowledge-base ../../phoenix_brain/
```

### B. Read and Write Workflow
- **Spec-to-Code**: Ask Hermes to read design specifications:
  > *"Hermes, read [[RFC-001 AI Orchestrator]] and write a new Go test in `phoenix_os/ai/orchestrator_test.go` that verifies feature execution ordering."*
- **Trace-to-Refinement**: Pipe output logs into a temporary note inside `phoenix_brain/01_CAPTURE/` and ask Hermes to analyze performance:
  > *"Analyze the latest replay execution log and update [[PhoenixOS Stage 1]] with the recommended Payoff function parameter tuning."*

---

## 4. Closing the Loop with PhoenixOS Run Loop

In your AI-first architecture, the Go orchestrator handles runtime deterministic ticks, while the offline Hermes Agent analyzes the ledger trace to modify policy weights dynamically.

1. **Dump Trace**: Execute PhoenixOS and record output hashes/logs.
2. **Consult Hermes**: Have Hermes scan the newly generated `phoenix_trace.db` alongside your system specifications in `phoenix_brain/`.
3. **Apply Feedback**: Hermes generates updated policy configuration constraints for `phoenix_os/arbiter/policy.go` and submits the patch automatically.
