# Package And Payload Ecosystem

## Scope

Applies to Linux-derived images built from Arch, Kali, Debian, LFS, or Buildroot.

## Repository Configuration

Arch:

```text
pacman.conf
mirrorlist
packages.x86_64
custom repo database
```

Kali/Debian:

```text
sources.list
apt preferences
package-lists/
includes.chroot/
hooks/
```

## Package Manifest

Every image build must emit:

```text
manifests/packages.txt
manifests/services.txt
manifests/filesystem.txt
manifests/image-metadata.txt
```

Minimum fields:

- Package name.
- Version.
- Repository.
- Architecture.
- Install reason.
- Checksum where available.

## Preinstalled Packages

Core:

- Kernel.
- Init system.
- Shell.
- Core utilities.
- Networking tools.
- Filesystem tools.
- Text editor.

Optional:

- Desktop environment.
- Installer.
- Security tools.
- Development tools.

Kali-specific:

- Prefer `kali-linux-core` and narrow tool groups.
- Avoid full metapackages until image size and legal scope are approved.

## Custom Configurations

Tracked payload areas:

```text
config/etc/
config/skel/
config/systemd/
config/themes/
config/security/
```

Examples:

- Dotfiles.
- `/etc/os-release`.
- Hostname.
- Systemd service overrides.
- Login banner.
- Firewall policy.
- Default shell profile.
- Theme assets.

## Service Policy

Default:

- No inbound network services.
- SSH server disabled unless explicitly required.
- Firewall enabled where available.
- Logging enabled.

## Signing And Integrity

Release images require:

- Package signature verification.
- ISO checksum.
- Manifest checksum.
- Immutable release artifact.

## Failure Considerations

- Missing package must fail the build.
- Unsigned package must fail release builds.
- Service enablement must be explicit.
- Embedded private material must fail release checks.

