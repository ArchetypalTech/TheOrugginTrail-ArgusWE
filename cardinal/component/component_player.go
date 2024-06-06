package component

import (
	"pkg.world.dev/world-engine/cardinal/types"
)

type Player struct {
	PlayerEntityID   types.EntityID // EntityID assigned to the player.
	PlayerName       string         // Name of the player, can be the username.
	PlayerID         uint32         `json:"player_id"`  // ID of the player.
	RoomID           uint32         `json:"room_id"`    // ID of the room where the playre is.
	ObjectIDs        [32]uint32     `json:"object_ids"` // Map of all the objects tha the player has. This should be the inventory.
	PlayerConnection bool           // Bool to notify if the player is connected or not.

}

func (Player) Name() string {
	return "Player"
}
