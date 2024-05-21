package component

import "pkg.world.dev/world-engine/cardinal/types"

type Player struct {
	PlayerEntityID   types.EntityID // EntityID assigned to the player.
	PlayerName       string         // Name of the player, can be the username.
	PlayerConnection bool           // Bool to notify if the player is connected or not.

}

func (Player) Name() string {
	return "Player"
}
