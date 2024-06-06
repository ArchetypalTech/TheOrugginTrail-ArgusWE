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
	setup.Init()
	initLogger()
	return nil
}

// GameSetup manages the setup of the game world.
type GameSetup struct {
	worldCtx       cardinal.WorldContext
	RoomStore      *component.RoomStore
	DirObjectStore *component.DirObjectStore
	ObjectStore    *component.ObjectStore
	ActionStore    *component.ActionStore
	TxtDefStore    *component.TxtDefStore
}

func NewGameSetup(worldCtx cardinal.WorldContext) *GameSetup {
	return &GameSetup{
		worldCtx:       worldCtx,
		RoomStore:      component.NewRoomStore(),
		DirObjectStore: component.NewDirObjectStore(),
		ObjectStore:    component.NewObjectStore(),
		ActionStore:    component.NewActionStore(),
		TxtDefStore:    component.NewTxtDefStore(),
	}
}

// Init initializes the GameSetup system.
func (s *GameSetup) Init() {
	s.setupWorld()
}

// setupWorld initializes the game world by setting up rooms.
func (s *GameSetup) setupWorld() {
	s.setupRooms()
	// Add setup functions for objects, actions, etc.
}

// setupRooms sets up various rooms in the game world.
func (s *GameSetup) setupRooms() {
	s.setupPlain()
	s.setupBarn()
	s.setupMountainPath()
}

// setupPlain sets up the plain in the game world.
func (s *GameSetup) setupPlain() {

	if isDevelopmentMode() {
		// Logging for setting up the plain
		logger.Debugf("\033[35mSetting up the plain...\033[0m")
	}

	// KPLAIN -> N, E
	open2Barn := s.createAction(enums.ActionTypeOpen, "the door opens with a farty noise\n"+
		"you can actually smell fart", true, true, true, 0, 0)

	plainBarn := [32]uint32{open2Barn}
	dObjs := [32]uint32{s.createDirObject(enums.DirectionTypeNorth, enums.RoomTypePlain,
		enums.DirObjectTypePath, enums.MaterialTypeDirt,
		"path", plainBarn)}

	open2Path := s.createAction(enums.ActionTypeOpen, "the door opens and a small hinge demon curses you\n"+
		"your nose is really itchy", true, true, true, 0, 0)

	plainPath := [32]uint32{open2Path}
	dObjs[1] = s.createDirObject(enums.DirectionTypeEast, enums.RoomTypeWoodCabin,
		enums.DirObjectTypePath, enums.MaterialTypeMud,
		"path", plainPath)

	kick := s.createAction(enums.ActionTypeKick, "the ball (such as it is)\n"+
		"bounces feebly\n then rolls into some fresh dog eggs\n"+
		"none the less you briefly feel a little better", true, false, true, 0, 0)

	ballActions := [32]uint32{kick}
	objs := [32]uint32{s.createObject(enums.ObjectTypeFootball, enums.MaterialTypeFlesh,
		"A slightly deflated knock off uefa football,\n"+
			"not quite spherical, it's "+
			"kickable though", "football", ballActions)}

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

	s.createPlace(roomID, enums.RoomTypePlain, dObjs, objs, tidPlain)

	if isDevelopmentMode() {
		// Logging for completing setup of the plain
		logger.Debugf("\033[32mPlain setup complete\033[0m")
	}

}

// setupBarn sets up the barn in the game world.
func (s *GameSetup) setupBarn() {
	if isDevelopmentMode() {
		// Logging for setting up the barn
		logger.Debugf("\033[35mSetting up the barn...\033[0m")
	}
	// KBARN -> S
	open2South := s.createAction(enums.ActionTypeOpen, "the door opens\n", true, true, true, 0, 0)
	barnPlain := [32]uint32{open2South}
	dObjs := [32]uint32{s.createDirObject(enums.DirectionTypeSouth, enums.RoomTypePlain,
		enums.DirObjectTypeDoor, enums.MaterialTypeWood,
		"door", barnPlain)}

	open2Forest := s.createAction(enums.ActionTypeOpen, "the window, glass and frame smashed\n"+
		"falls open", false, false, false, 0, 0)

	smashWindow := s.createAction(enums.ActionTypeBreak, "I love the sound of breaking glass\n"+
		"especially when I'm lonely , the panes and the frame shatter\n"+
		"satisfyingly spreading broken joy on the floor",
		true, false, false, open2Forest, 0)

	windowActions := [32]uint32{open2Forest, smashWindow}
	dObjs[1] = s.createDirObject(enums.DirectionTypeEast, enums.RoomTypeForge,
		enums.DirObjectTypeWindow, enums.MaterialTypeWood,
		"window", windowActions)

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

	s.createPlace(roomID, enums.RoomTypeWoodCabin, dObjs, [32]uint32{}, tidBarn) // Pass empty array instead of nil for the objects.

	if isDevelopmentMode() {
		// Logging for completing setup of the plain
		logger.Debugf("\033[32mBarn setup complete\033[0m")
	}
}

