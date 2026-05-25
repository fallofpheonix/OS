# Architectural Decision Log

This log lists critical architectural choices made during the PhoenixOS project.

## Decision 1: Concurrency Mutex vs. Channel-Only Synchronization
- **Status:** APPROVED.
- **Alternatives Considered:** Processing all SOC API requests sequentially via a single coordination channel.
- **Tradeoff Analysis:** While channel loops are clean, Warden FSM triggers need immediate microsecond-latency evaluation to protect platform integrity. Adding explicit Mutex locks on Warden and RWMutex locks on Ledger minimizes latency and guarantees map thread safety.

## Decision 2: TCS Sliding Window Dynamic Range Scanning
- **Status:** APPROVED.
- **Alternatives Considered:** Relying on slice boundaries `events[len-1] - events[0]` for sequence delta calculations.
- **Tradeoff Analysis:** Slice endpoints assume chronological events match sequence progression. Because network telemetry can arrive out of order, or system overflow events can inject negative SeqIDs, dynamically scanning active window events prevents underflows and caps loss rate estimates accurately.

## Decision 3: Advisory-Only LLM Loop
- **Status:** APPROVED.
- **Alternatives Considered:** Allowing the AI Orchestrator to directly trigger Warden state escalations if confidence > 90%.
- **Tradeoff Analysis:** Direct AI-actuated FSM control violates **Axiom 3** and introduces non-deterministic model failures. Keeping LLM outputs strictly advisory ensures audit logs are informative without compromising system predictability.
