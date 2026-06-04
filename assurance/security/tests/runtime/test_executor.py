"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
from runtime.shell.executor import execute


def test_success_pwd():
    res = execute("pwd")
    assert res.success is True
    assert res.exit_code == 0
    assert res.stdout.strip() != ""
    assert res.duration_ms >= 0


def test_stderr_ls_nonexistent():
    # ls is allowed but should produce stderr for a non-existent entry
    res = execute("ls definitely_does_not_exist_1234")
    assert res.success is False
    assert res.exit_code != 0
    assert res.stderr != ""
    assert res.duration_ms >= 0


def test_invalid_command():
    res = execute("this_command_is_not_whitelisted")
    assert res.success is False
    assert "not allowed" in res.stderr.lower()
    assert res.duration_ms >= 0


def test_timeout():
    # sleep is whitelisted for testing; force a very short timeout
    res = execute("sleep 2", timeout=0.1)
    assert res.success is False
    assert res.exit_code == -1
    assert "timed out" in res.stderr.lower()
    assert res.duration_ms >= 0


def test_whitelist_policy_blocks_extra_args_for_pwd():
    res = execute("pwd extra")
    assert res.success is False
    assert "too many arguments" in res.stderr.lower()


def test_echo_length_policy():
    res = execute("echo " + ("x" * 201))
    assert res.success is False
    assert "too long" in res.stderr.lower()
