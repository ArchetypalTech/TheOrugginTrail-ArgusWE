package system

import (
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/component"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"
	"pkg.world.dev/world-engine/cardinal"
)

// NGameSetupSystem initializes the game setup system.
func NGameSetupSystem(world cardinal.WorldContext) error {

	setup := NewGameSetup(world)
	setup.Init(world)

	// Log the number of entries and contents of TxtDefStore
	world.Logger().Debug().Msgf("TxtDefStore initialized with %d entries", len(setup.TxtDefStore.TxtDefs))
	for key, value := range setup.TxtDefStore.TxtDefs {
		world.Logger().Debug().Msgf("TxtDefStore entry - Key: %s, Value: %+v", key, value)
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
		ActionStore:    component.NewActionStore(worldCtx),
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
	s.setupForge(world)
	s.setupCellar(world)
}

// setupPlain sets up the plain in the game world.
func (s *GameSetup) setupPlain(world cardinal.WorldContext) {
	world.Logger().Debug().Msg("---->Setting up the plain.....")

	// KPLAIN -> N, E
	open2Barn := s.createAction(enums.ActionTypeOpen, " The door opens with a farty noise"+
		" you can actually smell fart", true, true, true, 0, 0)

	plainBarn := []uint32{open2Barn}
	dObjs := []component.Object{s.createDirObject(enums.DirectionTypeNorth, enums.RoomTypeBarn,
		enums.ObjectTypePath, enums.MaterialTypeDirt,
		"path", plainBarn, world)}

	open2Path := s.createAction(enums.ActionTypeOpen, " The door opens and a small hinge demon curses you"+
		" your nose is really itchy", true, true, true, 0, 0)

	plainPath := []uint32{open2Path}
	dObjs = append(dObjs, s.createDirObject(enums.DirectionTypeEast, enums.RoomTypeStoneCabin,
		enums.ObjectTypePath, enums.MaterialTypeMud,
		"path", plainPath, world))

	kick := s.createAction(enums.ActionTypeKick, " The ball (such as it is)"+
		" bounces feebly then rolls into some fresh dog eggs"+
		" none the less you briefly feel a little better", true, false, true, 0, 0)

	ballActions := []uint32{kick}
	objs := []component.Object{s.createObject(enums.ObjectTypeFootball, enums.MaterialTypeFlesh,
		"A slightly deflated knock off uefa football, not quite spherical, it's kickable though", enums.ObjectTypeFootball.String(), ballActions, true, world)}

	roomID := s.RoomStore.Add(component.Room{
		Description: "a windswept plain",
		RoomType:    enums.RoomTypePlain,
	})

	tidPlain := s._textGuid("a windswept plain")
	txtPlain := (" where the wind blowing is cold and" +
		" bison skulls in piles taller than houses" +
		" cover the plains as far as your eye can see" +
		" the air tastes of burnt grease and bensons")

	// Logs to verify the ID's in the setup
	world.Logger().Debug().Msgf("Object IDs for room %d: %v", roomID, objs)
	world.Logger().Debug().Msgf("Directional Object IDs for room %d: %v", roomID, dObjs)

	// Inspect Data Passed to Setup Functions
	world.Logger().Debug().Msgf("Data passed to createPlace for room %d: %+v", roomID, objs)
	world.Logger().Debug().Msgf("Data passed to createDirectionalObject for room %d: %+v", roomID, dObjs)

	s.createPlace(roomID, enums.RoomTypePlain, dObjs, objs, tidPlain, txtPlain, world)

	world.Logger().Debug().Msg("---->Plain setup complete")

}

// setupBarn sets up the barn in the game world.
func (s *GameSetup) setupBarn(world cardinal.WorldContext) {
	world.Logger().Debug().Msg("---->Setting up the barn.....")
	// KBARN -> S
	open2South := s.createAction(enums.ActionTypeOpen, "the door opens", true, true, true, 0, 0)
	barnPlain := []uint32{open2South}
	dObjs := []component.Object{s.createDirObject(enums.DirectionTypeSouth, enums.RoomTypePlain,
		enums.ObjectTypeDoor, enums.MaterialTypeWood,
		"door", barnPlain, world)}

	// KBARN -> E
	open2Forge := s.createAction(enums.ActionTypeOpen, " The window, glass and frame smashed"+
		" falls open.", false, false, false, 0, 0)

	smashWindow := s.createAction(enums.ActionTypeBreak, " I love the sound of breaking glass"+
		" especially when I'm lonely, the panes and the frame shatter"+
		" satisfyingly spreading broken joy on the floor.",
		true, false, false, open2Forge, 0)

	windowActions := []uint32{open2Forge, smashWindow}
	dObjs = append(dObjs, s.createDirObject(enums.DirectionTypeEast, enums.RoomTypeForge,
		enums.ObjectTypeWindow, enums.MaterialTypeWood,
		"window", windowActions, world))

	// KBARN -> DOWN
	open2Cellar := s.createAction(enums.ActionTypeOpen, " The hay having been burned quickly, reveals the stairs cleared up.", false, false, false, 0, 0)

	burnHay := s.createAction(enums.ActionTypeBurn, " I can hear the crispy sound of the hay burning"+
		" as it gets consumed by the dark fires."+
		" However, this enjoyable moment does not persist for too long.",
		true, false, true, open2Cellar, 0)

	hayActions := []uint32{open2Cellar, burnHay}
	dObjs = append(dObjs, s.createDirObject(enums.DirectionTypeDown, enums.RoomTypeCellar,
		enums.ObjectTypeStairs, enums.MaterialTypeStone,
		"stack of hay filling a stairs made of", hayActions, world))

	roomID := s.RoomStore.Add(component.Room{
		Description: "a barn",
		RoomType:    enums.RoomTypeBarn,
	})

	tidBarn := s._textGuid("a barn")
	txtBarn := (" and place is dusty and full of spiderwebs," +
		" something died in here, possibly your own self" +
		" plenty of corners and dark shadows.")

	// Logs to verify the ID's in the setup
	world.Logger().Debug().Msgf("Object IDs for room %d: %v", roomID, [32]uint32{})
	world.Logger().Debug().Msgf("Directional Object IDs for room %d: %v", roomID, dObjs)

	// Inspect Data Passed to Setup Functions
	world.Logger().Debug().Msgf("Data passed to createPlace for room %d: %+v", roomID, [32]uint32{})
	world.Logger().Debug().Msgf("Data passed to createDirectionalObject for room %d: %+v", roomID, dObjs)

	s.createPlace(roomID, enums.RoomTypeBarn, dObjs, nil, tidBarn, txtBarn, world) // Pass empty array instead of nil for the objects.

	world.Logger().Debug().Msg("---->Barn setup complete")
}

// setupMountainPath sets up the mountain path in the game world.
func (s *GameSetup) setupMountainPath(world cardinal.WorldContext) {
	world.Logger().Debug().Msg("---->Setting up the mountain path..")

	// KPATH -> W
	open2West := s.createAction(enums.ActionTypeOpen, "the path is passable", true, true, false, 0, 0)
	pathActions := []uint32{open2West}

	dObjs := []component.Object{s.createDirObject(enums.DirectionTypeWest, enums.RoomTypePlain,
		enums.ObjectTypePath, enums.MaterialTypeStone,
		"path", pathActions, world)}

	roomID := s.RoomStore.Add(component.Room{
		Description: "a high mountain pass",
		RoomType:    enums.RoomTypeStoneCabin,
	})

	// KPATH -> E
	open2ActII := s.createAction(enums.ActionTypeOpen, " The boulder, blasted to pieces"+
		" reveals a path to a new adventure."+
		" This is the end of ACT I, if you go East you will return to the plain", false, false, false, 0, 0)

	removeBoulder := s.createAction(enums.ActionTypeThrow, " The dynamite lands at the boulder while you keep running"+
		" and in just a second KBOOOM!.",
		true, false, false, open2ActII, 0)

	boulderActions := []uint32{open2ActII, removeBoulder}
	dObjs = append(dObjs, s.createDirObject(enums.DirectionTypeEast, enums.RoomTypePlain,
		enums.ObjectTypeBoulder, enums.MaterialTypeStone,
		"boulder", boulderActions, world))

	tidMpath := s._textGuid("a high mountain pass")
	txtMpath := (" where it winds through the mountains, the path is treacherous" +
		" toilet papered trees cover the steep valley sides below you." +
		" On closer inspection the TP might be the remains of a cricket team" +
		" or perhaps a lost and very dead KKK picnic group." +
		" It's brass monkeys.")

	// Logs to verify the ID's in the setup
	world.Logger().Debug().Msgf("Object IDs for room %d: %v", roomID, [32]uint32{})
	world.Logger().Debug().Msgf("Directional Object IDs for room %d: %v", roomID, dObjs)

	// Inspect Data Passed to Setup Functions
	world.Logger().Debug().Msgf("Data passed to createPlace for room %d: %+v", roomID, [32]uint32{})
	world.Logger().Debug().Msgf("Data passed to createDirectionalObject for room %d: %+v", roomID, dObjs)

	s.createPlace(roomID, enums.RoomTypeStoneCabin, dObjs, nil, tidMpath, txtMpath, world) // Pass empty array instead of nil for the objects.

	world.Logger().Debug().Msg("---->Mountain path setup complete")
}

// setupForge sets up the forge room in the game world.
func (s *GameSetup) setupForge(world cardinal.WorldContext) {
	world.Logger().Debug().Msg("---->Setting up the forge.....")
	// KFORGE -> W
	open2Barn := s.createAction(enums.ActionTypeOpen, " The broken window is still there"+
		" just don't cut yourself", true, true, true, 0, 0)
	plainBarn := []uint32{open2Barn}
	dObjs := []component.Object{s.createDirObject(enums.DirectionTypeWest, enums.RoomTypeBarn,
		enums.ObjectTypeWindow, enums.MaterialTypeWood,
		"window", plainBarn, world)}

	// Petrol Object
	burn := s.createAction(enums.ActionTypeBurn, " The petrol (with its unique funny smell)"+
		" burns fiercely with a scent that makes you feel dizzy", true, false, true, 0, 0)

	petrolActions := []uint32{burn}
	objs := []component.Object{s.createObject(enums.ObjectTypePetrol, enums.MaterialTypeIKEA,
		"a strange liquid that seems to be petrol, probably there is three liters of it, it's highly burnable though", enums.ObjectTypePetrol.String(), petrolActions, true, world)}

	// Matchsticks Object
	light := s.createAction(enums.ActionTypeLight, " The matchsticks (despite their small size)"+
		" lights enough to see a small distance. You have to use them quickly"+
		" or your fingers will be burn", true, false, true, 0, 0)

	matchesActions := []uint32{light}
	objs = append(objs, s.createObject(enums.ObjectTypeMatchsticks, enums.MaterialTypeWood,
		"a group of matchsticks that have survived the pass of time, you can probably light them up, as they seem to be in a good condition", enums.ObjectTypeMatchsticks.String(), matchesActions, true, world))

	roomID := s.RoomStore.Add(component.Room{
		Description: "a dusty forge",
		RoomType:    enums.RoomTypeForge,
	})

	tidForge := s._textGuid("a dusty forge")
	txtForge := (" and you can see that it has not been used in ages." +
		" There are many blood spots accross the forge and even the anvil is broken." +
		" You don't know what happened here.")

	// Logs to verify the ID's in the setup
	world.Logger().Debug().Msgf("Object IDs for room %d: %v", roomID, [32]uint32{})
	world.Logger().Debug().Msgf("Directional Object IDs for room %d: %v", roomID, dObjs)

	// Inspect Data Passed to Setup Functions
	world.Logger().Debug().Msgf("Data passed to createPlace for room %d: %+v", roomID, [32]uint32{})
	world.Logger().Debug().Msgf("Data passed to createDirectionalObject for room %d: %+v", roomID, dObjs)

	s.createPlace(roomID, enums.RoomTypeForge, dObjs, objs, tidForge, txtForge, world)

	world.Logger().Debug().Msg("---->Forge setup complete")
}

// setupCellar sets up the cellar room in the game world.
func (s *GameSetup) setupCellar(world cardinal.WorldContext) {
	// KCELLAR -> UP
	world.Logger().Debug().Msg("---->Setting up the cellar.....")

	open2Barn := s.createAction(enums.ActionTypeOpen, " The way is opened"+
		" just don't fall when going up", true, true, true, 0, 0)
	plainBarn := []uint32{open2Barn}
	dObjs := []component.Object{s.createDirObject(enums.DirectionTypeUp, enums.RoomTypeBarn,
		enums.ObjectTypeTrapdoor, enums.MaterialTypeStone,
		"trapdoor", plainBarn, world)}

	// Dynamite Object
	lightDynamite := s.createAction(enums.ActionTypeLight, " Seeing the fuse sparkling makes you remember the fireworks you loved as a child."+
		" Suddenly, you return back and see that there is almost no time."+
		" You have to do something or you will be flying into meat pieces!", true, false, true, 0, 0)

	throwDynamite := s.createAction(enums.ActionTypeThrow, " You throw quickly the dynamite"+
		" and start running for the hills as your life depens on it", true, false, true, 0, 0)

	dynamiteActions := []uint32{lightDynamite, throwDynamite}
	objs := []component.Object{s.createObject(enums.ObjectTypeDynamite, enums.MaterialTypeIKEA,
		"a high quality dynamite that can blast almost everything. It needs to be lighted up first", enums.ObjectTypeDynamite.String(), dynamiteActions, true, world)}

	// GLUE Object
	sniff := s.createAction(enums.ActionTypeSniff, " The smell of it is really relaxing,"+
		" but for now thats all", true, false, true, 0, 0)

	glueActions := []uint32{sniff}
	objs = append(objs, s.createObject(enums.ObjectTypeGlue, enums.MaterialTypeShit,
		"some glue of a strange name: &$#${+>@#!@. Don't knowing if it's your imagination, it is calling you", enums.ObjectTypeGlue.String(), glueActions, true, world))

	roomID := s.RoomStore.Add(component.Room{
		Description: "a small cellar",
		RoomType:    enums.RoomTypeForge,
	})

	tidForge := s._textGuid("a small cellar")
	txtForge := (" big enough to hide probably fifty people." +
		" It seems that it was constructed with great care as you can't find any cracks or holes in the walls." +
		" This place might be important.")

	// Logs to verify the ID's in the setup
	world.Logger().Debug().Msgf("Object IDs for room %d: %v", roomID, objs)
	world.Logger().Debug().Msgf("Directional Object IDs for room %d: %v", roomID, dObjs)

	// Inspect Data Passed to Setup Functions
	world.Logger().Debug().Msgf("Data passed to createPlace for room %d: %+v", roomID, objs)
	world.Logger().Debug().Msgf("Data passed to createDirectionalObject for room %d: %+v", roomID, dObjs)

	s.createPlace(roomID, enums.RoomTypeCellar, dObjs, objs, tidForge, txtForge, world)

	world.Logger().Debug().Msg("---->Forge setup complete")
}

// createDirObject creates a directional object in the game world.
func (s *GameSetup) createDirObject(dirType enums.DirectionType, dstID enums.RoomType, dOType enums.ObjectType,
	mType enums.MaterialType, desc string, actionObjects []uint32, world cardinal.WorldContext) component.Object {
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
		Description:     txtID,
		ObjectActionIDs: actionObjects,
		CanBePickedUp:   false,
	}

	// Add the directional object data to the DirObjectStore
	dirObjID := s.ObjectStore.Add(directionObjData)

	// Update the DirObject data with the assigned ID
	directionObjData.ObjectID = dirObjID

	world.Logger().Debug().Msgf("Directional object created - ID: %d, Type: %s, Destination Room Type: %s, Material: %s, Description: %s",
		dirObjID, dirType, dstID, mType, desc)

	return directionObjData
}

