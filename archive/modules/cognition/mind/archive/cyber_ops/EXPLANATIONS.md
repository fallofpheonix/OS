# Cyber Operations Explanations

This directory contains the "Red Team" (attack) and "Blue Team" (defense) agents that validate the PhoenixOS security posture.

## blue_team.py

### Beginner
The Blue Team is like a digital security guard. It constantly watches the system for "drift" (unusual behavior). If it sees something suspicious, it tries to figure out what happened (forensics) and then asks the system's "Warden" to stop the problem, like isolating a suspicious process.

### Intermediate
The `BlueTeamAgent` monitors real-time telemetry and calculates threat levels using a `ThreatProjector`. It implements a feedback loop where critical drift triggers causal forensic analysis. If a threat is confirmed, it issues an actuation request to the Warden and logs the event in a tamper-evident ledger for auditability.

### Expert
The Blue Team implements the defensive half of the "Cognitive Security Loop." It utilizes Stackelberg models in the `ThreatProjector` to anticipate adversarial maneuvers. The agent coordinates distributed defense by mapping high-entropy telemetry to causal paths and requesting hardware-level isolation (e.g., PID isolation) via the Warden's deterministic consent protocol.

---

## red_team.py

### Beginner
The Red Team is a "friendly attacker." Its job is to simulate real hacks and problems so we can see if the Blue Team is actually working. It pulls ideas from a "Failure Library" (a list of things that went wrong in the past) and tries to make them happen again in a safe way.

### Intermediate
The `RedTeamAgent` performs "Heat Injection" by replaying historical failure patterns. It can programmatically create and execute exploit payloads in the system's memory regions to trigger telemetry drift, allowing developers to verify that the monitoring and defense systems respond correctly to known threat vectors.

### Expert
The Red Team acts as the adversarial validation engine. It automates the "Exploitation" phase of the security lifecycle by replaying failure modes stored in the `06_FAILURE_LIBRARY`. It facilitates "Heat Injection" by dynamically generating executable payloads and setting up the environmental conditions necessary for failure reproduction, ensuring that the system's fractal architecture remains resilient against regression.

---

## orchestrator.py

### Beginner
The Orchestrator is the referee. It starts the Red Team's attack and then checks to see if the Blue Team caught it. It runs through different "scenarios" to make sure the whole system is safe.

### Intermediate
The Orchestrator manages the interaction between the `RedTeamAgent` and the `BlueTeamAgent`. It sets up specific security scenarios, such as "executable code in brain memory," and verifies that the defense pipeline (Detection -> Forensics -> Actuation) completes successfully.

### Expert
The `Master Orchestrator` provides the end-to-end validation harness for Cyber Operations. It manages the state machine of a security simulation, ensuring that adversarial heat injection is correctly followed by defensive monitoring and deterministic actuation, thereby closing the validation loop for the PhoenixOS containment architecture.

---

### Code Review
| File | Risk Score | Complexity Score |
| :--- | :--- | :--- |
| blue_team.py | 4/10 | 5/10 |
| red_team.py | 8/10 | 6/10 |
| orchestrator.py | 3/10 | 4/10 |

*Note: red_team.py has a high risk score because it executes arbitrary code for simulation purposes.*
