package repair

type RiskLevel string

const (
	LOW       RiskLevel = "LOW"
	MEDIUM    RiskLevel = "MEDIUM"
	HIGH      RiskLevel = "HIGH"
	DANGEROUS RiskLevel = "DANGEROUS"
)

func ClassifyRisk(impact int) RiskLevel {
	if impact > 80 {
		return DANGEROUS
	}
	return LOW
}
