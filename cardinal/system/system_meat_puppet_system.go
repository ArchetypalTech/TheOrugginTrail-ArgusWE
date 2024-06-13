package system

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/message"
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
		func(messageData message.TxData[msg.ProcessCommandsMsg]) (msg.ProcessCommandsReply, error) {
			playerEntity, err := findExistingPlayer(world, messageData.Msg.PlayerName)
			if err != nil {
				if isDevelopmentMode() {
					logger.Errorf("\033[31mError searching for Player entity: %v\033[0m", err)
				}

				return msg.ProcessCommandsReply{
					Success: false,
					Message: fmt.Sprintf("Error searching for Player entity: %v", err),
				}, err
			}

			player, err := getPlayerEntity(world, playerEntity)
			if err != nil {
				if isDevelopmentMode() {
					logger.Errorf("\033[31mError getting Player: %v\033[0m", err)
				}

				return msg.ProcessCommandsReply{
					Success: false,
					Message: fmt.Sprintf("Error getting Player: %v", err),
				}, err
			}

			// variables to use
			pID := player.PlayerID
			// rID := player.RoomID ---> Not used YET
			var tokens = messageData.Msg.Tokens
			if isDevelopmentMode() {
				logger.Infof("\033[33mTokens are: %v\033[0m", tokens)
			}
			var er uint8
			var move bool
			//var nxt uint32 ---> Not used YET

			ts = NewTokeniserSystem()

			if uint8(len(tokens)) > constants.MAX_TOK { // Check length of the tokens againts the max number.
				er = constants.ER_PR_TK_CX
			}

			var tok1 string
			tok1 = tokens[0]
			if isDevelopmentMode() {
				logger.Debugf("\033[35m---->CMD: %s\033[0m", tok1)
			}
			tokD := ts.GetDirectionType(tok1)

			if tokD != enums.DirectionTypeNone {
				move = true
				if isDevelopmentMode() {
					logger.Infof("\033[35m---->1-HERE GOES GET NEXT ROOM - DIRECTION SYSTEM\033[0m")
				}
				// HERE GOES GET NEXT ROOM - DIRECTION SYSTEM
			} else if ts.GetActionType(tok1) != enums.ActionTypeNone {
				if uint8(len(tokens)) >= constants.MIN_TOK {
					if isDevelopmentMode() {
						logger.Debugf("\033[35m---->tok.len %d\033[0m", len(tokens))
					}

					if ts.GetActionType(tok1) == enums.ActionTypeGo {
						// GO: form
						move = true
						// HERE GOES GET NEXT ROOM - DIRECTION SYSTEM
						if isDevelopmentMode() {
							logger.Infof("\033[35m---->2-HERE GOES GET NEXT ROOM - DIRECTION SYSTEM\033[0m")
						}
					} else {
						// VERB: form
						er = handleVerb(tokens, pID)
						move = false
					}

				} else {
					er = handleAlias(tokens, pID)
					move = false
				}
			} else {
				er = constants.ER_PR_NOP
			}

			// we have gone through the TOKENS, give err feedback if needed
			if er != 0 {
				if isDevelopmentMode() {
					logger.Errorf("\033[31m---->PCR_ERR: %v:\033[0m", er)
				}
				var errMsg string
				errMsg = insultMeat(er, "")
				// HERE GOES OUTPUT SET
				return msg.ProcessCommandsReply{
					Success: false,
					Message: fmt.Sprintf("%v", errMsg),
				}, err
			} else {
				// either a do something or move rooms command
				if move {
					// Here Goes Enter Room
					if isDevelopmentMode() {
						logger.Infof("\033[35m---->GOING TO ROOM\033[0m")
					}
				} else {
					// hit look libs_ perhaps?
					if isDevelopmentMode() {
						logger.Infof("\033[35m---->hit look libs_ perhaps?\033[0m")
					}
				}
			}

			if isDevelopmentMode() {
				logger.Infof("\033[32mProcessing tokens completed\033[0m")
			}

			ts.FishTokens(tokens)

			return msg.ProcessCommandsReply{
				Success: true,
				Message: "Processing tokens completed",
			}, nil

		},
	)
}

