#!/bin/bash

# Define the pairs "Repo:Domain"
REPOS="PhoenixCore:Phoenix.Nucleus PhoenixDistributed:Phoenix.Nucleus PhoenixFormal:Phoenix.Nucleus PhoenixGuard:Phoenix.Nucleus PhoenixKernel:Phoenix.Nucleus PhoenixTrace:Phoenix.Nucleus PhoenixTruth:Phoenix.Nucleus PhoenixValidation:Phoenix.Nucleus PhoenixMind:Phoenix.Cognition PhoenixMemoryLab:Phoenix.Cognition PhoenixVirtualizer:Phoenix.Crucible PhoenixChampions:Phoenix.Crucible PhoenixStimulation:Phoenix.Crucible ParticleStimulator:Phoenix.Crucible PhoenixRedteam:Phoenix.Crucible PhoenixDashboard:Phoenix.Terminus PhoenixExternal:Phoenix.Terminus PhoenixOS:Phoenix.Terminus pheonixos:Phoenix.Terminus PhoenixDocs:Phoenix.Terminus PhoenixResearch:Phoenix.Terminus"

for pair in $REPOS; do
  repo="${pair%%:*}"
  domain="${pair##*:}"
  
  if [ -d "$domain/$repo" ]; then
    echo "Checking out main in $repo..."
    cd "$domain/$repo"
    git fetch origin main || true
    git checkout main || true
    git pull origin main || true
    cd ../..
  fi
done
