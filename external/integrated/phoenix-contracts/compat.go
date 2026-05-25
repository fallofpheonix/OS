package contracts

// GlobalCompatibilityMatrix defines the system-wide support levels.
var GlobalCompatibilityMatrix = CompatibilityMatrix{
	SupportedLevels: []int{1},
	MinVersion: Version{
		Major: 0,
		Minor: 1,
		Patch: 0,
	},
}

// IsCompatible checks if a version/level is supported by the current runtime.
func IsCompatible(apiLevel int, v Version) bool {
	for _, l := range GlobalCompatibilityMatrix.SupportedLevels {
		if l == apiLevel {
			return true
		}
	}
	return false
}
