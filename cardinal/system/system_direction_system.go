package system

import (
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/component"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/constants"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"
	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/types"
)

func DirectionSystem(world cardinal.WorldContext) error {

	return nil
}

func CanMove(dirObject component.Object, world cardinal.WorldContext) bool {
	var canMove bool
	// Accessing the instance of the ActionStore that was created when the game was setup
	actionStore := component.NewActionStore()

	for _, actionID := range dirObject.ObjectActionIDs {
		if actionID != 0 {
			action, found := actionStore.Get(actionID)
			if !found {
				world.Logger().Error().Msgf("Action with ID %d not found in ActionStore", actionID)
				continue
			}

			switch action.ActionType {
			case enums.ActionTypeOpen:
				canMove = action.Enabled && action.DBit
				world.Logger().Debug().Msgf("canMove_OPEN: e: %v, d: %v", action.Enabled, action.DBit)
			case enums.ActionTypeLock:
				canMove = action.Enabled && !action.DBit
				world.Logger().Debug().Msgf("canMove_LOCK: e: %v, d: %v", action.Enabled, action.DBit)
			}
		}
	}
	return canMove
}

func DirectionCheck(roomID uint32, direction enums.DirectionType, ts *TokeniserSystem, world cardinal.WorldContext) (bool, component.Object) {
	var exitObject component.Object

	room, err := GetRoom(types.EntityID(roomID), world)
	if err != nil {
		world.Logger().Error().Msgf("Error getting the room for the Direction System at direction check: %v", err)
		return false, component.Object{}
	}

	for _, directionCheck := range room.DirObjs {

		dirObject := room.DirObjs[int(directionCheck.ObjectID)]
		if dirObject.DirType == direction {
			if CanMove(dirObject, world) == true {
				exitObject = dirObject
				return true, exitObject
			}
		}
	}

	return false, component.Object{}
}

func FishDirectionTok(tokens []string, ts *TokeniserSystem) (string, uint8) {
	var tok string
	var err uint8

	if ts.GetDirectionType(tokens[0]) != enums.DirectionTypeNone {
		/* Direction form
		 *
		 * dir = n | e | s | w
		 *
		 */
		tok = tokens[0]
	} else if ts.GetActionType(tokens[0]) != enums.ActionTypeNone {
		/* GO form
		 *
		 * go_cmd = go, [(pp da)], dir | obj
		 * pp = "to";
		 * da = "the";
		 * dir = n | e | s | w
		 */
		if len(tokens) >= 4 {
			/* GO long form
			* go_cmd = go, ("to" "the"), dir|obj
			 */
			tok = tokens[3] // dir | obj
		} else if len(tokens) == 2 {
			/* Go short form
			* go_cmd = go, dir|obj
			 */
			tok = tokens[1] // dir | obj
		}

		if ts.GetDirectionType(tok) != enums.DirectionTypeNone {
			err = 0
		} else {
			err = constants.ErrDirectionRoutineND.Code
		}
	}

	return tok, err
}
