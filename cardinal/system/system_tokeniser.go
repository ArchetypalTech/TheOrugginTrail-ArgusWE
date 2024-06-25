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
	vrbLookup        map[string]enums.ActionType             // Lookup table for command strings to ActionType
	dirLookup        map[string]enums.DirectionType          // Lookup table for direction strings to DirectionType
	dirObjLookup     map[string]enums.ObjectType             // Lookup table for directional object strings ObjectType
	objLookup        map[string]enums.ObjectType             // Lookup table for object strings to ObjectType
	grammarLookup    map[string]enums.GrammarType            // Lookup table for GrammarType
	responseLookup   map[enums.ActionType][]enums.ActionType // Lookup table for action responses
	revVrbLookup     map[string]enums.ActionType             // Lookup table for command string to ObjectType in lowercase
	revObjLookup     map[string]enums.ObjectType             //Lookup table for the object string in lowercase
	revDirObjLookup  map[string]enums.ObjectType             // Lookup table for directional object strings ObjectType in lowercase
	revGrammarLookup map[string]enums.GrammarType            // Lookup table for the grammar type in lowercase
	revMat           map[string]enums.MaterialType           // Material type map in lowercase
	revDirLookup     map[string]enums.DirectionType          // ookup table for direction object strings to DirObjectType in lowercase
}

