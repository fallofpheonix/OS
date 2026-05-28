package entity_runtime

import (
	"testing"
	"github.com/fallofpheonix/phoenixmind-validator/truth/evidence"
)

func TestUpdateEntity(t *testing.T) {
	entityID := "test-entity-1"
	
	// Test initial update
	entity := UpdateEntity(entityID, evidence.VALIDATED)
	if entity == nil {
		t.Fatalf("Expected entity, got nil")
	}
	if entity.ID != entityID {
		t.Errorf("Expected ID %s, got %s", entityID, entity.ID)
	}
	if entity.State != evidence.VALIDATED {
		t.Errorf("Expected state VALIDATED, got %s", entity.State)
	}

	// Test state change
	updatedEntity := UpdateEntity(entityID, evidence.ESCALATED)
	if updatedEntity.State != evidence.ESCALATED {
		t.Errorf("Expected state ESCALATED, got %s", updatedEntity.State)
	}
}
