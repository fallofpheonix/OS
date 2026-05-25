package state

// RuntimeState defines the authoritative states for the PhoenixOS runtime.
type RuntimeState string

const (
	Safe     RuntimeState = "SAFE"
	Watch    RuntimeState = "WATCH"
	Alert    RuntimeState = "ALERT"
	Contain  RuntimeState = "CONTAIN"
	Recovery RuntimeState = "RECOVERY"
)

// Legacy aliases for backward compatibility.
// These are only allowed to exist here in the state package.
const (
	StateNormal     RuntimeState = Safe
	StateSuspicious RuntimeState = Watch
	StateContained  RuntimeState = Contain
)

// ValidTransitions defines the allowed state movements to ensure FSM integrity.
var ValidTransitions = map[RuntimeState][]RuntimeState{
	Safe:     {Watch, Alert, Contain},
	Watch:    {Safe, Alert, Contain},
	Alert:    {Watch, Contain, Recovery},
	Contain:  {Recovery},
	Recovery: {Safe},
}
