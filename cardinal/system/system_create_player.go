package system

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/message"
	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/component"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/msg"
)

// CreatePlayerSystem creates a new player based on the "CreatePlayerMsg" transaction.
func CreatePlayerSystem(world cardinal.WorldContext) error {
	return cardinal.EachMessage[msg.CreatePlayerMsg, msg.CreatePlayerReply](
		world,
		func(createPlayerData message.TxData[msg.CreatePlayerMsg]) (msg.CreatePlayerReply, error) {
			// Search for an existing player by name
			existingPlayerEntityID, err := findExistingPlayer(world, createPlayerData.Msg.PlayersName)
			if err != nil {
				if isDevelopmentMode() {
					logger.Errorf("\033[31mError searching for Player entities: %v\033[0m", err)
				}

				return msg.CreatePlayerReply{
					Success: false,
					Message: fmt.Sprintf("Error searching for Player entities: %v", err),
				}, err
			}

			// If an existing player is found, return the existing entity ID.
			if existingPlayerEntityID != 0 {
				if isDevelopmentMode() {
					logger.Warnf("\033[33mPlayer with name: %v already exists.\033[0m", createPlayerData.Msg.PlayersName)
				}

				return msg.CreatePlayerReply{
					Success:        false,
					Message:        fmt.Sprintf("Player with name: %v already exists.", createPlayerData.Msg.PlayersName),
					PlayerEntityID: existingPlayerEntityID,
				}, nil
			}

			// Create a new player entity
			playerManagerID, err := createNewPlayer(world, createPlayerData.Msg.PlayersName)
			if err != nil {
				if isDevelopmentMode() {
					logger.Errorf("\033[31mFailed to create the Player Entity: %v\033[0m", err)
				}

				return msg.CreatePlayerReply{
					Success: false,
					Message: fmt.Sprintf("Failed to create the Player Entity: %v", err),
				}, err
			}

			playerID := uint32(playerManagerID) // Convert EntityID to uint32

			// Assign the player to the specified room
			roomID := types.EntityID(createPlayerData.Msg.RoomID)
			if err := assignPlayerToRoom(world, playerID, roomID); err != nil {
				if isDevelopmentMode() {
					logger.Errorf("\033[31mFailed to assign player to the room: %v\033[0m", err)
				}

				return msg.CreatePlayerReply{
					Success: false,
					Message: fmt.Sprintf("Failed to assign player to the room: %v", err),
				}, err
			}

			// Update the player's room ID
			if err := updatePlayerRoomID(world, playerManagerID, roomID); err != nil {
				if isDevelopmentMode() {
					logger.Errorf("\033[31mFailed to update player's room ID: %v\033[0m", err)
				}

				return msg.CreatePlayerReply{
					Success: false,
					Message: fmt.Sprintf("Failed to update player's room ID: %v", err),
				}, err
			}

			if isDevelopmentMode() {
				// Log player entity created successfully.
				logger.Infof("\033[32mPlayer entity created successfully\033[0m")
			}

			return msg.CreatePlayerReply{
				Success:        true,
				Message:        fmt.Sprintf("Player: %v was created successfully. It's entity ID is: %v and has been placed in room: %v", createPlayerData.Msg.PlayersName, playerManagerID, roomID),
				PlayerEntityID: playerManagerID,
			}, nil
		},
	)
}

// Search for existing player by name
func findExistingPlayer(world cardinal.WorldContext, playerName string) (types.EntityID, error) {
	var existingPlayerEntityID types.EntityID
	err := cardinal.NewSearch().Entity(filter.Exact(filter.Component[component.Player]())).
		Each(world, func(id types.EntityID) bool {
			playerManager, err := cardinal.GetComponent[component.Player](world, id)
			if err != nil {
				if isDevelopmentMode() {
					logger.Errorf("\033[31mError getting Player Component: %v\033[0m", err)
				}
				return true
			}

			if playerManager.PlayerName == playerName {
				existingPlayerEntityID = id
				return false // Stop the iteration since the player already exists.
			}

			return true // Continue with the iteration.
		})
	return existingPlayerEntityID, err
}

// Create a new player entity
func createNewPlayer(world cardinal.WorldContext, playerName string) (types.EntityID, error) {
	playerManagerID, err := cardinal.Create(world, &component.Player{})
	if err != nil {
		if isDevelopmentMode() {
			logger.Errorf("\033[31mFailed to create player entity: %v\033[0m", err)
		}
		return 0, err
	}

	playerID := uint32(playerManagerID) // Convert EntityID to uint32
	if err := cardinal.SetComponent[component.Player](world, playerManagerID, &component.Player{
		PlayerEntityID:   playerManagerID,
		PlayerName:       playerName,
		PlayerID:         playerID,
		PlayerConnection: true,
	}); err != nil {
		if isDevelopmentMode() {
			logger.Errorf("\033[31mError updating the Player entity: %v\033[0m", err)
		}
		return 0, err
	}

	return playerManagerID, nil
}

// Assign player to a specified room
func assignPlayerToRoom(world cardinal.WorldContext, playerID uint32, roomID types.EntityID) error {
	// Get the Room based on the roomID
	selectedRoom, err := cardinal.GetComponent[component.Room](world, roomID)
	if err != nil {
		if isDevelopmentMode() {
			logger.Errorf("\033[31mFailed to retrieve room component: %v\033[0m", err)
		}
		return err
	}

	// Add the player to the room using the player ID
	for i := 0; i < len(selectedRoom.Players); i++ {
		if selectedRoom.Players[i] == 0 {
			selectedRoom.Players[i] = playerID
			break
		}
	}

	// Update the room entity
	if err := cardinal.SetComponent[component.Room](world, roomID, selectedRoom); err != nil {
		if isDevelopmentMode() {
			logger.Errorf("\033[31mFailed to update room component: %v\033[0m", err)
		}
		return err
	}

	return nil
}

// Update player's room ID
func updatePlayerRoomID(world cardinal.WorldContext, playerManagerID types.EntityID, roomID types.EntityID) error {
	playerManager, err := cardinal.GetComponent[component.Player](world, playerManagerID)
	if err != nil {
		if isDevelopmentMode() {
			logger.Errorf("\033[31mError getting Player Component: %v\033[0m", err)
		}
		return err
	}

	playerManager.RoomID = uint32(roomID)
	if err := cardinal.SetComponent[component.Player](world, playerManagerID, playerManager); err != nil {
		if isDevelopmentMode() {
			logger.Errorf("\033[31mError updating the Player entity: %v\033[0m", err)
		}
		return err
	}

	return nil
}
