package system

import (
	"fmt"
	"strings"

	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/component"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"
)

// TokeniserSystem structure
type TokeniserSystem struct {
	cmdLookup      map[string]enums.ActionType             // Lookup table for command strings to ActionType
	dirLookup      map[string]enums.DirectionType          // Lookup table for direction strings to DirectionType
	dirObjLookup   map[string]enums.DirObjectType          // Lookup table for direction object strings to DirObjectType
	objLookup      map[string]enums.ObjectType             // Lookup table for object strings to ObjectType
	responseLookup map[enums.ActionType][]enums.ActionType // Lookup table for action responses
}

// NewTokeniserSystem creates a new instance of TokeniserSystem and initializes lookup tables
func NewTokeniserSystem() *TokeniserSystem {
	ts := &TokeniserSystem{
		cmdLookup:      make(map[string]enums.ActionType),
		dirLookup:      make(map[string]enums.DirectionType),
		dirObjLookup:   make(map[string]enums.DirObjectType),
		objLookup:      make(map[string]enums.ObjectType),
		responseLookup: make(map[enums.ActionType][]enums.ActionType),
	}
	ts.initLUTS()
	return ts
}

// initLUTS initializes the lookup tables by calling the setup functions
func (ts *TokeniserSystem) initLUTS() {
	ts.setupCmds()
	ts.setupObjects()
	ts.setupDirs()
	ts.setupDirObjs()
	ts.setupVrbAct()
}

// setupCmds initializes the command lookup table with predefined actions
func (ts *TokeniserSystem) setupCmds() {
	ts.cmdLookup["GO"] = enums.ActionTypeGo
	ts.cmdLookup["MOVE"] = enums.ActionTypeMove
	ts.cmdLookup["LOOT"] = enums.ActionTypeLoot
	ts.cmdLookup["DESCRIBE"] = enums.ActionTypeDescribe
	ts.cmdLookup["TAKE"] = enums.ActionTypeTake
	ts.cmdLookup["KICK"] = enums.ActionTypeKick
	ts.cmdLookup["LOCK"] = enums.ActionTypeLock
	ts.cmdLookup["UNLOCK"] = enums.ActionTypeUnlock
	ts.cmdLookup["OPEN"] = enums.ActionTypeOpen
	ts.cmdLookup["LOOK"] = enums.ActionTypeLook
	ts.cmdLookup["CLOSE"] = enums.ActionTypeClose
	ts.cmdLookup["BREAK"] = enums.ActionTypeBreak
	ts.cmdLookup["THROW"] = enums.ActionTypeThrow
	ts.cmdLookup["DROP"] = enums.ActionTypeDrop
	ts.cmdLookup["INVENTORY"] = enums.ActionTypeInventory
	ts.cmdLookup["BURN"] = enums.ActionTypeBurn
	ts.cmdLookup["LIGHT"] = enums.ActionTypeLight
}

// setupObjects initializes the object lookup table with predefined objects
func (ts *TokeniserSystem) setupObjects() {
	ts.objLookup["FOOTBALL"] = enums.ObjectTypeFootball
	ts.objLookup["BALL"] = enums.ObjectTypeFootball
	ts.objLookup["KEY"] = enums.ObjectTypeKey
	ts.objLookup["KNIFE"] = enums.ObjectTypeKnife
	ts.objLookup["BOTTLE"] = enums.ObjectTypeBottle
}

// setupDirs initializes the direction lookup table with predefined directions
func (ts *TokeniserSystem) setupDirs() {
	ts.dirLookup["NORTH"] = enums.DirectionTypeNorth
	ts.dirLookup["SOUTH"] = enums.DirectionTypeSouth
	ts.dirLookup["EAST"] = enums.DirectionTypeEast
	ts.dirLookup["WEST"] = enums.DirectionTypeWest
	ts.dirLookup["UP"] = enums.DirectionTypeUp
	ts.dirLookup["DOWN"] = enums.DirectionTypeDown
	ts.dirLookup["FORWARD"] = enums.DirectionTypeForward
	ts.dirLookup["BACKWARD"] = enums.DirectionTypeBackward
}

