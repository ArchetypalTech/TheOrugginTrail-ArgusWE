package system

import (
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/component"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"
	"pkg.world.dev/world-engine/cardinal"

	"github.com/sirupsen/logrus"
)

// Initialize logrus logger
var (
	logger = logrus.New()
)

func initLogger() {
	// Set log level
	logger.SetLevel(logrus.DebugLevel)

	// Define a custom formatter with colors
	formatter := &logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}

	logger.SetFormatter(formatter)
}

// NGameSetupSystem initializes the game setup system.
func NGameSetupSystem(world cardinal.WorldContext) error {

	setup := NewGameSetup(world)
	setup.Init(world)
	initLogger()

	if isDevelopmentMode() {
		// LOG FOR TESTING TxtDefStore entries
		// Log the number of entries and contents of TxtDefStore
		logger.Debugf("\033[34mTxtDefStore initialized with %d entries\033[0m", len(setup.TxtDefStore.TxtDefs))
		for key, value := range setup.TxtDefStore.TxtDefs {
			logger.Debugf("\033[34mTxtDefStore entry - Key: %s, Value: %+v\033[0m", key, value)
		}
	}

	return nil
}

// GameSetup manages the setup of the game world.
type GameSetup struct {
	worldCtx       cardinal.WorldContext
	RoomStore      *component.RoomStore
	DirObjectStore *component.ObjectStore
	ObjectStore    *component.ObjectStore
	ActionStore    *component.ActionStore
	TxtDefStore    *component.TxtDefStore
}

func NewGameSetup(worldCtx cardinal.WorldContext) *GameSetup {
	return &GameSetup{
		worldCtx:       worldCtx,
		RoomStore:      component.NewRoomStore(),
		DirObjectStore: component.NewObjectStore(),
		ObjectStore:    component.NewObjectStore(),
		ActionStore:    component.NewActionStore(),
		TxtDefStore:    component.NewTxtDefStore(),
	}
}

// Init initializes the GameSetup system.
func (s *GameSetup) Init(world cardinal.WorldContext) {
	s.setupWorld(world)
}

// setupWorld initializes the game world by setting up rooms.
func (s *GameSetup) setupWorld(world cardinal.WorldContext) {
	s.setupRooms(world)
	// Add setup functions for objects, actions, etc.
}

// setupRooms sets up various rooms in the game world.
func (s *GameSetup) setupRooms(world cardinal.WorldContext) {
	s.setupPlain(world)
	s.setupBarn(world)
	s.setupMountainPath(world)
}

// setupPlain sets up the plain in the game world.
func (s *GameSetup) setupPlain(world cardinal.WorldContext) {
	world.Logger().Debug().Msg("---->Setting up the plain.....")

	// KPLAIN -> N, E
	open2Barn := s.createAction(enums.ActionTypeOpen, "the door opens with a farty noise\n"+
		"you can actually smell fart", true, true, true, 0, 0)

	plainBarn := [32]uint32{open2Barn}
	dObjs := [32]uint32{s.createDirObject(enums.DirectionTypeNorth, enums.RoomTypePlain,
		enums.ObjectTypePath, enums.MaterialTypeDirt,
		"path", plainBarn, world)}

	open2Path := s.createAction(enums.ActionTypeOpen, "the door opens and a small hinge demon curses you\n"+
		"your nose is really itchy", true, true, true, 0, 0)

	plainPath := [32]uint32{open2Path}
	dObjs[1] = s.createDirObject(enums.DirectionTypeEast, enums.RoomTypeWoodCabin,
		enums.ObjectTypePath, enums.MaterialTypeMud,
		"path", plainPath, world)

	kick := s.createAction(enums.ActionTypeKick, "the ball (such as it is)\n"+
		"bounces feebly\n then rolls into some fresh dog eggs\n"+
		"none the less you briefly feel a little better", true, false, true, 0, 0)

	ballActions := [32]uint32{kick}
	objs := [32]uint32{s.createObject(enums.ObjectTypeFootball, enums.MaterialTypeFlesh,
		"A slightly deflated knock off uefa football,\n"+
			"not quite spherical, it's "+
			"kickable though", "football", ballActions, world)}

	roomID := s.RoomStore.Add(component.Room{
		Description: "a windswept plain",
		RoomType:    enums.RoomTypePlain,
	})

	tidPlain := s._textGuid("a windsept plain")
	s.TxtDefStore.Set(tidPlain, enums.TxtDefTypePlace, "the wind blowing is cold and\n"+
		"bison skulls in piles taller than houses\n"+
		"cover the plains as far as your eye can see\n"+
		"the air tastes of burnt grease and bensons.",
	)

	// Logs to verify the ID's in the setup
	world.Logger().Debug().Msgf("Object IDs for room %d: %v\033[0m", roomID, objs)
	world.Logger().Debug().Msgf("Directional Object IDs for room %d: %v", roomID, dObjs)

	// Inspect Data Passed to Setup Functions
	world.Logger().Debug().Msgf("ata passed to createPlace for room %d: %+v\033[0m", roomID, objs)
	world.Logger().Debug().Msgf("Data passed to createDirectionalObject for room %d: %+v\033[0m", roomID, dObjs)

	s.createPlace(roomID, enums.RoomTypePlain, dObjs, objs, tidPlain, world)

	world.Logger().Debug().Msg("---->Plain setup complete")

}

