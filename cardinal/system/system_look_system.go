package system

import (
	"fmt"

	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/component"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/types"
)

func LookSystem(world cardinal.WorldContext) error {

	return nil
}

func Stuff(tokens []string, curRmId uint32, playerId uint32, world cardinal.WorldContext) uint8 {
	world.Logger().Debug().Msgf("---->SEE T:%s, R:%d\n", tokens[0], curRmId)
	vrb := ts.GetActionType(tokens[0])
	var gObj enums.GrammarType
	var err uint8

	// we know it is an action because the commandProcessors has pre-parsed for us
	// so we dont need to test for a garbage vrb token
	if vrb == enums.ActionTypeLook {
		world.Logger().Debug().Msgf("---->LK RM:%d\n", curRmId)

		if len(tokens) > 1 {
			gObj = ts.GetGrammarType(tokens[len(tokens)-1])
			if gObj != enums.GrammarTypeAdverb {
				err := lookAround(curRmId, playerId, world)
				world.Logger().Debug().Msgf("-->_LA:%d", err)
				return err
			}
		} else {
			err := lookAround(curRmId, playerId, world)
			world.Logger().Debug().Msgf("-->_LOOK:%d", err)
			return err
		}
	} else if vrb == enums.ActionTypeDescribe || vrb == enums.ActionTypeLook {
		world.Logger().Debug().Msgf("---->DESC\n")
	}
	world.Logger().Debug().Msgf("---->_ERR:%d", err)
	return 0

}

func lookAround(rId uint32, playerId uint32, world cardinal.WorldContext) uint8 {
	output := genDescText(playerId, rId, world)
	world.Logger().Debug().Msgf("ROOM DESCRIPTION IS: %s", output)
	return 0
}

func genDescText(playerId uint32, id uint32, world cardinal.WorldContext) string {

	desc := "You are standing "
	rID := types.EntityID(id)
	room, err := getRoom(rID, world)
	if err != nil {
		world.Logger().Error().Msgf("Error2 getting Room Component: %v", err)
	}

	if room.RoomType == enums.RoomTypePlain {
		desc += fmt.Sprintf("on %s\n", room.Description)
	} else {
		desc += fmt.Sprintf("in %s\n", room.Description)
	}
	return desc
}

func getRoom(rID types.EntityID, world cardinal.WorldContext) (component.Room, error) {
	var exisingRoom component.Room
	err := cardinal.NewSearch().Entity(filter.Exact(filter.Component[component.Room]())).
		Each(world, func(id types.EntityID) bool {
			room, err := cardinal.GetComponent[component.Room](world, rID)
			if err != nil {
				world.Logger().Error().Msgf("Error getting Room Component: %v", err)
				return true
			}

			if room.ID == uint32(rID) {
				exisingRoom = *room
				return false
			}

			return true
		})

	return exisingRoom, err
}
