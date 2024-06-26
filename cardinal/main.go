package main

import (
	"errors"

	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/component"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/msg"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/system"
	"github.com/rs/zerolog/log"
	"pkg.world.dev/world-engine/cardinal"
)

func main() {
	w, err := cardinal.NewWorld(cardinal.WithDisableSignatureVerification())
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	MustInitWorld(w)

	Must(w.StartGame())
}

// MustInitWorld registers all components, messages, queries, and systems. This initialization happens in a helper
// function so that this can be used directly in tests.
func MustInitWorld(w *cardinal.World) {
	// Register components
	// NOTE: You must register your components here for it to be accessible.
	Must(
		cardinal.RegisterComponent[component.Action](w),
		cardinal.RegisterComponent[component.ActionOutput](w),
		cardinal.RegisterComponent[component.Description](w),
		cardinal.RegisterComponent[component.DirObject](w),
		cardinal.RegisterComponent[component.Object](w),
		cardinal.RegisterComponent[component.Output](w),
		cardinal.RegisterComponent[component.Player](w),
		cardinal.RegisterComponent[component.Room](w),
		cardinal.RegisterComponent[component.TxtDef](w),
	)

	// Register messages (user action)
	// NOTE: You must register your transactions here for it to be executed.
	Must(

		cardinal.RegisterMessage[msg.CreatePlayerMsg, msg.CreatePlayerReply](w, "create-player"),
		cardinal.RegisterMessage[msg.CreateRoomMsg, msg.CreateRoomReply](w, "create-room"),
	)

	// Register queries
	// NOTE: You must register your queries here for it to be accessible.

	Must()

	// Each system executes deterministically in the order they are added.
	// This is a neat feature that can be strategically used for systems that depends on the order of execution.
	// For example, you may want to join the match before changing the role of the player
	// that way, the roles are confirmed before starting the game.
	Must(cardinal.RegisterSystems(w,
		system.CreatePlayerSystem,
	))

}

func Must(err ...error) {
	e := errors.Join(err...)
	if e != nil {
		log.Fatal().Err(e).Msg("")
	}
}
