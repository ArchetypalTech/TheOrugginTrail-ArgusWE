package system

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/component"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/constants"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/msg"
)

// Define a variable to hold the TokeniserSystem instance
var ts *TokeniserSystem

func ProcessCommandsTokens(world cardinal.WorldContext) error {
	return cardinal.EachMessage[msg.ProcessCommandsMsg, msg.ProcessCommandsReply](
		world,
		func(messageData cardinal.TxData[msg.ProcessCommandsMsg]) (msg.ProcessCommandsReply, error) {
			playerEntity, err := findExistingPlayer(world, messageData.Msg.PlayerName)
			if err != nil {
				return msg.ProcessCommandsReply{
					Success: false,
					Message: fmt.Sprintf("Error searching for Player entity: %v", err),
				}, err
			}

			player, err := getPlayerEntity(world, playerEntity)
			if err != nil {
				return msg.ProcessCommandsReply{
					Success: false,
					Message: fmt.Sprintf("Error getting Player: %v", err),
				}, err
			}

			output, move, er := ProcessCommandsTokensLogic(messageData.Msg.Tokens, player, world)

			// we have gone through the TOKENS, give err feedback if needed
			if er != 0 {
				world.Logger().Error().Msgf("---->PCE: PCR_ERR: %v:\033[0m", er)
				var errMsg string
				errMsg = InsultMeat(er, "")
				// HERE GOES OUTPUT SET
				return msg.ProcessCommandsReply{
					Success: false,
					Message: fmt.Sprintf("Error: %v processing the commands: %s", er, errMsg),
					Result:  errMsg,
				}, err
			} else {
				// either a do something or move rooms command
				if move {
					// Here Goes Enter Room
					world.Logger().Debug().Msg("---->GOING TO ROOM")
				} else {
					// hit look libs_ perhaps?
					world.Logger().Debug().Msg("---->hit look libs_ perhaps?")
				}
			}

			world.Logger().Debug().Msg("---->Processing tokens completed")

			verbData := ts.FishTokens(messageData.Msg.Tokens)
			world.Logger().Debug().Msgf("P--->d.dobj:%s iobj:%s vrb:%s", verbData.DirectObject, verbData.IndirectObject, verbData.Verb)
			if verbData.ErrCode != constants.NOERR {
				world.Logger().Error().Msgf("E---err:%s", verbData.ErrCode)
			}

			return msg.ProcessCommandsReply{
				Success: true,
				Message: "Processing tokens completed",
				Result:  output,
			}, nil

		},
	)
}

// Returns the player entity component based on the id
func getPlayerEntity(world cardinal.WorldContext, pEID types.EntityID) (component.Player, error) {
	var exisingPlayer component.Player
	err := cardinal.NewSearch().Entity(filter.Exact(filter.Component[component.Player]())).
		Each(world, func(id types.EntityID) bool {
			player, err := cardinal.GetComponent[component.Player](world, id)
			if err != nil {
				world.Logger().Error().Msgf("GetPlayerEntity: Error getting Player Component: %v", err)
				return true
			}

			if player.PlayerID == uint32(pEID) {
				exisingPlayer = *player
				return false
			}

			return true
		})

	return exisingPlayer, err
}

