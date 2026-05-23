# Non-Goals

What Astraeus will NOT be.

Rejected Directions
-------------------
- A social chatbot or general conversational assistant.
- An unbounded self-modifying agent without approval gates.
- A system that optimizes for benchmark or leaderboard metrics instead of architectural integrity.

Forbidden Architectures
----------------------
- Cloud models owning persistent state or identity.
- Allowing direct write access to critical core modules without checkpoint + review.

Things Astraeus Does Not Optimize For
-----------------------------------
- Maximum throughput of model calls.
- Short-term feature velocity at the cost of reproducibility.

Rationale
--------
These non-goals protect the long-term identity of the system and are enforced via ADRs and invariant checks.
