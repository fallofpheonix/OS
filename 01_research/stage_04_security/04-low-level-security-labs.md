# Low-Level Security Labs

## Scope

Controlled exploit-engineering and mitigation-analysis labs for understanding memory safety, kernel attack surface, and defensive hardening.

Use only in isolated lab environments.

## 1. Memory Corruption Basics

### Topics

- Stack overflow.
- Buffer overflow.
- Saved return address corruption.
- Saved frame pointer corruption.
- Heap exploitation.
- Use-after-free.
- Double-free.
- Heap metadata corruption.
- C++ `vptr` overwrite.
- Race conditions.
- TOCTOU.

### Tools

- `pwndbg`.
- GEF.
- GDB.
- AddressSanitizer.
- QEMU snapshots.

### Lab

Create controlled vulnerable binaries:

- Stack overwrite.
- Heap use-after-free.
- Double-free.
- TOCTOU file check.

Debug objectives:

- Inspect `RIP`.
- Inspect saved `RBP`.
- Inspect stack frame layout.
- Inspect heap metadata.
- Confirm crash reason.

### Exit Criteria

- Each crash is reproducible.
- Root cause is documented.
- Mitigation behavior is recorded.
- No payload targets real systems.

## 2. Code-Reuse And Exploit Chains

### Topics

- ROP.
- JOP.
- Gadgets.
- `ret` chains.
- Indirect jumps.
- Dispatch tables.
- `mprotect`.
- `read`.
- `write`.
- `syscall`.

### Tools

- `ROPgadget`.
- `pwndbg`.
- GEF.
- GDB.

### Lab

Controlled binary only:

1. Compile target with known mitigations.
2. Identify gadgets with `ROPgadget`.
3. Build a ROP chain in a lab script.
4. Step through the chain in GDB.
5. Record how `RIP` changes after each `ret`.

### Exit Criteria

- Gadget list is documented.
- Stack layout is diagrammed.
- Mitigation state is recorded.
- Chain execution is explained instruction-by-instruction.

## 3. Modern Mitigations

### Topics

- ASLR.
- KASLR.
- DEP.
- NX.
- SMEP.
- SMAP.
- Stack canaries.
- Kernel stack cookies.
- RELRO.
- PIE.

### Tools

- `checksec`.
- `pwndbg`.
- GEF.
- QEMU.
- `/proc/sys/kernel/randomize_va_space`.

### Lab

Exercises:

- Trigger `__stack_chk_fail`.
- Compare binary layout with and without PIE.
- Observe NX behavior.
- Compare addresses across ASLR-enabled runs.
- Test exploit assumptions under mitigations.

### Exit Criteria

- Mitigation matrix exists.
- Each mitigation has a concrete observed effect.
- Bypass discussion remains lab-bound and defensive.

## 4. Kernel-Level Attacks

### Topics

- Kernel exploitation.
- Driver bugs.
- Filesystem bugs.
- IPC bugs.
- Scheduler bugs.
- Privilege escalation.
- `cred` structure abuse.
- Rootkits.
- Bootkits.
- Hypervisor attacks.

### Tools

- QEMU.
- KGDB.
- GDB.
- GEF or `pwndbg`.
- Kernel debug symbols.

### Lab

Use QEMU-based intentionally vulnerable kernels or toy drivers.

Study:

- TOCTOU.
- Race-based UAF.
- Faulty ioctl.
- Incorrect copy-to/from-user handling.
- Controlled credential-change demonstration in a toy kernel or vulnerable lab target.

### Exit Criteria

- Runs only in VM.
- Kernel image is disposable.
- Snapshot restore is tested.
- Crash dump or serial log is captured.

## 5. Tools In Practice

### pwndbg And GEF

Use for:

- Registers.
- Stack.
- Heap.
- Disassembly.
- Canary location.
- Memory protections.
- Context views.

### ROPgadget

Use for:

- Gadget discovery.
- libc gadget inventory.
- Offline gadget database per known binary/libc version.

### QEMU

Use for:

- Isolated execution.
- Multi-architecture testing.
- Snapshot rollback.
- Kernel GDB stub.

## Suggested Repo Structure

```text
low_level/
├── 01_memory_corruption/
│   ├── README.md
│   ├── stack_overflow/
│   ├── heap_uaf/
│   └── notes/
├── 02_rop_jop/
│   ├── README.md
│   ├── targets/
│   ├── gadgets/
│   └── traces/
├── 03_mitigations/
│   ├── README.md
│   ├── checksec_reports/
│   └── experiments/
├── 04_kernel_exploit/
│   ├── README.md
│   ├── qemu/
│   ├── vulnerable_driver/
│   └── kgdb_notes/
└── 05_rootkits_bootkits/
    ├── README.md
    ├── detection_notes.md
    └── defensive_analysis.md
```

## 8-Week Lab Plan

| Week | Focus | Output |
|---:|---|---|
| 1 | GDB, pwndbg, GEF | debugger workflow |
| 2 | Stack corruption | crash trace and stack diagram |
| 3 | Heap corruption | heap metadata notes |
| 4 | ROP basics | gadget inventory and trace |
| 5 | Mitigations | mitigation matrix |
| 6 | Race and TOCTOU | reproducible lab race |
| 7 | Kernel driver bugs | QEMU/KGDB lab |
| 8 | Defensive summary | hardening checklist |

## Safety Boundary

These modules are for defensive understanding, exploit mitigation, detection engineering, and lab-only reproduction.

Do not run exploit code against systems you do not own or administer with explicit authorization.

