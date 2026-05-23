# Philosophy

Intellectual foundations and reasoning behind Astraeus' architecture.

Problems With Current AI
------------------------
- Stateless prompt wrappers produce brittle, untraceable edits.
- Heavy reliance on model outputs without grounding leads to hallucination.

Cognition vs Automation
-----------------------
Automation executes tasks; cognition models, simulates, and learns. Astraeus aims to provide the latter for engineering work.

Why Stateful Systems Matter
--------------------------
- Persistence enables consolidation, error attribution, and causal learning.
- Stateful beliefs let the system form and test hypotheses about repository structure.

Why Memory Matters
------------------
- Memory consolidation turns repeated experiences into reusable strategies and reduces repeated failure modes.

Why Temporal Cognition Matters
-----------------------------
- Long-horizon planning prevents destructive short-term fixes that propagate damage.

Why Invariants Matter
---------------------
- Invariants are the system's laws; encoding them into checks prevents identity loss during evolution.
