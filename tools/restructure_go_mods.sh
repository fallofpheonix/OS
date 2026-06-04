#!/bin/bash

# Define the pairs "Repo:Domain"
REPOS="PhoenixCore:Phoenix.Nucleus PhoenixDistributed:Phoenix.Nucleus PhoenixFormal:Phoenix.Nucleus PhoenixGuard:Phoenix.Nucleus PhoenixKernel:Phoenix.Nucleus PhoenixTrace:Phoenix.Nucleus PhoenixTruth:Phoenix.Nucleus PhoenixValidation:Phoenix.Nucleus PhoenixMind:Phoenix.Cognition PhoenixMemoryLab:Phoenix.Cognition PhoenixVirtualizer:Phoenix.Crucible PhoenixChampions:Phoenix.Crucible PhoenixStimulation:Phoenix.Crucible ParticleStimulator:Phoenix.Crucible PhoenixRedteam:Phoenix.Crucible PhoenixDashboard:Phoenix.Terminus PhoenixExternal:Phoenix.Terminus PhoenixOS:Phoenix.Terminus pheonixos:Phoenix.Terminus PhoenixDocs:Phoenix.Terminus PhoenixResearch:Phoenix.Terminus"

# Find all go.mod files
find Phoenix.* -name go.mod | while read modfile; do
  
  # For each repo, update its replace directives
  for pair in $REPOS; do
    repo=$(echo $pair | cut -d: -f1)
    domain=$(echo $pair | cut -d: -f2)
    
    # Replace old relative paths
    sed -i '' "s|=> \.\./$repo|=> ../../$domain/$repo|g" "$modfile"
  done
done

# Create global go.work
echo "go 1.26" > go.work
echo "" >> go.work
echo "use (" >> go.work
find Phoenix.* -name go.mod -exec dirname {} \; | sed 's/^/	.\//' | sort >> go.work
echo ")" >> go.work
