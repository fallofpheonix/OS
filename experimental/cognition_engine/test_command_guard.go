package main

import (
	"experimental/cognition_engine/sandbox"
	"fmt"
)

func main() {
	guard := &sandbox.CommandGuard{}
	decision := guard.Allow("kernel_access")
	if decision != "REJECT" {
		panic(fmt.Sprintf("Expected REJECT, got %s", decision))
	}
	fmt.Println("Command guard tests passed.")
}
