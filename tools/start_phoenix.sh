#!/bin/bash

# PhoenixOS Working Model Start Script

# Colors for logging
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}Starting PhoenixOS Substrate Working Model...${NC}"

# 1. Start System Service (Go)
echo -e "${GREEN}[1/3] Starting System Service (Go)...${NC}"
cd Phoenix.UI/Service
go run main.go &
SERVICE_PID=$!
cd ../..

# 2. Wait for Service to be ready
echo -e "${BLUE}Waiting for System Service to initialize...${NC}"
until curl -s http://localhost:8080/api/health > /dev/null; do
  sleep 1
done

# 3. Start UI Shell (React)
echo -e "${GREEN}[2/3] Starting UI Shell (React)...${NC}"
cd Phoenix.UI/Shell
npm run dev &
SHELL_PID=$!
cd ../..

echo -e "${BLUE}[3/3] PhoenixOS is now active.${NC}"
echo -e "${GREEN}UI Shell: http://localhost:5173${NC}"
echo -e "${GREEN}System API: http://localhost:8080/api${NC}"
echo -e "${BLUE}Press Ctrl+C to shutdown all components.${NC}"

# Handle shutdown
trap "kill $SERVICE_PID $SHELL_PID; echo -e '\n${BLUE}PhoenixOS Shutdown Complete.${NC}'; exit" INT

wait
