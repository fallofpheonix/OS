package ai

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type PredictiveAdvisor struct {
	Threshold float64
	VMax      float64
	LogFile   *os.File
}

func NewPredictiveAdvisor(threshold, vMax float64) (*PredictiveAdvisor, error) {
	f, err := os.OpenFile("/tmp/predictive_advisor.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &PredictiveAdvisor{Threshold: threshold, VMax: vMax, LogFile: f}, nil
}

func (pa *PredictiveAdvisor) CalculateWeight(v float64) float64 {
	if v <= pa.Threshold {
		return 1.0
	}
	weight := 1.0 - (v-pa.Threshold)/(pa.VMax-pa.Threshold)
	if weight < 0.1 {
		return 0.1
	}
	return weight
}

func (pa *PredictiveAdvisor) LogPrediction(eventID string, drift float64, weight float64) {
	entry := map[string]interface{}{
		"timestamp": time.Now().Format(time.RFC3339),
		"event_id":  eventID,
		"drift":     drift,
		"weight":    weight,
		"action":    "PREDICTED_THROTTLE",
	}
	data, _ := json.Marshal(entry)
	_, err := pa.LogFile.Write(append(data, '\n'))
	if err != nil {
		fmt.Printf("[DEBUG] Failed to write to predictive log: %v\n", err)
	}
	pa.LogFile.Sync() // Ensure it hits the disk
}
