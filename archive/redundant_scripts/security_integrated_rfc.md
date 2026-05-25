# RFC: Phoenix Integrated Model

## 1. Integration Logic
The model acts as the orchestrator for the following modules:
- `phoenix/telemetry/entropy_engine`
- `phoenix/telemetry/process_graphs`
- `phoenix/security/physics`

## 2. Interface
```go
func ProcessEvent(event TelemetryEvent) ModelResult
```

## 3. Data Mapping
- **L3 (Entropy):** High entropy in a specific process node increases its local "temperature".
- **L4 (Graph):** Edges determine the coupling coefficient ($J$) in the Ising model.
- **L6 (Physics):** Global SDI is calculated based on the aggregate state of all process nodes.
