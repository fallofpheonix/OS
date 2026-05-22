package monitor

// RefineDriftDetection updates the monitor to use a dynamic gain adjustment for the Kalman filter
// based on current system entropy.
func (m *MonitorService) RefineDriftDetection(currentEntropy float64) {
	// Dynamic adjustment of the Kalman gain based on system entropy
	// Higher entropy = more uncertainty in telemetry = lower gain
	newGain := 0.1 / (1.0 + currentEntropy)
	m.kalman.SetGain(newGain)
}
