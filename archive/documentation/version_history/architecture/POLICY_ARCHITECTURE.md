# PhoenixOS: Policy Architecture

`Policy` interfaces define evaluation rules.
`PolicyContext` supplies necessary inputs.
Arbiter evaluates policies and returns `PolicyResult` (Score, Confidence, Class).
Policies are immutable during execution.
