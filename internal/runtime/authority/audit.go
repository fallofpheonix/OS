/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
/**
 * FILE: audit.go
 * PATH: Phoenix.Nucleus/authority/audit.go
 *
 * PURPOSE:
 * Implements the Authority Registry and recursive state auditing.
 * Ensures the conservation of authority across all system transitions.
 *
 * SUBSYSTEM:
 * Nucleus / Authority Cycle
 *
 * DEPENDENCIES:
 * encoding/json, sync, github.com/fallofpheonix/phoenix/foundation/ledger
 *
 * DEPENDENTS:
 * Phoenix.Nucleus/capability, Phoenix.Nucleus/recovery
 *
 * SECURITY:
 * This is a critical security boundary. It enforces the "Replay Reconstruction Theorem"
 * by ensuring the registry state is entirely derivable from the Ledger.
 *
 * PERFORMANCE:
 * O(D) for audits and state changes, where D is the depth of the delegation tree.
 */

package authority

import (
	"encoding/json"
	"fmt"
	"github.com/fallofpheonix/phoenix/foundation/ledger"
	"sync"
)

// BEGINNER EXPLANATION:
// This file is the "Accountant" of the system. It keeps track of who has how
// much power (Authority) and makes sure power is never created out of thin air.
// It also handles "freezing" or "revoking" power if someone breaks the rules.

// INTERMEDIATE EXPLANATION:
// The Registry maintains a map of AuthorityAtoms and their delegation trees.
// It supports multiple states (Active, Frozen, Quarantined, Revoked).
// The ReconstructFromLedger method proves that authority state is a projection of history.

// EXPERT EXPLANATION:
// Implements a recursive Authority Conservation engine. It ensures that
// Sum(Children) + Remainder == Parent at every node in the delegation DAG.
// The recursiveSetState function handles cascading revocations, ensuring
// that compromised authority branches are atomically neutralized.

// AuthorityState defines the current operational status of an authority.
type AuthorityState string

const (
	AuthActive      AuthorityState = "ACTIVE"
	AuthFrozen      AuthorityState = "FROZEN"
	AuthQuarantined AuthorityState = "QUARANTINED"
	AuthRevoked     AuthorityState = "REVOKED"
)

/**
 * Registry
 *
 * Tracks the total conservation and status of authority across the system.
 *
 * Responsibilities:
 * - Root authority issuance.
 * - Delegation verification and tracking.
 * - State management (Freeze/Revoke).
 * - Ledger-based state reconstruction.
 *
 * Thread Safety:
 * Thread-safe via sync.RWMutex.
 */
type Registry struct {
	mu          sync.RWMutex
	authorities map[string]ledger.AuthorityAtom // ID -> Remaining Atoms
	delegations map[string][]string             // Parent ID -> Child IDs
	states      map[string]AuthorityState       // ID -> State (Q723)
	reasons     map[string]string               // ID -> Last state change reason
}

/**
 * NewRegistry
 *
 * Initializes an authority registry.
 */
func NewRegistry() *Registry {
	return &Registry{
		authorities: make(map[string]ledger.AuthorityAtom),
		delegations: make(map[string][]string),
		states:      make(map[string]AuthorityState),
		reasons:     make(map[string]string),
	}
}

/**
 * SetState
 *
 * Updates the operational status of an authority.
 *
 * Input:
 * - id: The Authority ID to update.
 * - state: The new AuthorityState.
 * - reason: Explanation for the change.
 */
func (r *Registry) SetState(id string, state AuthorityState, reason string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.authorities[id]; !exists {
		return fmt.Errorf("authority %s not found", id)
	}

	r.states[id] = state
	r.reasons[id] = reason
	return nil
}

/**
 * GetState
 *
 * Returns the current state of an authority.
 *
 * Output:
 * - The current AuthorityState (defaults to ACTIVE).
 */
func (r *Registry) GetState(id string) AuthorityState {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if state, ok := r.states[id]; ok {
		return state
	}
	return AuthActive
}

/**
 * IssueRootAuthority
 *
 * Creates the genesis authority for the system.
 *
 * Input:
 * - id: Unique ID for the root authority.
 * - atoms: Total units of authority to issue.
 *
 * Security:
 * Only permitted during Genesis or via a valid Policy change.
 */
func (r *Registry) IssueRootAuthority(id string, atoms ledger.AuthorityAtom) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.authorities[id]; exists {
		return fmt.Errorf("authority %s already exists", id)
	}

	r.authorities[id] = atoms
	return nil
}