func getPlayerEntity(world cardinal.WorldContext, pEID types.EntityID) (component.Player, error) {
	var exisingPlayer component.Player
	err := cardinal.NewSearch().Entity(filter.Exact(filter.Component[component.Player]())).
		Each(world, func(id types.EntityID) bool {
			player, err := cardinal.GetComponent[component.Player](world, id)
			if err != nil {
				if isDevelopmentMode() {
					logger.Errorf("\033[31mError getting Player Component: %v\033[0m", err)
				}
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

func handleAlias(tokens []string, playerID uint32) (err uint8) {
	vrb := ts.GetActionType(tokens[0])
	var e uint8
	if vrb == enums.ActionTypeInventory {
		// HERE GOES INVENTORY FROM INVENTORY SYSTEM
		if isDevelopmentMode() {
			logger.Infof("\033[35m---->HANDLE ALIAS: NOW SHOULD BE GOING TO INVENTORY FROM INVENTORY SYSTEM\033[0m")
		}
		e = 0
	} else if vrb == enums.ActionTypeLook {
		// HERE GOES STUFF FROM LOOK SYSTEM
		if isDevelopmentMode() {
			logger.Infof("\033[35m---->HANDLE ALIAS: NOW SHOULD BE GOING TO STUFF FROM LOOK SYSTEM\033[0m")
		}
		e = 0
	}
	return e
}

func handleVerb(tokens []string, playerID uint32) (err uint8) {
	vrb := ts.GetActionType(tokens[0])
	var e uint8
	var resultStr string

	//cmdData := ts.FishTokens(tokens) ---> Not used YET
	if vrb == enums.ActionTypeLook || vrb == enums.ActionTypeDescribe {
		if isDevelopmentMode() {
			logger.Infof("\033[35m---->HANDLE VERB: NOW SHOULD BE GOING TO STUFF FROM LOOK SYSTEM\033[0m")
		}
		e = 0
	} else if vrb == enums.ActionTypeTake {
		if isDevelopmentMode() {
			logger.Infof("\033[35m---->HANDLE VERB: NOW SHOULD BE GOING TO TAKE FROM INVENTORY SYSTEM\033[0m")
		}
		e = 0
	} else if vrb == enums.ActionTypeDrop {
		if isDevelopmentMode() {
			logger.Infof("\033[35m---->HANDLE VERB: NOW SHOULD BE GOING TO DROP FROM INVENTORY SYSTEM\033[0m")
		}
		e = 0
	} else {
		if isDevelopmentMode() {
			logger.Infof("\033[35m---->HANDLE VERB: NOW SHOULD BE GOING TO ACT FROM ACTION SYSTEM\033[0m")
		}
		e, resultStr = 0, "testing"
		if isDevelopmentMode() {
			logger.Infof("\033[35m---->HANDLE VERB:resultStr: %s\033[0m", resultStr)
		}
	}
	return e
}

func insultMeat(cErr uint8, badCmd string) string {
	var eMsg string
	if cErr == constants.ER_PR_TK_CX {
		eMsg = "WTF, slow down cowboy, your gonna hurt yourself"
	} else if cErr == constants.ER_PR_NOP || cErr == constants.ER_PR_TK_C1 {
		eMsg = "Nope, gibberish\n" +
			"Stop breathing with your mouth."
	} else if cErr == constants.ER_PR_ND || cErr == constants.ER_DR_ND {
		eMsg = "Go where pilgrim?"
	} else if cErr == constants.ER_DR_NOP {
		eMsg = "Go " + badCmd + " is nowhere I know of bellend"
	} else if cErr == constants.GO_NO_EXIT {
		eMsg = "Can't go that away " + badCmd
	}
	return eMsg
}

func ProcessCommandsTokens2(Tokens []string, PlayerID uint32) error {
	pID := PlayerID
	// rID := player.RoomID ---> Not used YET
	tokens := Tokens
	var er uint8
	var move bool
	//var nxt uint32 ---> Not used YET

	if uint8(len(tokens)) > constants.MAX_TOK {
		er = constants.ER_PR_TK_CX
	}

	var tok1 string
	tok1 = tokens[0]
	if isDevelopmentMode() {
		logger.Debugf("\033[35m---->CMD: %s\033[0m", tok1)
	}
	tokD := ts.GetDirectionType(tok1)

	if tokD != enums.DirectionTypeNone {
		move = true
		// HERE GOES GET NEXT ROOM - DIRECTION SYSTEM
	} else if ts.GetActionType(tok1) != enums.ActionTypeNone {
		if uint8(len(tokens)) >= constants.MIN_TOK {
			if isDevelopmentMode() {
				logger.Debugf("\033[35m---->tok.len %d\033[0m", len(tokens))
			}

			if ts.GetActionType(tok1) == enums.ActionTypeGo {
				// GO: form
				move = true
				// HERE GOES GET NEXT ROOM - DIRECTION SYSTEM
			} else {
				// VERB: form
				er = handleVerb(tokens, pID)
				move = false
			}

		} else {
			er = handleAlias(tokens, pID)
			move = false
		}
	} else {
		er = constants.ER_PR_NOP
	}

	// we have gone through the TOKENS, give err feedback if needed
	if er != 0 {
		if isDevelopmentMode() {
			logger.Errorf("\033[31m---->PCR_ERR: %v:\033[0m", er)
		}
		var errMsg string
		errMsg = insultMeat(er, "")
		// HERE GOES OUTPUT SET
		if isDevelopmentMode() {
			logger.Debugf("\033[35m---->err message is: %s\033[0m", errMsg)
		}
		return nil
	} else {
		// either a do something or move rooms command
		if move {
			// Here Goes Enter Room
			if isDevelopmentMode() {
				logger.Infof("\033[35m---->GOING TO ROOM\033[0m")
			}
		} else {
			// hit look libs_ perhaps?
			if isDevelopmentMode() {
				logger.Infof("\033[35m---->hit look libs_ perhaps?\033[0m")
			}
		}
	}

	if isDevelopmentMode() {
		logger.Infof("\033[32mProcessing tokens completed\033[0m")
	}

	return nil

}
