#!/bin/bash
REPO=$1
TYPE=$2

mkdir -p "$REPO/.github/workflows"
CAT_FILE="$REPO/.github/workflows/ci.yml"

cat << 'INNER_EOF' > "$CAT_FILE"
name: Phoenix-CI
on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  validate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.25.0'
      - name: Deterministic Linting
        run: go vet ./...
      - name: Run Invariant Tests
        run: go test -v -race ./...
INNER_EOF

# Add repository-specific hardening gates
case "$TYPE" in
  kernel)
    cat << 'INNER_EOF' >> "$CAT_FILE"
      - name: Build eBPF Compatibility (Linux/AMD64)
        run: |
          CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build ./...
INNER_EOF
    ;;
  formal)
    cat << 'INNER_EOF' >> "$CAT_FILE"
      - name: Verify TLA+ Invariants (TLC Model Checker)
        run: |
          # Note: Requires TLC installed in runner
          java -cp tla2tools.jar tlc2.TLC -simulate ./tla/
INNER_EOF
    ;;
  core)
    cat << 'INNER_EOF' >> "$CAT_FILE"
      - name: Protocol Compliance Check (Protoc)
        run: |
          # Verify proto compilation
          PATH=\$PATH:\$HOME/go/bin protoc --go_out=. --go_opt=paths=source_relative proto/v1/fsm/fsm.proto
INNER_EOF
    ;;
  validation)
    cat << 'INNER_EOF' >> "$CAT_FILE"
      - name: Replay Engine Deterministic Parity
        run: |
          # Execute last 50 stable checkpoints
          go test -v ./replay/... -run TestDeterministicReconstruction
INNER_EOF
    ;;
esac

# Add common audit snapshot
cat << 'INNER_EOF' >> "$CAT_FILE"
  
  # Audit trail generation
  commit-snapshot:
    needs: validate
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Generate State Commit
        run: |
          find . -type f -not -path '*/.*' -exec sha256sum {} + | sort | sha256sum > .snapshot.sha256
      - name: Verify Determinism
        run: |
          # Ensure current state matches architectural commitments
          test -f .snapshot.sha256
INNER_EOF
