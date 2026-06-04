"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
from __future__ import annotations

import unittest

from runtime.shell.executor import execute
from runtime.tracing import RuntimeTrace


class AdversarialShellSemanticTests(unittest.TestCase):
    """Tests that prove shell runtime semantics remain deterministic."""

    def test_policy_rejection_emits_trace(self):
        result = execute("this_command_is_not_whitelisted")

        self.assertFalse(result.success)
        self.assertIsNotNone(result.trace)
        self.assertFalse(result.trace.success)
        self.assertEqual(result.trace.runtime_category, "shell")

    def test_repeated_policy_rejection_consistent(self):
        results = [execute("forbidden_command") for _ in range(3)]

        self.assertTrue(all(not r.success for r in results))
        self.assertTrue(all(isinstance(r.trace, RuntimeTrace) for r in results))
        self.assertTrue(all(not r.trace.success for r in results))
        self.assertTrue(all("not allowed" in r.stderr.lower() for r in results))

    def test_timeout_failure_emits_trace(self):
        result = execute("sleep 10", timeout=0.01)

        self.assertFalse(result.success)
        self.assertEqual(result.exit_code, -1)
        self.assertIsNotNone(result.trace)
        self.assertFalse(result.trace.success)

    def test_argument_policy_violation_emits_trace(self):
        result = execute("pwd arg1 arg2")

        self.assertFalse(result.success)
        self.assertIsNotNone(result.trace)
        self.assertFalse(result.trace.success)
        self.assertIn("too many arguments", result.stderr.lower())

    def test_trace_duration_consistent_with_result(self):
        result = execute("pwd")

        self.assertEqual(result.duration_ms, result.trace.duration_ms)

    def test_trace_is_immutable_under_failure(self):
        result = execute("forbidden")
        trace = result.trace

        with self.assertRaises(Exception):
            trace.success = True  # type: ignore[misc]


if __name__ == "__main__":
    unittest.main()
