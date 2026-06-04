#!/bin/bash

REPOS="PhoenixCore:Phoenix.Nucleus PhoenixDistributed:Phoenix.Nucleus PhoenixFormal:Phoenix.Nucleus PhoenixGuard:Phoenix.Nucleus PhoenixKernel:Phoenix.Nucleus PhoenixTrace:Phoenix.Nucleus PhoenixTruth:Phoenix.Nucleus PhoenixValidation:Phoenix.Nucleus PhoenixMind:Phoenix.Cognition PhoenixMemoryLab:Phoenix.Cognition PhoenixVirtualizer:Phoenix.Crucible PhoenixChampions:Phoenix.Crucible PhoenixStimulation:Phoenix.Crucible ParticleStimulator:Phoenix.Crucible PhoenixRedteam:Phoenix.Crucible PhoenixDashboard:Phoenix.Terminus PhoenixExternal:Phoenix.Terminus PhoenixOS:Phoenix.Terminus PhoenixDocs:Phoenix.Terminus PhoenixResearch:Phoenix.Terminus"

for pair in $REPOS; do
  repo="${pair%%:*}"
  domain="${pair##*:}"
  
  if [ ! -d "$domain/$repo" ]; then
    echo "Cloning $repo into $domain..."
    git clone "https://github.com/fallofpheonix/$repo" "$domain/$repo" || true
  fi
done
