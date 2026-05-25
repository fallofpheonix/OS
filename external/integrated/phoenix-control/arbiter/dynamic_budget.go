package arbiter

func CalculateDynamicBudget(base int, threatLevel float64) int {
	return int(float64(base) * (1.0 - threatLevel))
}
