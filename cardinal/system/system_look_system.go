package system

import (
	"fmt"
	"strings"

	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/component"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/types"
)

func LookSystem(world cardinal.WorldContext) error {

	return nil
}

func Stuff(tokens []string, curRmId uint32, playerId uint32, ts *TokeniserSystem, world cardinal.WorldContext) (string, uint8) {
	world.Logger().Debug().Msgf("LS - S: ---->SEE T:%s, R:%d\n", tokens[0], curRmId)
	vrb := ts.GetActionType(tokens[0])
	var gObj enums.GrammarType
	var err uint8
	var output string

	// we know it is an action because the commandProcessors has pre-parsed for us
	// so we dont need to test for a garbage vrb token
	if vrb == enums.ActionTypeLook {
		world.Logger().Debug().Msgf("LS - S: ---->LK RM:%d\n", curRmId)

		if len(tokens) > 1 {
			gObj = ts.GetGrammarType(tokens[len(tokens)-1])
			if gObj != enums.GrammarTypeAdverb {
				output, err := LookAround(curRmId, playerId, ts, world)
				world.Logger().Debug().Msgf("LS - S: ---->_LA:%d", err)
				return output, err
			}
		} else {
			output, err := LookAround(curRmId, playerId, ts, world)
			world.Logger().Debug().Msgf("LS - S: ---->_LOOK:%d", err)
			return output, err
		}
	} else if vrb == enums.ActionTypeDescribe || vrb == enums.ActionTypeLook {
		world.Logger().Debug().Msgf("LS - S: ---->DESC")
	}
	world.Logger().Debug().Msgf("LS - S: ---->_ERR:%d", err)
	return output, 0
}

func LookAround(rId uint32, playerId uint32, ts *TokeniserSystem, world cardinal.WorldContext) (string, uint8) {
	output := GenDescText(playerId, rId, ts, world)
	output += GenRoomTxt(rId, playerId, world)
	world.Logger().Debug().Msgf("LS - LA: Room description is: %s", output)
	return output, 0
}

// Generates the description on that will be shown
func GenDescText(playerID uint32, id uint32, ts *TokeniserSystem, world cardinal.WorldContext) string {

	desc := "You are standing "
	rID := types.EntityID(id)
	room, err := GetRoom(rID, world)
	if err != nil {
		world.Logger().Error().Msgf("LS - GDT: Error getting Room Component: %v", err)
	}

	if room.RoomType == enums.RoomTypePlain {
		desc += fmt.Sprintf("on %s", room.Description)
	} else {
		desc += fmt.Sprintf("in %s", room.Description)
	}

	return desc
}

func GenRoomTxt(roomID uint32, playerID uint32, world cardinal.WorldContext) string {
	var desc string

	rID := types.EntityID(roomID)
	room, err := GetRoom(rID, world)
	if err != nil {
		world.Logger().Error().Msgf("LS - GRT: Error getting Room Component: %v", err)
	}

	desc += room.RoomTxt
	desc += ObjectDescription(room, world)
	desc += DirObjectDescription(room, ts, world)
	desc += GetPlayersPresence(room, playerID, world)

	return desc
}

// Gets the room
func GetRoom(rID types.EntityID, world cardinal.WorldContext) (component.Room, error) {
	var exisingRoom component.Room
	err := cardinal.NewSearch().Entity(filter.Exact(filter.Component[component.Room]())).
		Each(world, func(id types.EntityID) bool {
			room, err := cardinal.GetComponent[component.Room](world, rID)
			if err != nil {
				world.Logger().Error().Msgf("LS - GR: Error getting Room Component: %v", err)
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

// Gets the Objects descriptions that exists on the room
func ObjectDescription(room component.Room, world cardinal.WorldContext) string {
	var descriptions []string
	isFirst := true

	for _, lookingObjects := range room.Objects {
		if lookingObjects.ObjectID != 0 {
			object := room.Objects[int(lookingObjects.ObjectID)]
			var description string
			if isFirst {
				description = ". You see " + fmt.Sprintf(object.Description)
				isFirst = false
			} else {
				description = "and " + fmt.Sprintf(object.Description)
			}
			world.Logger().Debug().Msgf("LS - DOD: Descriptions for dirObject with ID: %d is: %v", lookingObjects.ObjectID, description)
			descriptions = append(descriptions, description)
		}
	}

	return strings.Join(descriptions, " ")
}

// Gets the DirectionalObjects descriptions that exists on the room
func DirObjectDescription(room component.Room, ts *TokeniserSystem, world cardinal.WorldContext) string {
	var descriptions []string
	isFirst := true

	for _, lookingDirObject := range room.DirObjs {
		if lookingDirObject.ObjectID != 0 {
			dirObject := room.DirObjs[int(lookingDirObject.ObjectID)]
			var description string
			if isFirst {
				description = ". There is a " + fmt.Sprintf(dirObject.Description) +
					GenMaterialDesc(dirObject.MaterialType.String(), dirObject.ObjectType, ts) +
					"to the" + " " + dirObject.DirType.String()
				isFirst = false
			} else {
				description = "and there is a " + fmt.Sprintf(dirObject.Description) +
					GenMaterialDesc(dirObject.MaterialType.String(), dirObject.ObjectType, ts) +
					"to the" + " " + dirObject.DirType.String()
			}
			world.Logger().Debug().Msgf("LS - DOD: Descriptions for dirObject with ID: %d is: %v", lookingDirObject.ObjectID, description)
			descriptions = append(descriptions, description)
		}
	}

	return strings.Join(descriptions, " ")
}

func GenMaterialDesc(material string, dirObj enums.ObjectType, ts *TokeniserSystem) string {
	var description string
	if dirObj == enums.ObjectTypePath || dirObj == enums.ObjectTypeTrail {
		description = " made mainly from" + " " + ts.GetRevMaterialType(material).String() + " "

	} else {
		description = " " + ts.GetRevMaterialType(material).String() + " "
	}
	return description
}

func GetPlayersPresence(room component.Room, playerID uint32, world cardinal.WorldContext) string {
	var descriptions []string
	isFirst := true
	for _, lookingPlayer := range room.Players {
		if lookingPlayer.PlayerID != playerID && lookingPlayer.PlayerID != 0 {
			player := room.Players[int(lookingPlayer.PlayerID)]
			var description string
			if isFirst {
				description = ". In this room is " + player.PlayerName
				isFirst = false
			} else {
				description = player.PlayerName
			}
			world.Logger().Debug().Msgf("LS - GPP: Players found in the room are: %s", player.PlayerName)
			descriptions = append(descriptions, description)
		}
	}

	if len(descriptions) == 0 {
		return ". There is no other poor soul here apart from you."
	}

	// Handle proper punctuation for multiple players
	if len(descriptions) == 1 {
		return descriptions[0]
	}

	return descriptions[0] + ", and " + strings.Join(descriptions[1:], ", ")
}
