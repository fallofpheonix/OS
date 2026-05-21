# Kernel Design

## Kernel Type

Hybrid initially.

Rationale:

- Simpler driver integration than microkernel.
- More modular than strict monolithic design.
- Suitable for staged OS research.

## Subsystems

### Memory Manager

- Physical page allocator.
- Virtual memory.
- Kernel heap.
- User address spaces.
- Guard pages.

### Process Manager

- Process ID allocation.
- Thread model.
- Context switching.
- User/kernel transition.
- Process lifecycle.

### Scheduler

- Round-robin initially.
- Priority queues later.
- Timer interrupt driven.
- No SMP until locking and interrupt routing are defined.

### IPC

- Shared memory.
- Message passing.
- Kernel-mediated handles.

### Drivers

- Serial.
- Keyboard.
- Mouse.
- Display.
- Storage.
- Network.

## Failure Considerations

- Panic path must work before heap allocation.
- Page fault handler must dump fault address and CPU state.
- Interrupt handlers must avoid unbounded allocation.
- Kernel stacks require overflow detection when paging is stable.

