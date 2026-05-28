package sandbox

import "strings"

type PathFilter struct{}

func (f *PathFilter) IsSafe(path string) bool {
	// Block kernel and warden paths
	if strings.Contains(path, "kernel") || strings.Contains(path, "warden") {
		return false
	}
	return true
}
