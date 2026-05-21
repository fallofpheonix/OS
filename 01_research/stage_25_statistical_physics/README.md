# Stage 25: Statistical Physics in Pheonix

## 1. Core Objective
To model OS state transition behaviors using the laws of statistical mechanics and thermodynamics. By mapping processes to interacting particles, system configurations to thermodynamic states, and host disturbance to disorder, Pheonix defines a Security Disorder Index (SDI) and uses phase transition modeling (Ising model) to detect cooperative, system-wide compromises.

---

## 2. Mathematical Formulations & Subsystem Mapping

### 2.1. Thermodynamic Modeling of Host Security
We map OS components directly to thermodynamic constructs:
*   **Process** $\to$ Particle in a potential field.
*   **Thread State** $\to$ Microstate of a particle.
*   **Operating System Host** $\to$ Macro thermodynamic system.
*   **Global Threat Level** $\to$ System temperature ($\theta_T$).

#### System Security Energy ($E$):
The energy of a given system state $s \in \mathcal{S}$ is defined as a sum of security risks:
$$E(s) = \sum_{c \in \text{components}} \text{risk}(c)$$
Where secure configurations represent low-energy states, and compromised states represent high-energy configurations.

#### System Security Entropy ($S$):
The statistical entropy represents the diversity or unexpectedness of active microstates:
$$S = -k_B \sum_{s \in \mathcal{S}} P(s) \ln P(s)$$
Where $P(s) = \frac{1}{Z} e^{-\frac{E(s)}{\theta_T}}$ is the Boltzmann probability and $Z$ is the partition function.

#### Security Disorder Index (SDI):
The SDI defines the deviation of the operating system from its organized ground state:
$$\text{SDI} = S \cdot E$$
Under normal conditions, SDI is low. An intrusion elevates both entropy (unexpected processes) and energy (unauthorized privileges), causing a spike in SDI.

---

### 2.2. Ising Model for Cooperative Phase Transitions
The host is modeled as a lattice of interacting components (processes, file paths, sockets) with spin configurations $\sigma_i \in \{-1, +1\}$ representing secure ($+1$) or compromised/anomalous ($-1$) states.

The Hamiltonian is formulated as:
$$\mathcal{H}(\boldsymbol{\sigma}) = -J \sum_{\langle i, j \rangle} \sigma_i \sigma_j - H_{sec} \sum_{i} \sigma_i$$

Where:
*   $J$: Coupling coefficient defining dependencies between adjacent nodes (e.g. parent-child processes).
*   $H_{sec}$: The external stabilizing field representing active security enforcement boundaries (firewalls, LSM profiles).
*   **Phase Transition:** When system-wide threat temperature $\theta_T$ crosses a critical threshold $\theta_C$, net magnetization $M = \frac{1}{N} \sum \sigma_i$ collapses rapidly, representing a cascade compromise (e.g., automated lateral propagation).

---

### 2.3. Arrhenius Sandbox Barrier
The rate of an attacker escaping a sandbox container is modeled as a barrier crossing:
$$k_{compromise} = A e^{-\frac{\Delta E_{barrier}}{\theta_T}}$$
Where:
*   $\Delta E_{barrier}$: Activation energy barrier of the sandbox boundary (strengthened by Seccomp, AppArmor, user namespaces).
*   $A$: Attempt execution rate of the exploit.
*   $\theta_T$: Threat temperature. As threat temperature rises, the engine dynamically increases $\Delta E_{barrier}$ to clamp the escape rate.

---

## 3. Subsystem Mapping
*   **Engine Directory:** `07_security/physics/` (and the core engine at `07_security/security_physics/`)
*   **Components:**
    *   `entropy/`: Measures overall system microstate configurations.
    *   `energy/`: Computes privilege heights and configuration risks.
    *   `disorder/`: Combines entropy and energy to publish the Security Disorder Index (SDI).

---

## 4. Experiment Backlog

### Experiment R025: Disorder growth under active intrusion
*   **Objective:** Execute normal server workloads and contrast them with an active malware load (e.g. privilege escalation attempt followed by shell spawning). Calculate and plot SDI over time.
*   **Telemetry Source:** eBPF process lineage, network sockets, file configurations.
*   **Metrics:**
    *   Phase transition detection latency $\le 500$ ms.
    *   SDI sensitivity (true positive rate) $\ge 95\%$.
*   **Integration Target:** `07_security/physics/disorder`.
