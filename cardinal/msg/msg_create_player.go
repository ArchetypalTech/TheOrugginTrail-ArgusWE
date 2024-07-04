/* This message is used to communicate the intention of creating a player.
It uses the player name to create the player. Creates an ID for the player entity.
Returns a Success bool, a Message, the PlayerEntity ID, and the room description where he is created.
*/

package msg

import (
	"pkg.world.dev/world-engine/cardinal/types"
)

type CreatePlayerMsg struct {
	PlayersName string `json:"PlayerName"` // Name of the player, can be the username.
	RoomID      uint32 `json:"RoomID"`     // Room ID that indicates on which room the player will start at.
}

type CreatePlayerReply struct {
	Success         bool           `json:"Success"`         // Indicates whether the move was successful or not.
	Message         string         `json:"Message"`         // Optional message providing additional information.
	PlayerEntityID  types.EntityID `json:"PlayerEntityID"`  // Identifier for the player.
	RoomDescription string         `json:"RoomDescription"` // Description of the room where the player has been crated.
}
