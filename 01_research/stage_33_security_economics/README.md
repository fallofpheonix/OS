# Stage 33: Security Economics in Pheonix

## 1. Core Objective
To apply Mechanism Design and Security Economics—specifically Vickrey-Clarke-Groves (VCG) auctions—to schedule and allocate host and network resources among competing processes, containers, or tenants. By aligning individual strategic incentives with global system efficiency, Pheonix prevents resource-starvation attacks, mitigates denial-of-service attempts, and ensures truthful resource requirement reporting.

---

## 2. Mathematical Formulations

### 2.1. Vickrey-Clarke-Groves (VCG) Resource Allocation
Let:
*   $\mathcal{N} = \{1, 2, \dots, n\}$: Set of agents (processes, containers, or tenants).
*   $\mathcal{X}$: Set of feasible resource allocation configurations (e.g., allocation of CPU shares, memory pages, network queue slots).
*   $v_i(x) \in \mathbb{R}$: Private valuation of agent $i$ for allocation $x \in \mathcal{X}$, reflecting the agent's actual operational utility or priority.
*   $\hat{v}_i(x)$: Valuation reported (bid) by agent $i$ to the scheduler. An agent may misreport ($\hat{v}_i \neq v_i$) to hoard resources.

The Pheonix economic scheduler executes a two-step mechanism:

#### Step 1: Efficient Allocation Rule
Compute the socially optimal allocation $x^*$ that maximizes the sum of reported valuations:
$$x^* = \arg\max_{x \in \mathcal{X}} \sum_{j \in \mathcal{N}} \hat{v}_j(x)$$

#### Step 2: Externality Pricing Rule
Each agent $i$ is charged a payment $p_i$ representing the social cost (externality) they impose on the system by participating:
$$p_i = \max_{y \in \mathcal{X}} \sum_{j \neq i} \hat{v}_j(y) - \sum_{j \neq i} \hat{v}_j(x^*)$$

Where:
*   $\max_{y \in \mathcal{X}} \sum_{j \neq i} \hat{v}_j(y)$: The maximum social welfare achievable by all agents other than $i$ if agent $i$ did not exist.
*   $\sum_{j \neq i} \hat{v}_j(x^*)$: The actual social welfare achieved by all agents other than $i$ under the optimal allocation $x^*$.

#### Dominant-Strategy Incentive Compatibility (DSIC):
Since the utility of agent $i$ is $u_i = v_i(x^*) - p_i$, substituting $p_i$ yields:
$$u_i = v_i(x^*) + \sum_{j \neq i} \hat{v}_j(x^*) - \max_{y \in \mathcal{X}} \sum_{j \neq i} \hat{v}_j(y)$$
Maximizing $u_i$ is mathematically equivalent to maximizing $v_i(x^*) + \sum_{j \neq i} \hat{v}_j(x^*)$. The agent achieves this if and only if they report their true valuation ($\hat{v}_i = v_i$), making the mechanism strategy-proof.

---

### 2.2. Network QoS Auctions
For outbound network traffic, containers bid for priority slots in the queue. Let $B$ be the available bandwidth.
*   Each container $i$ bids a valuation curve $\hat{v}_i(b_i)$ for bandwidth $b_i$.
*   The network allocator solves:
    $$\max \sum_{i \in \mathcal{N}} \hat{v}_i(b_i) \quad \text{subject to} \quad \sum_{i \in \mathcal{N}} b_i \le B$$
*   Bandwidth throttling is applied based on the resulting VCG prices $p_i$, taxing high-bandwidth containers that degrade network quality of service for other tenants.

---

## 3. Subsystem Mapping
*   **Subsystem Directory:** `07_security/economics/`
    *   `vcg/`: VCG allocation solver and externality computation engine.
    *   `pricing/`: Resource-pricing regulator mapping cgroups parameters.

---

## 4. Experiment Backlog

### Experiment R033: VCG Truthful Resource Allocation
*   **Objective:** Simulate multiple containers bidding for CPU cycles under resource constraints. Introduce a malicious container attempting to starve other nodes by bidding artificially high values. Verify that VCG pricing charges the attacker an externality fee proportional to its disruption, rendering the starvation attack economically unviable.
*   **Telemetry Source:** Container CPU usage, bid histories, pricing records.
*   **Metrics:**
    *   Incentive Compatibility: Truthful bidding utility $\ge$ Strategic misreporting utility.
    *   Allocation Efficiency: Social welfare under VCG $\ge 90\%$ of theoretical maximum.
    *   Solver Latency: Pricing calculation time $\le 5$ ms.
*   **Integration Target:** `07_security/economics/vcg`.
