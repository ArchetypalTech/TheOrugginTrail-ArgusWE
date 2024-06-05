package system

import (
	"fmt"
	"log"
	"math/rand"
	"time"

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

					// Compare the Player's name from the entity with the one from the CreatePlayerMsg.
					if playerManager.PlayerName == createPlayerData.Msg.PlayersName {
						// Log for when the name already exists.
						log.Printf("Player with name: %v already exists.\n", createPlayerData.Msg.PlayersName)
						existingPlayerEntityID = id
						return false // Stop the iteration since the player already exists.
					}

					return true // Continue with the iteration.
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
				// Log that the player already exists.
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

			// Retrieve all rooms and select a random one.
			var roomIDs []types.EntityID
			err = cardinal.NewSearch().Entity(filter.Exact(filter.Component[component.Room]())).Each(world, func(id types.EntityID) bool {
				roomIDs = append(roomIDs, id)
				return true
			})

			if err != nil || len(roomIDs) == 0 {
				err := fmt.Errorf("no rooms available to add the player")
				log.Println(err)
				return msg.CreatePlayerReply{
					Success: false,
					Message: err.Error(),
				}, err
			}

			rand.Seed(time.Now().UnixNano())
			selectedRoomID := roomIDs[rand.Intn(len(roomIDs))]

			playerID := uint32(playerManagerID) // Convert EntityID to uint32

			// Set the Player Entity variables of the created player entity.
			if err := cardinal.SetComponent[component.Player](world, playerManagerID, &component.Player{
				PlayerEntityID:   playerManagerID,
				PlayerName:       createPlayerData.Msg.PlayersName,
				PlayerID:         playerID,
				RoomID:           uint32(selectedRoomID),
				PlayerConnection: true,
			}); err != nil {
				// Log error updating the Player entity.
				return msg.CreatePlayerReply{
					Success: false,
					Message: fmt.Sprintf("Error updating the Player Entity: %v", err),
				}, err
			}

			// Retrieve the room component and add the player to it.
			selectedRoom, err := cardinal.GetComponent[component.Room](world, selectedRoomID)
			if err != nil {
				log.Printf("Failed to retrieve room component: %v", err)
				return msg.CreatePlayerReply{
					Success: false,
					Message: fmt.Sprintf("Failed to retrieve the Room component: %v", err),
				}, err
			}

			for i := 0; i < len(selectedRoom.Players); i++ {
				if selectedRoom.Players[i] == 0 {
					selectedRoom.Players[i] = playerID
					break
				}
			}

			// Update the RoomComponent - Entity
			if err := cardinal.SetComponent[component.Room](world, selectedRoomID, selectedRoom); err != nil {
				log.Printf("Failed to update room component: %v", err)
				return msg.CreatePlayerReply{
					Success: false,
					Message: fmt.Sprintf("Failed to update the Room component: %v", err),
				}, err
			}

			// Log player entity created successfully.
			fmt.Printf("Player entity created successfully")

			return msg.CreatePlayerReply{
				Success:        true,
				Message:        fmt.Sprintf("Player: %v was created successfully. It's entity ID is: %v and has been placed at room: %v whose ID is: %v", createPlayerData.Msg.PlayersName, playerManagerID, selectedRoom.RoomType.String(), selectedRoomID),
				PlayerEntityID: playerManagerID,
			}, nil
		},
	)
}
