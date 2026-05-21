# Theory to OS Subsystem Map

| Theory | Formula / Core Concept | OS Subsystem | Runtime Target | Experiment | Status |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **Information Theory** | Shannon Entropy, KL Divergence | Telemetry / File Analysis | `09_telemetry/entropy_engine` | R002, R021 | ACTIVE |
| **Graph Theory** | Directed Acyclic Graphs (DAG), Centrality | Intelligence / Lineage | `09_telemetry/process_graphs` | R003, R022 | ACTIVE |
| **Signal Processing** | Kalman Filters, Wavelet Transforms | Telemetry Noise Reduction | `09_telemetry/math_filters` | R026 | IMPLEMENT_LATER |
| **Statistical Physics** | Ising Lattices, Arrhenius Eq, SDI | Phoenix Sentinel | `07_security/physics` | R024 | ACTIVE |
| **Game Theory** | Strong Stackelberg Eq, VCG Auctions | Decision Engine / Scheduler | `07_security/game` | R027, R028, R030 | IMPLEMENT_LATER |
| **Control Systems** | PID, Linear-Quadratic (LQ) Games | Actuation / Containment | `07_security/control` | R023, R031 | ACTIVE |
| **Dynamical Systems** | State-Space Trajectory Drift | Anomaly Detection | `06_ai/dynamics` | R025 | RESEARCH_ONLY |
| **Complex Systems** | Phase Transitions | Cascading Failure Prediction | `07_security/physics` | R024 | RESEARCH_ONLY |
| **Optimization** | Convex Optimization, Minimax | Cost Minimization | `07_security/game` | R023 | ACTIVE |
| **Evolutionary Systems** | ESS (Evolutionarily Stable Strategies) | N/A | N/A | N/A | REMOVE |

*Note: Unconnected theories like general Evolutionary Systems without specific runtime bindings have been marked REMOVE.*