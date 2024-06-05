package system

import (
	"log"

	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/component"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"
	"pkg.world.dev/world-engine/cardinal"
)

// NGameSetupSystem initializes the game setup system.
func NGameSetupSystem(world cardinal.WorldContext) error {
	setup := NewGameSetup(world)
	setup.Init()
	// LOG FOR TESTING TxDefStore entries
	// Log the number of entries and contents of TxtDefStore
	log.Printf("TxtDefStore initialized with %d entries", len(setup.TxtDefStore.TxtDefs))
	for key, value := range setup.TxtDefStore.TxtDefs {
		log.Printf("TxtDefStore entry - Key: %s, Value: %+v", key, value)
	}
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
	// Logging for setting up the plain
	log.Println("Setting up the plain...")

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

	/*  LOGS FOR TESTING VERIFICATION!
	// Logs to verify the ID's in the setup
	log.Printf("Object IDs for room %d: %v", roomID, objs)
	log.Printf("Directional Object IDs for room %d: %v", roomID, dObjs)
	// Inspect Data Passed to Setup Functions
	log.Printf("Data passed to createPlace for room %d: %+v", roomID, objs)
	log.Printf("Data passed to createDirectionalObject for room %d: %+v", roomID, dObjs)
	*/

	s.createPlace(roomID, enums.RoomTypePlain, dObjs, objs, tidPlain)

	// Logging for completing setup of the plain
	log.Println("Plain setup complete")
}

// setupBarn sets up the barn in the game world.
func (s *GameSetup) setupBarn() {
	// Logging for setting up the barn
	log.Println("Setting up the barn...")

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

	/* LOGS FOR TESTING VERIFICATION!
	// Logs to verify the ID's in the setup
	log.Printf("Object IDs for room %d: %v", roomID, [32]uint32{})
	log.Printf("Directional Object IDs for room %d: %v", roomID, dObjs)
	// Inspect Data Passed to Setup Functions
	log.Printf("Data passed to createPlace for room %d: %+v", roomID, [32]uint32{})
	log.Printf("Data passed to createDirectionalObject for room %d: %+v", roomID, dObjs)
	*/

	s.createPlace(roomID, enums.RoomTypeWoodCabin, dObjs, [32]uint32{}, tidBarn) // Pass empty array instead of nil for the objects.

	// Logging for completing setup of the barn
	log.Println("Barn setup complete")
}

// setupMountainPath sets up the mountain path in the game world.
func (s *GameSetup) setupMountainPath() {
	// Logging for setting up the mountain path
	log.Println("Setting up the mountain path...")

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

	/* LOGS FOR TESTING VERIFICATION!
	// Logs to verify the ID's in the setup
	log.Printf("Object IDs for room %d: %v", roomID, [32]uint32{})
	log.Printf("Directional Object IDs for room %d: %v", roomID, dObjs)
	// Inspect Data Passed to Setup Functions
	log.Printf("Data passed to createPlace for room %d: %+v", roomID, [32]uint32{})
	log.Printf("Data passed to createDirectionalObject for room %d: %+v", roomID, dObjs)
	*/

	s.createPlace(roomID, enums.RoomTypeStoneCabin, dObjs, [32]uint32{}, tidMpath) // Pass empty array instead of nil for the objects.

	// Logging for completing setup of the mountain path
	log.Println("Mountain path setup complete")
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

	// Log the directional object creation
	log.Printf("Directional object created - ID: %d, Type: %s, Destination Room Type: %s, Material: %s, Description: %s", dirObjID, dirType, dstID, mType, desc)

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

	// Log the object creation
	log.Printf("Object created - ID: %d, Type: %s, Material: %s, Description: %s", objID, objType, mType, desc)

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
		log.Fatalf("Failed to create entity for room with ID %d: %v", roomID, err)
	}

	// Log successful creation
	log.Printf("Room with ID %d created successfully, entity ID: %d", roomID, entityID)

	/* LOGS FOR TESTING THAT THE ROOMS HAS THEIR OBJECTS, ETC. NEED TO BE POLISHED!
	// Log objects added to the room
	for _, objID := range objs {
		if objID != 0 {
			objDef, found := s.TxtDefStore.Get(fmt.Sprintf("txt-%d", objID))
			if found {
				log.Printf("Object added to the room: %s", objDef.Description)
			} else {
				log.Printf("Failed to find description for object with ID %d", objID)
			}
		}
	}

	// Log directional objects added to the room
	for _, dirObjID := range dObjs {
		if dirObjID != 0 {
			dirObjDef, found := s.TxtDefStore.Get(fmt.Sprintf("txt-%d", dirObjID))
			if found {
				log.Printf("Directional object added to the room: %s", dirObjDef.Description)
			} else {
				log.Printf("Failed to find description for directional object with ID %d", dirObjID)
			}
		}
	}
	*/
}
