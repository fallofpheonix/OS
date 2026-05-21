from agents.surface.orchestrator.minimal_brain import BrainSample, build_minimal_brain


def test_minimal_brain_escalates_high_risk_signal():
    brain = build_minimal_brain()
    decision = brain.evaluate(BrainSample(signal="telemetry-spike", score=0.92, evidence_count=3))

    assert decision.action == "escalate"
    assert decision.risk == "high"
    assert "score=0.92" in decision.explanation
    assert any("decision action=escalate" in line for line in decision.trace)


def test_minimal_brain_debug_trace_is_readable():
    brain = build_minimal_brain()
    trace = brain.debug_trace(BrainSample(signal="steady-state", score=0.12, evidence_count=1))

    assert "thresholds investigate=0.40 escalate=0.80" in trace
    assert "decision action=observe risk=low" in trace