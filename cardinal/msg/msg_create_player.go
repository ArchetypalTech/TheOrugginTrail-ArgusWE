/* This message is used to communicate the intention of creating a player.
It uses the player name to create the player. Creates an ID for the player entity.
Returns a Success bool, a Message and the PlayerEntity ID.
*/

package msg

import (
	"pkg.world.dev/world-engine/cardinal/types"
)

type CreatePlayerMsg struct {
	PlayersName string `json:"PlayerName"` // Name of the player, can be the username.
}

type CreatePlayerReply struct {
	Success        bool           `json:"Success"`        // Indicates whether the move was successful or not.
	Message        string         `json:"Message"`        // Optional message providing additional information.
	PlayerEntityID types.EntityID `json:"PlayerEntityID"` // Identifier for the player.
}
