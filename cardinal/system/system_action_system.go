package system

import (
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/component"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/constants"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"
	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/types"
)

func ActionSystem(world cardinal.WorldContext) error {

	return nil
}

func Act(cmdData component.VerbData, roomID uint32, playerID uint32, ts *TokeniserSystem, world cardinal.WorldContext) (string, uint8) {
	var err uint8
	var bitCnt uint8
	var actResponseStr string
	var actionText string
	var bitTxt string
	var finalTxt string

	world.Logger().Debug().Msgf("AS - ACT: Verb data - vrb: %v, DObj: %v, IObj: %v, ERR:%v", cmdData.Verb, cmdData.DirectObject, cmdData.IndirectObject, cmdData.ErrCode)

	room, er := GetRoom(types.EntityID(roomID), world)
	if er != nil {
		world.Logger().Error().Msgf("AS - ACT: Error getting Room Component: %v", err)
	}
	world.Logger().Debug().Msgf("AS - ACT: RoomID is t: %v", room.ID)

	objects := FetchObjsForType(cmdData.Verb, room, playerID, ts, world)
	world.Logger().Debug().Msgf("AS - ACT: Number of objects got: %v", len(objects))

	dirObjects := FetchDirObjsForType(cmdData.Verb, room, ts, world)
	world.Logger().Debug().Msgf("AS - ACT: Number of dirObjects got: %v", len(dirObjects))

	if len(objects) > 0 {
		if cmdData.IndirectObject == enums.ObjectTypeNone {
			actionText, err = HandleBaseAction(cmdData, objects, world)
		}
	} else {
		actionText = ""
		err = 135
	}
	world.Logger().Debug().Msgf("AS - ACT: PRE-PRE-Final error is: %d", err)

	if len(dirObjects) > 0 {
		if cmdData.IndirectObject != enums.ObjectTypeNone {
			bitTxt, err, bitCnt = SetActionBits(cmdData, dirObjects, true, world)
		}
	}
	world.Logger().Debug().Msgf("AS - ACT: PRE-Final error is: %d", err)

	if err == 0 && bitCnt >= 0 {
		actResponseStr = GetResponseStr(cmdData, ts, world)
	}

	finalTxt = actResponseStr + actionText + bitTxt
	world.Logger().Debug().Msgf("AS - ACT: Final text is: %s, error sent is: %d", finalTxt, err)

	return finalTxt, err

}

func HandleBaseAction(cmd component.VerbData, objects []component.Object, world cardinal.WorldContext) (string, uint8) {
	var response string
	var err uint8

	// Access the singleton ActionStore
	actionStore := component.GetActionStore(world)

	world.Logger().Debug().Msgf("AS - HBA: Verb data - vrb: %v, DObj: %v, IObj: %v, ERR:%v", cmd.Verb, cmd.DirectObject, cmd.IndirectObject, cmd.ErrCode)
	if cmd.DirectObject != enums.ObjectTypeNone && cmd.IndirectObject == enums.ObjectTypeNone {
		for _, object := range objects {
			world.Logger().Debug().Msgf("AS - HBA: Processing object with ID: %d", object.ObjectID)
			if object.ObjectType == cmd.DirectObject {
				for _, actionID := range object.ObjectActionIDs {
					if actionID != 0 {
						world.Logger().Debug().Msgf("AS - HBA: Processing action ID: %d for object ID: %d", actionID, object.ObjectID)
						action, found := actionStore.Get(actionID)
						if !found {
							world.Logger().Error().Msgf("AS - HBA: Action with ID %d not found in ActionStore", actionID)
							continue
						}

						if actionID == action.ID {
							world.Logger().Debug().Msgf("AS - HBA: Action with ID %d is equal to action.ID:: %d", actionID, action.ID)
						} else {
							world.Logger().Debug().Msgf("AS - HBA: Action with ID %d is NOT equal to action.ID:: %d", actionID, action.ID)
						}

						if action.ActionType == cmd.Verb {
							response = action.DBitTxt + "."
							err = 0
							world.Logger().Debug().Msgf("AS - HBA: Action with ID %d has the following text: %s", actionID, response)
							break
						} else {
							world.Logger().Debug().Msgf("AS - HBA: Action with ID %d has actionType: %s (expected: %s)", actionID, action.ActionType, cmd.Verb)
						}
					}
				}
			} else {
				response = ""
				err = 135
			}

		}

	} else {
		world.Logger().Debug().Msgf("AS - HBA: DirectObject or IndirectObject conditions not met.")
	}

	world.Logger().Debug().Msgf("AS - HBA: Response sent: %s, err: %d", response, err)
	return response, err
}

