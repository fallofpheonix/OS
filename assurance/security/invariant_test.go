/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 */
package security

import (
	"fmt"
	"testing"
)

func TestContextualInvariant_NamespaceBreakout(t *testing.T) {
	// 1. Mock the GraphProvider
	mockProvider := &MockGraphProvider{}

	// 2. Initialize the hardened Invariant
	invariant := &ContextualInvariant{Provider: mockProvider}

	// 3. Construct a malicious escalation request (namespace breakout)
	req := AuthorityEscalationRequest{
		EventID:       "test-event-001",
		TargetNsproxy: 9999, // Malicious target namespace
		TargetTgid:    1,
		GraphProof: &GraphProof{
			Path:            []string{"node1", "node2"},
			ExpectedNsproxy: 1234, // Legitimate expected namespace
		},
	}

	snap := PostureSnapshot{State: StateSafe}

	// 4. Verify that the invariant vetoes the request
	err := invariant.Verify(req, snap)
	if err == nil {
		t.Fatal("Expected invariant violation due to namespace mismatch, but request was authorized!")
	}
	fmt.Printf("Successfully caught breakout attempt: %v\n", err)
}

// MockGraphProvider implements the GraphProvider interface
type MockGraphProvider struct{}

func (m *MockGraphProvider) VerifyPath(path []string) (bool, error) {
	return true, nil // Path is valid, but the namespace will be checked by ContextualInvariant
}
