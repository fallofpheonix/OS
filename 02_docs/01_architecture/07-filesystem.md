# Filesystem

## Phases

| Phase | Target |
|---|---|
| 1 | FAT32 boot partition |
| 2 | Read-only initramfs or ramfs |
| 3 | EXT4 support for Linux-derived path |
| 4 | Custom filesystem research |

## Required Features

- Path resolution.
- File descriptors.
- Mount table.
- Read/write operations.
- Directory iteration.
- Permission checks.

## Future Features

- Journaling.
- Snapshots.
- Compression.
- Encryption.
- Quotas.

## Failure Considerations

- Metadata writes must be ordered.
- Mount failure must drop to recovery shell or panic with reason.
- Filesystem driver must reject malformed structures.
- Read-only root should be supported for live images.