// Gets the description string for the state changes on the objectss tree
func GetResponseStr(cmd component.VerbData, ts *TokeniserSystem, world cardinal.WorldContext) string {
	world.Logger().Debug().Msg("----->AS - GRS: Get Response String")
	var response string = "You"
	var verb string = ts.GetActionType(cmd.Verb.String()).String()
	var dObj string = ts.GetObjectType(cmd.DirectObject.String()).String()
	var iObj string = ts.GetObjectType(cmd.IndirectObject.String()).String()
	response += " " + verb
	if cmd.DirectObject != enums.ObjectTypeNone {
		response += " " + "the" + " " + dObj
	}

	if cmd.IndirectObject != enums.ObjectTypeNone {
		response += " " + "at the" + " " + iObj + "."
	} else {
		response += "."
	}
	world.Logger().Debug().Msgf("----->AS - GRS: Response is: %s", response)
	return response
}

// FetchObjsForType fetches objects for a given type and action type within a room.
func FetchObjsForType(actType enums.ActionType, room component.Room, playerID uint32, ts *TokeniserSystem, world cardinal.WorldContext) []component.Object {
	var matchedObjects []component.Object

	// Access the singleton ActionStore
	actionStore := component.GetActionStore(world)

	if len(room.Objects) > 0 {
		for _, object := range room.Objects {
			if object.ObjectID != 0 {
				world.Logger().Debug().Msgf("AS - FOBJ: ObjectGot has ID: %v", object.ObjectID)

				for _, actionID := range object.ObjectActionIDs {
					if actionID != 0 {
						action, found := actionStore.Get(actionID)
						if !found {
							world.Logger().Error().Msgf("AS - FOBJ: Action with ID %d not found in ActionStore", actionID)
							continue
						}
						world.Logger().Debug().Msgf("AS - FOBJ: Action description: %s is for Object with ID: %v", action.DBitTxt, object.ObjectID)

						if action.ActionType == enums.ActionTypeNone {
							world.Logger().Error().Msgf("AS - FOBJ: Action with ID %d is of type: %v", actionID, action.ActionType)
							break
						}

						responses := ts.GetResponseForVerb(action.ActionType)

						if len(responses) > 0 {
							for _, response := range responses {
								if response == actType {
									matchedObjects = append(matchedObjects, object)
									break // Move to the next object after finding a match
								}
							}
						}

					}
				}
			}
		}
	} else {
		matchedObjectsInventory := FetchObjsForTypeInventory(actType, playerID, ts, world)
		matchedObjects = append(matchedObjects, matchedObjectsInventory...)
	}

	return matchedObjects
}

