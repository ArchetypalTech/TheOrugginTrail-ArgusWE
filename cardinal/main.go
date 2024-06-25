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
		cardinal.RegisterComponent[component.ActionStore](w),
		cardinal.RegisterComponent[component.ActionOutput](w),
		cardinal.RegisterComponent[component.Description](w),
		cardinal.RegisterComponent[component.Object](w),
		cardinal.RegisterComponent[component.ObjectStore](w),
		cardinal.RegisterComponent[component.Output](w),
		cardinal.RegisterComponent[component.Player](w),
		cardinal.RegisterComponent[component.Room](w),
		cardinal.RegisterComponent[component.RoomStore](w),
		cardinal.RegisterComponent[component.TxtDef](w),
	)

	// Register messages (user action)
	// NOTE: You must register your transactions here for it to be executed.
	Must(
		cardinal.RegisterMessage[msg.CreatePlayerMsg, msg.CreatePlayerReply](w, "create-player"),
		cardinal.RegisterMessage[msg.ProcessCommandsMsg, msg.ProcessCommandsReply](w, "process-commands"),
	)

	// Register queries
	// NOTE: You must register your queries here for it to be accessible.

	Must()

	// Each system executes deterministically in the order they are added.
	// This is a neat feature that can be strategically used for systems that depends on the order of execution.
	Must(cardinal.RegisterSystems(w,
		system.CreatePlayerSystem,
		system.ProcessCommandsTokens,
		system.NTokeniserSystem,
		system.LookSystem,
		system.InventorySystem,
	))

	// Register the init system when the world is initiated.
	Must(cardinal.RegisterInitSystems(w,
		system.NGameSetupSystem,
	))

}

func Must(err ...error) {
	e := errors.Join(err...)
	if e != nil {
		log.Fatal().Err(e).Msg("")
	}
}