// NewTokeniserSystem creates a new instance of TokeniserSystem and initializes lookup tables
func NewTokeniserSystem() *TokeniserSystem {
	ts := &TokeniserSystem{
		vrbLookup:        make(map[string]enums.ActionType),
		dirLookup:        make(map[string]enums.DirectionType),
		dirObjLookup:     make(map[string]enums.ObjectType),
		objLookup:        make(map[string]enums.ObjectType),
		grammarLookup:    make(map[string]enums.GrammarType),
		responseLookup:   make(map[enums.ActionType][]enums.ActionType),
		revVrbLookup:     make(map[string]enums.ActionType),
		revObjLookup:     make(map[string]enums.ObjectType),
		revDirObjLookup:  make(map[string]enums.ObjectType),
		revGrammarLookup: make(map[string]enums.GrammarType),
		revMat:           make(map[string]enums.MaterialType),
		revDirLookup:     make(map[string]enums.DirectionType),
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

	// lowercase
	ts.revVrbLookup["go"] = enums.ActionTypeGo
	ts.revVrbLookup["move"] = enums.ActionTypeMove
	ts.revVrbLookup["loot"] = enums.ActionTypeLoot
	ts.revVrbLookup["describe"] = enums.ActionTypeDescribe
	ts.revVrbLookup["take"] = enums.ActionTypeTake
	ts.revVrbLookup["kick"] = enums.ActionTypeKick
	ts.revVrbLookup["lock"] = enums.ActionTypeLock
	ts.revVrbLookup["unlock"] = enums.ActionTypeUnlock
	ts.revVrbLookup["open"] = enums.ActionTypeOpen
	ts.revVrbLookup["look"] = enums.ActionTypeLook
	ts.revVrbLookup["close"] = enums.ActionTypeClose
	ts.revVrbLookup["break"] = enums.ActionTypeBreak
	ts.revVrbLookup["throw"] = enums.ActionTypeThrow
	ts.revVrbLookup["drop"] = enums.ActionTypeDrop
	ts.revVrbLookup["inventory"] = enums.ActionTypeInventory
	ts.revVrbLookup["burn"] = enums.ActionTypeBurn
	ts.revVrbLookup["light"] = enums.ActionTypeLight
}

// setupObjects initializes the object lookup table with predefined objects
func (ts *TokeniserSystem) setupObjects() {
	ts.objLookup["FOOTBALL"] = enums.ObjectTypeFootball
	ts.objLookup["BALL"] = enums.ObjectTypeFootball
	ts.objLookup["KEY"] = enums.ObjectTypeKey
	ts.objLookup["KNIFE"] = enums.ObjectTypeKnife
	ts.objLookup["BOTTLE"] = enums.ObjectTypeBottle

	// lowercase
	ts.revObjLookup["football"] = enums.ObjectTypeFootball
	ts.revObjLookup["ball"] = enums.ObjectTypeFootball
	ts.revObjLookup["key"] = enums.ObjectTypeKey
	ts.revObjLookup["knife"] = enums.ObjectTypeKnife
	ts.revObjLookup["bottle"] = enums.ObjectTypeBottle
	ts.revObjLookup["Football"] = enums.ObjectTypeFootball
	ts.revObjLookup["Ball"] = enums.ObjectTypeFootball
	ts.revObjLookup["Key"] = enums.ObjectTypeKey
	ts.revObjLookup["Knife"] = enums.ObjectTypeKnife
	ts.revObjLookup["Bottle"] = enums.ObjectTypeBottle
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

	// lowercase
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

	// lowercase
	ts.revObjLookup["door"] = enums.ObjectTypeDoor
	ts.revObjLookup["window"] = enums.ObjectTypeWindow
	ts.revObjLookup["stairs"] = enums.ObjectTypeStairs
	ts.revObjLookup["ladder"] = enums.ObjectTypeLadder
	ts.revObjLookup["path"] = enums.ObjectTypePath
	ts.revObjLookup["trail"] = enums.ObjectTypeTrail
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

	// lowercase
	ts.revGrammarLookup["the"] = enums.GrammarTypeDefinitionArticle
	ts.revGrammarLookup["to"] = enums.GrammarTypePreposition
	ts.revGrammarLookup["at"] = enums.GrammarTypePreposition
	ts.revGrammarLookup["with"] = enums.GrammarTypePreposition
	ts.revGrammarLookup["aroundD"] = enums.GrammarTypeAdverb
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

// lookupTokenVRB checks both the primary and reverse maps for a given token and returns the corresponding value and a boolean indicating success.
func LookupTokenVRB(primary map[string]enums.ActionType, reverse map[string]enums.ActionType, token string) (enums.ActionType, bool) {
	if value, ok := primary[token]; ok {
		return value, true
	}
	if value, ok := reverse[token]; ok {
		return value, true
	}
	return enums.ActionTypeNone, false
}

// lookupTokenDOBJ checks both the primary and reverse maps for a given token and returns the corresponding value and a boolean indicating success.
func LookupTokenDOBJ(primary map[string]enums.ObjectType, reverse map[string]enums.ObjectType, token string) (enums.ObjectType, bool) {
	if value, ok := primary[token]; ok {
		return value, true
	}
	if value, ok := reverse[token]; ok {
		return value, true
	}
	return enums.ObjectTypeNone, false
}

// lookupTokenIOB checks both the primary and reverse maps for a given token and returns the corresponding value and a boolean indicating success.
func LookupTokenIOBJ(primary map[string]enums.ObjectType, reverse map[string]enums.ObjectType, token string) (enums.ObjectType, bool) {
	if value, ok := primary[token]; ok {
		return value, true
	}
	if value, ok := reverse[token]; ok {
		return value, true
	}
	return enums.ObjectTypeNone, false
}

// FishTokens processes the tokenized command and returns VerbData
func (ts *TokeniserSystem) FishTokens(tokens []string) component.VerbData {
	var data component.VerbData // Initialize VerbData to store the result
	lenTokens := len(tokens) - 1

	// Look up the verb from the tokens
	var VRB enums.ActionType
	VRB, _ = LookupTokenVRB(ts.vrbLookup, ts.revVrbLookup, tokens[0])

	// Look up the direct object from the tokens
	var DObj enums.ObjectType
	DObj, _ = LookupTokenDOBJ(ts.objLookup, ts.revObjLookup, tokens[lenTokens])

	// Look up the indirect object from the tokens
	var IObj enums.ObjectType
	IObj, _ = LookupTokenIOBJ(ts.objLookup, ts.revObjLookup, tokens[lenTokens])

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
			DObj, _ = LookupTokenDOBJ(ts.objLookup, ts.revObjLookup, tokens[1])
			if DObj == enums.ObjectTypeNone {
				DObj, _ = LookupTokenDOBJ(ts.objLookup, ts.revObjLookup, tokens[2])
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
	if object, ok := ts.objLookup[key]; ok { // Upercase
		return object
	} else if revObject, ok := ts.revObjLookup[key]; ok { // lowercase
		return revObject
	}
	return enums.ObjectTypeNone
}

// GetDirectionType returns the DirectionType for a given direction key
func (ts *TokeniserSystem) GetDirectionType(key string) enums.DirectionType {
	if direction, ok := ts.dirLookup[key]; ok { // upercase
		return direction
	} else if revDirection, ok := ts.revDirLookup[key]; ok { // lowercase
		return revDirection
	}
	return enums.DirectionTypeNone
}

// GetActionType returns the ActionType for a given action key
func (ts *TokeniserSystem) GetActionType(key string) enums.ActionType {
	//return ts.cmdLookup[key]
	if action, ok := ts.vrbLookup[key]; ok {
		return action
	} else if revAction, ok := ts.revVrbLookup[key]; ok {
		return revAction
	}
	return enums.ActionTypeNone
}

// GetGrammarType returns the GrammarType for a given grammar key
func (ts *TokeniserSystem) GetGrammarType(key string) enums.GrammarType {
	if grammar, ok := ts.grammarLookup[key]; ok { // Upercase
		return grammar
	} else if revGrammar, ok := ts.revGrammarLookup[key]; ok { // lowercase
		return revGrammar
	}
	return enums.GrammarTypeNone
}

// GetMaterial type returns the material type for a given material key
func (ts *TokeniserSystem) GetRevMaterialType(key string) enums.MaterialType {
	if material, ok := ts.revMat[key]; ok {
		return material
	}
	return enums.MaterialTypeNone
}