// FetchObjsForTypeInventory fetches objects for a given type and action type within the player inventory.
func FetchObjsForTypeInventory(actType enums.ActionType, playerID uint32, ts *TokeniserSystem, world cardinal.WorldContext) []component.Object {
	var matchedObjects []component.Object

	// Access the singleton ActionStore
	actionStore := component.GetActionStore(world)

	// Get Player
	player, err := GetPlayer(types.EntityID(playerID), world)
	if err != nil {
		world.Logger().Error().Msgf("AS - FOBJ-I: Error getting Player Component: %v for playerID: %d", err, playerID)
	}

	for _, object := range player.Inventory {
		if object.ObjectID != 0 {
			world.Logger().Debug().Msgf("AS - FOBJ-I: ObjectGot has ID: %v", object.ObjectID)

			for _, actionID := range object.ObjectActionIDs {
				if actionID != 0 {
					action, found := actionStore.Get(actionID)
					if !found {
						world.Logger().Error().Msgf("AS - FOBJ-I: Action with ID %d not found in ActionStore", actionID)
						continue
					}
					world.Logger().Debug().Msgf("AS - FOBJ-I: Action description: %s is for Object with ID: %v", action.DBitTxt, object.ObjectID)

					if action.ActionType == enums.ActionTypeNone {
						world.Logger().Error().Msgf("AS - FOBJ-I: Action with ID %d is of type: %v", actionID, action.ActionType)
						break
					}

					responses := ts.GetResponseForVerb(action.ActionType)

					if len(responses) > 0 {
						for _, response := range responses {
							if response == actType {
								matchedObjects = append(matchedObjects, object)
								break // Move to the next object after finding a match
							}
						}
					}
				}
			}
		}
	}

	return matchedObjects
}

// FetchDirObjsForType fetches directional objects for a given type and action type within a room.
func FetchDirObjsForType(actType enums.ActionType, room component.Room, ts *TokeniserSystem, world cardinal.WorldContext) []component.Object {
	var matchedDirObjects []component.Object

	// Access the singleton ActionStore
	actionStore := component.GetActionStore(world)

	for _, dirObject := range room.DirObjs {
		if dirObject.ObjectID != 0 {
			world.Logger().Debug().Msgf("AS - FDIROBJ: ObjectGot has ID: %v", dirObject.ObjectID)

			for _, actionID := range dirObject.ObjectActionIDs {
				if actionID != 0 {
					action, found := actionStore.Get(actionID)
					if !found {
						world.Logger().Error().Msgf("AS - FDIROBJ: Action with ID %d not found in ActionStore", actionID)
						continue
					}
					world.Logger().Debug().Msgf("AS - FDIROBJ: Action description: %s is for Object with ID: %v", action.DBitTxt, dirObject.ObjectID)

					if action.ActionType == enums.ActionTypeNone {
						world.Logger().Error().Msgf("AS - FDIROBJ: Action with ID %d is of type: %v", actionID, action.ActionType)
						break
					}

					responses := ts.GetResponseForVerb(action.ActionType)
					print(len(responses))

					if len(responses) > 0 {
						for _, response := range responses {
							if response == action.ActionType {
								matchedDirObjects = append(matchedDirObjects, dirObject)
								break // Move to the next dirObject after finding a match
							}
						}
					}

				}
			}
		}
	}

	return matchedDirObjects
}

