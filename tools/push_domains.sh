#!/bin/bash

DOMAINS="Phoenix.Nucleus:https://github.com/fallofpheonix/Phoenix.Nucleus Phoenix.Cognition:https://github.com/fallofpheonix/Phoenix.Cognition Phoenix.Crucible:https://github.com/fallofpheonix/Phoenix.Crucible Phoenix.Terminus:https://github.com/fallofpheonix/Phoenix.Terminus"

for pair in $DOMAINS; do
  domain="${pair%%:*}"
  url="${pair##*:}"
  
  echo "-------------------------------------"
  echo "Pushing $domain to $url..."
  cd "$domain"
  
  # Remove previous failed attempts if necessary
  rm -rf .git || true
  
  git init
  git checkout -b main
  git remote add origin "$url"
  git add .
  git commit -m "chore: initial push of Sovereign Domain with architectural annotations"
  git push -u origin main --force
  
  cd ..
done
