package system

import (
	"fmt"
	"strings"

	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/component"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"
	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/types"
)

func InventorySystem(world cardinal.WorldContext) error {

	return nil
}

func Inventory(playerID uint32, ts *TokeniserSystem, world cardinal.WorldContext) (string, uint8) {
	var descriptions []string
	var description string
	var errInv uint8

	// Change the type of variable from uint32 to types.EntityID
	pEID := types.EntityID(playerID)
	// Gets the player based on its entityID
	player, err := GetPlayer(pEID, world)
	if err != nil {
		world.Logger().Error().Msgf("Error getting the player for the Inventory System: %v", err)
		return "Error getting the player", 1
	}

	// Check the length of items in the inventory
	if len(player.Inventory) == 0 {
		description = "Your carrier bag doesn't even have a spiderweb."
		errInv = 0
	} else {
		// Iterate over the inventory objects and build the description
		for _, inInventoryObject := range player.Inventory {
			object := player.Inventory[int(inInventoryObject.ObjectID)]
			objectDescription := ts.GetObjectType(object.ObjectType.String()).String()
			if len(descriptions) == 0 {
				// First item
				description = fmt.Sprintf("You have a %s", objectDescription)
			} else {
				// Subsequent items
				description = fmt.Sprintf("and a %s", objectDescription)
			}
			descriptions = append(descriptions, description)
		}

		// Join all descriptions into a single string
		description = fmt.Sprintf("%s.", strings.Join(descriptions, " "))
		errInv = 0
		world.Logger().Debug().Msgf("In Inventory: %s", description)
	}

	return description, errInv
}

func Take(tokens []string, playerID uint32, roomID uint32, ts *TokeniserSystem, world cardinal.WorldContext) (string, uint8) {
	var tok_err uint8
	var hasBeenPickedUp bool
	var tok string
	var description string
	tok = tokens[1]

	// Get the room
	room, err := GetRoom(types.EntityID(roomID), world)
	if err != nil {
		world.Logger().Error().Msgf("Error getting the room with ID: %v, for the Inventory System %v", roomID, err)
	}

	// Get the player
	player, err := GetPlayer(types.EntityID(playerID), world)
	if err != nil {
		world.Logger().Error().Msgf("Error getting the player wiht ID: %v, for the Inventory System: %v", playerID, err)
	}

	// Get the Object Type
	ObjectType := ts.GetObjectType(tok)

	if ObjectType != enums.ObjectTypeNone {
		for _, lookingObject := range room.Objects {
			if lookingObject.ObjectID != 0 && lookingObject.CanBePickedUp {
				object := room.Objects[int(lookingObject.ObjectID)]
				if object.ObjectName == ObjectType.String() {
					// Add the object to the player inventory
					player.Inventory[int(object.ObjectID)] = object
					// Remove the object from the room
					delete(room.Objects, int(object.ObjectID))

					room.Players[int(player.PlayerID)] = player

					// Update the player entity
					if err := cardinal.SetComponent[component.Player](world, types.EntityID(playerID), &player); err != nil {
						world.Logger().Debug().Msgf("Error updating the room entity: %v, after taking an object in the inventory system", err)
					}

					// Update the room entity
					if err := cardinal.SetComponent[component.Room](world, types.EntityID(roomID), &room); err != nil {
						world.Logger().Debug().Msgf("Error updating the room entity: %v, after taking an object in the inventory system", err)
					}

					description = fmt.Sprintf("You picked up a %s.", object.ObjectName)
					tok_err = 0
					hasBeenPickedUp = true
					break
				}
			}
		}

	}

	if hasBeenPickedUp == false {
		description = "Can't pick something that doesn't exists, right?"
		tok_err = 0
	}

	return description, tok_err
}