/*
		@notice flip the action bit i.e make it a past participle, broken/smashed/opened
	    @return bitCount, records the number of bite flipped
		@return er, error code. 0 for success or an error code

	    The logic is that we check for an enabled bit and then change the state
	    and then follow any linked actions which allows us to then build puzzle chains
*/
func SetActionBits(cmd component.VerbData, objects []component.Object, isD bool, world cardinal.WorldContext) (string, uint8, uint8) {
	var ct uint32 = uint32(len(objects))
	var bc uint8 = 0
	var dBit_Txt string
	var err uint8 = 0
	var actionResultStr string
	// Accessing the instance of the ActionStore that was created when the game was setup
	actionStore := component.GetActionStore(world)

	world.Logger().Debug().Msgf("AS - SAB: ----->sz: %v", ct)

	// Track processed actions to avoid duplicates
	processedActions := make(map[uint32]bool)

	for _, object := range objects {
		if object.ObjectType == cmd.IndirectObject {
			// set the action bit on the object for the verb
			if isD {
				if cmd.IndirectObject != enums.ObjectTypeNone {
					for _, actionID := range object.ObjectActionIDs {
						if actionID != 0 {
							// Skip already processed actions
							if processedActions[actionID] {
								continue
							}

							action, found := actionStore.Get(actionID)
							if !found {
								world.Logger().Error().Msgf("AS - SAB: Action with ID %d not found in ActionStore", actionID)
								continue
							}

							if action.Enabled {
								bc++
								if action.Revert && action.DBit {
									action.DBit = !action.DBit
								} else if !action.DBit {
									action.DBit = !action.DBit
									actionStore.Set(action.ID, action)
									dBit_Txt = action.DBitTxt
									actionResultStr += dBit_Txt
								}

								linkedActionID := action.AffectsActionID
								world.Logger().Debug().Msgf("AS - SAB: ActionID: %d, links affects actionID: %d", actionID, linkedActionID)
								if linkedActionID != 0 {
									var linkedActions []uint32
									linkedActions = append(linkedActions, linkedActionID)
									FollowLinkedActions(linkedActionID, &linkedActions, world)

									for _, linkedActionData := range linkedActions {
										// Skip already processed linked actions
										if processedActions[linkedActionData] {
											continue
										}

										bc++
										linkedAction, found := actionStore.Get(linkedActionData)
										if !found {
											world.Logger().Error().Msgf("AS - SAB: LinkedAction with ID %d not found in ActionStore", linkedAction.ID)
											break
										}
										world.Logger().Debug().Msgf("AS - SAB: linkedActionID: %d, enabled value is: %v", linkedAction.ID, linkedAction.Enabled)
										world.Logger().Debug().Msgf("AS - SAB: linkedActionID: %d, DBIT value is: %v", linkedAction.ID, linkedAction.DBit)
										linkedAction.Enabled = !linkedAction.Enabled
										linkedAction.DBit = !linkedAction.DBit
										actionResultStr += linkedAction.DBitTxt
										world.Logger().Debug().Msgf("AS - SAB: linkedActionID: %d, NEW enabled value is: %v", linkedAction.ID, linkedAction.Enabled)
										world.Logger().Debug().Msgf("AS - SAB: linkedActionID: %d, NEW DBIT value is: %v", linkedAction.ID, linkedAction.DBit)
										actionStore.Set(linkedAction.ID, linkedAction)

										// Mark linked action as processed
										processedActions[linkedActionData] = true
									}
								}

								// Mark original action as processed
								processedActions[actionID] = true
							}
						}
					}
				}
			} else {
				// handle for objects
				// :TODO
			}
		} else {
			err = constants.ErrActionHandleBadCommand0.Code
		}

	}

	if bc > 0 {
		err = 0
		world.Logger().Debug().Msgf("AS - SAB: returning the following text: %s, the following err: %d, and the following bc: %d", actionResultStr, 0, bc)
		return actionResultStr, err, bc
	} else {
		world.Logger().Debug().Msgf("AS - SAB: returning the following text: %s, the following err: %d, and the following bc: %d", "", constants.ErrActionHandleBadCommand0.Code, bc)
		return "", err, bc
	}
}

// FollowLinkedActions recursively follows linked actions, adding them to the ids slice.
func FollowLinkedActions(top uint32, ids *[]uint32, world cardinal.WorldContext) uint8 {
	actionStore := component.GetActionStore(world)
	action, found := actionStore.Get(top)

	if !found {
		world.Logger().Error().Msgf("AS - FLA: Action with ID %d not found in ActionStore", top)
		return 1 // Return an error code indicating the action was not found
	}

	// Add the current action to the list
	*ids = append(*ids, top)
	world.Logger().Debug().Msgf("AS - FLA: Added action ID %d to linked actions", top)

	// Get the next linked action ID
	nextID := action.AffectsActionID

	if nextID == 0 {
		return 0 // No further linked actions
	}

	world.Logger().Debug().Msgf("AS - FLA: Following link to action ID %d", nextID)
	return FollowLinkedActions(nextID, ids, world)
}
