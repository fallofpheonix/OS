# Research backlog: R002 - Real-time File Write Entropy Detection

## 1. Metadata
*   **Experiment ID:** R002
*   **Subsystem:** Filesystem / Anomaly Detection
*   **Priority:** P0 (Entropy Detection Foundation)
*   **Validation Status:** Proposed (Performance gate pending)
*   **Risk Level:** Medium (Context switching and buffer copies from kernel to user space)
*   **Pilot Mapping:** Phoenix Correlation & IPS Engine (`07_security/ips/`)

---

## 2. Objectives
Develop a real-time, low-overhead file write monitoring pipeline to detect ransomware attacks by calculating Shannon entropy on data buffers before they are written to disk.

---

## 3. Threat Path
*   **Attack Vectors:** Ransomware encrypting user directories, data exfiltration via encrypted channels, and wiper attacks replacing system files with high-entropy noise.
*   **Threat Assumptions:** Attackers will attempt to evade detection by encrypting files in small chunk intervals, executing slow-encryption loops, or using custom encoding to reduce apparent entropy.

---

## 4. Telemetry Source
*   **Production hooks:** eBPF kprobes (`kprobe/vfs_write`) or LSM hooks (`lsm/file_permission`).
*   **Simulation hooks:** macOS local file modification event generator sending mock high-entropy bytes.

---

## 5. Experiment Proposal
Build a prototype analyzer that:
1.  Hooks write system calls and reads the head (first 4KB) of the buffer being written.
2.  Applies Shannon's entropy algorithm to compute randomness:
    $$H(X) = -\sum_{i=1}^{n} P(x_i) \log_2 P(x_i)$$
3.  If the calculated entropy exceeds **7.5** (indicative of compression or encryption) and the rate of writes across different files exceeds a threshold (e.g. >10 files/sec), flag the PID as highly anomalous.

---

## 6. Metrics & Performance Budget
*   **Entropy Calculation Overhead:** Must take **< 15 microseconds** per 4KB block.
*   **Resource Budget:** CPU usage must not exceed **5%** during active file copy operations.
*   **False Positive Rate:** **< 0.1%** during normal development activities (e.g., git commits, GCC compilations).

---

## 7. Evidence Required
*   Incident alerts in the Phoenix Bus containing: PID, target file path, calculated entropy score, write rate, and timestamp.
*   Verified containment trigger showing the suspension of the offending process.

---

## 8. Validation Gate
1.  **Sensitivity:** Compiling a standard text file should register low entropy (<5.0). Writing an encrypted archive (e.g., zip, gpg) or AES-encrypted stream must register high entropy (>7.8) and trigger an alert.
2.  **Mitigation:** Verify that the system call returns a blocked error or suspends the target PID immediately if threshold rules are violated.

---

## 9. Integration Target
*   [07_security/ips/](file:///Users/fallofPhoenix/os/07_security/ips/) (IPS blocking logic)
*   [06_ai/anomaly_detection/](file:///Users/fallofPhoenix/os/06_ai/anomaly_detection/) (Adaptive anomaly thresholds)
