package v1

import "fmt"

// EventVersion represents the semantic version of an event schema.
type EventVersion struct {
	Major uint32
	Minor uint32
	Patch uint32
}

// String returns the string representation of the event version.
func (v EventVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}
