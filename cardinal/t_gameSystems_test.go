package main

import (
	"fmt"
	"testing"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/receipt"
	"pkg.world.dev/world-engine/cardinal/testutils"
	"pkg.world.dev/world-engine/cardinal/types"

	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/constants"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/msg"

	"github.com/stretchr/testify/assert"
)

const (
	processTokenskMsgName = "game.process-commands"
	createMsgName         = "game.create-player"
)

// #region ProcessCommandTokens Test
func TestSystem_ProcessCommands_Success(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)

	const playerName = "Hueyu"
	const roomSpawn = 0
	var tokens = []string{"LOOK", "WITH", "BOTTLE", "AT", "THE", "WINDOW"}
	expectedOut := "You are standing on a windsept plain where the wind blowing is cold and bison skulls in piles taller than houses cover the plains as far as your eye can see the air tastes of burnt grease and bensons.\n You see a A slightly deflated knock off uefa football, not quite spherical, it's kickable though. You see a path"

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	// Process Commands
	processTxHash := tf.AddTransaction(getProcessMsgID(t, tf.World), msg.ProcessCommandsMsg{
		PlayerName: playerName,
		Tokens:     tokens,
	})
	tf.DoTick()

	// Make sure process was successful
	processReceipt := getReceiptFromPastTick(t, tf.World, processTxHash)
	if errs := processReceipt.Errs; len(errs) > 0 {
		t.Fatalf("expected no errors when processing the commands; got %v", errs)
	}

	// Type assert the Result field to msg.ProcessCommandsReply
	processCommandsReply, ok := processReceipt.Result.(msg.ProcessCommandsReply)
	if !ok {
		t.Fatalf("expected processReceipt.Result to be of type msg.ProcessCommandsReply; got %T", processReceipt.Result)
	}
	// Access the fields
	assert.Equal(t, true, processCommandsReply.Success)
	assert.Equal(t, "Processing tokens completed", processCommandsReply.Message)
	assert.Equal(t, expectedOut, processCommandsReply.Result)
}

func TestSystem_ProcessCommands_Failure(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)

	const playerName = "Hueyu"
	const roomSpawn = 0
	var tokens = []string{"LOOK", "WITH", "BOTTLE", "AT", "THE", "WINDOW", "LOOK", "WITH", "BOTTLE", "AT", "THE", "WINDOW", "LOOK", "WITH", "BOTTLE", "AT", "THE", "WINDOW"}
	er := constants.ErrParserRoutineTKCX.Code
	expectedOut := "WTF, slow down cowboy, you're gonna hurt yourself"

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	// Process Commands
	processTxHash := tf.AddTransaction(getProcessMsgID(t, tf.World), msg.ProcessCommandsMsg{
		PlayerName: playerName,
		Tokens:     tokens,
	})
	tf.DoTick()

	// Make sure process was successful
	processReceipt := getReceiptFromPastTick(t, tf.World, processTxHash)
	if errs := processReceipt.Errs; len(errs) > 0 {
		t.Fatalf("expected no errors when processing the commands; got %v", errs)
	}

	// Type assert the Result field to msg.ProcessCommandsReply
	processCommandsReply, ok := processReceipt.Result.(msg.ProcessCommandsReply)
	if !ok {
		t.Fatalf("expected processReceipt.Result to be of type msg.ProcessCommandsReply; got %T", processReceipt.Result)
	}
	// Access the fields
	assert.Equal(t, false, processCommandsReply.Success)
	assert.Equal(t, (fmt.Sprintf("Error: %v processing the commands: %s", er, expectedOut)), processCommandsReply.Message)
	assert.Equal(t, expectedOut, processCommandsReply.Result)
}

// THIS NEEDS TO GET THE SYSTEMS
func TestHandleVerb(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)

	const playerName = "Hueyu"
	const roomSpawn = 0
	/*
		tokens := []string{"LOOK", "TO", "THE", "WINDOW"}
		print(mts)
		sut = handleVerb()
	*/
}

// #endregion ProcessCommandTokens Test

func getCreateMsgID(t *testing.T, world *cardinal.World) types.MessageID {
	return getMsgID(t, world, createMsgName)
}

func getProcessMsgID(t *testing.T, world *cardinal.World) types.MessageID {
	return getMsgID(t, world, processTokenskMsgName)
}

func getMsgID(t *testing.T, world *cardinal.World, fullName string) types.MessageID {
	msg, ok := world.GetMessageByFullName(fullName)
	if !ok {
		t.Fatalf("failed to get %q message", fullName)
	}
	return msg.ID()
}

// getReceiptFromPastTick search past ticks for a txHash that matches the given txHash. An error will be returned if
// the txHash cannot be found in Cardinal's history.
func getReceiptFromPastTick(t *testing.T, world *cardinal.World, txHash types.TxHash) receipt.Receipt {
	tick := world.CurrentTick()
	for {
		tick--
		receipts, err := world.GetTransactionReceiptsForTick(tick)
		if err != nil {
			t.Fatal(err)
		}
		for _, r := range receipts {
			if r.TxHash == txHash {
				return r
			}
		}
	}
}
