package main

import (
	"math"
)

type StateVector []int8

func CalculateSDI(states StateVector) float64 {
	if len(states) == 0 { return 0 }
	counts := make(map[int8]int)
	for _, s := range states {
		counts[s]++
	}
	var sdi float64
	n := float64(len(states))
	for _, count := range counts {
		p := float64(count) / n
		if p > 0 {
			sdi -= p * math.Log(p)
		}
	}
	return sdi
}

func main() {
	// Boot check
	CalculateSDI(StateVector{1, -1})
}
