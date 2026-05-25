import React, { useState, useEffect, useRef } from "react";
import { Play, Pause, RotateCcw, AlertTriangle, ShieldCheck } from "lucide-react";

export default function ReplayTimeline({
  events,
  currentSeqID,
  onChangeSeqID,
  wardenState
}) {
  const [isPlaying, setIsPlaying] = useState(false);
  const [playSpeed, setPlaySpeed] = useState(200); // ms per tick
  const timerRef = useRef(null);
  const logEndRef = useRef(null);

  // Compute maximum seq id
  const maxSeqID = events.length > 0 ? events[events.length - 1].seq_id : 1;
  const currentTick = currentSeqID;

  // Filter events up to selected sequence ID
  const activeEvents = events.filter((e) => e.seq_id <= currentSeqID);

  // Playback timer
  useEffect(() => {
    if (isPlaying) {
      timerRef.current = setInterval(() => {
        onChangeSeqID((prev) => {
          if (prev >= maxSeqID) {
            setIsPlaying(false);
            return prev;
          }
          return prev + 1;
        });
      }, playSpeed);
    } else {
      if (timerRef.current) {
        clearInterval(timerRef.current);
      }
    }
    return () => {
      if (timerRef.current) clearInterval(timerRef.current);
    };
  }, [isPlaying, playSpeed, maxSeqID]);

  // Scroll logs to bottom
  useEffect(() => {
    if (logEndRef.current) {
      logEndRef.current.scrollIntoView({ behavior: "smooth" });
    }
  }, [activeEvents.length]);

  return (
    <div className="border border-cyber-border rounded-lg bg-cyber-panel p-4 flex flex-col gap-4">
      {/* Playback Controls */}
      <div className="flex flex-col md:flex-row md:items-center justify-between gap-4">
        <div className="flex items-center gap-3">
          <button
            onClick={() => setIsPlaying(!isPlaying)}
            className={`p-2.5 rounded-lg border flex items-center justify-center transition ${
              isPlaying
                ? "bg-cyber-accent/10 border-cyber-accent text-cyber-accent hover:bg-cyber-accent/20"
                : "bg-slate-800 border-slate-700 text-slate-300 hover:bg-slate-700"
            }`}
          >
            {isPlaying ? <Pause size={18} /> : <Play size={18} />}
          </button>
          <button
            onClick={() => {
              setIsPlaying(false);
              onChangeSeqID(1);
            }}
            className="p-2.5 rounded-lg border bg-slate-800 border-slate-700 text-slate-300 hover:bg-slate-700 transition"
          >
            <RotateCcw size={18} />
          </button>
          <div className="font-mono text-xs">
            <div className="text-cyber-muted">TICK PROGRESS</div>
            <div className="text-cyber-text font-bold">
              {currentTick} / {maxSeqID}
            </div>
          </div>
        </div>

        {/* Timeline Slider */}
        <div className="flex-1 px-4 flex items-center">
          <input
            type="range"
            min="1"
            max={maxSeqID || 1}
            value={currentSeqID}
            onChange={(e) => {
              setIsPlaying(false);
              onChangeSeqID(parseInt(e.target.value, 10));
            }}
            className="w-full h-1.5 rounded-lg appearance-none cursor-pointer bg-slate-800 accent-cyber-accent"
          />
        </div>

        {/* Speed and State */}
        <div className="flex items-center gap-4">
          <div className="flex items-center gap-1.5">
            <span className="text-xs text-cyber-muted font-mono">SPEED:</span>
            <select
              value={playSpeed}
              onChange={(e) => setPlaySpeed(parseInt(e.target.value, 10))}
              className="bg-slate-800 border border-slate-700 text-cyber-text text-xs rounded-md p-1 px-2 font-mono outline-none focus:border-cyber-accent"
            >
              <option value="500">0.5s/tick</option>
              <option value="200">0.2s/tick</option>
              <option value="50">0.05s/tick</option>
            </select>
          </div>

          {/* Warden State Shield */}
          <div className="flex items-center gap-2 border border-cyber-border rounded-md px-3 py-1 bg-slate-950 font-mono">
            <span className="text-xs text-cyber-muted">WARDEN:</span>
            <span
              className={`text-xs font-bold flex items-center gap-1.5 ${
                wardenState === "NORMAL"
                  ? "text-cyber-success"
                  : wardenState === "SUSPICIOUS"
                  ? "text-cyber-warning"
                  : "text-cyber-danger"
              }`}
            >
              {wardenState === "NORMAL" ? (
                <ShieldCheck size={14} />
              ) : (
                <AlertTriangle size={14} />
              )}
              {wardenState}
            </span>
          </div>
        </div>
      </div>

      {/* Telemetry Event Logs Table */}
      <div className="flex-1 flex flex-col min-h-[220px] max-h-[300px]">
        <div className="flex items-center justify-between pb-2 border-b border-cyber-border">
          <span className="text-xs font-bold text-cyber-text font-mono">DETERMINISTIC TELEMETRY EVENT STREAM</span>
          <span className="text-[10px] text-cyber-muted font-mono">Showing {activeEvents.length} logs</span>
        </div>
        <div className="flex-1 overflow-y-auto mt-2 pr-1 border border-slate-900 rounded bg-slate-950/50">
          <table className="w-full text-left font-mono text-xs border-collapse">
            <thead>
              <tr className="text-cyber-muted border-b border-slate-900 bg-slate-950 sticky top-0">
                <th className="p-2 w-16">SEQ</th>
                <th className="p-2 w-28">EVENT TYPE</th>
                <th className="p-2 w-16">PID</th>
                <th className="p-2 w-16 text-center">ENTROPY</th>
                <th className="p-2 w-24">SOURCE</th>
                <th className="p-2">DETAILS</th>
              </tr>
            </thead>
            <tbody>
              {activeEvents.map((ev) => {
                const entropy = ev.payload?.entropy_score || 3.20;
                let textClass = "text-slate-400";
                if (entropy > 7.5) textClass = "text-cyber-danger font-bold";
                else if (entropy > 6.0) textClass = "text-cyber-warning";

                return (
                  <tr
                    key={ev.seq_id}
                    className="border-b border-slate-900/50 hover:bg-slate-900/30 transition"
                  >
                    <td className="p-2 text-cyber-accent font-bold">#{ev.seq_id}</td>
                    <td className="p-2 uppercase text-slate-300 font-semibold">{ev.event_type}</td>
                    <td className="p-2 text-slate-300">{ev.pid}</td>
                    <td className={`p-2 text-center ${textClass}`}>
                      {entropy.toFixed(2)}
                    </td>
                    <td className="p-2 text-slate-500 text-[10px]">{ev.category}</td>
                    <td className="p-2 text-slate-300 truncate max-w-xs" title={ev.exe_path}>
                      {ev.comm} → <span className="text-slate-500">{ev.exe_path}</span>
                    </td>
                  </tr>
                );
              })}
              <tr ref={logEndRef}></tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
}
