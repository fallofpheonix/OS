package repair

type RepairState string

const (
    Success   RepairState = "SUCCESS"
    Failed    RepairState = "FAILED"
    Regressed RepairState = "REGRESSED"
    Blocked   RepairState = "BLOCKED"
    Unknown   RepairState = "UNKNOWN"
)

type RiskLevel string

const (
    Low       RiskLevel = "LOW"
    Medium    RiskLevel = "MEDIUM"
    High      RiskLevel = "HIGH"
    Dangerous RiskLevel = "DANGEROUS"
)
