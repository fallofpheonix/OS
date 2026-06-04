# PhoenixOS Shell: Integrated Model

A sophisticated, high-performance React-based shell for PhoenixOS, designed to provide a human-centric interface for the underlying substrate.

## Core Applications

### 1. Terminal
- **Purpose**: Direct operator interface for system interaction.
- **Features**: Mock command system (`help`, `status`, `domains`, `neofetch`), scrolling history, and auto-focus.
- **Aesthetic**: Minimalist JetBrains Mono typography with syntax highlighting.

### 2. System Monitor
- **Purpose**: Real-time telemetry for the 6 core domains.
- **Metrics**: CPU Load, Memory Usage, Authority Level, and Reconstruction Integrity.
- **Domain Health**: Individual status tracking for Nucleus, Cognition, Crucible, Terminus, UI, and Arbiter.

### 3. Control Center
- **Purpose**: System-wide configuration and governance.
- **Categories**: 
  - **Governance**: Policy enforcement and authority delegation.
  - **Security**: Formal verification status and audit logging.
  - **Cognition**: Model provider selection and semantic cache management.
  - **Substrate**: Ledger integrity and fracture isolation.

## UX & Visuals

- **Substrate Rendering**: GPU-accelerated animated gradient background using `@firecms/neat`.
- **Windowing System**: Multi-window support with z-index depth management and smooth state transitions.
- **Glassmorphism**: High-blur backdrop filters and semi-transparent surfaces for a modern "OS" feel.
- **Launcher & Taskbar**: Centralized app access with active-state tracking and system telemetry.

## Technical Stack

- **Framework**: React 18 (TypeScript)
- **Bundler**: Vite
- **Styling**: Vanilla CSS (Post-CSS compliant)
- **Animations**: Framer Motion
- **Icons**: Lucide React
- **Rendering**: WebGL (via `@firecms/neat`)

## Usage

```bash
cd Phoenix.UI/Shell
npm install
npm run dev
```
