package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/v1/chat/completions", func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]interface{}{
			"choices": []map[string]interface{}{
				{
					"message": map[string]string{
						"content": `{"command": "ISOLATE_PID", "confidence_score": 0.95, "reasoning": "High entropy anomaly detected in execve lineage.", "graph_proof": ["evt-1", "evt-2"]}`,
					},
				},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	fmt.Println("Mock Oracle listening on :7860...")
	http.ListenAndServe(":7860", nil)
}
