package control

import (
	"fmt"
	"sync"
	"time"

	"phoenix/agents/internal/types"
)

type ControlAgent interface {
	EnforceStrategy(s types.Strategy, currentThreatTemp float64) error
	GetPIDMetrics() types.PIDMetrics
	GetActionHistory() []string
}

type Agent struct {
	mu            sync.Mutex
	kp            float64
	ki            float64
	kd            float64
	integral      float64
	lastError     float64
	setpoint      float64 // Target maximum threat temperature (e.g., 2.0)
	lastUpdateTime time.Time
	metrics       types.PIDMetrics
	actionHistory []string
}

func NewControlAgent(kp, ki, kd, setpoint float64) *Agent {
	return &Agent{
		kp:             kp,
		ki:             ki,
		kd:             kd,
		setpoint:       setpoint,
		lastUpdateTime: time.Now(),
		metrics: types.PIDMetrics{
			Setpoint: setpoint,
		},
		actionHistory: make([]string, 0),
	}
}

// EnforceStrategy runs a PID step using the currentThreatTemp compared against the setpoint,
// and then executes containment actions for the specified target PIDs in the strategy.
func (a *Agent) EnforceStrategy(s types.Strategy, currentThreatTemp float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	now := time.Now()
	dt := now.Sub(a.lastUpdateTime).Seconds()
	if dt <= 0 {
		dt = 0.001 // prevent division by zero or negative time step
	}
	a.lastUpdateTime = now

	// Calculate error: current - setpoint
	errVal := currentThreatTemp - a.setpoint

	// Calculate integral with anti-windup clamping (limit to -10 to 10)
	a.integral += errVal * dt
	if a.integral > 10.0 {
		a.integral = 10.0
	} else if a.integral < -10.0 {
		a.integral = -10.0
	}

	// Calculate derivative
	derivative := (errVal - a.lastError) / dt
	a.lastError = errVal

	// Calculate PID Output
	output := (a.kp * errVal) + (a.ki * a.integral) + (a.kd * derivative)

	// Keep track of metrics
	a.metrics = types.PIDMetrics{
		Setpoint:  a.setpoint,
		Measured:  currentThreatTemp,
		Output:    output,
		LastError: errVal,
		Integral:  a.integral,
	}

	// Action selection pipeline based on Strategy Level (0-5) and PID output
	// Level 0: Observe
	// Level 1-2: Limit (throttling)
	// Level 3: Freeze (suspend)
	// Level 4: Isolate (sandbox/network block)
	// Level 5: Kill (SIGKILL)
	level := s.ContainmentLevel

	// PID output can escalate containment level if output is very high
	if output > 5.0 && level < 5 {
		level = 5
	} else if output > 2.0 && level < 3 {
		level = 3
	}

	actionTime := now.Format(time.RFC3339)
	for _, pid := range s.TargetPIDs {
		switch level {
		case 0:
			msg := fmt.Sprintf("[%s] ACTION: OBSERVE PID %d (PID output: %.2f)", actionTime, pid, output)
			a.actionHistory = append(a.actionHistory, msg)
		case 1, 2:
			msg := fmt.Sprintf("[%s] ACTION: LIMIT CPU/IO for PID %d by %.1f%% (PID output: %.2f)", actionTime, pid, mathMin(100.0, mathMax(10.0, output*20.0)), output)
			a.actionHistory = append(a.actionHistory, msg)
		case 3:
			msg := fmt.Sprintf("[%s] ACTION: FREEZE PID %d (SIGSTOP sent) (PID output: %.2f)", actionTime, pid, output)
			a.actionHistory = append(a.actionHistory, msg)
		case 4:
			msg := fmt.Sprintf("[%s] ACTION: ISOLATE PID %d (Network isolated) (PID output: %.2f)", actionTime, pid, output)
			a.actionHistory = append(a.actionHistory, msg)
		case 5:
			msg := fmt.Sprintf("[%s] ACTION: KILL PID %d (SIGKILL sent) (PID output: %.2f)", actionTime, pid, output)
			a.actionHistory = append(a.actionHistory, msg)
		}
	}

	return nil
}

func (a *Agent) GetPIDMetrics() types.PIDMetrics {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.metrics
}

func (a *Agent) GetActionHistory() []string {
	a.mu.Lock()
	defer a.mu.Unlock()
	// Return a copy
	res := make([]string, len(a.actionHistory))
	copy(res, a.actionHistory)
	return res
}

// Helpers
func mathMax(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

func mathMin(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}
