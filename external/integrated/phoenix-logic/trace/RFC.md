# RFC: Phoenix Trace

## 1. Description
Phoenix Trace provides L4 telemetry. It receives events from the Phoenix Bus (fork, exec, exit) and updates the global causality graph.

## 2. Specification
- **Nodes:** Process, File, Network.
- **Edges:** forked, executed, wrote, connected.

## 3. Interface
```go
func (g *Graph) AddEvent(e Event)
func (g *Graph) TraceAncestors(pid string) []string
```