// setupBarn sets up the barn in the game world.
func (s *GameSetup) setupBarn(world cardinal.WorldContext) {
	world.Logger().Debug().Msg("---->Setting up the barn.....")
	// KBARN -> S
	open2South := s.createAction(enums.ActionTypeOpen, "the door opens\n", true, true, true, 0, 0)
	barnPlain := [32]uint32{open2South}
	dObjs := [32]uint32{s.createDirObject(enums.DirectionTypeSouth, enums.RoomTypePlain,
		enums.ObjectTypeDoor, enums.MaterialTypeWood,
		"door", barnPlain, world)}

	open2Forest := s.createAction(enums.ActionTypeOpen, "the window, glass and frame smashed\n"+
		"falls open", false, false, false, 0, 0)

	smashWindow := s.createAction(enums.ActionTypeBreak, "I love the sound of breaking glass\n"+
		"especially when I'm lonely , the panes and the frame shatter\n"+
		"satisfyingly spreading broken joy on the floor",
		true, false, false, open2Forest, 0)

	windowActions := [32]uint32{open2Forest, smashWindow}
	dObjs[1] = s.createDirObject(enums.DirectionTypeEast, enums.RoomTypeForge,
		enums.ObjectTypeWindow, enums.MaterialTypeWood,
		"window", windowActions, world)

	roomID := s.RoomStore.Add(component.Room{
		Description: "a barn",
		RoomType:    enums.RoomTypeWoodCabin,
	})

	tidBarn := s._textGuid("a barn")
	s.TxtDefStore.Set(tidBarn, enums.TxtDefTypePlace,
		"The place is dusty and full of spiderwebs,\n"+
			"something died in here, possibly your own self\n"+
			"plenty of corners and dark shadows",
	)

	// Logs to verify the ID's in the setup
	world.Logger().Debug().Msgf("Object IDs for room %d: %v\033[0m", roomID, [32]uint32{})
	world.Logger().Debug().Msgf("Directional Object IDs for room %d: %v", roomID, dObjs)

	// Inspect Data Passed to Setup Functions
	world.Logger().Debug().Msgf("ata passed to createPlace for room %d: %+v\033[0m", roomID, [32]uint32{})
	world.Logger().Debug().Msgf("Data passed to createDirectionalObject for room %d: %+v\033[0m", roomID, dObjs)

	s.createPlace(roomID, enums.RoomTypeWoodCabin, dObjs, [32]uint32{}, tidBarn, world) // Pass empty array instead of nil for the objects.

	world.Logger().Debug().Msg("---->Barn setup complete")

}

// setupMountainPath sets up the mountain path in the game world.
func (s *GameSetup) setupMountainPath(world cardinal.WorldContext) {
	world.Logger().Debug().Msg("---->Setting up the mountain path..")

	// KPATH -> W
	open2West := s.createAction(enums.ActionTypeOpen, "the path is passable", true, true, false, 0, 0)
	pathActions := [32]uint32{open2West}

	dObjs := [32]uint32{s.createDirObject(enums.DirectionTypeWest, enums.RoomTypePlain,
		enums.ObjectTypePath, enums.MaterialTypeStone,
		"path", pathActions, world)}

	roomID := s.RoomStore.Add(component.Room{
		Description: "a high mountain pass",
		RoomType:    enums.RoomTypeStoneCabin,
	})

	tidMpath := s._textGuid("a high mountain pass")
	s.TxtDefStore.Set(tidMpath, enums.TxtDefTypePlace,
		"it winds through the mountains, the path is treacherous\n"+
			"toilet papered trees cover the steep \n valley sides below you.\n"+
			"On closer inspection the TP might \nbe the remains of a cricket team\n"+
			"or perhaps a lost and very dead KKK picnic group.\n"+
			"It's brass monkeys.",
	)

	// Logs to verify the ID's in the setup
	world.Logger().Debug().Msgf("Object IDs for room %d: %v\033[0m", roomID, [32]uint32{})
	world.Logger().Debug().Msgf("Directional Object IDs for room %d: %v", roomID, dObjs)

	// Inspect Data Passed to Setup Functions
	world.Logger().Debug().Msgf("ata passed to createPlace for room %d: %+v\033[0m", roomID, [32]uint32{})
	world.Logger().Debug().Msgf("Data passed to createDirectionalObject for room %d: %+v\033[0m", roomID, dObjs)

	s.createPlace(roomID, enums.RoomTypeStoneCabin, dObjs, [32]uint32{}, tidMpath, world) // Pass empty array instead of nil for the objects.

	world.Logger().Debug().Msg("---->Mountain path setup complete")
}

