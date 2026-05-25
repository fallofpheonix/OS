# Operations: Build Runbook

This document details the build sequence and commands required to compile all PhoenixOS node subsystems.

## 1. System Requirements
- Go $\ge 1.25$
- Clang/LLVM (for compiling eBPF C probes)
- Make
- Node.js (for the web console)

## 2. Syncing Workspace
PhoenixOS uses a Go workspace (`go.work`) to manage local dependencies. Always synchronize before compiling:
```bash
go mod tidy
go work sync
```

## 3. Compiling the Go Binaries
Compile the six core services in `cmd/`:
```bash
# Compile all subsystems
go build -o bin/kernel_agent cmd/kernel_agent/main.go
go build -o bin/truth_service cmd/truth/main.go
go build -o bin/trace_engine cmd/trace/main.go
go build -o bin/warden cmd/warden/main.go
go build -o bin/arbiter cmd/arbiter/main.go
go build -o bin/recovery_agent cmd/recovery/main.go
```

## 4. Compiling eBPF Probes (Kernel Layer)
To rebuild the kernel-level telemetry collectors:
```bash
cd phoenix_os/kernel/probes
clang -O2 -target bpf -c tracepoint.c -o tracepoint.lf.o
```

## 5. Web Console Compilation
Compile the React/Vite dashboard:
```bash
cd phoenix_web
npm install
npm run build
```
