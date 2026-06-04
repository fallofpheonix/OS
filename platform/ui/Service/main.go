/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type SystemStatus struct {
	CPU          float64 `json:"cpu"`
	MemoryUsed   uint64  `json:"memory_used"`
	MemoryTotal  uint64  `json:"memory_total"`
	Uptime       float64 `json:"uptime"`
	Identity     string  `json:"identity"`
	Architecture string  `json:"architecture"`
	OS           string  `json:"os"`
}

type DomainStatus struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Load   string `json:"load"`
	Health int    `json:"health"`
}

type TerminalResponse struct {
	Output string `json:"output"`
}

var startTime = time.Now()

func main() {
	http.HandleFunc("/api/status", handleStatus)
	http.HandleFunc("/api/domains", handleDomains)
	http.HandleFunc("/api/terminal", handleTerminal)
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Add CORS support for local development
	handler := corsMiddleware(http.DefaultServeMux)

	port := "8080"
	fmt.Printf("PhoenixOS System Service starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	status := SystemStatus{
		CPU:          2.5, // Mock CPU usage
		MemoryUsed:   m.Alloc / 1024 / 1024,
		MemoryTotal:  16384,
		Uptime:       time.Since(startTime).Seconds(),
		Identity:     "PHOENIX-GENESIS-0",
		Architecture: runtime.GOARCH,
		OS:           runtime.GOOS,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

func handleDomains(w http.ResponseWriter, r *http.Request) {
	domains := []DomainStatus{
		{Name: "Nucleus", Status: "Optimal", Load: "12%", Health: 100},
		{Name: "Cognition", Status: "Healthy", Load: "45%", Health: 98},
		{Name: "Crucible", Status: "Idle", Load: "0%", Health: 100},
		{Name: "Terminus", Status: "Active", Load: "28%", Health: 95},
		{Name: "UI", Status: "Running", Load: "18%", Health: 100},
		{Name: "Arbiter", Status: "Optimal", Load: "5%", Health: 100},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(domains)
}

func handleTerminal(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload struct {
		Command string `json:"command"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	args := strings.Fields(payload.Command)
	if len(args) == 0 {
		json.NewEncoder(w).Encode(TerminalResponse{Output: ""})
		return
	}

	var output string
	cmd := args[0]

	// Strictly restricted command set for security
	switch cmd {
	case "ls":
		out, _ := exec.Command("ls", args[1:]...).Output()
		output = string(out)
	case "pwd":
		out, _ := exec.Command("pwd").Output()
		output = string(out)
	case "echo":
		output = strings.Join(args[1:], " ")
	case "date":
		output = time.Now().Format(time.RFC1123)
	case "whoami":
		output = "operator@phoenix"
	case "substrate":
		output = "PHOENIX-SUBSTRATE-L3-ACTIVE"
	default:
		output = fmt.Sprintf("Command restricted or not found: %s", cmd)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TerminalResponse{Output: output})
}
