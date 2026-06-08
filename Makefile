# Master Phoenix Ignition Makefile
#
# All build/test targets iterate over every module listed in /Users/fallofpheonix/os/go.work
# so the workspace is always validated end-to-end. Add a module to go.work and
# it is automatically included in `make build`, `make test`, etc.

WORKSPACE := $(shell pwd)
MODULES   := $(shell grep -oE '\./[^ ]+' go.work)

.PHONY: build build-one test test-race vet lint fmt tidy cover ci clean ignite logs help

# ---------------------------------------------------------------------------
# Build
# ---------------------------------------------------------------------------

build:
	@for mod in $(MODULES); do \
		echo "=== $$mod ==="; \
		(go build ./$$mod/...) || exit 1; \
	done

build-one:
	@if [ -z "$(MODULE)" ]; then \
		echo "MODULE is required, e.g. make build-one MODULE=Phoenix.Nucleus/PhoenixGuard"; \
		exit 1; \
	fi
	go build ./$(MODULE)/...

# ---------------------------------------------------------------------------
# Hygiene & Analysis
# ---------------------------------------------------------------------------

fmt:
	@for mod in $(MODULES); do \
		echo "=== Formatting $$mod ==="; \
		(cd $$mod && go fmt ./...) || exit 1; \
	done

tidy:
	@for mod in $(MODULES); do \
		echo "=== Tidying $$mod ==="; \
		(cd $$mod && go mod tidy) || exit 1; \
	done

# ---------------------------------------------------------------------------
# Test
# ---------------------------------------------------------------------------

test:
	@for mod in $(MODULES); do \
		echo "=== $$mod ==="; \
		(go test -count=1 ./$$mod/...) || exit 1; \
	done

test-race:
	@for mod in $(MODULES); do \
		echo "=== $$mod ==="; \
		(go test -race -count=1 ./$$mod/...) || exit 1; \
	done

# ---------------------------------------------------------------------------
# Repository invariant checks
# ---------------------------------------------------------------------------

check-invariants:
	@echo "=== PACKAGE-002: Scanning for silent stub patterns ==="; \
	TMPFILE=$$(mktemp); \
	for mod in $(MODULES); do \
		find "$$mod" -name '*.go' ! -name '*_test.go' ! -name '*.pb.go' -exec grep -l "STATUS: STUB" {} \; 2>/dev/null >> "$$TMPFILE"; \
	done; \
	STUB_COUNT=$$(sort -u "$$TMPFILE" | wc -l | tr -d ' '); \
	echo "  Files marked STATUS: STUB: $$STUB_COUNT"; \
	for f in $$(sort -u "$$TMPFILE"); do \
		echo "    $$f"; \
	done; \
	rm -f "$$TMPFILE"; \
	echo "  OK"; \
	echo ""; \
	echo "=== PACKAGE-001: Testing stub count ==="; \
	TMPFILE2=$$(mktemp); \
	for mod in $(MODULES); do \
		find "$$mod" -name '*_test.go' -exec dirname {} \; 2>/dev/null | sort -u >> "$$TMPFILE2"; \
	done; \
	TEST_ONLY=0; \
	for dir in $$(sort -u "$$TMPFILE2"); do \
		non_test=$$(find "$$dir" -maxdepth 1 -name '*.go' ! -name '*_test.go' 2>/dev/null | wc -l | tr -d ' '); \
		if [ "$$non_test" -eq 0 ]; then \
			echo "    $$dir (test-only)"; \
			TEST_ONLY=$$((TEST_ONLY + 1)); \
		fi; \
	done; \
	rm -f "$$TMPFILE2"; \
	echo "  Test-only packages: $$TEST_ONLY"; \
	echo "  OK"

# ---------------------------------------------------------------------------
# Static analysis
# ---------------------------------------------------------------------------

vet:
	@for mod in $(MODULES); do \
		echo "=== $$mod ==="; \
		(go vet ./$$mod/...) || exit 1; \
	done

lint:
	@GOLANGCI_BIN="$$(command -v golangci-lint 2>/dev/null || echo "$$(go env GOPATH 2>/dev/null)/bin/golangci-lint")"; \
	if [ ! -x "$$GOLANGCI_BIN" ]; then \
		echo "golangci-lint not installed. Install with:"; \
		echo "  brew install golangci-lint"; \
		echo "  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
		exit 1; \
	fi; \
	for mod in $(MODULES); do \
		echo "=== Linting $$mod ==="; \
		(cd $$mod && $$GOLANGCI_BIN run --timeout 5m --exclude-dirs node_modules ./... 2>&1) || exit 1; \
	done

# ---------------------------------------------------------------------------
# Coverage
# ---------------------------------------------------------------------------

cover:
	@mkdir -p coverage
	@for mod in $(MODULES); do \
		out=coverage/$$(echo $$mod | tr / _).out; \
		echo "=== $$mod -> $$out ==="; \
		(go test -count=1 -coverprofile=$$out -covermode=atomic ./$$mod/...) || exit 1; \
	done
	@echo "Coverage profiles written to coverage/"

# ---------------------------------------------------------------------------
# Full CI suite
# ---------------------------------------------------------------------------

ci: vet test-race lint check-invariants
	@echo "--- Full CI suite passed ---"

# ---------------------------------------------------------------------------
# Container / runtime
# ---------------------------------------------------------------------------

clean:
	@echo "Cleaning up workspace..."
	-docker compose down -v
	rm -rf bin/
	rm -rf build_logs/
	rm -f *.out
	rm -f all_md_files.txt all_packages.txt all_referenced_packages.txt cycles_analysis.txt dependency_analysis.txt dependency_raw.txt
	find . -name "*.log" -type f -not -path '*/.*' -delete
	find . -name "*.out" -type f -not -path '*/.*' -delete
	@echo "--- Cleanup complete ---"

ignite:
	docker compose up -d --build

logs:
	docker compose logs -f

# ---------------------------------------------------------------------------
# Help
# ---------------------------------------------------------------------------

help:
	@echo "Phoenix workspace Makefile"
	@echo ""
	@echo "Modules (from go.work):"
	@for mod in $(MODULES); do echo "  $$mod"; done
	@echo ""
	@echo "Targets:"
	@echo "  build        Build every module"
	@echo "  build-one    Build a single module (set MODULE=...)"
	@echo "  test         Run go test for every module"
	@echo "  test-race    Run go test -race for every module"
	@echo "  vet          Run go vet for every module"
	@echo "  lint         Run golangci-lint for every module"
	@echo "  cover        Run tests with coverage, write profiles to coverage/"
	@echo "  check-invariants  Check for silent stubs and empty packages"
	@echo "  ci           Run check-invariants + vet + test-race + lint"
	@echo "  clean        Stop docker compose stack"
	@echo "  ignite       Start docker compose stack"
	@echo "  logs         Tail docker compose logs"