// setupMountainPath sets up the mountain path in the game world.
func (s *GameSetup) setupMountainPath() {
	if isDevelopmentMode() {
		// Logging for setting up the mountain path
		logger.Debugf("\033[35mSetting up the mountain path...\033[0m")
	}

	// KPATH -> W
	open2West := s.createAction(enums.ActionTypeOpen, "the path is passable", true, true, false, 0, 0)
	pathActions := [32]uint32{open2West}

	dObjs := [32]uint32{s.createDirObject(enums.DirectionTypeWest, enums.RoomTypePlain,
		enums.DirObjectTypePath, enums.MaterialTypeStone,
		"path", pathActions)}

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

	s.createPlace(roomID, enums.RoomTypeStoneCabin, dObjs, [32]uint32{}, tidMpath) // Pass empty array instead of nil for the objects.

	if isDevelopmentMode() {
		// Logging for completing setup of the plain
		logger.Debugf("\033[32mMountain path setup complete\033[0m")
	}
}

// createDirObject creates a directional object in the game world.
func (s *GameSetup) createDirObject(dirType enums.DirectionType, dstID enums.RoomType, dOType enums.DirObjectType,
	mType enums.MaterialType, desc string, actionObjects [32]uint32) uint32 {
	// Generate a text GUID for the description
	txtID := s._textGuid(desc)

	// Set the text definition for the directional object and add it to the TxtDefStore
	s.TxtDefStore.Set(txtID, enums.TxtDefTypeDirObject, desc)

	// Create directional object data
	dirObjData := component.DirObject{
		DirType:         dirType,
		DestID:          dstID,
		ObjType:         dOType,
		MatType:         mType,
		TxtDefID:        txtID,
		ObjectActionIDs: actionObjects,
	}

	// Add the directional object data to the DirObjectStore
	dirObjID := s.DirObjectStore.Add(dirObjData)

	// Update the DirObject data with the assigned ID
	dirObjData.ID = dirObjID

	if isDevelopmentMode() {
		// Log the directional object creation
		logger.Debugf("\033[33mDirectional object created - ID: %d, Type: %s, Destination Room Type: %s, Material: %s, Description: %s\033[0m",
			dirObjID, dirType, dstID, mType, desc)
	}

	return dirObjID
}

// createObject creates an object in the game world.
func (s *GameSetup) createObject(objType enums.ObjectType, mType enums.MaterialType, desc, objName string,
	actionObjects [32]uint32) uint32 {

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
	objData.ID = objID

	if isDevelopmentMode() {
		// Log the object creation
		logger.Debugf("\033[33mObject created - ID: %d, Type: %s, Material: %s, Description: %s\033[0m", objID, objType, mType, desc)
	}

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
	return "txt-" + desc
}

// createPlace creates a room in the game world and populates it with objects and directional objects.
func (s *GameSetup) createPlace(roomID uint32, roomType enums.RoomType, dObjs [32]uint32, objs [32]uint32, tid string) {
	// Create an entity representing the room with its components
	entityID, err := cardinal.Create(s.worldCtx,
		component.Room{
			ID:          roomID,
			Description: tid,
			RoomType:    roomType,
			ObjectIDs:   objs,  // Populate ObjectIDs with the IDs of objects
			DirObjIDs:   dObjs, // Populate DirObjIDs with the IDs of directional objects
		},
	)

	// Check for errors during entity creation
	if err != nil {
		logger.Errorf("\033[31mFailed to create entity for room with ID %d: %v\033[0m", roomID, err)
	}

	if isDevelopmentMode() {
		// Log successful creation
		logger.Debugf("\033[32mRoom with ID %d created successfully, entity ID: %d%s", roomID, entityID, "\033[0m")
	}
}

// isDevelopmentMode returns true if the application is running in development/debug mode.
func isDevelopmentMode() bool {

	return true // Change this based on if you run in dev or production.
}
