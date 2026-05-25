package arbiter

import (
	"github.com/fallofpheonix/phoenix-control/warden"
)

// Base costs from cost_matrix.md
const (
	ACCritical = 1000.0
	ACHigh     = 500.0
	ACMedium   = 100.0
	ACLow      = 10.0
	ACInfo     = 1.0

	CCClassObserve         = 0.0
	CCClassLog             = 1.0
	CCClassThrottle        = 50.0
	CCClassLocalIsolate    = 200.0
	CCClassClusterIsolate  = 800.0
	CCClassKernelEmergency = 2000.0
)

// CalculateAttackCost determines the potential damage cost of the threat.
// Formula: AC_final = BaseCost * (1 + SystemLoad)
func CalculateAttackCost(ctx SystemContext) float64 {
	base := ACInfo
	z := ctx.ThreatScore.ZScore

	// Deterministic mapping of Z-Score to Severity Level
	switch {
	case z >= 10.0:
		base = ACCritical
	case z >= 5.0:
		base = ACHigh
	case z >= 3.0:
		base = ACMedium
	case z >= 1.5:
		base = ACLow
	}

	return base * (1.0 + ctx.SystemLoad)
}

// GetContainmentCost returns the fixed cost for a given actuation class as defined in the cost matrix.
func GetContainmentCost(class warden.ActuationClass) float64 {
	switch class {
	case warden.ClassObserve:
		return CCClassObserve
	case warden.ClassLog:
		return CCClassLog
	case warden.ClassThrottle:
		return CCClassThrottle
	case warden.ClassLocalIsolate:
		return CCClassLocalIsolate
	case warden.ClassClusterIsolate:
		return CCClassClusterIsolate
	case warden.ClassKernelEmergency:
		return CCClassKernelEmergency
	default:
		return 9999.0 // Extremely high cost for unknown/invalid actions
	}
}
