# Stage 1: The Linux Compilation Bridge
FROM golang:1.22-alpine AS builder

# Install LLVM, Clang, and Linux headers required for eBPF
RUN apk add --no-cache clang llvm make libbpf-dev linux-headers

WORKDIR /phoenix

# Copy the entire source tree
COPY . .

# Compile the C eBPF hooks into bytecode
RUN clang -O2 -g -target bpf -D__TARGET_ARCH_x86 
    -I/usr/include 
    -c phoenix_os/warden/ebpf/src/egress_policy.c -o egress_policy.o

RUN clang -O2 -g -target bpf -D__TARGET_ARCH_x86 
    -I/usr/include 
    -c phoenix_os/warden/ebpf/src/xdp_ingress.c -o xdp_ingress.o

# Compile the Go Warden Orchestrator (statically linked)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 
    go build -ldflags="-w -s" -o phoenix-warden phoenix_os/warden/ebpf/orchestrator/main.go

# Stage 2: The Minimalist Execution Layer
FROM scratch
COPY --from=builder /phoenix/phoenix-warden /phoenix-warden
COPY --from=builder /phoenix/egress_policy.o /egress_policy.o
COPY --from=builder /phoenix/xdp_ingress.o /xdp_ingress.o

# Set the Go binary as the absolute entrypoint
ENTRYPOINT ["/phoenix-warden"]
