package system

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/component"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/msg"
)

// CreatePlayerSystem creates a new player based on the "CreatePlayerMsg" transaction.
func CreatePlayerSystem(world cardinal.WorldContext) error {
	return cardinal.EachMessage[msg.CreatePlayerMsg, msg.CreatePlayerReply](
		world,
		func(createPlayerData cardinal.TxData[msg.CreatePlayerMsg]) (msg.CreatePlayerReply, error) {
			// Search for an existing player by name
			existingPlayerEntityID, err := FindExistingPlayer(world, createPlayerData.Msg.PlayersName)
			if err != nil {
				return msg.CreatePlayerReply{
					Success: false,
					Message: fmt.Sprintf("Error searching for Player entities: %v", err),
				}, err
			}

			// If an existing player is found, return the existing entity ID.
			if existingPlayerEntityID != 0 {
				world.Logger().Debug().Msgf("Player with name: %v already exists", createPlayerData.Msg.PlayersName)

				return msg.CreatePlayerReply{
					Success:        false,
					Message:        fmt.Sprintf("Player with name: %v already exists.", createPlayerData.Msg.PlayersName),
					PlayerEntityID: existingPlayerEntityID,
				}, nil
			}

			// Create a new player entity
			playerManagerID, err := CreateNewPlayer(world, createPlayerData.Msg.PlayersName)
			if err != nil {
				return msg.CreatePlayerReply{
					Success: false,
					Message: fmt.Sprintf("Failed to create the Player Entity: %v", err),
				}, err
			}

			player, err := GetPlayer(playerManagerID, world)

			// Assign the player to the specified room
			roomID := types.EntityID(createPlayerData.Msg.RoomID)
			if err := AssignPlayerToRoom(world, player, roomID); err != nil {
				world.Logger().Debug().Msgf("Failed to assign player to the room: %v", err)

				return msg.CreatePlayerReply{
					Success: false,
					Message: fmt.Sprintf("Failed to assign player to the room: %v", err),
				}, err
			}

			// Update the player's room ID
			if err := updatePlayerRoomID(world, playerManagerID, roomID); err != nil {
				world.Logger().Debug().Msgf("Failed to update player's room ID: %v", err)

				return msg.CreatePlayerReply{
					Success: false,
					Message: fmt.Sprintf("Failed to update player's room ID: %v", err),
				}, err
			}

			world.Logger().Info().Msgf("Player entity created successfully with ID: %v", playerManagerID)

			return msg.CreatePlayerReply{
				Success:        true,
				Message:        fmt.Sprintf("Player: %v was created successfully. It's entity ID is: %v and has been placed in room: %v", createPlayerData.Msg.PlayersName, playerManagerID, roomID),
				PlayerEntityID: playerManagerID,
			}, nil
		},
	)
}

// Search for existing player by name
func FindExistingPlayer(world cardinal.WorldContext, playerName string) (types.EntityID, error) {
	var existingPlayerEntityID types.EntityID
	err := cardinal.NewSearch().Entity(filter.Exact(filter.Component[component.Player]())).
		Each(world, func(id types.EntityID) bool {
			playerManager, err := cardinal.GetComponent[component.Player](world, id)
			if err != nil {
				world.Logger().Debug().Msgf("Error getting Player Component: %v", err)
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
func CreateNewPlayer(world cardinal.WorldContext, playerName string) (types.EntityID, error) {
	playerManagerID, err := cardinal.Create(world, &component.Player{})
	if err != nil {
		world.Logger().Debug().Msgf("Failed to create player entity: %v", err)
		return 0, err
	}

	playerID := uint32(playerManagerID) // Convert EntityID to uint32
	if err := cardinal.SetComponent[component.Player](world, playerManagerID, &component.Player{
		PlayerEntityID:   playerManagerID,
		PlayerName:       playerName,
		PlayerID:         playerID,
		PlayerConnection: true,
	}); err != nil {
		world.Logger().Debug().Msgf("Error updating the Player entity: %v when creating", err)
		return 0, err
	}

	return playerManagerID, nil
}

// Assign player to a specified room
func AssignPlayerToRoom(world cardinal.WorldContext, player component.Player, roomID types.EntityID) error {
	// Get the Room based on the roomID
	selectedRoom, err := cardinal.GetComponent[component.Room](world, roomID)
	if err != nil {
		world.Logger().Debug().Msgf("Failed to retrieve room component: %v", err)
		return err
	}

	// Add the player to the room using the player ID
	selectedRoom.Players[int(player.PlayerEntityID)] = player

	// Update the room entity
	if err := cardinal.SetComponent[component.Room](world, roomID, selectedRoom); err != nil {
		world.Logger().Debug().Msgf("Failed to update room component: %v", err)
		return err
	}

	return nil
}

// Update player's room ID
func updatePlayerRoomID(world cardinal.WorldContext, playerManagerID types.EntityID, roomID types.EntityID) error {
	playerManager, err := cardinal.GetComponent[component.Player](world, playerManagerID)
	if err != nil {
		world.Logger().Debug().Msgf("Error getting Player Component: %v", err)
		return err
	}

	playerManager.RoomID = uint32(roomID)
	if err := cardinal.SetComponent[component.Player](world, playerManagerID, playerManager); err != nil {
		world.Logger().Debug().Msgf("Error updating the Player entity: %v", err)
		return err
	}

	return nil
}

// Gets the just created player
func GetPlayer(pID types.EntityID, world cardinal.WorldContext) (component.Player, error) {
	var existingPlayer component.Player
	err := cardinal.NewSearch().Entity(filter.Exact(filter.Component[component.Player]())).
		Each(world, func(id types.EntityID) bool {
			player, err := cardinal.GetComponent[component.Player](world, pID)
			if err != nil {
				world.Logger().Error().Msgf("Error getting Player Component: %v", err)
				return true
			}

			if player.PlayerID == uint32(pID) {
				existingPlayer = *player
				return false
			}

			return true
		})

	return existingPlayer, err
}
