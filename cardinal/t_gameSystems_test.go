package main

import (
	"fmt"
	"testing"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/receipt"
	"pkg.world.dev/world-engine/cardinal/testutils"
	"pkg.world.dev/world-engine/cardinal/types"

	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/component"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/constants"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/msg"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/system"

	"github.com/stretchr/testify/assert"
)

const (
	processTokenskMsgName = "game.process-commands"
	createMsgName         = "game.create-player"
)

// #region Create Player Test
func TestSystem_CreatePlayer_Success(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)

	const playerName = "Hueyu"
	const roomSpawn = 1
	const expectedMessage = "Player: Hueyu was created successfully. It's entity ID is: 4 and has been placed in room: 1"
	const expectedRoomDesc = "You are standing on a windswept plain"

	// Create an initial player
	processTxHash := tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	// Make sure process was successful
	processReceipt := getReceiptFromPastTick(t, tf.World, processTxHash)
	if errs := processReceipt.Errs; len(errs) > 0 {
		t.Fatalf("expected no errors when creating the player; got %v", errs)
	}

	// Type assert the Result field to msg.ProcessCommandsReply
	createPlayerReply, ok := processReceipt.Result.(msg.CreatePlayerReply)
	if !ok {
		t.Fatalf("expected processReceipt.Result to be of type msg.CreatePlayerReply; got %T", processReceipt.Result)
	}
	// Access the fields
	assert.Equal(t, true, createPlayerReply.Success)
	assert.Equal(t, expectedMessage, createPlayerReply.Message)
	assert.Equal(t, expectedRoomDesc, createPlayerReply.RoomDescription)
}

func TestFindExistingPlayer_NoFind(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)

	const playerName = "Hueyu"
	const roomSpawn = 1
	var expectedID types.EntityID
	var expectedErr error = nil

	id, err := system.FindExistingPlayer(cardinal.NewReadOnlyWorldContext(tf.World), playerName)

	assert.Equal(t, expectedID, id)
	assert.Equal(t, expectedErr, err)
}

func TestFindExistingPlayer_Find(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)

	const playerName = "Hueyu"
	const roomSpawn = 1
	var expectedID types.EntityID = 4
	var expectedErr error

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	id, err := system.FindExistingPlayer(cardinal.NewReadOnlyWorldContext(tf.World), playerName)

	assert.Equal(t, expectedID, id)
	assert.Equal(t, expectedErr, err)
}

func TestCreateNewPlayer_Success(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	tf.DoTick()

	const playerName = "Hueyu"
	const roomSpawn = 1
	var expectedID types.EntityID = 4
	var expectedErr error

	id, err := system.CreateNewPlayer(cardinal.NewWorldContext(tf.World), playerName)
	tf.DoTick()

	assert.Equal(t, expectedID, id)
	assert.Equal(t, expectedErr, err)
}

func TestGetPlayer_Success(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	tf.DoTick()

	const playerName = "Hueyu"
	const roomSpawn = 1
	var pID types.EntityID = 4
	var expectedErr error = nil

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	player, err := system.GetPlayer(pID, cardinal.NewWorldContext(tf.World))
	playerID := player.PlayerID
	tf.DoTick()

	assert.Equal(t, uint32(pID), playerID)
	assert.Equal(t, expectedErr, err)
}

func TestAssignPlayerToRoom_Success(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	tf.DoTick()

	const playerName = "Hueyu"
	const roomSpawn = 1
	var pID types.EntityID = 4
	var expectedErr error = nil

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	player, err := system.GetPlayer(pID, cardinal.NewWorldContext(tf.World))

	tf.DoTick()
	rID := types.EntityID(roomSpawn)

	err = system.AssignPlayerToRoom(cardinal.NewWorldContext(tf.World), player, rID)

	assert.Equal(t, expectedErr, err)
}

// #endregion Create Player Test

