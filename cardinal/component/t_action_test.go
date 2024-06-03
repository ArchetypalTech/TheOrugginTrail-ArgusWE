package component

import (
	"testing"

	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"
)

func TestActionStore_AddAndGet(t *testing.T) {
	store := NewActionStore()

	// Test Add method
	actionID := store.Add(Action{
		ActionType: enums.ActionType(1),
		Enabled:    true,
	})

	// Test Get method
	action, found := store.Get(actionID)
	if !found {
		t.Errorf("Expected to find action with ID %d, but not found", actionID)
	}
	if action.ActionType != enums.ActionType(1) {
		t.Errorf("Expected action type to be 1, got %d", action.ActionType)
	}
	if !action.Enabled {
		t.Errorf("Expected action to be enabled, but it's disabled")
	}
}

func TestActionStore_Set(t *testing.T) {
	store := NewActionStore()
	actionID := store.Add(Action{
		ActionType: enums.ActionType(1),
	})

	// Test Set method
	store.Set(actionID, Action{
		ActionType: enums.ActionType(2),
		Enabled:    true,
	})

	// Retrieve the action and check if it's updated
	action, _ := store.Get(actionID)
	if action.ActionType != enums.ActionType(2) {
		t.Errorf("Expected action type to be 2, got %d", action.ActionType)
	}
	if !action.Enabled {
		t.Errorf("Expected action to be enabled, but it's disabled")
	}
}

// tries to retrieve an action with an ID that doesn't exist in the store and expects it not to be found.
func TestActionStore_Get_NotFound(t *testing.T) {
	store := NewActionStore()

	// Try to get an action that does not exist
	_, found := store.Get(999)
	if found {
		t.Error("Expected action not to be found, but it was found")
	}
}

// tries to set an action with an ID that doesn't exist in the store and expects this operation to fail.
func TestActionStore_Set_NotFound(t *testing.T) {
	store := NewActionStore()

	// Try to set an action that does not exist
	store.Set(999, Action{
		ActionType: enums.ActionType(1),
		Enabled:    true,
	})

	// If no panic or error occurs, consider the test failed
	t.Error("Expected setting action to fail, but it succeeded")
}
