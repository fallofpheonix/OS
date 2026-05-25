package main

import (
	"experimental/cognition_engine/repair"
	"fmt"
)

func main() {
	eval := &repair.Evaluator{}
	decision := eval.Evaluate(repair.Dangerous)
	if decision != "HUMAN_ONLY" {
		panic(fmt.Sprintf("Expected HUMAN_ONLY, got %s", decision))
	}
	fmt.Println("Repair evaluator tests passed.")
}
