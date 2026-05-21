# Boot And Initialization Sequence

## Purpose

Define early startup order so hangs can be debugged with deterministic checkpoints.

## Firmware Interface

| Interface | Status | Notes |
|---|---|---|
| BIOS | Prototype only | Simple boot-sector experiments |
| UEFI | Primary | Required for modern boot path |

UEFI-specific requirements:

- Handle memory map.
- Handle Graphics Output Protocol if framebuffer is used.
- Preserve ACPI table pointers when needed.
- Exit boot services only after required boot data is captured.

## Bootloader Decision

| Loader | Status | Use |
|---|---|---|
| Limine | Preferred for scratch kernel | Modern protocol, low setup cost |
| GRUB | Supported | Multiboot workflows and ISO boot |
| Custom assembly loader | Prototype only | BIOS learning path |

## Scratch Boot Flow

```text
Firmware
  -> Bootloader
  -> Kernel entry
  -> Stack setup
  -> GDT load
  -> IDT load
  -> Memory map validation
  -> Paging enable
  -> Physical memory manager
  -> Kernel heap
  -> Timer
  -> Scheduler
  -> Init task
  -> Ring 3 userspace
```

## Early Kernel Initialization

Required sequence:

1. Validate boot protocol structure.
2. Initialize serial logging.
3. Validate CPU features.
4. Install GDT.
5. Install IDT.
6. Parse memory map.
7. Reserve kernel and boot regions.
8. Enable paging.
9. Initialize physical memory manager.
10. Initialize kernel heap.
11. Initialize interrupt controller and timer.
12. Start scheduler.

## User Space Transition

Ring 0 to Ring 3 requires:

- User code segment.
- User data segment.
- User stack.
- User page table mappings.
- Syscall or interrupt gate entry path.
- Controlled return using `iretq`, `sysret`, or equivalent architecture mechanism.

Initial userspace target:

```text
init -> shell -> halt/reboot
```

## Debug Checkpoints

Emit a serial log marker after each step:

```text
[boot] entry
[boot] serial
[boot] gdt
[boot] idt
[boot] mmap
[boot] paging
[boot] pmm
[boot] heap
[boot] timer
[boot] scheduler
[boot] userspace
```

## Failure Policy

- Early failure must print to serial if initialized.
- If serial is unavailable, halt in a known loop.
- Page fault handler must print fault address and error code.
- Triple faults must be reduced by validating GDT/IDT before interrupts.

