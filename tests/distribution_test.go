package tests

import (
	"os"
	"testing"
)

func TestDistributionStructure(t *testing.T) {
	requiredDirs := []string{
		"arch_base",
		"branding",
		"build_system",
		"iso",
		"kali_base",
		"lfs",
		"packages",
	}

	for _, dir := range requiredDirs {
		path := "../11_distribution/" + dir
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("Required distribution directory %s not found", path)
		}
	}
}
