#!/bin/bash
set -e

echo "=== Running Boundary Check ==="

VIOLATIONS=0

# Rule 1: Contracts cannot import concrete implementations
echo "Rule 1: Verifying contracts do not import implementations..."
IMPORTS=$(go list -f '{{join .Imports "\n"}}' ./core/Phoenix.Nucleus/PhoenixCore/contracts/... 2>/dev/null | sort -u)
for imp in $IMPORTS; do
    if [[ "$imp" == *"PhoenixValidation"* || "$imp" == *"PhoenixGuard"* ]]; then
        echo "  [VIOLATION] Contract imports implementation: $imp"
        VIOLATIONS=$((VIOLATIONS + 1))
    fi
done

# Rule 2: Production PhoenixCore (excluding adapters) cannot import PhoenixValidation or PhoenixGuard
echo "Rule 2: Verifying PhoenixCore production packages do not import Validation/Guard..."
# Find all packages under PhoenixCore, filtering out adapters, contracts, and test files
PACKAGES=$(go list ./core/Phoenix.Nucleus/PhoenixCore/... 2>/dev/null | grep -v "/adapters" | grep -v "/contracts" | grep -v "_test" | grep -v "/testing" | grep -v "/proto")
for pkg in $PACKAGES; do
    IMPORTS=$(go list -f '{{join .Imports "\n"}}' "$pkg" 2>/dev/null || true)
    for imp in $IMPORTS; do
        if [[ "$imp" == *"PhoenixValidation"* || "$imp" == *"PhoenixGuard"* ]]; then
            echo "  [VIOLATION] Package $pkg imports: $imp"
            VIOLATIONS=$((VIOLATIONS + 1))
        fi
    done
done

# Rule 3: Guard cannot import Validation
echo "Rule 3: Verifying Guard packages do not import Validation..."
IMPORTS_GUARD=$(go list -f '{{join .Imports "\n"}}' ./core/Phoenix.Nucleus/PhoenixGuard/... 2>/dev/null | sort -u)
for imp in $IMPORTS_GUARD; do
    if [[ "$imp" == *"PhoenixValidation"* ]]; then
        echo "  [VIOLATION] Guard imports Validation: $imp"
        VIOLATIONS=$((VIOLATIONS + 1))
    fi
done

# Rule 4: Guard cannot import Runtime internals (excluding common config, bus, or adapters)
echo "Rule 4: Verifying Guard packages do not import internal Core runtime packages..."
for imp in $IMPORTS_GUARD; do
    if [[ "$imp" == *"PhoenixCore"* ]]; then
        if [[ "$imp" != *"PhoenixCore/contracts"* && "$imp" != *"PhoenixCore/bus"* && "$imp" != *"PhoenixCore/common/config"* ]]; then
            echo "  [VIOLATION] Guard imports internal Core runtime package: $imp"
            VIOLATIONS=$((VIOLATIONS + 1))
        fi
    fi
done

# Rule 5: Security contracts cannot import Guard
echo "Rule 5: Verifying Security contracts do not import Guard..."
IMPORTS_SEC_CONTRACTS=$(go list -f '{{join .Imports "\n"}}' ./core/Phoenix.Nucleus/PhoenixCore/contracts/security/... 2>/dev/null | sort -u)
for imp in $IMPORTS_SEC_CONTRACTS; do
    if [[ "$imp" == *"PhoenixGuard"* ]]; then
        echo "  [VIOLATION] Security contract imports Guard: $imp"
        VIOLATIONS=$((VIOLATIONS + 1))
    fi
done

if [ $VIOLATIONS -eq 0 ]; then
    echo "=== Boundary Check PASSED ==="
    exit 0
else
    echo "=== Boundary Check FAILED with $VIOLATIONS violations ==="
    exit 1
fi