// #region ProcessCommandTokens System Test
func TestSystem_ProcessCommands_Success(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	tf.DoTick()

	const playerName = "Hueyu"
	const roomSpawn = 1
	var tokens = []string{"LOOK", "WITH", "BOTTLE", "AT", "THE", "WINDOW"}
	var expectedOut string = ("You are standing on a windswept plain where the wind blowing is cold and bison skulls in piles taller than houses cover the plains as far as your eye can see" +
		" the air tastes of burnt grease and bensons." +
		" You see a A slightly deflated knock off uefa football, not quite spherical, it's kickable though." +
		" There is a path made mainly from dirt to the North and there is a path made mainly from mud to the East." +
		" There is no other poor soul here apart from you.")
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

func TestSystem_ProcessCommands_Success_GoNextRoom(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	tf.DoTick()

	const playerName1 = "Hueyu"
	const playerName2 = "GOD"
	const roomSpawn1 = 3
	const roomSpawn2 = 1
	var tokens = []string{"go", "west"}
	var expectedOut string = ("You are standing on a windswept plain")
	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName1,
		RoomID:      roomSpawn1,
	})
	tf.DoTick()

	// Create second player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName2,
		RoomID:      roomSpawn2,
	})
	tf.DoTick()

	// Process Commands
	processTxHash := tf.AddTransaction(getProcessMsgID(t, tf.World), msg.ProcessCommandsMsg{
		PlayerName: playerName1,
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

func TestSystem_ProcessCommands_Failure_GoNextRoom(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	tf.DoTick()

	const playerName1 = "Hueyu"
	const playerName2 = "GOD"
	const roomSpawn1 = 3
	const roomSpawn2 = 1
	var tokens = []string{"go", "south"}
	var expectedOut string = ("Can't go that way ")
	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName1,
		RoomID:      roomSpawn1,
	})
	tf.DoTick()

	// Create second player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName2,
		RoomID:      roomSpawn2,
	})
	tf.DoTick()

	// Process Commands
	processTxHash := tf.AddTransaction(getProcessMsgID(t, tf.World), msg.ProcessCommandsMsg{
		PlayerName: playerName1,
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
	assert.Equal(t, "Error: 8 processing the commands: Can't go that way ", processCommandsReply.Message)
	assert.Equal(t, expectedOut, processCommandsReply.Result)
}

var ts *system.TokeniserSystem

func setup() {
	ts = system.NewTokeniserSystem()
}

func TestHandleVerb_Success(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 1

	tokens := []string{"TAKE", "ball"}
	var expectedOutTake string = ("You picked up a Football.")

	tokens2 := []string{"DROP", "ball"}
	expectedOutStDrop := "You dropped the Football."
	var expectedOutErr uint8 = 0

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	outputTake, er := system.HandleVerb(tokens, roomSpawn, playerID, ts, cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	outputDrop, er := system.HandleVerb(tokens2, roomSpawn, playerID, ts, cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	assert.Equal(t, expectedOutTake, outputTake)
	assert.Equal(t, expectedOutStDrop, outputDrop)
	assert.Equal(t, expectedOutErr, er)
}

/*
The HandleVerb does not have a failure itself as the erros come from the function that calls this function one
from the functions that this function calls like stuff from the look system.
The switch case has a default case which returns a value of 0 for the error so no error and the verb.
*/
func TestHandleVerb_Failure(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 1

	tokens := []string{"KICK"}
	expectedOutSt := "You Kick."
	var expectedOutErr uint8 = 0

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	output, er := system.HandleVerb(tokens, roomSpawn, playerID, ts, cardinal.NewReadOnlyWorldContext(tf.World))

	assert.Equal(t, expectedOutSt, output)
	assert.Equal(t, expectedOutErr, er)
}

func TestHandleAlias_Success(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 1

	tokens := []string{"LOOK"}
	var expectedOutSt string = ("You are standing on a windswept plain where the wind blowing is cold and bison skulls in piles taller than houses cover the plains as far as your eye can see" +
		" the air tastes of burnt grease and bensons." +
		" You see a A slightly deflated knock off uefa football, not quite spherical, it's kickable though." +
		" There is a path made mainly from dirt to the North and there is a path made mainly from mud to the East." +
		" There is no other poor soul here apart from you.")
	var expectedOutErr uint8 = 0

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	output, er := system.HandleAlias(tokens, roomSpawn, playerID, ts, cardinal.NewReadOnlyWorldContext(tf.World))

	assert.Equal(t, expectedOutSt, output)
	assert.Equal(t, expectedOutErr, er)
}

/*
The HandleAlias does not have a failure itself as the erros come from the function that calls this function one
or from the functions that this function calls like stuff from the look system.
The switch case is either a look or inventory type. This is testing if the inventory returns a string if there is no item inside it..
*/
func TestHandleAlias_Failure(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 1

	tokens := []string{"INVENTORY"}
	expectedOutSt := "Your carrier bag doesn't even have a spiderweb."
	var expectedOutErr uint8 = 0

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	output, er := system.HandleAlias(tokens, roomSpawn, playerID, ts, cardinal.NewReadOnlyWorldContext(tf.World))

	assert.Equal(t, expectedOutSt, output)
	assert.Equal(t, expectedOutErr, er)
}

func TestInsultMeat_Success(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	const erroR = 130
	const badCommand = "SEE"
	expectedOutSt := "Nope, gibberish. Stop breathing with your mouth."

	output := system.InsultMeat(erroR, badCommand)

	assert.Equal(t, expectedOutSt, output)
}

/*
The insult meat does not have an error itself, it recieves an uint8 and depending on tha value it will use a swtich case.
For this test we are sending a value that is part of the default and returns the following string: "What are you doing?!?!"
*/
func TestInsultMeat_Failure(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	var erroR uint8
	const badCommand = "SEE"
	expectedOutSt := "What are you doing?!?!"

	output := system.InsultMeat(erroR, badCommand)

	assert.Equal(t, expectedOutSt, output)
}

// #endregion ProcessCommandTokens System Test

// #region Tokeniser System Test
func TestFishTokens_Success(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	var tokens = []string{"look", "WITH", "BOTTLE", "AT", "THE", "window"}
	const expectedVrb = "Look"
	const expectedDObj = "Bottle"
	const expectedIObj = "Window"

	tf.DoTick()

	data := system.NewTokeniserSystem().FishTokens(tokens)

	assert.Equal(t, expectedVrb, data.Verb.String())
	assert.Equal(t, expectedDObj, data.DirectObject.String())
	assert.Equal(t, expectedIObj, data.IndirectObject.String())
}

func TestFishTokens_NoDObj(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	var tokens = []string{"burn", "AT", "THE", "window"}
	const expectedVrb = "Burn"
	const expectedDObj = "None"
	const expectedIObj = "Window"

	tf.DoTick()

	data := system.NewTokeniserSystem().FishTokens(tokens)

	assert.Equal(t, expectedVrb, data.Verb.String())
	assert.Equal(t, expectedDObj, data.DirectObject.String())
	assert.Equal(t, expectedIObj, data.IndirectObject.String())
}

func TestFishTokens_NoIObj(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	var tokens = []string{"KICK", "THE", "window"}
	const expectedVrb = "Kick"
	const expectedDObj = "Window"
	const expectedIObj = "None"

	tf.DoTick()

	data := system.NewTokeniserSystem().FishTokens(tokens)

	assert.Equal(t, expectedVrb, data.Verb.String())
	assert.Equal(t, expectedDObj, data.DirectObject.String())
	assert.Equal(t, expectedIObj, data.IndirectObject.String())
}

// #endregion Tokeniser System Test

// #region Look System Test

func TestStuff_Success(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 3
	const roomSpawn = 0

	tokens := []string{"LOOK"}
	var expectedOutSt string = ("You are standing on a windsept plain where the wind blowing is cold and bison skulls in piles taller than houses cover the plains as far as your eye can see" +
		" the air tastes of burnt grease and bensons.\n " +
		"You see a A slightly deflated knock off uefa football, not quite spherical, it's kickable though." +
		" There is a path made mainly from mud to the East and there is a path made mainly from dirt to the North." +
		" There is no other poor soul here apart from you.")
	var expectedOutErr uint8 = 0

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	output, er := system.Stuff(tokens, roomSpawn, playerID, ts, cardinal.NewReadOnlyWorldContext(tf.World))

	assert.Equal(t, expectedOutSt, output)
	assert.Equal(t, expectedOutErr, er)
}

func TestLookAround_Success(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 3
	const roomSpawn = 0

	var expectedOutSt string = ("You are standing on a windsept plain where the wind blowing is cold and bison skulls in piles taller than houses cover the plains as far as your eye can see" +
		" the air tastes of burnt grease and bensons.\n " +
		"You see a A slightly deflated knock off uefa football, not quite spherical, it's kickable though." +
		" There is a path made mainly from dirt to the North and there is a path made mainly from mud to the East." +
		" There is no other poor soul here apart from you.")
	var expectedOutErr uint8 = 0

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	output, er := system.LookAround(roomSpawn, playerID, ts, cardinal.NewReadOnlyWorldContext(tf.World))

	assert.Equal(t, expectedOutSt, output)
	assert.Equal(t, expectedOutErr, er)
}

func TestDescriptionText_Success(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 3
	const roomSpawn = 0

	var expectedOutSt string = ("You are standing on a windsept plain where the wind blowing is cold and bison skulls in piles taller than houses cover the plains as far as your eye can see" +
		" the air tastes of burnt grease and bensons.\n " +
		"You see a A slightly deflated knock off uefa football, not quite spherical, it's kickable though." +
		" There is a path made mainly from dirt to the North and there is a path made mainly from mud to the East." +
		" There is no other poor soul here apart from you.")
	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	output := system.GenDescText(playerID, roomSpawn, ts, cardinal.NewReadOnlyWorldContext(tf.World))

	assert.Equal(t, expectedOutSt, output)
}

func TestDirObjsDesc_Success(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	tf.DoTick()
	setup()

	const roomID = 1
	rID := types.EntityID(roomID)
	expectedOutSt := "There is a path made mainly from dirt to the North and there is a path made mainly from mud to the East."

	room, err := system.GetRoom(rID, cardinal.NewReadOnlyWorldContext(tf.World))
	if err != nil {
		fmt.Printf("Error getting Room Component: %v", err)
	}
	tf.DoTick()

	output := system.DirObjectDescription(room, ts, cardinal.NewReadOnlyWorldContext(tf.World))

	assert.Equal(t, expectedOutSt, output)
}

func TestGetRoom_Success(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	tf.DoTick()
	setup()

	const roomID = 1
	rID := types.EntityID(roomID)
	const expectedOutSt = "RoomID is: 1"
	var output string

	room, err := system.GetRoom(rID, cardinal.NewReadOnlyWorldContext(tf.World))
	if err != nil {
		fmt.Printf("Error getting Room Component: %v", err)

	} else {
		output = fmt.Sprintf("RoomID is: %v", room.ID)
	}

	assert.Equal(t, expectedOutSt, output)
}
func TestGetRoom_Failure(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	tf.DoTick()
	setup()

	const roomID = 5
	rID := types.EntityID(roomID)

	room, err := system.GetRoom(rID, cardinal.NewReadOnlyWorldContext(tf.World))
	if err != nil {
		fmt.Printf("Error getting Room Component: %v", err)

	} else {
		print(fmt.Sprintf("RoomID is: %v", room.ID))
	}
}
func TestGetMaterial_Success(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	tf.DoTick()
	setup()

	const material = "mud"
	const dirObj = enums.ObjectTypePath
	expectedOutSt := " made mainly from mud"

	matDes := system.GenMaterialDesc(material, dirObj, ts)
	assert.Equal(t, expectedOutSt, matDes)
}

func TestGetPlayersPresence_NoOtherPlayerAround(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	tf.DoTick()
	setup()

	const roomID = 1
	const playerName = "Hueyu"
	const pID = 4
	rID := types.EntityID(roomID)
	expectedOutSt := "There is no other poor soul here apart from you."

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomID,
	})
	tf.DoTick()

	room, err := system.GetRoom(rID, cardinal.NewReadOnlyWorldContext(tf.World))
	if err != nil {
		fmt.Printf("Error getting Room Component: %v", err)
	}
	tf.DoTick()

	output := system.GetPlayersPresence(room, pID, cardinal.NewReadOnlyWorldContext(tf.World))

	assert.Equal(t, expectedOutSt, output)
}
func TestGetPlayersPresence_OtherPlayerAround(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	tf.DoTick()
	setup()

	const roomID = 1
	const playerName1 = "Hueyu"
	const playerName2 = "Tatron"
	const playerName3 = "GOD"

	const pID = 4
	rID := types.EntityID(roomID)
	expectedOutSt := " In this room is Tatron, and GOD"

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName1,
		RoomID:      roomID,
	})
	tf.DoTick()

	// Create an second player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName2,
		RoomID:      roomID,
	})
	tf.DoTick()

	// Create an second player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName3,
		RoomID:      roomID,
	})
	tf.DoTick()

	room, err := system.GetRoom(rID, cardinal.NewReadOnlyWorldContext(tf.World))
	if err != nil {
		fmt.Printf("Error getting Room Component: %v", err)
	}
	tf.DoTick()

	output := system.GetPlayersPresence(room, pID, cardinal.NewReadOnlyWorldContext(tf.World))

	assert.Equal(t, expectedOutSt, output)
}

