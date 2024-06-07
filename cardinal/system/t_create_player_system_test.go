package system

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"pkg.world.dev/world-engine/cardinal/types"

	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/component"
)

// Mock setPlayer function signature for testing.
type setPlayerFunc func(types.EntityID, *component.Player) error

// Define setPlayer function for testing.
func setPlayer(entityID types.EntityID, player *component.Player) error {
	// Implement your setPlayer logic here.
	return nil
}

// Mock getPlayer function signature for testing.
type getPlayerFunc func(types.EntityID) (*component.Player, error)

// Mock getPlayer function to return a player.
func getPlayer(entityID types.EntityID) (*component.Player, error) {
	// Mock a player with ID 1.
	return &component.Player{}, nil
}

// Mock getRoom function signature for testing.
type getRoomFunc func(types.EntityID) (*component.Room, error)

// Mock setRoom function signature for testing.
type setRoomFunc func(types.EntityID, *component.Room) error

func TestFindExistingPlayer(t *testing.T) {
	// Mock getPlayer function to return an existing player.
	getPlayer := func(entityID types.EntityID) (*component.Player, error) {
		// Mock a player with ID 1.
		player := &component.Player{
			PlayerName: "ExistingPlayer",
		}
		return player, nil
	}

	// Test finding an existing player.
	foundEntityID, err := findExistingPlayerTest("ExistingPlayer", getPlayer)
	assert.NoError(t, err)
	assert.NotEqual(t, types.EntityID(0), foundEntityID)

	// Test not finding a player.
	foundEntityID, err = findExistingPlayerTest("NonExistingPlayer", getPlayer)
	assert.NoError(t, err)
	assert.Equal(t, types.EntityID(0), foundEntityID)
}

func TestCreateNewPlayer(t *testing.T) {
	// Mock createEntity function to return a new player entity ID.
	createEntity := func(entity interface{}) (types.EntityID, error) {
		// Mock a new player with ID 1.
		return types.EntityID(1), nil
	}

	// Test creating a new player.
	playerEntityID, err := createNewPlayerTest("NewPlayer", createEntity)
	assert.NoError(t, err)
	assert.NotEqual(t, types.EntityID(0), playerEntityID)
}

func TestAssignPlayerToRoom(t *testing.T) {
	// Mock getRoom function to return a room.
	getRoom := func(roomID types.EntityID) (*component.Room, error) {
		// Mock a room with ID 1.
		return &component.Room{}, nil
	}

	// Mock setRoom function to perform assignment without error.
	setRoom := func(roomID types.EntityID, room *component.Room) error {
		return nil
	}

	// Test assigning player to room.
	err := assignPlayerToRoomTest(1, 5, getRoom, setRoom)
	assert.NoError(t, err)
}

func TestUpdatePlayerRoomID(t *testing.T) {
	// Mock getPlayer function to return a player.
	getPlayer := func(entityID types.EntityID) (*component.Player, error) {
		// Mock a player with ID 1.
		return &component.Player{}, nil
	}

	// Mock setPlayer function to perform update without error.
	setPlayer := func(entityID types.EntityID, player *component.Player) error {
		return nil
	}

	// Test updating player's room ID.
	err := updatePlayerRoomIDTest(1, 1, getPlayer, setPlayer)
	assert.NoError(t, err)
}

// findExistingPlayer searches for an existing player by name.
func findExistingPlayerTest(playerName string, getPlayer getPlayerFunc) (types.EntityID, error) {
	var existingPlayerEntityID types.EntityID

	// Replace the search logic with the provided function to get player details.
	playerManager, err := getPlayer(existingPlayerEntityID)
	if err != nil {
		return 0, err
	}

	// Check if the player with the given name exists.
	if playerManager.PlayerName == playerName {
		existingPlayerEntityID = types.EntityID(1) // or the ID of the existing player
	}

	return existingPlayerEntityID, nil
}

// createNewPlayer creates a new player entity.
func createNewPlayerTest(playerName string, createEntity func(entity interface{}) (types.EntityID, error)) (types.EntityID, error) {
	// Create a new player entity.
	playerManagerID, err := createEntity(&component.Player{})
	if err != nil {
		if isDevelopmentMode() {
			logger.Errorf("\033[31mFailed to create player entity: %v\033[0m", err)
		}
		return 0, err
	}

	playerID := uint32(playerManagerID) // Convert EntityID to uint32

	// Set player component details.
	if err := setPlayer(playerManagerID, &component.Player{
		PlayerEntityID:   playerManagerID,
		PlayerName:       playerName,
		PlayerID:         playerID,
		PlayerConnection: true,
	}); err != nil {
		return 0, err
	}

	return playerManagerID, nil
}

// assignPlayerToRoom assigns a player to a specified room.
func assignPlayerToRoomTest(playerID uint32, roomID types.EntityID, getRoom getRoomFunc, setRoom setRoomFunc) error {
	// Get the Room based on the roomID.
	selectedRoom, err := getRoom(roomID)
	if err != nil {
		if isDevelopmentMode() {
			logger.Errorf("\033[31mFailed to retrieve room component: %v\033[0m", err)
		}
		return err
	}

	// Add the player to the room using the player ID.
	for i := 0; i < len(selectedRoom.Players); i++ {
		if selectedRoom.Players[i] == 0 {
			selectedRoom.Players[i] = playerID
			break
		}
	}

	// Update the room entity.
	if err := setRoom(roomID, selectedRoom); err != nil {
		if isDevelopmentMode() {
			logger.Errorf("\033[31mFailed to update room component: %v\033[0m", err)
		}
		return err
	}

	return nil
}

// updatePlayerRoomID updates player's room ID.
func updatePlayerRoomIDTest(playerManagerID types.EntityID, roomID types.EntityID, getPlayer getPlayerFunc, setPlayer setPlayerFunc) error {
	playerManager, err := getPlayer(playerManagerID)
	if err != nil {
		if isDevelopmentMode() {
			logger.Errorf("\033[31mError getting Player Component: %v\033[0m", err)
		}
		return err
	}

	playerManager.RoomID = uint32(roomID)
	if err := setPlayer(playerManagerID, playerManager); err != nil {
		if isDevelopmentMode() {
			logger.Errorf("\033[31mError updating the Player entity: %v\033[0m", err)
		}
		return err
	}

	return nil
}
