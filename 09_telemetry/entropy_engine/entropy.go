package entropy_engine

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"os"
)

// Result represents the calculation output
type Result struct {
	Entropy      float64 `json:"entropy"`
	KLDivergence float64 `json:"kl_divergence"`
	IsAnomaly    bool    `json:"is_anomaly"`
}

// Event represents a telemetry event for replay
type Event struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

// Calculate Shannon Entropy and KL Divergence
func Calculate(data []byte, baseline []float64) Result {
	if len(data) == 0 {
		return Result{}
	}

	counts := make([]int, 256)
	for _, b := range data {
		counts[b]++
	}

	var entropy float64
	var klDiv float64
	n := float64(len(data))

	for i := 0; i < 256; i++ {
		p := float64(counts[i]) / n
		if p > 0 {
			entropy -= p * math.Log2(p)
			
			if baseline != nil && i < len(baseline) {
				q := baseline[i]
				if q > 0 {
					klDiv += p * math.Log2(p/q)
				} else {
					klDiv += p * 8.0 
				}
			}
		}
	}

	isAnomaly := entropy > 7.5 || klDiv > 4.0

	return Result{
		Entropy:      entropy,
		KLDivergence: klDiv,
		IsAnomaly:    isAnomaly,
	}
}

func main() {
	sim := flag.Bool("sim", false, "Run simulation")
	replay := flag.String("replay", "", "Replay file")
	flag.Parse()

	if *sim {
		fmt.Println("Simulating Entropy calculation on random vs structured data...")
		structured := []byte("This is a structured text file with some patterns.")
		random := make([]byte, 1024)
		for i := range random {
			random[i] = byte(i % 256) 
		}

		res1 := Calculate(structured, nil)
		res2 := Calculate(random, nil)

		fmt.Printf("Structured: %+v\n", res1)
		fmt.Printf("Random:     %+v\n", res2)
		return
	}

	if *replay != "" {
		f, err := os.Open(*replay)
		if err != nil {
			fmt.Printf("Error opening replay file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()

		var events []Event
		if err := json.NewDecoder(f).Decode(&events); err != nil {
			fmt.Printf("Error decoding JSON: %v\n", err)
			os.Exit(1)
		}

		var results []Result
		for _, e := range events {
			data, _ := base64.StdEncoding.DecodeString(e.Data)
			results = append(results, Calculate(data, nil))
		}

		out, _ := json.MarshalIndent(results, "", "  ")
		fmt.Println(string(out))
		return
	}
}
