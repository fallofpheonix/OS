package state
type SystemState string
const (
    StateSafe     SystemState = "SAFE"
    StateWatch    SystemState = "WATCH"
    StateAlert    SystemState = "ALERT"
    StateContain  SystemState = "CONTAIN"
    StateRecovery SystemState = "RECOVERY"
)