// #endregion Look System Test

// #region Inventory System Test
func TestInventory_WithItem(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 1

	tokens := []string{"take", "ball"}
	var expectedOutInventory string = ("You have a Football.")
	var expectedOutTake string = ("You picked up a Football.")
	var expectedOutErr uint8 = 0

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	outputTake, er := system.Take(tokens, playerID, roomSpawn, ts, cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	output, er := system.Inventory(playerID, ts, cardinal.NewReadOnlyWorldContext(tf.World))

	assert.Equal(t, expectedOutInventory, output)
	assert.Equal(t, expectedOutTake, outputTake)
	assert.Equal(t, expectedOutErr, er)
}

func TestInventory_WithNOItem(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 1
	var expectedOutInventory string = ("Your carrier bag doesn't even have a spiderweb.")
	var expectedOutErr uint8 = 0

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	output, er := system.Inventory(playerID, ts, cardinal.NewReadOnlyWorldContext(tf.World))

	assert.Equal(t, expectedOutInventory, output)

	assert.Equal(t, expectedOutErr, er)
}

func TestTake_Sucess(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 1

	tokens := []string{"take", "ball"}
	var expectedOutTake string = ("You picked up a Football.")
	var expectedOutErr uint8 = 0

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	outputTake, er := system.Take(tokens, playerID, roomSpawn, ts, cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	assert.Equal(t, expectedOutTake, outputTake)
	assert.Equal(t, expectedOutErr, er)
}

func TestTake_Failure(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 1

	tokens := []string{"take", "car"}
	var expectedOutTake string = ("Can't pick something that doesn't exists, right?")
	var expectedOutErr uint8 = 0

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	outputTake, er := system.Take(tokens, playerID, roomSpawn, ts, cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	assert.Equal(t, expectedOutTake, outputTake)
	assert.Equal(t, expectedOutErr, er)
}

func TestDrop_Sucess(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 1

	tokens1 := []string{"take", "ball"}
	tokens2 := []string{"drop", "ball"}
	var expectedOutTake string = ("You picked up a Football.")
	var expectedOutDrop string = ("You dropped the Football.")
	var expectedOutErr uint8 = 0

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	outputTake, er := system.Take(tokens1, playerID, roomSpawn, ts, cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	outputDrop, er := system.Drop(tokens2, playerID, roomSpawn, ts, cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	assert.Equal(t, expectedOutTake, outputTake)
	assert.Equal(t, expectedOutDrop, outputDrop)
	assert.Equal(t, expectedOutErr, er)
}

func TestDrop_Failure(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 1

	tokens := []string{"drop", "ball"}
	var expectedOutDrop string = ("Can't drop something that you don't even have, right?")
	var expectedOutErr uint8 = 0

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	outputDrop, er := system.Drop(tokens, playerID, roomSpawn, ts, cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	assert.Equal(t, expectedOutDrop, outputDrop)
	assert.Equal(t, expectedOutErr, er)
}

// #endregion Inventory System Test

// #region Direction System Test
func TestGetNexRoom_Success(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 1

	tokens := []string{"GO", "NORTH"}
	var expectedOutGNR uint32 = uint32(enums.RoomTypeBarn)
	var expectedOutErr uint8 = 0

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	outputGNR, er := system.GetNextRoom(tokens, roomSpawn, ts, cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	assert.Equal(t, expectedOutGNR, outputGNR)
	assert.Equal(t, expectedOutErr, er)
}

func TestGetNexRoom_Failure(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 1

	tokens := []string{"south"}
	var expectedOutGNR uint32 = 0x0
	var expectedOutErr uint8 = constants.ErrNoExit.Code
	var expectedOutErrMsg string = "Can't go that way south"

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	outputGNR, er := system.GetNextRoom(tokens, roomSpawn, ts, cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	outputGNR_ERR_msg := system.InsultMeat(er, tokens[0])

	tf.DoTick()

	assert.Equal(t, expectedOutGNR, outputGNR)
	assert.Equal(t, expectedOutErr, er)
	assert.Equal(t, expectedOutErrMsg, outputGNR_ERR_msg)
}

func TestFishDirectionToken_1Token(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 1

	tokens := []string{"north"}
	var expectedOutTake string = ("north")
	var expectedOutErr uint8 = 0

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	outputToken, er := system.FishDirectionTok(tokens, ts)

	tf.DoTick()

	assert.Equal(t, expectedOutTake, outputToken)
	assert.Equal(t, expectedOutErr, er)
}

func TestFishDirectionToken_ShortForm(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 1

	tokens := []string{"go", "backward"}
	var expectedOutTake string = ("backward")
	var expectedOutErr uint8 = 0

	tf.DoTick()

	outputToken, er := system.FishDirectionTok(tokens, ts)

	tf.DoTick()

	assert.Equal(t, expectedOutTake, outputToken)
	assert.Equal(t, expectedOutErr, er)
}

func TestFishDirectionToken_LongForm(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 1

	tokens := []string{"GO", "to", "the", "SOUTH"}
	var expectedOutTake string = ("SOUTH")
	var expectedOutErr uint8 = 0

	tf.DoTick()

	outputToken, er := system.FishDirectionTok(tokens, ts)

	tf.DoTick()

	assert.Equal(t, expectedOutTake, outputToken)
	assert.Equal(t, expectedOutErr, er)
}

func TestDirectionCheck_CanMove_Success(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 1
	const direction = enums.DirectionTypeNorth

	var expectedOutBool bool = true

	tf.DoTick()

	outputDirCheck, component := system.DirectionCheck(roomSpawn, direction, ts, cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	print(component.Description)

	assert.Equal(t, expectedOutBool, outputDirCheck)
}

func TestDirectionCheck_CanMove_Failure(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 3
	const direction = enums.DirectionTypeEast

	var expectedOutBool bool = false

	tf.DoTick()

	outputDirCheck, component := system.DirectionCheck(roomSpawn, direction, ts, cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	print(component.Description)

	assert.Equal(t, expectedOutBool, outputDirCheck)
}

// #endregion Direction System Test

// #region Action System Test

func TestFetchObjsForType_Success(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 1
	const actType = enums.ActionTypeKick

	expectedObject := component.Object{
		ObjectID:        3,
		ObjectName:      "Football",
		ObjectType:      enums.ObjectTypeFootball,
		MaterialType:    enums.MaterialTypeFlesh,
		DirType:         enums.DirectionTypeNone,
		DestID:          enums.RoomTypeNone,
		Description:     "A slightly deflated knock off uefa football, not quite spherical, it's kickable though.",
		ObjectActionIDs: []uint32{3},
		CanBePickedUp:   true,
	}

	tf.DoTick()

	room, roomErr := system.GetRoom(types.EntityID(roomSpawn), cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	fetchedObjs := system.FetchObjsForType(actType, room, playerID, ts, cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	assert.Equal(t, roomErr, nil)
	assert.Equal(t, []component.Object([]component.Object{expectedObject}), fetchedObjs)
}

func TestFetchObjsForType_Failure(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 1
	const actType = enums.ActionTypeLock

	expectedObject := component.Object{
		ObjectID:        4,
		ObjectName:      "Knife",
		ObjectType:      enums.ObjectTypeKnife,
		MaterialType:    enums.MaterialTypeIron,
		DirType:         enums.DirectionTypeNone,
		DestID:          enums.RoomTypeNone,
		Description:     "A rusty iron knife that is missing some edge.",
		ObjectActionIDs: []uint32{3},
		CanBePickedUp:   true,
	}

	tf.DoTick()

	room, roomErr := system.GetRoom(types.EntityID(roomSpawn), cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	fetchedObjs := system.FetchObjsForType(actType, room, playerID, ts, cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	assert.Equal(t, roomErr, nil)
	assert.NotEqual(t, []component.Object([]component.Object{expectedObject}), fetchedObjs)
}

func TestFetchDirObjsForType_Success(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 2
	const actType = enums.ActionTypeKick

	expectedObject1 := component.Object{
		ObjectID:        4,
		ObjectName:      "",
		ObjectType:      enums.ObjectTypeDoor,
		MaterialType:    enums.MaterialTypeWood,
		DirType:         enums.DirectionTypeSouth,
		DestID:          enums.RoomTypePlain,
		Description:     "door",
		ObjectActionIDs: []uint32{4},
		CanBePickedUp:   false,
	}

	expectedObject2 := component.Object{
		ObjectID:        5,
		ObjectName:      "",
		ObjectType:      enums.ObjectTypeWindow,
		MaterialType:    enums.MaterialTypeWood,
		DirType:         enums.DirectionTypeEast,
		DestID:          enums.RoomTypePlain,
		Description:     "window",
		ObjectActionIDs: []uint32{5, 6},
		CanBePickedUp:   false,
	}

	tf.DoTick()

	room, roomErr := system.GetRoom(types.EntityID(roomSpawn), cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	fetchedObjs := system.FetchDirObjsForType(actType, room, ts, cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	assert.Equal(t, roomErr, nil)
	assert.Equal(t, []component.Object([]component.Object{expectedObject1, expectedObject2, expectedObject2}), fetchedObjs)
}

func TestFetchDirObjsForType_Failure(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 2
	const actType = enums.ActionTypeKick

	expectedObject1 := component.Object{
		ObjectID:        4,
		ObjectName:      "",
		ObjectType:      enums.ObjectTypeDoor,
		MaterialType:    enums.MaterialTypeStone,
		DirType:         enums.DirectionTypeWest,
		DestID:          enums.RoomTypePlain,
		Description:     "path",
		ObjectActionIDs: []uint32{4},
		CanBePickedUp:   false,
	}

	expectedObject2 := component.Object{
		ObjectID:        5,
		ObjectName:      "",
		ObjectType:      enums.ObjectTypeWindow,
		MaterialType:    enums.MaterialTypeWood,
		DirType:         enums.DirectionTypeEast,
		DestID:          enums.RoomTypePlain,
		Description:     "window",
		ObjectActionIDs: []uint32{5, 6},
		CanBePickedUp:   false,
	}

	tf.DoTick()

	room, roomErr := system.GetRoom(types.EntityID(roomSpawn), cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	fetchedObjs := system.FetchDirObjsForType(actType, room, ts, cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	assert.Equal(t, roomErr, nil)
	assert.NotEqual(t, []component.Object([]component.Object{expectedObject1, expectedObject2, expectedObject2}), fetchedObjs)
}

func TestGetResponseStr_WithIObj(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	cmdData := component.VerbData{
		Verb:           enums.ActionTypeBurn,
		DirectObject:   enums.ObjectTypeFootball,
		IndirectObject: enums.ObjectTypeDoor,
		ErrCode:        constants.NOERR,
	}
	var expectedStr = "You Burn the Football at the Door."

	tf.DoTick()

	responseGot := system.GetResponseStr(cmdData, ts, cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	assert.Equal(t, expectedStr, responseGot)
}

func TestGetResponseStr_WithNOIObj(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	cmdData := component.VerbData{
		Verb:           enums.ActionTypeKick,
		DirectObject:   enums.ObjectTypeBottle,
		IndirectObject: enums.ObjectTypeNone,
		ErrCode:        constants.NOERR,
	}
	var expectedStr = "You Kick the Bottle."

	tf.DoTick()

	responseGot := system.GetResponseStr(cmdData, ts, cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	assert.Equal(t, expectedStr, responseGot)
}

func TestHandleBaseAction_Success(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 1
	const actType = enums.ActionTypeKick

	expectedObject := component.Object{
		ObjectID:        3,
		ObjectName:      "Football",
		ObjectType:      enums.ObjectTypeFootball,
		MaterialType:    enums.MaterialTypeFlesh,
		DirType:         enums.DirectionTypeNone,
		DestID:          enums.RoomTypeNone,
		Description:     "A slightly deflated knock off uefa football, not quite spherical, it's kickable though.",
		ObjectActionIDs: []uint32{3},
		CanBePickedUp:   true,
	}

	cmdData := component.VerbData{
		Verb:           enums.ActionTypeKick,
		DirectObject:   enums.ObjectTypeFootball,
		IndirectObject: enums.ObjectTypeNone,
		ErrCode:        constants.NOERR,
	}

	var expectedStr = (" The ball (such as it is)" +
		" bounces feebly then rolls into some fresh dog eggs" +
		" none the less you briefly feel a little better.")

	tf.DoTick()

	room, roomErr := system.GetRoom(types.EntityID(roomSpawn), cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	fetchedObjs := system.FetchObjsForType(actType, room, playerID, ts, cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	handleStr, handleErr := system.HandleBaseAction(cmdData, fetchedObjs, cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	assert.Equal(t, roomErr, nil)
	assert.Equal(t, []component.Object([]component.Object{expectedObject}), fetchedObjs)
	assert.Equal(t, handleErr, uint8(0))
	assert.Equal(t, expectedStr, handleStr)
}

func TestHandleBaseAction_Failure(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 1
	const actType = enums.ActionTypeKick

	expectedObject := component.Object{
		ObjectID:        3,
		ObjectName:      "Football",
		ObjectType:      enums.ObjectTypeFootball,
		MaterialType:    enums.MaterialTypeFlesh,
		DirType:         enums.DirectionTypeNone,
		DestID:          enums.RoomTypeNone,
		Description:     "A slightly deflated knock off uefa football, not quite spherical, it's kickable though.",
		ObjectActionIDs: []uint32{3},
		CanBePickedUp:   true,
	}

	cmdData := component.VerbData{
		Verb:           enums.ActionTypeKick,
		DirectObject:   enums.ObjectTypeDoor,
		IndirectObject: enums.ObjectTypeNone,
		ErrCode:        constants.NOERR,
	}

	var expectedStr = ""

	tf.DoTick()

	room, roomErr := system.GetRoom(types.EntityID(roomSpawn), cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	fetchedObjs := system.FetchObjsForType(actType, room, playerID, ts, cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	handleStr, handleErr := system.HandleBaseAction(cmdData, fetchedObjs, cardinal.NewWorldContext(tf.World))
	print(handleErr)
	tf.DoTick()

	assert.Equal(t, roomErr, nil)
	assert.Equal(t, []component.Object([]component.Object{expectedObject}), fetchedObjs)
	assert.Equal(t, expectedStr, handleStr)
}

func TestSystem_ProcessCommands_ActionSystemComplete_Success_NoBarn(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	tf.DoTick()

	const playerName = "Hueyu"
	const roomSpawn = 1
	var tokens = []string{"kick", "the", "ball", "AT", "THE", "path"}
	var expectedOut string = ("You Kick the Football at the Path.")
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

func TestSystem_ProcessCommands_ActionSystemComplete_Success_WithBarn(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	tf.DoTick()

	const playerName = "Hueyu"
	const roomSpawn = 1
	var tokens1 = []string{"take", "BALL"}
	var expectedOut1 string = ("You picked up a Football.")
	var tokens2 = []string{"go", "NORTH"}
	var expectedOut2 string = ("You are standing in a barn and place is dusty and full of spiderwebs, something died in here, possibly your own self plenty of corners and dark shadows." +
		" There is a door wood to the South and there is a window wood to the East.There is no other poor soul here apart from you.")
	var tokens3 = []string{"drop", "ball"}
	var expectedOut3 string = ("You dropped the Football.")
	var tokens4 = []string{"kick", "the", "ball", "to", "the", "window"}
	var expectedOut4 string = ("You Kick the Football at the Window. I love the sound of breaking glass especially when I'm lonely, the panes and the frame shatter satisfyingly spreading broken joy on the floor. The window, glass and frame smashed falls open.")

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	// Process Commands
	processTxHash1 := tf.AddTransaction(getProcessMsgID(t, tf.World), msg.ProcessCommandsMsg{
		PlayerName: playerName,
		Tokens:     tokens1,
	})
	tf.DoTick()

	// Make sure process was successful
	processReceipt1 := getReceiptFromPastTick(t, tf.World, processTxHash1)
	if errs := processReceipt1.Errs; len(errs) > 0 {
		t.Fatalf("expected no errors when processing the commands; got %v", errs)
	}

	// Type assert the Result field to msg.ProcessCommandsReply
	processCommandsReply1, ok := processReceipt1.Result.(msg.ProcessCommandsReply)
	if !ok {
		t.Fatalf("expected processReceipt.Result to be of type msg.ProcessCommandsReply; got %T", processReceipt1.Result)
	}

	assert.Equal(t, expectedOut1, processCommandsReply1.Result)
	tf.DoTick()

	// Process Commands2
	processTxHash2 := tf.AddTransaction(getProcessMsgID(t, tf.World), msg.ProcessCommandsMsg{
		PlayerName: playerName,
		Tokens:     tokens2,
	})
	tf.DoTick()

	// Make sure process was successful
	processReceipt2 := getReceiptFromPastTick(t, tf.World, processTxHash2)
	if errs := processReceipt2.Errs; len(errs) > 0 {
		t.Fatalf("expected no errors when processing the commands; got %v", errs)
	}

	// Type assert the Result field to msg.ProcessCommandsReply
	processCommandsReply2, ok := processReceipt2.Result.(msg.ProcessCommandsReply)
	if !ok {
		t.Fatalf("expected processReceipt.Result to be of type msg.ProcessCommandsReply; got %T", processReceipt2.Result)
	}

	assert.Equal(t, expectedOut2, processCommandsReply2.Result)
	tf.DoTick()

	// Process Commands3
	processTxHash3 := tf.AddTransaction(getProcessMsgID(t, tf.World), msg.ProcessCommandsMsg{
		PlayerName: playerName,
		Tokens:     tokens3,
	})
	tf.DoTick()

	// Make sure process was successful
	processReceipt3 := getReceiptFromPastTick(t, tf.World, processTxHash3)
	if errs := processReceipt3.Errs; len(errs) > 0 {
		t.Fatalf("expected no errors when processing the commands; got %v", errs)
	}

	// Type assert the Result field to msg.ProcessCommandsReply
	processCommandsReply3, ok := processReceipt3.Result.(msg.ProcessCommandsReply)
	if !ok {
		t.Fatalf("expected processReceipt.Result to be of type msg.ProcessCommandsReply; got %T", processReceipt3.Result)
	}

	assert.Equal(t, expectedOut3, processCommandsReply3.Result)
	tf.DoTick()

	// Process Commands4
	processTxHash4 := tf.AddTransaction(getProcessMsgID(t, tf.World), msg.ProcessCommandsMsg{
		PlayerName: playerName,
		Tokens:     tokens4,
	})
	tf.DoTick()

	// Make sure process was successful
	processReceipt4 := getReceiptFromPastTick(t, tf.World, processTxHash4)
	if errs := processReceipt4.Errs; len(errs) > 0 {
		t.Fatalf("expected no errors when processing the commands; got %v", errs)
	}

	// Type assert the Result field to msg.ProcessCommandsReply
	processCommandsReply4, ok := processReceipt4.Result.(msg.ProcessCommandsReply)
	if !ok {
		t.Fatalf("expected processReceipt.Result to be of type msg.ProcessCommandsReply; got %T", processReceipt4.Result)
	}

	assert.Equal(t, expectedOut4, processCommandsReply4.Result)
	tf.DoTick()
}

func TestFetchObjsForTypeInventory_Success(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 1
	const actType = enums.ActionTypeKick

	tokens := []string{"take", "ball"}
	var expectedOutInventory string = ("You have a Football.")
	var expectedOutTake string = ("You picked up a Football.")

	expectedObject := component.Object{
		ObjectID:        3,
		ObjectName:      "Football",
		ObjectType:      enums.ObjectTypeFootball,
		MaterialType:    enums.MaterialTypeFlesh,
		DirType:         enums.DirectionTypeNone,
		DestID:          enums.RoomTypeNone,
		Description:     "A slightly deflated knock off uefa football, not quite spherical, it's kickable though.",
		ObjectActionIDs: []uint32{3},
		CanBePickedUp:   true,
	}

	tf.DoTick()

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})

	tf.DoTick()

	outputTake, takeErr := system.Take(tokens, playerID, roomSpawn, ts, cardinal.NewWorldContext(tf.World))
	assert.Equal(t, takeErr, uint8(0))
	assert.Equal(t, expectedOutTake, outputTake)
	tf.DoTick()

	output, inventoryErr := system.Inventory(playerID, ts, cardinal.NewReadOnlyWorldContext(tf.World))
	assert.Equal(t, expectedOutInventory, output)
	assert.Equal(t, inventoryErr, uint8(0))
	tf.DoTick()

	fetchedObjs := system.FetchObjsForTypeInventory(actType, playerID, ts, cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	assert.Equal(t, []component.Object([]component.Object{expectedObject}), fetchedObjs)
}

func TestFetchObjsForTypeInventory_Failure(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 1
	const actType = enums.ActionTypeKick

	expectedObject := component.Object{
		ObjectID:        3,
		ObjectName:      "Football",
		ObjectType:      enums.ObjectTypeFootball,
		MaterialType:    enums.MaterialTypeFlesh,
		DirType:         enums.DirectionTypeNone,
		DestID:          enums.RoomTypeNone,
		Description:     "A slightly deflated knock off uefa football, not quite spherical, it's kickable though.",
		ObjectActionIDs: []uint32{3},
		CanBePickedUp:   true,
	}

	tf.DoTick()

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})

	tf.DoTick()

	fetchedObjs := system.FetchObjsForTypeInventory(actType, playerID, ts, cardinal.NewWorldContext(tf.World))

	tf.DoTick()

	assert.NotEqual(t, []component.Object([]component.Object{expectedObject}), fetchedObjs)
	assert.Equal(t, []component.Object(nil), fetchedObjs)
}

func TestFetchObjsForType_RoomToInventory_Success(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 1
	const actType = enums.ActionTypeKick

	tokens := []string{"take", "ball"}
	var expectedOutInventory string = ("You have a Football.")
	var expectedOutTake string = ("You picked up a Football.")

	expectedObject := component.Object{
		ObjectID:        3,
		ObjectName:      "Football",
		ObjectType:      enums.ObjectTypeFootball,
		MaterialType:    enums.MaterialTypeFlesh,
		DirType:         enums.DirectionTypeNone,
		DestID:          enums.RoomTypeNone,
		Description:     "A slightly deflated knock off uefa football, not quite spherical, it's kickable though.",
		ObjectActionIDs: []uint32{3},
		CanBePickedUp:   true,
	}

	tf.DoTick()

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	outputTake, takeErr := system.Take(tokens, playerID, roomSpawn, ts, cardinal.NewWorldContext(tf.World))
	assert.Equal(t, takeErr, uint8(0))
	assert.Equal(t, expectedOutTake, outputTake)
	tf.DoTick()

	output, inventoryErr := system.Inventory(playerID, ts, cardinal.NewReadOnlyWorldContext(tf.World))
	assert.Equal(t, expectedOutInventory, output)
	assert.Equal(t, inventoryErr, uint8(0))
	tf.DoTick()

	room, roomErr := system.GetRoom(types.EntityID(roomSpawn), cardinal.NewWorldContext(tf.World))
	assert.Equal(t, roomErr, nil)
	tf.DoTick()

	fetchedObjs := system.FetchObjsForType(actType, room, playerID, ts, cardinal.NewWorldContext(tf.World))
	tf.DoTick()

	assert.Equal(t, []component.Object([]component.Object{expectedObject}), fetchedObjs)
}

func TestFetchObjsForType_RoomToInventory_Empty(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	setup()

	const playerName = "Hueyu"
	const playerID = 4
	const roomSpawn = 2
	const actType = enums.ActionTypeKick

	tokens := []string{"kick", "the", "ball"}
	var expectedOut string = ("That object is only in your imagination!")

	tf.DoTick()

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
	assert.Equal(t, "Error: 135 processing the commands: That object is only in your imagination!", processCommandsReply.Message)
	assert.Equal(t, expectedOut, processCommandsReply.Result)
}

// #endregion Action System Test

// #region REST ROOMS ACT I Test
func TestUsingForgeObjects(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	tf.DoTick()

	const playerName = "Hueyu"
	const roomSpawn = 4
	var tokens1 = []string{"light", "matchsticks"}
	var expectedOut1 string = ("You Light the Matchsticks. The matchsticks (despite their small size) lights enough to see a small distance. You have to use them quickly or your fingers will be burn.")
	var tokens2 = []string{"burn", "petrol"}
	var expectedOut2 string = ("You Burn the Petrol. The petrol (with its unique funny smell) burns fiercely with a scent that makes you feel dizzy.")

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	// Process Commands
	processTxHash1 := tf.AddTransaction(getProcessMsgID(t, tf.World), msg.ProcessCommandsMsg{
		PlayerName: playerName,
		Tokens:     tokens1,
	})
	tf.DoTick()

	// Make sure process was successful
	processReceipt1 := getReceiptFromPastTick(t, tf.World, processTxHash1)
	if errs := processReceipt1.Errs; len(errs) > 0 {
		t.Fatalf("expected no errors when processing the commands; got %v", errs)
	}

	// Type assert the Result field to msg.ProcessCommandsReply
	processCommandsReply1, ok := processReceipt1.Result.(msg.ProcessCommandsReply)
	if !ok {
		t.Fatalf("expected processReceipt.Result to be of type msg.ProcessCommandsReply; got %T", processReceipt1.Result)
	}

	assert.Equal(t, expectedOut1, processCommandsReply1.Result)
	tf.DoTick()

	// Process Commands2
	processTxHash2 := tf.AddTransaction(getProcessMsgID(t, tf.World), msg.ProcessCommandsMsg{
		PlayerName: playerName,
		Tokens:     tokens2,
	})
	tf.DoTick()

	// Make sure process was successful
	processReceipt2 := getReceiptFromPastTick(t, tf.World, processTxHash2)
	if errs := processReceipt2.Errs; len(errs) > 0 {
		t.Fatalf("expected no errors when processing the commands; got %v", errs)
	}

	// Type assert the Result field to msg.ProcessCommandsReply
	processCommandsReply2, ok := processReceipt2.Result.(msg.ProcessCommandsReply)
	if !ok {
		t.Fatalf("expected processReceipt.Result to be of type msg.ProcessCommandsReply; got %T", processReceipt2.Result)
	}

	assert.Equal(t, expectedOut2, processCommandsReply2.Result)
	tf.DoTick()
}

func TestBurningPetrolAtStairsWithHay(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	tf.DoTick()

	const playerName = "Hueyu"
	const roomSpawn = 4
	var tokens1 = []string{"take", "petrol"}
	var expectedOut1 string = ("You picked up a Petrol.")
	var tokens2 = []string{"take", "matchsticks"}
	var expectedOut2 string = ("You picked up a Matchsticks.")
	var tokens3 = []string{"go", "west"}
	var expectedOut3 string = ("You are standing in a barn")
	var tokens4 = []string{"burn", "petrol", "at", "the", "stairs"}
	var expectedOut4 string = ("You Burn the Petrol at the Stairs. I can hear the crispy sound of the hay burning as it gets consumed by the dark fires." +
		" However, this enjoyable moment does not persist for too long. The hay having been burned quickly, reveals the stairs cleared up.")

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	// Process Commands
	processTxHash1 := tf.AddTransaction(getProcessMsgID(t, tf.World), msg.ProcessCommandsMsg{
		PlayerName: playerName,
		Tokens:     tokens1,
	})
	tf.DoTick()

	// Make sure process was successful
	processReceipt1 := getReceiptFromPastTick(t, tf.World, processTxHash1)
	if errs := processReceipt1.Errs; len(errs) > 0 {
		t.Fatalf("expected no errors when processing the commands; got %v", errs)
	}

	// Type assert the Result field to msg.ProcessCommandsReply
	processCommandsReply1, ok := processReceipt1.Result.(msg.ProcessCommandsReply)
	if !ok {
		t.Fatalf("expected processReceipt.Result to be of type msg.ProcessCommandsReply; got %T", processReceipt1.Result)
	}

	assert.Equal(t, expectedOut1, processCommandsReply1.Result)
	tf.DoTick()

	// Process Commands2
	processTxHash2 := tf.AddTransaction(getProcessMsgID(t, tf.World), msg.ProcessCommandsMsg{
		PlayerName: playerName,
		Tokens:     tokens2,
	})
	tf.DoTick()

	// Make sure process was successful
	processReceipt2 := getReceiptFromPastTick(t, tf.World, processTxHash2)
	if errs := processReceipt2.Errs; len(errs) > 0 {
		t.Fatalf("expected no errors when processing the commands; got %v", errs)
	}

	// Type assert the Result field to msg.ProcessCommandsReply
	processCommandsReply2, ok := processReceipt2.Result.(msg.ProcessCommandsReply)
	if !ok {
		t.Fatalf("expected processReceipt.Result to be of type msg.ProcessCommandsReply; got %T", processReceipt2.Result)
	}

	assert.Equal(t, expectedOut2, processCommandsReply2.Result)
	tf.DoTick()

	// Process Commands3
	processTxHash3 := tf.AddTransaction(getProcessMsgID(t, tf.World), msg.ProcessCommandsMsg{
		PlayerName: playerName,
		Tokens:     tokens3,
	})
	tf.DoTick()

	// Make sure process was successful
	processReceipt3 := getReceiptFromPastTick(t, tf.World, processTxHash3)
	if errs := processReceipt3.Errs; len(errs) > 0 {
		t.Fatalf("expected no errors when processing the commands; got %v", errs)
	}

	// Type assert the Result field to msg.ProcessCommandsReply
	processCommandsReply3, ok := processReceipt3.Result.(msg.ProcessCommandsReply)
	if !ok {
		t.Fatalf("expected processReceipt.Result to be of type msg.ProcessCommandsReply; got %T", processReceipt3.Result)
	}

	assert.Equal(t, expectedOut3, processCommandsReply3.Result)
	tf.DoTick()

	// Process Commands4
	processTxHash4 := tf.AddTransaction(getProcessMsgID(t, tf.World), msg.ProcessCommandsMsg{
		PlayerName: playerName,
		Tokens:     tokens4,
	})
	tf.DoTick()

	// Make sure process was successful
	processReceipt4 := getReceiptFromPastTick(t, tf.World, processTxHash4)
	if errs := processReceipt4.Errs; len(errs) > 0 {
		t.Fatalf("expected no errors when processing the commands; got %v", errs)
	}

	// Type assert the Result field to msg.ProcessCommandsReply
	processCommandsReply4, ok := processReceipt4.Result.(msg.ProcessCommandsReply)
	if !ok {
		t.Fatalf("expected processReceipt.Result to be of type msg.ProcessCommandsReply; got %T", processReceipt4.Result)
	}

	assert.Equal(t, expectedOut4, processCommandsReply4.Result)
	tf.DoTick()
}

func TestUsingCellarObjects(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	tf.DoTick()

	const playerName = "Hueyu"
	const roomSpawn = 5
	var tokens1 = []string{"sniff", "glue"}
	var expectedOut1 string = ("You Sniff the Glue. The smell of it is really relaxing, but for now thats all.")
	var tokens2 = []string{"light", "dynamite"}
	var expectedOut2 string = ("You Light the Dynamite. Seeing the fuse sparkling makes you remember the fireworks you loved as a child." +
		" Suddenly, you return back and see that there is almost no time. You have to do something or you will be flying into meat pieces!.")
	var tokens3 = []string{"throw", "dynamite"}
	var expectedOut3 string = ("You Throw the Dynamite. You throw quickly the dynamite and start running for the hills as your life depens on it.")

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	// Process Commands
	processTxHash1 := tf.AddTransaction(getProcessMsgID(t, tf.World), msg.ProcessCommandsMsg{
		PlayerName: playerName,
		Tokens:     tokens1,
	})
	tf.DoTick()

	// Make sure process was successful
	processReceipt1 := getReceiptFromPastTick(t, tf.World, processTxHash1)
	if errs := processReceipt1.Errs; len(errs) > 0 {
		t.Fatalf("expected no errors when processing the commands; got %v", errs)
	}

	// Type assert the Result field to msg.ProcessCommandsReply
	processCommandsReply1, ok := processReceipt1.Result.(msg.ProcessCommandsReply)
	if !ok {
		t.Fatalf("expected processReceipt.Result to be of type msg.ProcessCommandsReply; got %T", processReceipt1.Result)
	}

	assert.Equal(t, expectedOut1, processCommandsReply1.Result)
	tf.DoTick()

	// Process Commands2
	processTxHash2 := tf.AddTransaction(getProcessMsgID(t, tf.World), msg.ProcessCommandsMsg{
		PlayerName: playerName,
		Tokens:     tokens2,
	})
	tf.DoTick()

	// Make sure process was successful
	processReceipt2 := getReceiptFromPastTick(t, tf.World, processTxHash2)
	if errs := processReceipt2.Errs; len(errs) > 0 {
		t.Fatalf("expected no errors when processing the commands; got %v", errs)
	}

	// Type assert the Result field to msg.ProcessCommandsReply
	processCommandsReply2, ok := processReceipt2.Result.(msg.ProcessCommandsReply)
	if !ok {
		t.Fatalf("expected processReceipt.Result to be of type msg.ProcessCommandsReply; got %T", processReceipt2.Result)
	}

	assert.Equal(t, expectedOut2, processCommandsReply2.Result)
	tf.DoTick()

	// Process Commands2
	processTxHash3 := tf.AddTransaction(getProcessMsgID(t, tf.World), msg.ProcessCommandsMsg{
		PlayerName: playerName,
		Tokens:     tokens3,
	})
	tf.DoTick()

	// Make sure process was successful
	processReceipt3 := getReceiptFromPastTick(t, tf.World, processTxHash3)
	if errs := processReceipt3.Errs; len(errs) > 0 {
		t.Fatalf("expected no errors when processing the commands; got %v", errs)
	}

	// Type assert the Result field to msg.ProcessCommandsReply
	processCommandsReply3, ok := processReceipt3.Result.(msg.ProcessCommandsReply)
	if !ok {
		t.Fatalf("expected processReceipt.Result to be of type msg.ProcessCommandsReply; got %T", processReceipt3.Result)
	}

	assert.Equal(t, expectedOut3, processCommandsReply3.Result)
	tf.DoTick()
}

func TestThrowingDynamiteAtBoulder(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)
	tf.DoTick()

	const playerName = "Hueyu"
	const roomSpawn = 5
	var tokens1 = []string{"take", "dynamite"}
	var expectedOut1 string = ("You picked up a Dynamite.")
	var tokens2 = []string{"go", "up"}
	var expectedOut2 string = ("You are standing in a barn")
	var tokens3 = []string{"go", "south"}
	var expectedOut3 string = ("You are standing on a windswept plain")
	var tokens4 = []string{"go", "east"}
	var expectedOut4 string = ("You are standing in a high mountain pass")
	var tokens5 = []string{"throw", "petrol", "at", "the", "boulder"}
	var expectedOut5 string = ("You Throw the Petrol at the Boulder. The dynamite lands at the boulder while you keep running and in just a second KBOOOM!." +
		" The boulder, blasted to pieces reveals a path to a new adventure. This is the end of ACT I, if you go East you will return to the plain")

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		PlayersName: playerName,
		RoomID:      roomSpawn,
	})
	tf.DoTick()

	// Process Commands
	processTxHash1 := tf.AddTransaction(getProcessMsgID(t, tf.World), msg.ProcessCommandsMsg{
		PlayerName: playerName,
		Tokens:     tokens1,
	})
	tf.DoTick()

	// Make sure process was successful
	processReceipt1 := getReceiptFromPastTick(t, tf.World, processTxHash1)
	if errs := processReceipt1.Errs; len(errs) > 0 {
		t.Fatalf("expected no errors when processing the commands; got %v", errs)
	}

	// Type assert the Result field to msg.ProcessCommandsReply
	processCommandsReply1, ok := processReceipt1.Result.(msg.ProcessCommandsReply)
	if !ok {
		t.Fatalf("expected processReceipt.Result to be of type msg.ProcessCommandsReply; got %T", processReceipt1.Result)
	}

	assert.Equal(t, expectedOut1, processCommandsReply1.Result)
	tf.DoTick()

	// Process Commands2
	processTxHash2 := tf.AddTransaction(getProcessMsgID(t, tf.World), msg.ProcessCommandsMsg{
		PlayerName: playerName,
		Tokens:     tokens2,
	})
	tf.DoTick()

	// Make sure process was successful
	processReceipt2 := getReceiptFromPastTick(t, tf.World, processTxHash2)
	if errs := processReceipt2.Errs; len(errs) > 0 {
		t.Fatalf("expected no errors when processing the commands; got %v", errs)
	}

	// Type assert the Result field to msg.ProcessCommandsReply
	processCommandsReply2, ok := processReceipt2.Result.(msg.ProcessCommandsReply)
	if !ok {
		t.Fatalf("expected processReceipt.Result to be of type msg.ProcessCommandsReply; got %T", processReceipt2.Result)
	}

	assert.Equal(t, expectedOut2, processCommandsReply2.Result)
	tf.DoTick()

	// Process Commands3
	processTxHash3 := tf.AddTransaction(getProcessMsgID(t, tf.World), msg.ProcessCommandsMsg{
		PlayerName: playerName,
		Tokens:     tokens3,
	})
	tf.DoTick()

	// Make sure process was successful
	processReceipt3 := getReceiptFromPastTick(t, tf.World, processTxHash3)
	if errs := processReceipt3.Errs; len(errs) > 0 {
		t.Fatalf("expected no errors when processing the commands; got %v", errs)
	}

	// Type assert the Result field to msg.ProcessCommandsReply
	processCommandsReply3, ok := processReceipt3.Result.(msg.ProcessCommandsReply)
	if !ok {
		t.Fatalf("expected processReceipt.Result to be of type msg.ProcessCommandsReply; got %T", processReceipt3.Result)
	}

	assert.Equal(t, expectedOut3, processCommandsReply3.Result)
	tf.DoTick()

	// Process Commands4
	processTxHash4 := tf.AddTransaction(getProcessMsgID(t, tf.World), msg.ProcessCommandsMsg{
		PlayerName: playerName,
		Tokens:     tokens4,
	})
	tf.DoTick()

	// Make sure process was successful
	processReceipt4 := getReceiptFromPastTick(t, tf.World, processTxHash4)
	if errs := processReceipt4.Errs; len(errs) > 0 {
		t.Fatalf("expected no errors when processing the commands; got %v", errs)
	}

	// Type assert the Result field to msg.ProcessCommandsReply
	processCommandsReply4, ok := processReceipt4.Result.(msg.ProcessCommandsReply)
	if !ok {
		t.Fatalf("expected processReceipt.Result to be of type msg.ProcessCommandsReply; got %T", processReceipt4.Result)
	}

	assert.Equal(t, expectedOut4, processCommandsReply4.Result)
	tf.DoTick()

	// Process Commands4
	processTxHash5 := tf.AddTransaction(getProcessMsgID(t, tf.World), msg.ProcessCommandsMsg{
		PlayerName: playerName,
		Tokens:     tokens5,
	})
	tf.DoTick()

	// Make sure process was successful
	processReceipt5 := getReceiptFromPastTick(t, tf.World, processTxHash5)
	if errs := processReceipt5.Errs; len(errs) > 0 {
		t.Fatalf("expected no errors when processing the commands; got %v", errs)
	}

	// Type assert the Result field to msg.ProcessCommandsReply
	processCommandsReply5, ok := processReceipt5.Result.(msg.ProcessCommandsReply)
	if !ok {
		t.Fatalf("expected processReceipt.Result to be of type msg.ProcessCommandsReply; got %T", processReceipt5.Result)
	}

	assert.Equal(t, expectedOut5, processCommandsReply5.Result)
	tf.DoTick()
}

// #endregion REST ROOMS ACT I Test

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
