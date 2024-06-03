package component

import (
	"testing"

	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"
)

func TestRoomStore_AddAndGet(t *testing.T) {
	store := NewRoomStore()

	// Test Add method
	roomID := store.Add(Room{
		Description: "Test Room",
		RoomType:    enums.RoomType(1),
	})

	// Test Get method
	room, found := store.Get(roomID)
	if !found {
		t.Errorf("Expected to find room with ID %d, but not found", roomID)
	}
	if room.Description != "Test Room" {
		t.Errorf("Expected description to be 'Test Room', got %s", room.Description)
	}
	if room.RoomType != enums.RoomType(1) {
		t.Errorf("Expected room type to be 1, got %d", room.RoomType)
	}
}

func TestRoomStore_SetDescription(t *testing.T) {
	store := NewRoomStore()
	roomID := store.Add(Room{
		Description: "Test Room",
	})

	// Test SetDescription method
	store.SetDescription(roomID, "Updated Room Description")

	// Retrieve the room and check if the description is updated
	room, _ := store.Get(roomID)
	if room.Description != "Updated Room Description" {
		t.Errorf("Expected description to be 'Updated Room Description', got %s", room.Description)
	}
}

func TestRoomStore_SetRoomType(t *testing.T) {
	store := NewRoomStore()
	roomID := store.Add(Room{
		RoomType: enums.RoomType(1),
	})

	// Test SetRoomType method
	store.SetRoomType(roomID, enums.RoomType(2))

	// Retrieve the room and check if the room type is updated
	room, _ := store.Get(roomID)
	if room.RoomType != enums.RoomType(2) {
		t.Errorf("Expected room type to be 2, got %d", room.RoomType)
	}
}

func TestRoomStore_SetDescription_Failure(t *testing.T) {
	store := NewRoomStore()
	roomID := store.Add(Room{
		Description: "Test Room",
	})

	// Test SetDescription method with an invalid room ID
	invalidRoomID := roomID + 1 // Adding 1 to make it invalid
	store.SetDescription(invalidRoomID, "Updated Room Description")

	// Retrieve the original room and check if the description remains unchanged
	room, _ := store.Get(roomID)
	if roomID != invalidRoomID {
		t.Errorf("Room ID is incorrect as it should be %d but is %d", roomID, invalidRoomID)
	}
	if room.Description != "Test Room" {
		t.Errorf("Expected description to remain unchanged as 'Test Room', got %s", room.Description)
	}
}

// tries to retrieve a room with an ID that doesn't exist in the store and expects it not to be found.
func TestRoomStore_Get_NotFound(t *testing.T) {
	store := NewRoomStore()

	// Try to get a room that does not exist
	_, found := store.Get(999)
	if found {
		t.Error("Expected room not to be found, but it was found")
	}
}

// tries to set a room with an ID that doesn't exist in the store and expects this operation to fail.
func TestRoomStore_Set_NotFound(t *testing.T) {
	store := NewRoomStore()

	// Try to set a room that does not exist
	store.Set(999, Room{
		Description: "Test Room",
	})

	// If no panic or error occurs, consider the test failed
	t.Error("Expected setting room to fail, but it succeeded")
}
