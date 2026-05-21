# RFC: Sentinel Integrated Model

## 1. Integration Logic
The model acts as the orchestrator for the following modules:
- `sentinel/telemetry/entropy_engine`
- `sentinel/telemetry/process_graphs`
- `sentinel/security/physics`

## 2. Interface
```go
func ProcessEvent(event TelemetryEvent) ModelResult
```

## 3. Data Mapping
- **L3 (Entropy):** High entropy in a specific process node increases its local "temperature".
- **L4 (Graph):** Edges determine the coupling coefficient ($J$) in the Ising model.
- **L6 (Physics):** Global SDI is calculated based on the aggregate state of all process nodes.
