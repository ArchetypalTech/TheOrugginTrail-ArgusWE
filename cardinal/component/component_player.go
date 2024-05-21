package component

import "pkg.world.dev/world-engine/cardinal/types"

type Player struct {
	playerEntityID   types.EntityID // EntityID assigned to the player.
	playerName       string         // Name of the player, can be the username.
	playerConnection bool           // Bool to notify if the player is connected or not.

}

func (Player) Name() string {
	return "Player"
}
