package state_runtime

import (
	"testing"
)

func TestTaskLifecycle(t *testing.T) {
	lifecycle := &Lifecycle{}
	taskID := "task-123"
	entityID := "entity-abc"

	// Start task
	task := lifecycle.StartTask(taskID, entityID)
	if task.ID != taskID || task.EntityID != entityID || task.Status != "RUNNING" {
		t.Errorf("StartTask failed: got %+v", task)
	}
	// Note: StartTime is set within the function, not directly testable here
	// without mocking time.Time. We check if it's set to non-zero value in actual impl.

	// Complete task
	completedTask := lifecycle.CompleteTask(task, "success")
	if completedTask.Status != "COMPLETED" || completedTask.Result != "success" {
		t.Errorf("CompleteTask failed: got %+v", completedTask)
	}
	// Note: EndTime is set within the function.

	// Fail task
	task = lifecycle.StartTask("task-456", "entity-def") // Start another task to fail
	failedTask := lifecycle.FailTask(task, "timeout")
	if failedTask.Status != "FAILED" || failedTask.Result != "timeout" {
		t.Errorf("FailTask failed: got %+v", failedTask)
	}
	// Note: EndTime is set within the function.
}
