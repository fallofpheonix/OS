package drift

// ComputeDrift calculates the mean drift across modules and compares to baseline.
func ComputeDrift(history []DriftRecord, baseline Baseline) float64 {
	if len(history) == 0 {
		return 0.0
	}

	totalDrift := 0.0
	for _, rec := range history {
		baseDrift := baseline.Modules[rec.Module]
		totalDrift += (rec.Drift - baseDrift)
	}
	meanDrift := totalDrift / float64(len(history))
	return meanDrift
}
