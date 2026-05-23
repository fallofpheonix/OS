package game

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type PhysicsThreshold struct {
	State               string `json:"state"`
	Description         string `json:"description"`
	TransitionIndicator string `json:"transition_indicator"`
	Range               string `json:"range"`
}

type MathRegistryItem struct {
	Category string   `json:"category"`
	Subkeys  []string `json:"subkeys"`
}

type EcosystemConfig struct {
	PhysicsThresholds []PhysicsThreshold `json:"physics_thresholds"`
	MathRegistry      []MathRegistryItem `json:"math_registry"`
}

func LoadEcosystemConfig() (*EcosystemConfig, error) {
	cfg := &EcosystemConfig{
		PhysicsThresholds: []PhysicsThreshold{},
		MathRegistry:      []MathRegistryItem{},
	}

	// 1. Parse physics_thresholds.md
	physFile, err := os.Open("/Users/fallofpheonix/os/parts/engineering/physics-runtime/physics_thresholds.md")
	if err == nil {
		defer physFile.Close()
		scanner := bufio.NewScanner(physFile)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "|") {
				parts := strings.Split(line, "|")
				if len(parts) >= 4 {
					state := strings.TrimSpace(parts[1])
					// remove markdown formatting like **
					state = strings.ReplaceAll(state, "**", "")
					if state == "State" || strings.HasPrefix(state, "---") {
						continue
					}
					desc := strings.TrimSpace(parts[2])
					indicator := strings.TrimSpace(parts[3])

					rng := ""
					switch state {
					case "Stable":
						rng = "S_I < 0.40"
					case "Warning":
						rng = "0.40 <= S_I < 0.70"
					case "Critical":
						rng = "0.70 <= S_I < 0.85"
					case "Collapse":
						rng = "S_I >= 0.85"
					}

					cfg.PhysicsThresholds = append(cfg.PhysicsThresholds, PhysicsThreshold{
						State:               state,
						Description:         desc,
						TransitionIndicator: indicator,
						Range:               rng,
					})
				}
			}
		}
	}

	// 2. Parse math_registry.yaml
	mathFile, err := os.Open("/Users/fallofpheonix/os/parts/engineering/mathematics-engine/math_registry.yaml")
	if err == nil {
		defer mathFile.Close()
		scanner := bufio.NewScanner(mathFile)
		var currentCat string
		var currentSubkeys []string
		for scanner.Scan() {
			line := scanner.Text()
			trimmed := strings.TrimSpace(line)
			if trimmed == "" || strings.HasPrefix(trimmed, "#") || trimmed == "registry:" {
				continue
			}
			if strings.HasSuffix(trimmed, ":") {
				if currentCat != "" {
					cfg.MathRegistry = append(cfg.MathRegistry, MathRegistryItem{
						Category: currentCat,
						Subkeys:  currentSubkeys,
					})
				}
				currentCat = strings.TrimSuffix(trimmed, ":")
				currentSubkeys = []string{}
			} else if strings.Contains(trimmed, ":") {
				parts := strings.Split(trimmed, ":")
				subkey := strings.TrimSpace(parts[0])
				currentSubkeys = append(currentSubkeys, subkey)
			}
		}
		if currentCat != "" {
			cfg.MathRegistry = append(cfg.MathRegistry, MathRegistryItem{
				Category: currentCat,
				Subkeys:  currentSubkeys,
			})
		}
	}

	return cfg, nil
}

func (gs *GameServer) handleConfig(w http.ResponseWriter, r *http.Request) {
	cfg, err := LoadEcosystemConfig()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(cfg)
}