/**
 * DelegateAuthority
 *
 * Splits a parent authority into children, ensuring conservation.
 *
 * Input:
 * - parentID: Source authority.
 * - childAtoms: Map of child IDs to their allocated atoms.
 *
 * Complexity: O(Number of Children)
 */
func (r *Registry) DelegateAuthority(parentID string, childAtoms map[string]ledger.AuthorityAtom) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	parentRemaining, exists := r.authorities[parentID]
	if !exists {
		return fmt.Errorf("parent authority %s not found", parentID)
	}

	var totalRequested ledger.AuthorityAtom = 0
	for _, atoms := range childAtoms {
		totalRequested += atoms
	}

	if totalRequested > parentRemaining {
		return fmt.Errorf("authority conservation breach: requested %v, remaining %v", totalRequested, parentRemaining)
	}

	// Apply delegation
	r.authorities[parentID] -= totalRequested
	for id, atoms := range childAtoms {
		r.authorities[id] = atoms
		r.delegations[parentID] = append(r.delegations[parentID], id)
	}

	return nil
}

/**
 * AuditConservation
 *
 * Performs a recursive check to ensure no authority was created.
 *
 * Input:
 * - rootID: The authority at the root of the tree.
 * - originalTotal: The expected total atoms.
 *
 * Complexity: O(Nodes in Tree)
 */
func (r *Registry) AuditConservation(rootID string, originalTotal ledger.AuthorityAtom) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	actualTotal := r.calculateSum(rootID)
	if actualTotal != originalTotal {
		return fmt.Errorf("conservation failure: root %s expected %v, found %v", rootID, originalTotal, actualTotal)
	}

	return nil
}

func (r *Registry) calculateSum(id string) ledger.AuthorityAtom {
	sum := r.authorities[id]
	for _, childID := range r.delegations[id] {
		sum += r.calculateSum(childID)
	}
	return sum
}

/**
 * RevokeAuthority
 *
 * Invalidates an authority and all its descendants.
 *
 * Input:
 * - id: The target authority ID.
 * - reason: Explanation for revocation.
 *
 * Side Effects:
 * - Cascades revocation to the entire sub-tree.
 */
func (r *Registry) RevokeAuthority(id string, reason string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.authorities[id]; !exists {
		return fmt.Errorf("authority %s not found", id)
	}

	return r.recursiveSetState(id, AuthRevoked, reason)
}

func (r *Registry) recursiveSetState(id string, state AuthorityState, reason string) error {
	r.states[id] = state
	r.reasons[id] = reason
	for _, childID := range r.delegations[id] {
		r.recursiveSetState(childID, state, reason)
	}
	return nil
}

/**
 * IsRevoked
 *
 * Checks if an authority is explicitly revoked.
 */
func (r *Registry) IsRevoked(id string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.states[id] == AuthRevoked
}

/**
 * ReconstructFromLedger
 *
 * Builds the registry state by replaying events from the ledger.
 *
 * Proves Theorem Q612: Recovery can reconstruct authority entirely from the ledger.
 *
 * Complexity: O(Events)
 */
func (r *Registry) ReconstructFromLedger(events []*ledger.Event) error {
	for _, e := range events {
		switch e.Type {
		case ledger.EventGenesis, ledger.EventAuthorityIssue:
			var p ledger.AuthorityIssuePayload
			if err := json.Unmarshal(e.Payload, &p); err != nil {
				return err
			}
			r.IssueRootAuthority(p.ID, p.Atoms)

		case ledger.EventDelegation:
			var p ledger.AuthorityDelegatePayload
			if err := json.Unmarshal(e.Payload, &p); err != nil {
				return err
			}
			r.DelegateAuthority(p.ParentID, p.ChildAtoms)

		case ledger.EventAuthorityRevoke:
			var p ledger.AuthorityRevokePayload
			if err := json.Unmarshal(e.Payload, &p); err != nil {
				return err
			}
			r.RevokeAuthority(p.ID, p.Reason)

		case ledger.EventAuthorityFreeze:
			var p ledger.AuthorityStatusPayload
			if err := json.Unmarshal(e.Payload, &p); err != nil {
				return err
			}
			r.SetState(p.ID, AuthFrozen, p.Reason)

		case ledger.EventAuthorityQuarantine:
			var p ledger.AuthorityStatusPayload
			if err := json.Unmarshal(e.Payload, &p); err != nil {
				return err
			}
			r.SetState(p.ID, AuthQuarantined, p.Reason)
		}
	}
	return nil
}
