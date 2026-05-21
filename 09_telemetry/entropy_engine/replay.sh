#!/bin/bash
echo "Running Replay..."
go run src/main.go --replay replay/sample.json > artifacts/replay.json
