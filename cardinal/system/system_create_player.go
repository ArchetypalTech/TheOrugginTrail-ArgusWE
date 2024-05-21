package system

import (
	"fmt"
	"log"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/message"
	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/component"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/msg"
)

// CreatePlayerSystem creates a new player bases on the "CreatePlayerMsg" transaction.
func CreatePlayerSystem(world cardinal.WorldContext) error {
	return cardinal.EachMessage[msg.CreatePlayerMsg, msg.CreatePlayerReply](
		world,
		func(createPlayerData message.TxData[msg.CreatePlayerMsg]) (msg.CreatePlayerReply, error) {
			// Search for all entities with the Player component.
			var existingPlayerEntityID types.EntityID
			err := cardinal.NewSearch().Entity(filter.Exact(filter.Component[component.Player]())).
				Each(world, func(id types.EntityID) bool {
					// Get the Player component for the current entity.
					playerManager, err := cardinal.GetComponent[component.Player](world, id)
					if err != nil {
						// Log error getting Player component.
						log.Printf("Error getting Player Component: %v\n", err)
						return true
					}

					// Compare the Player's name from the entity with the one from the CreatePlayerMsg
					if playerManager.PlayerName == createPlayerData.Msg.PlayersName {
						// Log for when the name already exists
						log.Printf("Player with name: %v already exists.\n", createPlayerData.Msg.PlayersName)
						existingPlayerEntityID = id
						return false // Stop the iteration since the player already exists.
					}

					return true // continue with the iteration.
				})

			if err != nil {
				// Log search error.
				log.Printf("Error searching for Player entities: %v\n", err)
				return msg.CreatePlayerReply{
					Success: false,
					Message: fmt.Sprintf("Error searching for Player entities: %v", err),
				}, err
			}

			// If an entity with Player component was found, return false and the existing entity ID.
			if existingPlayerEntityID != 0 {
				// Log that the match already exists.
				fmt.Printf("Player with name: %v already exists.\n", createPlayerData.Msg.PlayersName)
				return msg.CreatePlayerReply{
					Success:        false,
					Message:        fmt.Sprintf("Player with name: %v already exists.", createPlayerData.Msg.PlayersName),
					PlayerEntityID: existingPlayerEntityID,
				}, nil
			}

			// If no Player entity was found, create a new one.
			playerManagerID, err := cardinal.Create(world, &component.Player{})

			if err != nil {
				// Log player creation failure.
				log.Printf("Failed to create player entity: %v", err)
				return msg.CreatePlayerReply{
					Success: false,
					Message: fmt.Sprintf("Failed to create the Player Entity: %s", err),
				}, err
			}

			// Set the Player Entity variables of the created player entity
			if err := cardinal.SetComponent[component.Player](world, playerManagerID, &component.Player{
				PlayerEntityID:   playerManagerID,
				PlayerName:       createPlayerData.Msg.PlayersName,
				PlayerConnection: true,
			}); err != nil {
				// Log error updating the Player entity.
				return msg.CreatePlayerReply{
					Success: false,
					Message: fmt.Sprintf("Error updating the Player Entity:%v", err),
				}, err
			}

			// Log player entity created successfully.
			fmt.Printf("Player entity created successfully")

			return msg.CreatePlayerReply{
				Success:        true,
				Message:        fmt.Sprintf("Player: %v was created successfully. It's entity ID is: %v", createPlayerData.Msg.PlayersName, playerManagerID),
				PlayerEntityID: playerManagerID,
			}, nil
		},
	)
}
