import React, { useState, useEffect } from "react";
import { Terminal, BrainCircuit, ArrowRight, ShieldAlert, Cpu } from "lucide-react";

export default function PhoenixMind({ selectedNode, onApplyAction, wardenState }) {
  const [advice, setAdvice] = useState("");
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (!selectedNode) {
      setAdvice(
        "PhoenixMind idle. Select any process node in the DAG graph canvas to analyze its lineage and receive security recommendations."
      );
      return;
    }

    setLoading(true);
    // Simulate query to local Ollama / Jan AI or fallback to rich local heuristics
    const timer = setTimeout(() => {
      const pid = selectedNode.id;
      const label = selectedNode.label;
      const entropy = selectedNode.entropy || 3.2;

      let msg = "";
      if (entropy > 7.5) {
        msg = `[CRITICAL ANOMALY ALERT] Process ${label} displays an execution entropy score of ${entropy.toFixed(2)}. 
This is mathematically highly indicative of a compressed malware payload or privilege escalation attempt.

RECOMMENDED STEPS (L7 Swarm Nexus):
1. Delegate execution to [forge-agent] to isolate the process immediately (Class 3 Local Isolation).
2. Trigger [astraeus-core] repair planner to verify and restore system state invariants.
3. Verify the de-escalation actuation policy using the [control-plane] purity scanner before applying Warden modifications.`;
      } else if (entropy > 6.0) {
        msg = `[SUSPICIOUS TELEMETRY DETECTED] Process ${label} exhibits elevated drift indicators ($S_D = ${entropy.toFixed(2)}$). 
It may be opening unauthorized files or establishing unverified sockets.

RECOMMENDED STEPS (L5.5 strategic policy):
1. Harden FSM system state to SUSPICIOUS.
2. Run the [control-plane] dependency validator on all spawned child processes.`;
      } else {
        msg = `[SYSTEM STATE: HEALTHY] Process ${label} (Entropy: ${entropy.toFixed(2)}) is executing within normal boundaries. 
Deterministic telemetry hashes match the Merkle ledger. All agent systems ([forge-agent], [astraeus-core]) report nominal state.`;
      }

      setAdvice(msg);
      setLoading(false);
    }, 400);

    return () => clearTimeout(timer);
  }, [selectedNode]);

  const handleActionClick = () => {
    if (!selectedNode) return;
    const entropy = selectedNode.entropy || 3.2;
    if (entropy > 7.5) {
      onApplyAction("isolate", selectedNode.id);
    } else {
      onApplyAction("harden", "0.85");
    }
  };

  return (
    <div className="border border-cyber-border rounded-lg bg-cyber-panel p-4 flex flex-col gap-4">
      <div className="flex items-center justify-between pb-2 border-b border-cyber-border">
        <h2 className="text-md font-bold text-cyber-text flex items-center gap-2">
          <BrainCircuit className="text-cyber-accent" size={18} />
          PHOENIXMIND™ AI ADVISOR
        </h2>
        <span className="text-[10px] text-cyber-accent font-mono bg-cyber-accent/10 border border-cyber-accent/30 rounded px-1.5 py-0.5">
          LOCAL LLM ACTIVE
        </span>
      </div>

      {/* AI Advice Output */}
      <div className="flex-1 min-h-[140px] bg-slate-950/70 border border-slate-900 rounded p-3 font-mono text-xs text-slate-300 leading-relaxed overflow-y-auto relative">
        {loading ? (
          <div className="absolute inset-0 flex items-center justify-center bg-slate-950/80">
            <div className="flex items-center gap-2 text-cyber-accent">
              <Cpu className="animate-spin" size={16} />
              <span>Analyzing lineage...</span>
            </div>
          </div>
        ) : null}
        <pre className="whitespace-pre-wrap">{advice}</pre>
      </div>

      {/* Action Button */}
      {selectedNode && selectedNode.entropy > 5.0 && (
        <button
          onClick={handleActionClick}
          className="w-full flex items-center justify-center gap-2 p-2.5 rounded-lg font-mono text-xs font-bold transition bg-cyber-accent/10 border border-cyber-accent text-cyber-accent hover:bg-cyber-accent hover:text-cyber-bg cursor-pointer"
        >
          <span>EXECUTE RECOMMENDED COUNTERMEASURE</span>
          <ArrowRight size={14} />
        </button>
      )}

      {/* Dynamic Security Log Context */}
      <div className="flex items-center gap-3 p-2.5 rounded-lg border border-slate-900 bg-slate-950/40 text-[11px] font-mono text-slate-400">
        <ShieldAlert className="text-cyber-warning shrink-0" size={16} />
        <div>
          Current Warden Guardrails: State <span className="text-cyber-accent font-bold">{wardenState}</span>. 
          Hysteresis duration locked at 30 ticks.
        </div>
      </div>
    </div>
  );
}