// Process the Commands tokens, this is the function dedicated to it.
func ProcessCommandsTokensLogic(Tokens []string, Player component.Player, world cardinal.WorldContext) (string, bool, uint8) {
	pID := Player.PlayerID
	rID := Player.RoomID
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
		world.Logger().Debug().Msgf("---->CMD: %s", tok1)
		tokD := ts.GetDirectionType(tok1)

		if tokD != enums.DirectionTypeNone {
			move = true
			// HERE GOES GET NEXT ROOM - DIRECTION SYSTEM
		} else if ts.GetActionType(tok1) != enums.ActionTypeNone {
			if uint8(len(tokens)) >= constants.MIN_TOK {
				world.Logger().Debug().Msgf("---->tok.len %d", len(tokens))
				if ts.GetActionType(tok1) == enums.ActionTypeGo {
					// GO: form
					move = true
					output = "GOING TO NEXT ROOM - DIRECTION SYSTEM - TO BE IMPLEMENTED"
					// HERE GOES GET NEXT ROOM - DIRECTION SYSTEM
				} else {
					// VERB: form
					output, er = HandleVerb(tokens, rID, pID, ts, world)
					move = false
				}

			} else {
				output, er = HandleAlias(tokens, rID, pID, ts, world)
				move = false
			}
		} else {
			er = constants.ErrParserRoutineNOP.Code
			output = constants.ErrParserRoutineNOP.Message
		}
	}

	return output, move, er

}

// handle if the token is an alias
func HandleAlias(tokens []string, roomID uint32, playerID uint32, ts *TokeniserSystem, world cardinal.WorldContext) (output string, err uint8) {
	vrb := ts.GetActionType(tokens[0])
	var resultStr string
	var e uint8
	if vrb == enums.ActionTypeInventory {
		// HERE GOES INVENTORY FROM INVENTORY SYSTEM
		world.Logger().Debug().Msg("---->HANDLE ALIAS: NOW SHOULD BE GOING TO INVENTORY FROM INVENTORY SYSTEM")
		resultStr = "---->HANDLE ALIAS: NOW SHOULD BE GOING TO INVENTORY FROM INVENTORY SYSTEM"
		e = 135 // This just for showing errors
	} else if vrb == enums.ActionTypeLook {
		// HERE GOES STUFF FROM LOOK SYSTEM
		world.Logger().Debug().Msg("--->HANDLE ALIAS: NOW SHOULD BE GOING TO STUFF FROM LOOK SYSTEM")
		resultStr, e = Stuff(tokens, roomID, playerID, ts, world)
	}
	return resultStr, e
}

// Handle if the token is a verb
func HandleVerb(tokens []string, roomID uint32, playerID uint32, ts *TokeniserSystem, world cardinal.WorldContext) (output string, err uint8) {
	vrb := ts.GetActionType(tokens[0])
	var e uint8
	var resultStr string

	//cmdData := ts.FishTokens(tokens) ---> Not used YET
	switch vrb {
	case enums.ActionTypeLook, enums.ActionTypeDescribe:
		world.Logger().Debug().Msg("---->HANDLE VERB: NOW SHOULD BE GOING TO STUFF FROM LOOK SYSTEM")
		resultStr, e = Stuff(tokens, roomID, playerID, ts, world)

	case enums.ActionTypeTake:
		world.Logger().Debug().Msg("---->HANDLE VERB: NOW SHOULD BE GOING TO TAKE FROM INVENTORY SYSTEM")
		resultStr = "---->HANDLE VERB: NOW SHOULD BE GOING TO TAKE FROM INVENTORY SYSTEM"
		e = 0

	case enums.ActionTypeDrop:
		world.Logger().Debug().Msg("---->HANDLE VERB: NOW SHOULD BE GOING TO DROP FROM INVENTORY SYSTEM")
		resultStr = "---->HANDLE VERB: NOW SHOULD BE GOING TO DROP FROM INVENTORY SYSTEM"
		e = 0

	default:
		world.Logger().Debug().Msg("---->HANDLE VERB: NOW SHOULD BE GOING TO ACT FROM ACTION SYSTEM")
		resultStr = "---->HANDLE VERB: NOW SHOULD BE GOING TO ACT FROM ACTION SYSTEM"
		e = 0
		world.Logger().Debug().Msgf("---->HANDLE VERB:resultStr: %s", resultStr)
	}

	return resultStr, e
}

// returns a string to be used on the message variable of the transaction/message
func InsultMeat(cErr uint8, badCmd string) string {
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
		eMsg = "What are you doing?!?!"
	}

	return eMsg
}
