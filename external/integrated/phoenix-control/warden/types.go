package warden

import "github.com/fallofpheonix/phoenix-contracts"

// ActuationClass defines the severity level of a containment action.
type ActuationClass = contracts.ActuationClass

const (
	ClassObserve         = contracts.ClassObserve
	ClassLog             = contracts.ClassLog
	ClassThrottle        = contracts.ClassThrottle
	ClassLocalIsolate    = contracts.ClassLocalIsolate
	ClassClusterIsolate  = contracts.ClassClusterIsolate
	ClassKernelEmergency = contracts.ClassKernelEmergency
)
