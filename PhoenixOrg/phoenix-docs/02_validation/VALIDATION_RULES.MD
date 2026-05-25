# PhoenixOS: Validation Rules (L7)

Validation rules ensure that input events are structured, bounded, and verified before they impact the strategic control layers. At the core of system validation is the **Telemetry Confidence Score (TCS)**, which measures sensor network integrity.

---

## 1. The Determinism Invariant

Every active runtime module must enforce the pure input-output determinism rule:
$$\delta(S_t, E_{t+1}) \rightarrow (S_{t+1}, O_t)$$
A sequence of identical inputs processed in identical logical ticks must yield byte-for-byte identical state transitions and outputs.

---

## 2. Telemetry Confidence Score (TCS)

The `tcs.SlidingWindow` evaluates the integrity of raw telemetry events over a sliding window (typically $60\text{s}$).

### TCS Mathematical Formulation
The confidence score $T_C \in [0.0, 1.0]$ is a weighted combination of packet loss and normalized jitter:
$$T_C = w_1 \times (1 - P_{\text{loss}}) + w_2 \times (1 - J_{\text{norm}})$$

Where:
- $w_1 = 0.7$ (70% weight on packet delivery).
- $w_2 = 0.3$ (30% weight on arrival timing stability).

#### 1. Packet Loss Rate ($P_{\text{loss}}$)
$$P_{\text{loss}} = \min\left(1.0, \frac{\text{DroppedPackets}}{MaxSequenceID - MinSequenceID}\right)$$

#### 2. Normalized Jitter ($J_{\text{norm}}$)
$$J_{\text{norm}} = \min\left(1.0, \frac{AverageJitterNS}{50,000,000\text{ ns}}\right)$$
*Timing jitter is normalized against a $50\text{ms}$ threshold. Any average jitter above $50\text{ms}$ fully degrades the timing confidence component to $0.0$.*

---

## 3. Telemetry Security Shields

To prevent malicious sensor manipulation or network spoofing, the sliding window implements two security shields:

### Logical Time Lock (Anti-Time Warp)
To prevent attackers from injecting telemetry from the far future to skew historical metrics or advance the clock artificially:
$$\text{If } e.\text{Timestamp} > \text{WindowLatest} + \text{WindowSize} \implies \text{Drop Event}$$
Events arriving further than 1 window-size into the future relative to the latest seen event timestamp are discarded immediately.

### Gap Capping
To prevent massive sequence jumps (e.g., from network partitioning or packet spoofing) from permanently pegging the confidence score to $0.0$, perceived packet loss is capped at $100$ per sequence gap:
$$\text{If } \Delta Seq > 1 \implies \text{perceived\_loss} = \min(100, \Delta Seq - 1)$$

---

## 4. Bounded Telemetry Window

Memory footprint is strictly bounded to prevent memory exhaustion:
- Events older than `WindowSize` (60s) are pruned on every write.
- During pruning, historical losses decay by half to ensure the system recovers confidence when the network stabilizes:
  $$\text{DroppedPackets}_{t+1} = \frac{\text{DroppedPackets}_t}{2}$$