// setupDirObjs initializes the directional object lookup table with predefined directional objects
func (ts *TokeniserSystem) setupDirObjs() {
	ts.dirObjLookup["DOOR"] = enums.DirObjectTypeDoor
	ts.dirObjLookup["WINDOW"] = enums.DirObjectTypeWindow
	ts.dirObjLookup["STAIRS"] = enums.DirObjectTypeStairs
	ts.dirObjLookup["LADDER"] = enums.DirObjectTypeLadder
	ts.dirObjLookup["PATH"] = enums.DirObjectTypePath
	ts.dirObjLookup["TRAIL"] = enums.DirObjectTypeTrail
}

// setupVrbAct initializes the verb action response lookup table with predefined responses
func (ts *TokeniserSystem) setupVrbAct() {
	ts.responseLookup[enums.ActionTypeKick] = []enums.ActionType{enums.ActionTypeBreak, enums.ActionTypeHit, enums.ActionTypeDamage}
	ts.responseLookup[enums.ActionTypeBurn] = []enums.ActionType{enums.ActionTypeBurn, enums.ActionTypeLight, enums.ActionTypeDamage}
	ts.responseLookup[enums.ActionTypeLight] = []enums.ActionType{enums.ActionTypeBurn, enums.ActionTypeLight, enums.ActionTypeDamage}
	ts.responseLookup[enums.ActionTypeOpen] = []enums.ActionType{enums.ActionTypeOpen}
	ts.responseLookup[enums.ActionTypeBreak] = []enums.ActionType{enums.ActionTypeBreak}
}

// FishTokens processes the tokenized command and returns VerbData
func (ts *TokeniserSystem) FishTokens(tokens []string) component.VerbData {
	var data component.VerbData // Initialize VerbData to store the result
	var err uint8               // Error code variable
	lenTokens := len(tokens) - 1

	// Look up the verb, object, and directional object from the tokens
	vrb, vrbExists := ts.cmdLookup[strings.ToUpper(tokens[0])]
	obj, objExists := ts.objLookup[strings.ToUpper(tokens[lenTokens])]
	dobj, dobjExists := ts.dirObjLookup[strings.ToUpper(tokens[lenTokens])]

	if !vrbExists {
		// Handle the case where the verb lookup fails
		fmt.Printf("Error: Verb '%s' not recognized\n", tokens[0])
		data.ErrCode = 1 // Set an appropriate error code
		return data
	}

	data.Verb = vrb // Set the verb in VerbData
	if !objExists && !dobjExists {
		data.ErrCode = 1 // Set error code if no object or directional object is found
		return data
	}

	// Handle cases where the object exists and the token length is less than or equal to 3
	if objExists && len(tokens) <= 3 {
		data.DirectNoun = obj
	} else if !objExists {
		err = 1 // Set error code if no object exists
	}

	// Handle cases where the token length is greater than 3
	if len(tokens) > 3 {
		if dobjExists {
			// Look for object in the second or third token if directional object exists
			obj, objExists = ts.objLookup[strings.ToUpper(tokens[1])]
			if !objExists {
				obj, objExists = ts.objLookup[strings.ToUpper(tokens[2])]
				if !objExists {
					err = 1 // Set error code if no object is found
				}
			}
		} else if objExists {
			// Return early if dealing with a form where the second object is not a directional object
			return data
		}
	}

	// Set the direct noun, indirect directional noun, and error code in VerbData
	data.DirectNoun = obj
	data.IndirectDirNoun = dobj
	data.ErrCode = err
	fmt.Printf("--->d.dobj:%s iobj:%s vrb:%s\n", data.DirectNoun, data.IndirectDirNoun, data.Verb)
	return data
}
