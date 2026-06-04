---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Validation — Control Flow

> Last verified: 2026-06-04

Defines the execution flow of the validation sandbox and tests runner.

## Exploit Simulation Control Flow

```mermaid
stateDiagram-v2
    [*] --> SetupSandbox: Create isolated namespace
    SetupSandbox --> LaunchTarget: Spawn target process
    LaunchTarget --> RunExploit: Inject malicious payload
    RunExploit --> AuditWarden: Read violation logs
    AuditWarden --> VerifyContainment: Assert process was killed within target timeout
    VerifyContainment --> TeardownSandbox: Clean up resources
    TeardownSandbox --> [*]
```
