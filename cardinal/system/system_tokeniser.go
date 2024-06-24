package system

import (
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/component"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/constants"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"
	"pkg.world.dev/world-engine/cardinal"
)

func NTokeniserSystem(world cardinal.WorldContext) error {

	return nil
}

// TokeniserSystem structure
type TokeniserSystem struct {
	vrbLookup      map[string]enums.ActionType             // Lookup table for command strings to ActionType
	dirLookup      map[string]enums.DirectionType          // Lookup table for direction strings to DirectionType
	dirObjLookup   map[string]enums.ObjectType             // Lookup table for direction object strings to DirObjectType
	objLookup      map[string]enums.ObjectType             // Lookup table for object strings to ObjectType
	grammarLookup  map[string]enums.GrammarType            // Lookup table for GrammarType
	responseLookup map[enums.ActionType][]enums.ActionType // Lookup table for action responses
	revMat         map[string]enums.MaterialType           // Material type map in lowercase
	revDirLookup   map[string]enums.DirectionType          // ookup table for direction object strings to DirObjectType in lowercase
}

// NewTokeniserSystem creates a new instance of TokeniserSystem and initializes lookup tables
func NewTokeniserSystem() *TokeniserSystem {
	ts := &TokeniserSystem{
		vrbLookup:      make(map[string]enums.ActionType),
		dirLookup:      make(map[string]enums.DirectionType),
		dirObjLookup:   make(map[string]enums.ObjectType),
		objLookup:      make(map[string]enums.ObjectType),
		grammarLookup:  make(map[string]enums.GrammarType),
		responseLookup: make(map[enums.ActionType][]enums.ActionType),
		revMat:         make(map[string]enums.MaterialType),
		revDirLookup:   make(map[string]enums.DirectionType),
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
	ts.setupGrammar()
	ts.setupVrbAct()
	ts.setupRevDirs()
	ts.setupMaterial()
}

// setupCmds initializes the command lookup table with predefined actions
func (ts *TokeniserSystem) setupCmds() {
	ts.vrbLookup["GO"] = enums.ActionTypeGo
	ts.vrbLookup["MOVE"] = enums.ActionTypeMove
	ts.vrbLookup["LOOT"] = enums.ActionTypeLoot
	ts.vrbLookup["DESCRIBE"] = enums.ActionTypeDescribe
	ts.vrbLookup["TAKE"] = enums.ActionTypeTake
	ts.vrbLookup["KICK"] = enums.ActionTypeKick
	ts.vrbLookup["LOCK"] = enums.ActionTypeLock
	ts.vrbLookup["UNLOCK"] = enums.ActionTypeUnlock
	ts.vrbLookup["OPEN"] = enums.ActionTypeOpen
	ts.vrbLookup["LOOK"] = enums.ActionTypeLook
	ts.vrbLookup["CLOSE"] = enums.ActionTypeClose
	ts.vrbLookup["BREAK"] = enums.ActionTypeBreak
	ts.vrbLookup["THROW"] = enums.ActionTypeThrow
	ts.vrbLookup["DROP"] = enums.ActionTypeDrop
	ts.vrbLookup["INVENTORY"] = enums.ActionTypeInventory
	ts.vrbLookup["BURN"] = enums.ActionTypeBurn
	ts.vrbLookup["LIGHT"] = enums.ActionTypeLight
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

func (ts *TokeniserSystem) setupRevDirs() {
	ts.revDirLookup["north"] = enums.DirectionTypeNorth
	ts.revDirLookup["south"] = enums.DirectionTypeSouth
	ts.revDirLookup["east"] = enums.DirectionTypeEast
	ts.revDirLookup["west"] = enums.DirectionTypeWest
	ts.revDirLookup["up"] = enums.DirectionTypeUp
	ts.revDirLookup["down"] = enums.DirectionTypeDown
	ts.revDirLookup["foward"] = enums.DirectionTypeForward
	ts.revDirLookup["backward"] = enums.DirectionTypeBackward
}

// setupDirObjs initializes the directional object lookup table with predefined directional objects
func (ts *TokeniserSystem) setupDirObjs() {
	ts.objLookup["DOOR"] = enums.ObjectTypeDoor
	ts.objLookup["WINDOW"] = enums.ObjectTypeWindow
	ts.objLookup["STAIRS"] = enums.ObjectTypeStairs
	ts.objLookup["LADDER"] = enums.ObjectTypeLadder
	ts.objLookup["PATH"] = enums.ObjectTypePath
	ts.objLookup["TRAIL"] = enums.ObjectTypeTrail
}

// setupVrbAct initializes the verb action response lookup table with predefined responses
func (ts *TokeniserSystem) setupVrbAct() {
	ts.responseLookup[enums.ActionTypeKick] = []enums.ActionType{enums.ActionTypeBreak, enums.ActionTypeHit, enums.ActionTypeDamage}
	ts.responseLookup[enums.ActionTypeBurn] = []enums.ActionType{enums.ActionTypeBurn, enums.ActionTypeLight, enums.ActionTypeDamage}
	ts.responseLookup[enums.ActionTypeLight] = []enums.ActionType{enums.ActionTypeBurn, enums.ActionTypeLight, enums.ActionTypeDamage}
	ts.responseLookup[enums.ActionTypeOpen] = []enums.ActionType{enums.ActionTypeOpen}
	ts.responseLookup[enums.ActionTypeBreak] = []enums.ActionType{enums.ActionTypeBreak}
}

// setupGrammar initializes the grammar response tookup table with predifined grammar
func (ts *TokeniserSystem) setupGrammar() {
	ts.grammarLookup["THE"] = enums.GrammarTypeDefinitionArticle
	ts.grammarLookup["TO"] = enums.GrammarTypePreposition
	ts.grammarLookup["AT"] = enums.GrammarTypePreposition
	ts.grammarLookup["WITH"] = enums.GrammarTypePreposition
	ts.grammarLookup["AROUND"] = enums.GrammarTypeAdverb
}

func (ts *TokeniserSystem) setupMaterial() {
	ts.revMat["wood"] = enums.MaterialTypeWood
	ts.revMat["stone"] = enums.MaterialTypeStone
	ts.revMat["iron"] = enums.MaterialTypeIron
	ts.revMat["shit"] = enums.MaterialTypeShit
	ts.revMat["IKEA"] = enums.MaterialTypeIKEA
	ts.revMat["flesh"] = enums.MaterialTypeFlesh
	ts.revMat["dirt"] = enums.MaterialTypeDirt
	ts.revMat["mud"] = enums.MaterialTypeMud
	ts.revMat["glass"] = enums.MaterialTypeGlass
}

// FishTokens processes the tokenized command and returns VerbData
func (ts *TokeniserSystem) FishTokens(tokens []string) component.VerbData {
	var data component.VerbData // Initialize VerbData to store the result
	lenTokens := len(tokens) - 1

	// Look up the verb, object, and directional object from the tokens
	var VRB enums.ActionType = ts.vrbLookup[(tokens[0])]
	var DObj enums.ObjectType = ts.objLookup[(tokens[lenTokens])]
	var IObj enums.ObjectType = ts.objLookup[(tokens[lenTokens])]

	data.Verb = VRB // Set the verb in VerbData
	data.ErrCode = constants.NOERR

	switch {
	case DObj == enums.ObjectTypeNone && IObj == enums.ObjectTypeNone:
		data.ErrCode = constants.ErrNoDirectObject
	case len(tokens) <= 3:
		switch {
		case DObj != enums.ObjectTypeNone:
			data.DirectObject = DObj
			IObj = enums.ObjectTypeNone
		case DObj == enums.ObjectTypeNone:
			data.ErrCode = constants.ErrNoDirectObject
		}
	case len(tokens) > 3:
		switch {
		case IObj != enums.ObjectTypeNone:
			// We have IOBJ, find DOBJ
			DObj = ts.objLookup[tokens[1]]
			if DObj == enums.ObjectTypeNone {
				DObj = ts.objLookup[tokens[2]]
				if DObj == enums.ObjectTypeNone {
					data.ErrCode = constants.ErrNoDirectObject
				}
			}
		case DObj != enums.ObjectTypeNone:
			// We aren't dealing with this type of structure right now
			// But we have a "throw thing1 at thing2" form where thing2
			// is not a direction object. Probably combat as it goes
			// so for now return
			return data
		}
	}

	// Set the direct noun, indirect directional noun, and error code in VerbData
	data.DirectObject = DObj
	data.IndirectObject = IObj
	//world.Logger().Debug().Msgf("P--->d.dobj:%s iobj:%s vrb:%s\033[0m", data.DirectObject, data.IndirectObject, data.Verb)
	return data
}

// GetResponseForVerb returns the response actions for a given verb
func (ts *TokeniserSystem) GetResponseForVerb(key enums.ActionType) []enums.ActionType {
	if verb, ok := ts.responseLookup[key]; ok {
		return verb
	}
	return nil
}

// GetObjectType returns the ObjectType for a given object key
func (ts *TokeniserSystem) GetObjectType(key string) enums.ObjectType {
	if object, ok := ts.objLookup[key]; ok {
		return object
	}
	return enums.ObjectTypeNone
}

// GetActionType returns the ActionType for a given action key
func (ts *TokeniserSystem) GetActionType(key string) enums.ActionType {
	//return ts.cmdLookup[key]
	if action, ok := ts.vrbLookup[key]; ok {
		return action
	}
	return enums.ActionTypeNone
}

// GetGrammarType returns the GrammarType for a given grammar key
func (ts *TokeniserSystem) GetGrammarType(key string) enums.GrammarType {
	if grammar, ok := ts.grammarLookup[key]; ok {
		return grammar
	}
	return enums.GrammarTypeNone
}

// GetDirectionType returns the DirectionType for a given direction key
func (ts *TokeniserSystem) GetDirectionType(key string) enums.DirectionType {
	if direction, ok := ts.dirLookup[key]; ok {
		return direction
	}
	return enums.DirectionTypeNone
}

// GetMaterial type returns the material type for a given material key
func (ts *TokeniserSystem) GetRevMaterialType(key string) enums.MaterialType {
	if material, ok := ts.revMat[key]; ok {
		return material
	}
	return enums.MaterialTypeNone
}

// GetMaterial type returns the material type for a given material key
func (ts *TokeniserSystem) GetRevDirectionType(key string) enums.DirectionType {
	if revDirection, ok := ts.revDirLookup[key]; ok {
		return revDirection
	}
	return enums.DirectionTypeNone
}