// createDirObject creates a directional object in the game world.
func (s *GameSetup) createDirObject(dirType enums.DirectionType, dstID enums.RoomType, dOType enums.ObjectType,
	mType enums.MaterialType, desc string, actionObjects [32]uint32, world cardinal.WorldContext) uint32 {
	// Generate a text GUID for the description
	txtID := s._textGuid(desc)

	// Set the text definition for the directional object and add it to the TxtDefStore
	s.TxtDefStore.Set(txtID, enums.TxtDefTypeDirObject, desc)

	// Create directional object data
	directionObjData := component.Object{
		DirType:         dirType,
		DestID:          dstID,
		ObjectType:      dOType,
		MaterialType:    mType,
		TxtDefID:        txtID,
		ObjectActionIDs: actionObjects,
	}

	// Add the directional object data to the DirObjectStore
	dirObjID := s.ObjectStore.Add(directionObjData)

	// Update the DirObject data with the assigned ID
	directionObjData.ObjectID = dirObjID

	world.Logger().Debug().Msgf("Directional object created - ID: %d, Type: %s, Destination Room Type: %s, Material: %s, Description: %s",
		dirObjID, dirType, dstID, mType, desc)

	return dirObjID
}

// createObject creates an object in the game world.
func (s *GameSetup) createObject(objType enums.ObjectType, mType enums.MaterialType, desc, objName string,
	actionObjects [32]uint32, world cardinal.WorldContext) uint32 {

	// Generate a text GUID for the description
	txtID := s._textGuid(desc)

	// Set the text definition for the object and add it to the TxtDefStore
	s.TxtDefStore.Set(txtID, enums.TxtDefTypeObject, desc)

	// Create object data
	objData := component.Object{
		ObjectName:      objName,
		ObjectType:      objType,
		MaterialType:    mType,
		TxtDefID:        txtID,
		ObjectActionIDs: actionObjects,
	}

	// Add the object data to the ObjectStore
	objID := s.ObjectStore.Add(objData)

	// Update the Object data with the assigned ID
	objData.ObjectID = objID

	world.Logger().Debug().Msgf("Object created - ID: %d, Type: %s, Material: %s, Description: %s", objID, objType, mType, desc)

	return objID
}

// createAction creates an action in the game world.
func (s *GameSetup) createAction(actionType enums.ActionType, desc string, enabled bool, dBit bool, revert bool, affectsID uint32, affectedByID uint32) uint32 {
	// Make sure s.ActionStore is properly initialized and not nil
	if s.ActionStore == nil {
		// Handle the error or panic accordingly
		panic("ActionStore is nil")
	}

	txtID := s._textGuid(desc)
	actionData := component.Action{ActionType: actionType, DBitTxt: txtID, Enabled: enabled, Revert: revert, DBit: dBit, AffectsActionID: affectsID, AffectedByActionID: affectedByID}
	actionID := s.ActionStore.Add(actionData)
	return actionID
}

// _textGuid generates a text GUID for the given description.
func (s *GameSetup) _textGuid(desc string) string {
	// Just a placeholder for actual GUID generation
	return desc
}

// createPlace creates a room in the game world and populates it with objects and directional objects.
func (s *GameSetup) createPlace(roomID uint32, roomType enums.RoomType, dObjs [32]uint32, objs [32]uint32, tid string, world cardinal.WorldContext) {
	// Create an entity representing the room with its components
	entityID, err := cardinal.Create(s.worldCtx,
		component.Room{
			ID:          roomID - 1,
			Description: tid,
			RoomType:    roomType,
			ObjectIDs:   objs,  // Populate ObjectIDs with the IDs of objects
			DirObjIDs:   dObjs, // Populate DirObjIDs with the IDs of directional objects
		},
	)

	// Check for errors during entity creation
	if err != nil {
		world.Logger().Debug().Msgf("Failed to create entity for room with ID %d: %v", roomID, err)
	}

	world.Logger().Debug().Msgf("Room with ID %d created successfully, entity ID: %d", roomID, entityID)

}

// isDevelopmentMode returns true if the application is running in development/debug mode.
func isDevelopmentMode() bool {

	return false // Change this based on if you run in dev or production.
}
