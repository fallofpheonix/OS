# Continue Dev Integration

## Goal
Use Continue.dev in VS Code as the front end for your local engineering AI stack, while keeping model behavior separated by task.

This note answers three questions:
1. Where do I give instructions?
2. How do models change for specific tasks?
3. How does this connect to the local AI stack you built?

## Recommended Architecture

Continue should be the interactive client.
Your local control-plane layer should decide which model handles which task.

Suggested split:
- Continue.dev: editor UI, chat, code actions, quick prompts
- astraeus-core: task routing, model selection, memory, planning
- repo-analyzer: code search, semantic context, repository structure
- Ollama: local model serving

## Where To Give Instructions

### 1. Global behavior
Put persistent behavior in your Continue config file:
- `~/.continue/config.yaml` for a user-wide setup
- project-level `.continue/config.yaml` if you want repo-specific behavior

Use this for:
- default model list
- instruction style
- model aliases
- local Ollama connection details
- task-specific model preferences

### 2. One-off task instructions
Put the immediate instruction in the chat box inside VS Code.
Use this for:
- a single refactor request
- a single debugging question
- a single documentation task
- a one-time repository analysis

### 3. Reusable task instructions
If you want repeatable behavior, keep task prompts in one of these places:
- a workspace note in `brain/`
- a prompt template in your project docs
- a custom Continue prompt or slash command if your version supports it

Use this for:
- "analyze this repo"
- "debug this stack trace"
- "extract reusable modules"
- "write ADR for this change"

## How Models Should Change For Specific Tasks

Continue itself usually does not magically infer and switch models by task.
You have two practical options:

### Option A: Manual model selection in Continue
Best when you want direct control.

Create separate model entries and choose the right one from the model picker.

Recommended mapping:
- Qwen2.5-coder:7b -> code generation, refactoring, architecture
- DeepSeek-coder:6.7b -> debugging, stack traces, concurrency, reasoning
- Mistral:7b -> documentation, ADRs, summaries
- CodeLlama:7b -> alternate implementations, synthesis, comparison
- Phi3:mini -> planning, decomposition, lightweight control-plane

Use this when:
- you want to stay inside Continue only
- you are working on one task at a time
- you want explicit model control

### Option B: Automatic routing through astraeus-core
Best when you want task-based model switching.

Flow:
1. Continue receives the user request
2. astraeus-core classifies the task
3. router assigns the best model
4. the model executes the task
5. results go back to Continue

This is the right pattern for your stack because it preserves separation:
- Continue = interface
- astraeus-core = brain
- Ollama = model runtime
- repo-analyzer = repo context

## Task To Model Routing

| Task Type | Model | Why |
|---|---|---|
| Code generation | Qwen2.5-coder | strongest general coding role |
| Debugging | DeepSeek-coder | stack traces, failure analysis, reasoning |
| Documentation | Mistral | concise explanations and ADRs |
| Synthesis | CodeLlama | alternatives and comparisons |
| Planning | Phi3:mini | lightweight decomposition and routing |
| Semantic context | nomic-embed-text | embeddings and retrieval |

## Practical Integration Pattern

### 1. Keep Continue simple
Use Continue for:
- chat
- file edits
- code completion
- selecting a model manually when needed

### 2. Keep instructions centralized
Store your task policy in one place:
- `brain/08_REFERENCE/configurations/Continue Dev Integration.md`
- `brain/10_META/` if you want operational policy
- repo docs if the behavior is project-specific

### 3. Let astraeus-core route by task
Use astraeus-core rules for:
- automatic model selection
- memory lookups
- task decomposition
- repo-aware reasoning

### 4. Use repo-analyzer for context
Before asking a model to change code, supply the relevant repo context:
- file structure
- search results
- module relationships
- semantic matches

## Example Workflow

### Example: debugging task
User asks:
"Why is this service failing under load?"

Flow:
1. Continue captures the request
2. astraeus-core classifies it as debugging
3. router sends it to DeepSeek-coder
4. failure memory is checked for similar incidents
5. repo-analyzer provides relevant code context
6. model returns diagnosis and fix plan

### Example: documentation task
User asks:
"Write an ADR for this architecture decision"

Flow:
1. Continue captures the request
2. astraeus-core classifies it as documentation
3. router sends it to Mistral
4. architecture memory is consulted
5. result is written as ADR-style documentation

### Example: code generation task
User asks:
"Extract this duplicated logic into a module"

Flow:
1. Continue captures the request
2. astraeus-core classifies it as code generation or extraction
3. router sends it to Qwen2.5-coder
4. repo-analyzer finds duplicate patterns
5. model proposes refactor and implementation steps

## Recommended Continue Config Shape

Keep your Continue config small and explicit.

Example structure:
```yaml
name: Local Config
version: 1.0.0
schema: v1
models:
  - name: Autodetect
    provider: ollama
    model: AUTODETECT
  - name: Qwen Coder
    provider: ollama
    model: qwen2.5-coder:7b
  - name: DeepSeek Debugger
    provider: ollama
    model: deepseek-coder:6.7b
  - name: Mistral Docs
    provider: ollama
    model: mistral:7b
```

Use the smaller models for the task they are best at, not all at once.

## Rules Of Thumb

- Do not use one model for every task
- Do not run two large 7B models at the same time on the M3 8GB machine
- Use Phi3 for planning, not heavy code generation
- Use Qwen for code edits and architecture
- Use DeepSeek for failures and runtime bugs
- Use Mistral for summaries and documentation
- Keep semantic memory separate from session memory
- Let routing live in the orchestrator, not in the chat client

## Short Answer

Where to give instructions:
- persistent rules: `~/.continue/config.yaml` or project `.continue/config.yaml`
- one-off requests: Continue chat box
- reusable task prompts: vault note or project doc

How models change by task:
- Continue can let you choose the model manually
- astraeus-core can switch models automatically based on task type
- model role changes should happen in the router, not by ad hoc prompt text

## Next Step

If you want, the next document should be:
- a concrete `Continue` config example for your local Ollama models
- a `astraeus-core` routing table document
- a combined workflow doc showing Continue -> astraeus-core -> Ollama -> repo-analyzer
