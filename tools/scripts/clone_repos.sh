#!/bin/bash
set -e

# Target directory
TARGET_DIR="/Users/fallofpheonix/os/parts"
mkdir -p "$TARGET_DIR"
cd "$TARGET_DIR"

# Repositories list: URL followed by folder name
clones=(
  "https://github.com/mit-pdos/xv6-riscv.git" "xv6-riscv"
  "https://github.com/redox-os/redox.git" "redox"
  "https://github.com/SerenityOS/serenity.git" "serenity"
  "https://github.com/cilium/cilium.git" "cilium"
  "https://github.com/cilium/tetragon.git" "tetragon"
  "https://github.com/nano-visor/nano-visor.git" "nano-visor"
  "https://github.com/imjoy-team/jan.git" "jan"
  "https://github.com/CoPaw/Copaw.git" "Copaw"
  "https://github.com/ironclaw/ironclaw.git" "ironclaw"
  "https://github.com/ggerganov/llama.cpp.git" "llama.cpp"
  "https://github.com/ollama/ollama.git" "ollama"
  "https://github.com/jacomyal/sigma.js.git" "sigma.js"
  "https://github.com/d3/d3.git" "d3"
  "https://github.com/visjs/vis-network.git" "vis-network"
  "https://github.com/grafana/grafana.git" "grafana"
  "https://github.com/madsim-rs/madsim.git" "madsim"
  "https://github.com/systems-group/anysystem.git" "anysystem"
  "https://github.com/openreplay/openreplay.git" "openreplay"
  "https://github.com/replayio/replay.git" "replay"
)

for ((i=0; i<${#clones[@]}; i+=2)); do
  url="${clones[i]}"
  dir="${clones[i+1]}"
  if [ ! -d "$dir" ]; then
    echo "Cloning $dir from $url..."
    # If the clone fails, log but continue so we don't break the entire run for one network timeout
    git clone --depth 1 "$url" "$dir" || echo "[WARN] Failed to clone $dir"
  else
    echo "$dir already exists, skipping."
  fi
done

echo "All clone operations completed."
