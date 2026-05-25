import React, { useState, useEffect } from "react";
import SocDash from "./pages/SocDash";
import ReplayTimeline from "./pages/ReplayTimeline";
import PhoenixMind from "./ai/PhoenixMind";
import { Shield, Trophy, Activity, Award, RefreshCw, AlertCircle, FileText } from "lucide-react";

export default function App() {
  const [events, setEvents] = useState([]);
  const [currentSeqID, setCurrentSeqID] = useState(1);
  const [graphData, setGraphData] = useState({ nodes: [], edges: [] });
  const [scoreState, setScoreState] = useState({
    score: 0,
    level: "Novice SOC Analyst",
    multiplier: 1.0,
    badges: [],
    completed_challenges: []
  });
  const [selectedNodeId, setSelectedNodeId] = useState(null);
  const [selectedNode, setSelectedNode] = useState(null);
  const [wardenState, setWardenState] = useState("NORMAL");
  const [alert, setAlert] = useState(null);
  const [ecosystemConfig, setEcosystemConfig] = useState({
    physics_thresholds: [],
    math_registry: []
  });

  const API_BASE = "http://localhost:8080";

  // Fetch initial data
  useEffect(() => {
    fetchEvents();
    fetchScore();
    fetchEcosystemConfig();
  }, []);

  // Fetch graph data whenever timeline scrolls
  useEffect(() => {
    fetchGraph(currentSeqID);
  }, [currentSeqID]);

  // Synchronize active node details
  useEffect(() => {
    if (selectedNodeId && graphData.nodes) {
      const node = graphData.nodes.find((n) => n.id === selectedNodeId);
      setSelectedNode(node || null);
    } else {
      setSelectedNode(null);
    }
  }, [selectedNodeId, graphData]);

  // Dynamically update simulated Warden state based on timeline position (heuristics match Go FSM)
  useEffect(() => {
    if (events.length === 0) return;
    const activeEvents = events.filter((e) => e.seq_id <= currentSeqID);
    
    // Find the latest high entropy event
    let maxEntropy = 0;
    activeEvents.forEach((e) => {
      const entropy = e.payload?.entropy_score || 3.2;
      if (entropy > maxEntropy) {
        maxEntropy = entropy;
      }
    });

    if (maxEntropy > 7.5) {
      // If malware executes, Warden escalates to CONTAINED
      setWardenState("CONTAINED");
    } else if (maxEntropy > 6.0) {
      // If suspicious activity occurs, Warden escalates to SUSPICIOUS
      setWardenState("SUSPICIOUS");
    } else {
      setWardenState("NORMAL");
    }
  }, [currentSeqID, events]);

  const fetchEvents = async () => {
    try {
      const res = await fetch(`${API_BASE}/events`);
      if (res.ok) {
        const data = await res.json();
        setEvents(data);
        if (data.length > 0) {
          setCurrentSeqID(data.length); // Start at the end of the scenario
        }
      }
    } catch (err) {
      console.error("Failed to fetch events", err);
    }
  };

  const fetchScore = async () => {
    try {
      const res = await fetch(`${API_BASE}/game/score`);
      if (res.ok) {
        const data = await res.json();
        setScoreState(data);
      }
    } catch (err) {
      console.error("Failed to fetch score", err);
    }
  };

  const fetchEcosystemConfig = async () => {
    try {
      const res = await fetch(`${API_BASE}/game/config`);
      if (res.ok) {
        const data = await res.json();
        setEcosystemConfig(data);
      }
    } catch (err) {
      console.error("Failed to fetch ecosystem config", err);
    }
  };

  const fetchGraph = async (seqID) => {
    try {
      const res = await fetch(`${API_BASE}/graph?seq_id=${seqID}`);
      if (res.ok) {
        const data = await res.json();
        setGraphData(data);
      }
    } catch (err) {
      console.error("Failed to fetch graph", err);
    }
  };

  const handleApplyAction = async (action, target) => {
    try {
      const res = await fetch(`${API_BASE}/game/action`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ action, target })
      });
      if (res.ok) {
        const data = await res.json();
        setAlert({
          type: data.success ? "success" : "error",
          message: data.message
        });
        fetchScore();
        if (data.success && action === "isolate") {
          // Temporarily visual feedback: change node class to normal/green to show neutralized
          setGraphData((prev) => {
            const updatedNodes = prev.nodes.map((n) => {
              if (n.id === target) {
                return { ...n, group: "normal", label: `${n.label} (CONTAINED)` };
              }
              return n;
            });
            return { ...prev, nodes: updatedNodes };
          });
        }
      }
    } catch (err) {
      console.error("Failed to apply action", err);
    }
  };

  const triggerReset = async () => {
    // Reset points
    setAlert(null);
    setSelectedNodeId(null);
    setCurrentSeqID(events.length > 0 ? events[events.length - 1].seq_id : 1);
    try {
      // Direct post to reset
      await fetch(`${API_BASE}/game/action`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ action: "reset" })
      });
      fetchScore();
    } catch (e) {
      console.error(e);
    }
  };

  return (
    <div className="min-h-screen bg-cyber-bg text-cyber-text flex flex-col p-4 md:p-6 gap-6 selection:bg-cyber-accent/30 font-sans">
      {/* Header Panel */}
      <header className="flex flex-col md:flex-row md:items-center justify-between border-b border-cyber-border pb-4 gap-4">
        <div className="flex items-center gap-3">
          <div className="p-2.5 rounded-lg bg-cyber-accent/10 border border-cyber-accent/30 text-cyber-accent">
            <Shield size={24} className="animate-pulse" />
          </div>
          <div>
            <h1 className="text-xl font-bold tracking-tight text-white flex items-center gap-2">
              PHOENIXOS <span className="text-xs font-mono text-cyber-accent border border-cyber-accent/30 rounded px-1.5 py-0.5 bg-cyber-accent/5">SOC SIMULATOR v1.0</span>
            </h1>
            <p className="text-xs text-cyber-muted font-mono">Thermodynamic Cybernetic Security Training Console</p>
          </div>
        </div>

        {/* Stats Summary */}
        <div className="flex items-center gap-6">
          <div className="flex items-center gap-2.5 font-mono border border-slate-900 bg-slate-950/60 rounded-lg p-2 px-4">
            <Trophy className="text-cyber-warning" size={18} />
            <div>
              <div className="text-[10px] text-cyber-muted">SCORE</div>
              <div className="text-md font-bold text-white leading-none">{scoreState.score} pts</div>
            </div>
          </div>

          <div className="flex items-center gap-2.5 font-mono border border-slate-900 bg-slate-950/60 rounded-lg p-2 px-4">
            <Activity className="text-cyber-accent" size={18} />
            <div>
              <div className="text-[10px] text-cyber-muted">LEVEL</div>
              <div className="text-md font-bold text-cyber-accent leading-none">{scoreState.level}</div>
            </div>
          </div>

          <div className="flex items-center gap-2.5 font-mono border border-slate-900 bg-slate-950/60 rounded-lg p-2 px-4">
            <Award className="text-cyber-success" size={18} />
            <div>
              <div className="text-[10px] text-cyber-muted">MULTIPLIER</div>
              <div className="text-md font-bold text-cyber-success leading-none">{scoreState.multiplier.toFixed(1)}x</div>
            </div>
          </div>

          <button
            onClick={triggerReset}
            className="p-2.5 rounded-lg border border-slate-800 bg-slate-900 text-slate-400 hover:text-white hover:bg-slate-800 transition cursor-pointer"
            title="Reset Game"
          >
            <RefreshCw size={18} />
          </button>
        </div>
      </header>

      {/* Main Alert Notification */}
      {alert && (
        <div
          className={`flex items-center justify-between p-3.5 rounded-lg border text-sm font-mono ${
            alert.type === "success"
              ? "bg-cyber-success/5 border-cyber-success/30 text-cyber-success"
              : "bg-cyber-danger/5 border-cyber-danger/30 text-cyber-danger"
          }`}
        >
          <div className="flex items-center gap-2">
            <AlertCircle size={16} />
            <span>{alert.message}</span>
          </div>
          <button
            onClick={() => setAlert(null)}
            className="text-[10px] uppercase font-bold tracking-wider hover:underline"
          >
            Dismiss
          </button>
        </div>
      )}

      {/* Layout Columns */}
      <main className="flex-1 grid grid-cols-1 lg:grid-cols-12 gap-6 min-h-0">
        {/* Left Column (Timeline & Logs) - takes 8 columns */}
        <section className="lg:col-span-8 flex flex-col gap-6 min-h-0">
          {/* Timeline Controls */}
          <ReplayTimeline
            events={events}
            currentSeqID={currentSeqID}
            onChangeSeqID={setCurrentSeqID}
            wardenState={wardenState}
          />

          {/* Process-DAG Visualizer */}
          <SocDash
            graphData={graphData}
            onSelectNode={setSelectedNodeId}
            selectedNodeId={selectedNodeId}
          />
        </section>

        {/* Right Column (PhoenixMind AI Advisory & Badges) - takes 4 columns */}
        <section className="lg:col-span-4 flex flex-col gap-6">
          <PhoenixMind
            selectedNode={selectedNode}
            onApplyAction={handleApplyAction}
            wardenState={wardenState}
          />

          {/* Ecosystem Specification Panel */}
          <div className="border border-cyber-border rounded-lg bg-cyber-panel p-4 flex flex-col gap-3">
            <h3 className="text-xs font-bold text-cyber-text tracking-wider uppercase font-mono flex items-center gap-2">
              <Cpu size={16} className="text-cyber-accent" />
              PHOENIX MATRIX SPECIFICATION
            </h3>
            <div className="text-[11px] font-mono text-slate-400 flex flex-col gap-3">
              {ecosystemConfig?.physics_thresholds?.length > 0 && (
                <div>
                  <div className="text-cyber-accent font-bold mb-1">L6: Physical Stability States</div>
                  <div className="grid grid-cols-1 gap-1.5 bg-slate-950/40 p-2 rounded border border-slate-900">
                    {ecosystemConfig.physics_thresholds.map((t, idx) => (
                      <div key={idx} className="flex flex-col border-b border-slate-900/50 pb-1 last:border-b-0 last:pb-0">
                        <span className="font-bold text-white text-xs">{t.state} Threshold</span>
                        <span className="text-[10px] text-cyber-muted">{t.range} &mdash; {t.description} ({t.transition_indicator})</span>
                      </div>
                    ))}
                  </div>
                </div>
              )}

              {ecosystemConfig?.math_registry?.length > 0 && (
                <div>
                  <div className="text-cyber-accent font-bold mb-1">L3: Mathematics Substrate Domains</div>
                  <div className="flex flex-wrap gap-1.5 p-2 rounded bg-slate-950/40 border border-slate-900">
                    {ecosystemConfig.math_registry.map((m, idx) => (
                      <span
                        key={idx}
                        className="text-[9px] px-1.5 py-0.5 rounded bg-slate-900 border border-slate-800 text-slate-300"
                        title={m.subkeys.join(", ")}
                      >
                        {m.category}
                      </span>
                    ))}
                  </div>
                </div>
              )}
            </div>
          </div>

          {/* Badges Panel */}
          <div className="border border-cyber-border rounded-lg bg-cyber-panel p-4 flex flex-col gap-3">
            <h3 className="text-xs font-bold text-cyber-text tracking-wider uppercase font-mono flex items-center gap-2">
              <Award size={16} className="text-cyber-success" />
              EARNED BADGES & AWARDS
            </h3>
            <div className="flex flex-wrap gap-2 mt-1">
              {scoreState.badges && scoreState.badges.length > 0 ? (
                scoreState.badges.map((badge, idx) => (
                  <span
                    key={idx}
                    className="text-[10px] font-mono font-bold px-2 py-1 rounded bg-cyber-success/10 border border-cyber-success/30 text-cyber-success flex items-center gap-1"
                  >
                    ✦ {badge}
                  </span>
                ))
              ) : (
                <span className="text-xs text-cyber-muted font-mono italic">
                  No badges earned yet. Defend the system to win achievements.
                </span>
              )}
            </div>
          </div>

          {/* Instructions Panel */}
          <div className="border border-cyber-border rounded-lg bg-cyber-panel p-4 flex flex-col gap-3">
            <h3 className="text-xs font-bold text-cyber-text tracking-wider uppercase font-mono flex items-center gap-2">
              <FileText size={16} className="text-cyber-accent" />
              TRAINING MISSION BRIEFING
            </h3>
            <div className="text-xs text-slate-400 font-mono flex flex-col gap-2">
              <p>
                1. Scrub the slider to advance or rewind telemetry logical ticks.
              </p>
              <p>
                2. Click on orange or red nodes in the lineage graph to analyze them using the AI advisor.
              </p>
              <p>
                3. Click "Execute Recommended Countermeasure" to isolate compromised components and earn score points.
              </p>
              <p>
                4. Avoid isolating normal system binaries to prevent performance penalties.
              </p>
            </div>
          </div>
        </section>
      </main>
    </div>
  );
}
