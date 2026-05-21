# RFC-010: Phoenix Kernel (L2) Validation Strategy

## Status
Draft

## 1. Objective
This RFC defines the comprehensive test and validation strategy for the Phoenix Kernel's eBPF-based telemetry probes. The goal is to ensure their stability, data integrity, and performance before they are integrated into the main OS image.

## 2. Test Environment
*   **Virtualization:** QEMU/KVM
*   **Base OS:** Arch Linux (latest stable)
*   **Kernel Version:** Linux 6.1 LTS (or current project target)
*   **Tooling:** `pytest` for orchestration, custom Go programs for generating load and verifying output.

## 3. Test Cases & Validation

### TC-K1: Probe Attachment & Verification
*   **Test:** The telemetry agent attempts to load and attach all eBPF probes (`sched_process_fork`, `sys_enter_openat`, etc.).
*   **Metric:** The agent must report 100% successful attachments.
*   **Pass/Fail:** Failure to attach any single probe results in a test failure.

### TC-K2: Data Integrity Validation
*   **Test:** A test script will generate a known sequence of system events (e.g., create a file, write "hello", execute a child process).
*   **Metric:** The telemetry data captured by the agent in userspace must exactly match the sequence and content of the generated events.
*   **Pass/Fail:** Any discrepancy in arguments, paths, or event order is a failure.

### TC-K3: Performance & Overhead
*   **Test:** A load generator will spawn 10,000 short-lived processes per second for 60 seconds.
*   **Metric:** CPU overhead attributable to the eBPF probes and the telemetry agent must remain below **2%** of a single core. Memory usage must not exceed **50MB**.
*   **Pass/Fail:** Exceeding either the CPU or memory budget is a failure.

### TC-K4: Probe Unload/Failure Detection
*   **Test:** A probe (e.g., `vfs_write`) will be forcibly detached from the kernel during a test run.
*   **Metric:** The telemetry agent must detect the probe's absence within 5 seconds and raise a "Sensor Blinding" critical alert.
*   **Pass/Fail:** Failure to raise the alert within the time window is a failure.

## 4. Automation
These tests will be automated in a separate GitHub Actions workflow (`.github/workflows/kernel_validation.yml`) that runs on a self-hosted runner with KVM capabilities.
