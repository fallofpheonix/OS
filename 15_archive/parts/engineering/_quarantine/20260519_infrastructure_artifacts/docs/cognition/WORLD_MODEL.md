# World Model

Purpose
-------
Defines the system's internal representation of reality used for prediction, simulation, and planning.

Structure
---------
- World State: repository_state, active_goals, unresolved_failures, architecture_health
- Repository State: symbol graph, file states, branch topology
- Task State: active tasks, timelines, expectations
- Environment State: external inputs, model capabilities, runtime resources
- Belief Graph: nodes, confidence scores, provenance
- Future Simulation: trajectories, outcome estimates
- State Transitions: allowable transitions and costs
- Reality Reconstruction: grounding observations into internal state

Next steps
----------
- Add schema examples and JSON/YAML artifact definitions.
