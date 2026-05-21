# Userspace

## Components

- Init system.
- Shell.
- Core utilities.
- Service manager.
- Package manager.
- Optional window manager.

## Planned Commands

```text
ps
ls
cat
mount
kill
pkg
```

## Execution Model

| Area | Initial Design |
|---|---|
| Binary format | ELF |
| Process creation | Spawn from init |
| Standard streams | Console-backed initially |
| Permissions | UID/GID or capability model |
| Configuration | Text files under `/etc` equivalent |

## Future Work

- GUI stack.
- Wayland compatibility for Linux-derived path.
- Container runtime.
- Multi-user session management.

