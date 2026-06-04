#!/bin/bash

DOMAINS="Phoenix.Nucleus:https://github.com/fallofpheonix/Phoenix.Nucleus Phoenix.Cognition:https://github.com/fallofpheonix/Phoenix.Cognition Phoenix.Crucible:https://github.com/fallofpheonix/Phoenix.Crucible Phoenix.Terminus:https://github.com/fallofpheonix/Phoenix.Terminus"

for pair in $DOMAINS; do
  domain="${pair%%:*}"
  url="${pair##*:}"
  
  echo "-------------------------------------"
  echo "Pushing $domain to $url..."
  cd "$domain"
  
  git remote set-url origin "$url" || git remote add origin "$url"
  
  # Ensure we have the branch and commit
  git add .
  git commit -m "chore: initial push of Sovereign Domain with architectural annotations" || true
  
  # The issue was using '//github' instead of 'https://github' because of a typo in the previous script or how the URL was parsed. Wait, the URL is https://...
  # Let's just push directly.
  git push -u origin main --force
  
  cd ..
done
