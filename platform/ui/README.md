# Phoenix.UI

The User Interface and UX Research domain for PhoenixOS.

## Overview

This directory houses the UI components, design systems, and external research repositories used to build the PhoenixOS interface. It focuses on high-performance rendering (GPU-accelerated), modern desktop environment UX, and modular design systems.

## Structure

- `External/`: Cloned open-source repositories for research and integration.
- `docs/`: (Planned) UI-specific documentation and design guidelines.

## External Repositories

The following repositories are currently cloned for research:

1. **Seelen UI**: Customizable desktop environment with tiling WM.
2. **Ghostty**: GPU-accelerated terminal emulator (Zig/Swift).
3. **WezTerm**: GPU-accelerated terminal and multiplexer (Rust).
4. **Plus UI**: Comprehensive design system for React/Tailwind.
5. **Desktop UI**: Native macOS/Windows look-and-feel components.
6. **Uiverse Galaxy**: Large library of community-built UI elements.
7. **macOS Vue**: Vue-based macOS simulation.
8. **ProzillaOS**: Web-based OS interface (React/Vite).
9. **TailAdmin**: Next.js admin dashboard template.
10. **React Desktop**: macOS and Windows component library for React.

## Integration Strategy

- Use **Seelen UI** patterns for the core desktop environment.
- Embed **Ghostty** via `libghostty` for high-performance terminal access.
- Utilize **Plus UI** for standardized component development across agent interfaces.
