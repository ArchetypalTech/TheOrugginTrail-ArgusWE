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
	cmdLookup      map[string]enums.ActionType             // Lookup table for command strings to ActionType
	dirLookup      map[string]enums.DirectionType          // Lookup table for direction strings to DirectionType
	dirObjLookup   map[string]enums.DirObjectType          // Lookup table for direction object strings to DirObjectType
	objLookup      map[string]enums.ObjectType             // Lookup table for object strings to ObjectType
	grammarLookup  map[string]enums.GrammarType            // Lookup table for GrammarType
	responseLookup map[enums.ActionType][]enums.ActionType // Lookup table for action responses
}

// NewTokeniserSystem creates a new instance of TokeniserSystem and initializes lookup tables
func NewTokeniserSystem() *TokeniserSystem {
	ts := &TokeniserSystem{
		cmdLookup:      make(map[string]enums.ActionType),
		dirLookup:      make(map[string]enums.DirectionType),
		dirObjLookup:   make(map[string]enums.DirObjectType),
		objLookup:      make(map[string]enums.ObjectType),
		grammarLookup:  make(map[string]enums.GrammarType),
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
	ts.setupGrammar()
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

// setupGrammar initializes the grammar response tookup table with predifined grammar
func (ts *TokeniserSystem) setupGrammar() {
	ts.grammarLookup["THE"] = enums.GrammarTypeDefinitionArticle
	ts.grammarLookup["TO"] = enums.GrammarTypePreposition
	ts.grammarLookup["AT"] = enums.GrammarTypePrepo
	ts.grammarLookup["Around"] = enums.GrammarTypeAdverb
}

// FishTokens processes the tokenized command and returns VerbData
func (ts *TokeniserSystem) FishTokens(tokens []string) component.VerbData {
	var data component.VerbData // Initialize VerbData to store the result
	var err uint8 = 0           // Error code variable
	lenTokens := len(tokens) - 1

	// Look up the verb, object, and directional object from the tokens
	var vrb enums.ActionType = ts.cmdLookup[(tokens[0])]
	var obj enums.ObjectType = ts.objLookup[(tokens[lenTokens])]
	var dobj enums.DirObjectType = ts.dirObjLookup[(tokens[lenTokens])]

	data.Verb = vrb // Set the verb in VerbData
	if obj == enums.ObjectTypeNone && dobj == enums.DirObjectTypeNone {
		data.ErrCode = constants.ER_TKPR_NO
		if isDevelopmentMode() {
			logger.Errorf("\033[31mE--->1err:%d\033[0m", data.ErrCode)
		}
	} else {
		// ? VRB, OBJ ? //
		if obj != enums.ObjectTypeNone && len(tokens) <= 3 {
			data.DirectNoun = obj
		} else if obj == enums.ObjectTypeNone && len(tokens) <= 3 {
			err = constants.ER_TKPR_NO
			if isDevelopmentMode() {
				logger.Errorf("\033[31mE--->2err:%d\033[0m", err)
			}
		}
		if len(tokens) > 3 {
			// ? VRB, [DA], OBJ, IOBJ ? //
			// dirObj ?
			if dobj != enums.DirObjectTypeNone {
				// we have IOBJ find DOBJ
				obj = ts.objLookup[tokens[1]]
				if obj == enums.ObjectTypeNone {
					obj = ts.objLookup[tokens[2]]
					if obj == enums.ObjectTypeNone {
						err = constants.ER_TKPR_NO
						if isDevelopmentMode() {
							logger.Errorf("\033[31mE--->3err:%d\033[0m", err)
						}
					}
				}
			} else if obj != enums.ObjectTypeNone {
				// we arent dealing with this type structure right now
				// but we have a "throw thing1 at thing2" form where thing2
				// is not a direction object. Probably combat as it goes
				// so for now return
				return data
			}
		}
	}

	// Set the direct noun, indirect directional noun, and error code in VerbData
	data.DirectNoun = obj
	data.IndirectDirNoun = dobj
	data.ErrCode = err
	//fmt.Printf("--->d.dobj:%s iobj:%s vrb:%s\n", data.DirectNoun, data.IndirectDirNoun, data.Verb)
	if isDevelopmentMode() {
		logger.Infof("\033[34mP--->d.dobj:%s iobj:%s vrb:%s\033[0m", data.DirectNoun, data.IndirectDirNoun, data.Verb)
	}
	return data
}

// GetResponseForVerb returns the response actions for a given verb
func (ts *TokeniserSystem) GetResponseForVerb(key enums.ActionType) []enums.ActionType {
	return ts.responseLookup[key]
}

// GetObjectType returns the ObjectType for a given object key
func (ts *TokeniserSystem) GetObjectType(key string) enums.ObjectType {
	return ts.objLookup[key]
}

// GetActionType returns the ActionType for a given action key
func (ts *TokeniserSystem) GetActionType(key string) enums.ActionType {
	return ts.cmdLookup[key]
}

// GetGrammarType returns the GrammarType for a given grammar key
func (ts *TokeniserSystem) GetGrammarType(key string) enums.GrammarType {
	return ts.grammarLookup[key]
}

// GetDirectionType returns the DirectionType for a given direction key
func (ts *TokeniserSystem) GetDirectionType(key string) enums.DirectionType {
	return ts.dirLookup[key]
}