// createObject creates an object in the game world.
func (s *GameSetup) createObject(objType enums.ObjectType, mType enums.MaterialType, desc, objName string,
	actionObjects []uint32, pickUp bool, world cardinal.WorldContext) component.Object {

	// Generate a text GUID for the description
	txtID := s._textGuid(desc)

	// Set the text definition for the object and add it to the TxtDefStore
	s.TxtDefStore.Set(txtID, enums.TxtDefTypeObject, desc)

	// Create object data
	objData := component.Object{
		ObjectName:      objName,
		ObjectType:      objType,
		MaterialType:    mType,
		Description:     desc,
		ObjectActionIDs: actionObjects,
		CanBePickedUp:   pickUp,
	}

	// Add the object data to the ObjectStore
	objID := s.ObjectStore.Add(objData)

	// Update the Object data with the assigned ID
	objData.ObjectID = objID

	world.Logger().Debug().Msgf("Object created - ID: %d, Type: %s, Material: %s, Description: %s", objID, objType, mType, desc)

	return objData
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
func (s *GameSetup) createPlace(roomID uint32, roomType enums.RoomType, dObjs []component.Object, objs []component.Object, tid string, roomTxt string, world cardinal.WorldContext) {
	// Create an entity representing the room with its components
	roomManagerID, err := cardinal.Create(s.worldCtx,
		component.Room{
			ID:          roomID,
			Description: tid,
			RoomTxt:     roomTxt,
			RoomType:    roomType,
			Objects:     make(map[int]component.Object),
			DirObjs:     make(map[int]component.Object),
			Players:     make(map[int]component.Player),
		},
	)

	// Check for errors during entity creation
	if err != nil {
		world.Logger().Error().Msgf("Failed to create entity for room with ID %d: %v", roomID-1, err)
	}

	roomManager, err := cardinal.GetComponent[component.Room](world, roomManagerID)
	if err != nil {
		world.Logger().Error().Msgf("Error getting Object Component: %v", err)
	}

	//Populate Objects map with the objects passed. Uses as identifier the object ID
	for _, Obj := range objs {
		if Obj.ObjectID != 0 {
			roomManager.Objects[int(Obj.ObjectID)] = Obj
		}
	}

	// Populate DirObjs map with the directional objects passed. Uses as identifier the Object ID
	for _, dirObj := range dObjs {
		if dirObj.ObjectID != 0 {
			roomManager.DirObjs[int(dirObj.ObjectID)] = dirObj
		}
	}

	world.Logger().Info().Msgf("Room with ID %d created successfully, entity ID: %d", roomID, roomManagerID)
}
