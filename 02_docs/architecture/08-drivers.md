# Drivers

## Initial Driver Order

1. Serial console.
2. Timer.
3. Interrupt controller.
4. Keyboard.
5. Framebuffer.
6. Block device.
7. Network device.

## Driver Model

| Field | Requirement |
|---|---|
| Initialization | Deterministic order |
| Errors | Explicit status codes |
| Memory | DMA-safe allocation rules |
| Interrupts | Deferred work support |
| Userspace access | Device nodes or syscall API |

## Hardware Scope

Initial target:

- QEMU virtual hardware.
- UEFI firmware.
- VirtIO devices where possible.

Avoid broad hardware support until VM target is stable.

## Failure Considerations

- Driver probe failure must not corrupt global state.
- Interrupt handlers must acknowledge hardware correctly.
- DMA buffers must be pinned and aligned.
- Device reset paths must be explicit.

