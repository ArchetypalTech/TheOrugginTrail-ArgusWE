package system

import (
	"fmt"
	"testing"

	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/component"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/constants"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"
	"github.com/stretchr/testify/assert"
)

func setup() {
	ts = NewTokeniserSystem()
}

// TestProcessCommandsTokensLogic tests the ProcessCommandsTokensLogic function
func TestProcessCommandsTokensLogic(t *testing.T) {
	tests := []struct {
		name         string
		tokens       []string
		player       component.Player
		expectedOut  string
		expectedMove bool
		expectedErr  uint8
	}{
		{
			name:         "Too many tokens",
			tokens:       []string{"THROW", "THE", "BALL", "TO", "THE", "WINDOW", "LOCK", "THE", "DOOR", "TO", "THE", "PATH", "BURN", "THE", "BOTTLE", "AT", "THE", "STAIRS"},
			player:       component.Player{PlayerID: 6, RoomID: 0},
			expectedOut:  constants.ErrParserRoutineTKCX.Message,
			expectedMove: false,
			expectedErr:  constants.ErrParserRoutineTKCX.Code,
		},
		{
			name:         "Go command",
			tokens:       []string{"GO", "TO", "PATH"},
			player:       component.Player{PlayerID: 6, RoomID: 0},
			expectedOut:  "GOING TO NEXT ROOM - DIRECTION SYSTEM - TO BE IMPLEMENTED",
			expectedMove: true,
			expectedErr:  0,
		},
		{
			name:         "Invalid command",
			tokens:       []string{"SEE", "UP"},
			player:       component.Player{PlayerID: 6, RoomID: 0},
			expectedOut:  constants.ErrParserRoutineNOP.Message,
			expectedMove: false,
			expectedErr:  constants.ErrParserRoutineNOP.Code,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, move, err := ProcessCommandsTokensLogicTest(tt.tokens, tt.player.PlayerID, tt.player.RoomID)

			assert.Equal(t, tt.expectedOut, output)
			assert.Equal(t, tt.expectedMove, move)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

// TestHandleAlias tests the handleAlias function
func TestHandleAlias(t *testing.T) {
	setup()
	tests := []struct {
		name        string
		tokens      []string
		playerID    uint32
		expectedErr uint8
	}{
		{
			name:        "Inventory alias",
			tokens:      []string{"INVENTORY"},
			playerID:    6,
			expectedErr: 135,
		},
		{
			name:        "Look alias",
			tokens:      []string{"LOOK", "UP"},
			playerID:    6,
			expectedErr: 130,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := handleAliasTest(tt.tokens, tt.playerID)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

// TestHandleVerb tests the handleVerb function
func TestHandleVerb(t *testing.T) {
	setup()
	tests := []struct {
		name        string
		tokens      []string
		roomID      uint32
		playerID    uint32
		expectedOut string
		expectedErr uint8
	}{
		{
			name:        "Look verb",
			tokens:      []string{"LOOK", "TO", "THE", "WINDOW"},
			roomID:      0,
			playerID:    6,
			expectedOut: "---->HANDLE VERB: NOW SHOULD BE GOING TO STUFF FROM LOOK SYSTEM",
			expectedErr: 0,
		},
		{
			name:        "Take verb",
			tokens:      []string{"TAKE", "THE", "BALL"},
			roomID:      0,
			playerID:    6,
			expectedOut: "---->HANDLE VERB: NOW SHOULD BE GOING TO TAKE FROM INVENTORY SYSTEM",
			expectedErr: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := handleVerbTest(tt.tokens, tt.roomID, tt.playerID)
			assert.Equal(t, tt.expectedOut, output)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

// TestInsultMeat tests the insultMeat function
func TestInsultMeat(t *testing.T) {
	tests := []struct {
		name        string
		cErr        uint8
		badCmd      string
		expectedMsg string
	}{
		{
			name:        "Too many tokens",
			cErr:        constants.ErrParserRoutineTKCX.Code,
			badCmd:      "",
			expectedMsg: "WTF, slow down cowboy, you're gonna hurt yourself",
		},
		{
			name:        "Gibberish command",
			cErr:        constants.ErrParserRoutineNOP.Code,
			badCmd:      "",
			expectedMsg: "Nope, gibberish. Stop breathing with your mouth.",
		},
		{
			name:        "No direction",
			cErr:        constants.ErrParserRoutineND.Code,
			badCmd:      "north",
			expectedMsg: "Go where pilgrim?",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := insultMeatTest(tt.cErr, tt.badCmd)
			assert.Equal(t, tt.expectedMsg, msg)
		})
	}
}

// Process the Commands tokens, this is the function dedicated to it.
func ProcessCommandsTokensLogicTest(Tokens []string, PlayerID uint32, roomID uint32) (string, bool, uint8) {
	pID := PlayerID
	rID := roomID
	tokens := Tokens
	var er uint8
	var move bool
	var output string
	//var nxt uint32 ---> Not used YET

	// Start a new token system
	ts = NewTokeniserSystem()

	if uint8(len(tokens)) > constants.MAX_TOK {
		er = constants.ErrParserRoutineTKCX.Code
		output = constants.ErrParserRoutineTKCX.Message
	} else {
		var tok1 string
		tok1 = tokens[0]
		tokD := ts.GetDirectionType(tok1)

		if tokD != enums.DirectionTypeNone {
			move = true
			// HERE GOES GET NEXT ROOM - DIRECTION SYSTEM
		} else if ts.GetActionType(tok1) != enums.ActionTypeNone {
			if uint8(len(tokens)) >= constants.MIN_TOK {

				if ts.GetActionType(tok1) == enums.ActionTypeGo {
					// GO: form
					move = true
					output = "GOING TO NEXT ROOM - DIRECTION SYSTEM - TO BE IMPLEMENTED"
					// HERE GOES GET NEXT ROOM - DIRECTION SYSTEM
				} else {
					// VERB: form
					output, er = handleVerbTest(tokens, rID, pID)
					move = false
				}

			} else {
				er = handleAliasTest(tokens, pID)
				move = false
				output = "VERB GOES TO HANDLE ALIAS. TO BE IMPLEMENTED"
			}
		} else {
			er = constants.ErrParserRoutineNOP.Code
			output = constants.ErrParserRoutineNOP.Message
		}
	}

	return output, move, er
}

func handleAliasTest(tokens []string, playerID uint32) (err uint8) {
	vrb := ts.GetActionType(tokens[0])
	//var e uint8
	if vrb == enums.ActionTypeInventory {
		fmt.Printf("---->HANDLE ALIAS: NOW SHOULD BE GOING TO INVENTORY FROM INVENTORY SYSTEM")
		err = 135 // This is just for showing errors
	} else if vrb == enums.ActionTypeLook {
		fmt.Printf("--->HANDLE ALIAS: NOW SHOULD BE GOING TO STUFF FROM LOOK SYSTEM")
		err = 130 // This is just for showing errors
	}
	return err
}

// Handle if the token is a verb
func handleVerbTest(tokens []string, roomID uint32, playerID uint32) (output string, err uint8) {
	vrb := ts.GetActionType(tokens[0])
	var e uint8
	var resultStr string

	switch vrb {
	case enums.ActionTypeLook, enums.ActionTypeDescribe:
		fmt.Printf("---->HANDLE VERB: NOW SHOULD BE GOING TO STUFF FROM LOOK SYSTEM")
		resultStr = ("---->HANDLE VERB: NOW SHOULD BE GOING TO STUFF FROM LOOK SYSTEM")
		e = 0

	case enums.ActionTypeTake:
		fmt.Printf("---->HANDLE VERB: NOW SHOULD BE GOING TO TAKE FROM INVENTORY SYSTEM")
		resultStr = "---->HANDLE VERB: NOW SHOULD BE GOING TO TAKE FROM INVENTORY SYSTEM"
		e = 0

	case enums.ActionTypeDrop:
		fmt.Printf("---->HANDLE VERB: NOW SHOULD BE GOING TO DROP FROM INVENTORY SYSTEM")
		resultStr = "---->HANDLE VERB: NOW SHOULD BE GOING TO DROP FROM INVENTORY SYSTEM"
		e = 0

	default:
		fmt.Printf("---->HANDLE VERB: NOW SHOULD BE GOING TO ACT FROM ACTION SYSTEM")
		resultStr = "---->HANDLE VERB: NOW SHOULD BE GOING TO ACT FROM ACTION SYSTEM"
		e = 0
		fmt.Printf("---->HANDLE VERB:resultStr: %s", resultStr)
	}

	return resultStr, e
}

// returns a string to be used on the message variable of the transaction/message
func insultMeatTest(cErr uint8, badCmd string) string {
	var eMsg string
	switch cErr {
	case constants.ErrParserRoutineTKCX.Code:
		eMsg = "WTF, slow down cowboy, you're gonna hurt yourself"

	case constants.ErrDirectionRoutineNOP.Code, constants.ErrParserRoutineNOP.Code, constants.ErrParserRoutineTKC1.Code, constants.ErrNoObjectsToHandle.Code, constants.ErrBadLookCommand.Code:
		eMsg = "Nope, gibberish. Stop breathing with your mouth."

	case constants.ErrParserRoutineND.Code, constants.ErrDirectionRoutineND.Code:
		eMsg = "Go where pilgrim?"

	case constants.ErrDirectionRoutineNOP.Code:
		eMsg = "Go " + badCmd + " is nowhere I know of bellend"

	case constants.ErrNoExit.Code:
		eMsg = "Can't go that away " + badCmd

	default:
		// Add a default case if needed for handling unexpected cErr values
	}

	return eMsg
}
