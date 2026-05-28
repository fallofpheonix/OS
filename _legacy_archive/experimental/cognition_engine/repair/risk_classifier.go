package repair

func ClassifyRisk(impact int) RiskLevel {
	if impact > 80 {
		return Dangerous
	}
	return Low
}
