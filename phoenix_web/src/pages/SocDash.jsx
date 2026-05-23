import React, { useEffect, useRef } from "react";
import { Network } from "vis-network/standalone";

export default function SocDash({ graphData, onSelectNode, selectedNodeId }) {
  const containerRef = useRef(null);
  const networkRef = useRef(null);

  useEffect(() => {
    if (!containerRef.current || !graphData) return;

    // Define custom styles for nodes matching the dark cybernetic system
    const data = {
      nodes: graphData.nodes || [],
      edges: graphData.edges || [],
    };

    const options = {
      nodes: {
        shape: "dot",
        size: 22,
        font: {
          color: "#E2E8F0",
          size: 12,
          face: "JetBrains Mono, monospace",
        },
        borderWidth: 2,
        shadow: {
          enabled: true,
          color: "rgba(0,0,0,0.5)",
          size: 10,
          x: 0,
          y: 4
        }
      },
      edges: {
        width: 1.5,
        color: {
          color: "#334155",
          highlight: "#06B6D4",
          hover: "#06B6D4"
        },
        arrows: {
          to: { enabled: true, scaleFactor: 0.8 }
        },
        smooth: {
          type: "cubicBezier",
          forceDirection: "none",
          roundness: 0.5
        }
      },
      groups: {
        normal: {
          color: { background: "#020617", border: "#10B981" } // Emerald green
        },
        suspicious: {
          color: { background: "#1E1502", border: "#F59E0B" } // Amber
        },
        critical: {
          color: { background: "#1A0505", border: "#EF4444" } // Red
        }
      },
      interaction: {
        hover: true,
        tooltipDelay: 100,
      },
      physics: {
        enabled: true,
        solver: "forceAtlas2Based",
        forceAtlas2Based: {
          gravitationalConstant: -50,
          centralGravity: 0.01,
          springConstant: 0.08,
          springLength: 80,
          damping: 0.4
        },
        stabilization: {
          iterations: 150,
          updateInterval: 25
        }
      }
    };

    // Instantiate vis network
    const network = new Network(containerRef.current, data, options);
    networkRef.current = network;

    // Click handler
    network.on("click", (params) => {
      if (params.nodes && params.nodes.length > 0) {
        onSelectNode(params.nodes[0]);
      } else {
        onSelectNode(null);
      }
    });

    return () => {
      if (networkRef.current) {
        networkRef.current.destroy();
        networkRef.current = null;
      }
    };
  }, [graphData]);

  // Handle programmatic selection highlight
  useEffect(() => {
    if (networkRef.current && selectedNodeId) {
      networkRef.current.selectNodes([selectedNodeId]);
    }
  }, [selectedNodeId]);

  return (
    <div className="flex-1 flex flex-col min-h-[400px] border border-cyber-border rounded-lg bg-cyber-panel p-4 relative overflow-hidden">
      <div className="flex justify-between items-center mb-2 z-10">
        <div>
          <h2 className="text-md font-bold text-cyber-text flex items-center gap-2">
            <span className="w-2 h-2 rounded-full bg-cyber-accent animate-pulse"></span>
            ACTIVE PROCESS LINEAGE (DAG)
          </h2>
          <p className="text-xs text-cyber-muted">Click nodes to investigate process execution trees</p>
        </div>
        <div className="flex items-center gap-4 text-xs">
          <span className="flex items-center gap-1.5"><span className="w-2.5 h-2.5 rounded-full border border-cyber-success bg-[#020617]"></span>Normal</span>
          <span className="flex items-center gap-1.5"><span className="w-2.5 h-2.5 rounded-full border border-cyber-warning bg-[#1E1502]"></span>Suspicious</span>
          <span className="flex items-center gap-1.5"><span className="w-2.5 h-2.5 rounded-full border border-cyber-danger bg-[#1A0505]"></span>Critical</span>
        </div>
      </div>
      <div ref={containerRef} className="flex-1 w-full h-full relative min-h-[350px]"></div>
    </div>
  );
}
