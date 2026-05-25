# MEMORY POISONING ANALYSIS

## Current Memory Architecture
Astraeus relies on:
1. `events.jsonl` (Event Bus)
2. `journal.jsonl` (Filesystem Journal)
3. `run.json` (State Memory)

## Poisoning Vectors

### 1. Unsigned Appends
- **Mechanism:** Files are plain text JSON lines. There is no cryptographic signature, HMAC, or external verification.
- **Impact:** Any component (or attacker) that achieves file write access can inject arbitrary memories, tricking the Replay Engine and the Rollback Engine.

### 2. Predictable UUIDs
- **Mechanism:** `uuid4().hex[:12]` is used for entry IDs.
- **Impact:** While UUID4 is pseudorandom, truncation to 12 characters increases collision risk. If an attacker can guess IDs, they can potentially forge artifacts.

### 3. Context Window Poisoning (Prompt Injection)
- **Mechanism:** The repository cognitive map reads source files and passes them into the LLM context window.
- **Impact:** An attacker who commits a file with adversarial instructions (`IGNORE ALL PREVIOUS PROMPTS`) will poison the cognitive map of the system. The planner will then act on the attacker's instructions instead of the user's intent.

## Conclusion
The memory subsystem is structurally vulnerable to both local filesystem poisoning and indirect prompt-injection poisoning.
