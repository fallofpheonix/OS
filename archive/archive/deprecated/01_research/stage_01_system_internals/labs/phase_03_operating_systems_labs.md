# Phase 3 Labs: Operating Systems

## Rule

Research stays in `01_research/`. Implementation evidence belongs in `14_experiments/poc/stage_01_system_internals/phase_03_operating_systems/`.

## Lab 01: Process Lifecycle

- Draw process state machine.
- Trace `fork`, `exec`, `wait`, and termination.
- Record zombie cleanup behavior.

## Lab 02: Scheduler

- Simulate FCFS, RR, priority aging, and MLFQ.
- Measure waiting time, turnaround time, and response time.
- Record fairness and throughput trade-offs.

## Lab 03: Context Switching

- Trace register/state save and restore.
- Record voluntary vs involuntary context-switch examples.
- Measure switch overhead where possible.

## Lab 04: Virtual Memory

- Walk page-table translation.
- Model page faults.
- Compare page replacement policies.
- Record CoW behavior.

## Lab 05: VFS And Mini Filesystem

- Implement inode metadata.
- Implement directory lookup.
- Implement open/read/write/delete.
- Persist to a disk image.

## Lab 06: xv6 Modification

- Add syscall.
- Add priority scheduling.
- Add simple file operation.
- Record source changes and QEMU output.

## Completion Record

| Lab | Status | Evidence path | Notes |
|---|---|---|---|
| Process lifecycle | TODO | TBD | states, fork/exec/wait |
| Scheduler | TODO | TBD | RR, priority, MLFQ |
| Context switching | TODO | TBD | save/restore and overhead |
| Virtual memory | TODO | TBD | paging and page faults |
| VFS and mini filesystem | TODO | TBD | inode and persistence |
| xv6 modification | TODO | TBD | syscall, scheduler, file op |

