package enums

/////////////// ENUMS///////////

// RoomType - Custom type for room types
type RoomType int

const (
	RoomTypeNone         RoomType = iota // Starts at 0
	RoomTypeBarn                         // 1
	RoomTypeStore                        // 2
	RoomTypeCavern                       // 3
	RoomTypeMountainPath                 // 4
	RoomTypeFort                         // 5
	RoomTypeRoom                         // 6
	RoomTypePlain                        // 7
	RoomTypeForge                        // 8
	RoomTypeCellar                       // 9
)

func (r RoomType) String() string {
	return toStringRoom[r]
}

// Map to store the string representations of the enum
var toStringRoom = map[RoomType]string{
	RoomTypeNone:         "none",
	RoomTypeBarn:         "barn",
	RoomTypeStore:        "store",
	RoomTypeCavern:       "cavern",
	RoomTypeMountainPath: "mountain path",
	RoomTypeFort:         "fort",
	RoomTypeRoom:         "room",
	RoomTypePlain:        "plain",
	RoomTypeForge:        "forge",
	RoomTypeCellar:       "cellar",
}

// Map to store the enum values by string
var toEnumRoom = map[string]RoomType{
	"none":          RoomTypeNone,
	"barn":          RoomTypeBarn,
	"store":         RoomTypeStore,
	"cavern":        RoomTypeCavern,
	"mountain path": RoomTypeMountainPath,
	"fort":          RoomTypeFort,
	"room":          RoomTypeRoom,
	"plain":         RoomTypePlain,
	"forge":         RoomTypeForge,
	"cellar":        RoomTypeCellar,
}

// DirectionType - Custom type for direction types
type DirectionType int

const (
	DirectionTypeNone     DirectionType = iota // Starts at 0
	DirectionTypeNorth                         // 1
	DirectionTypeSouth                         // 2
	DirectionTypeEast                          // 3
	DirectionTypeWest                          // 4
	DirectionTypeUp                            // 5
	DirectionTypeDown                          // 6
	DirectionTypeForward                       // 7
	DirectionTypeBackward                      // 8
)

func (d DirectionType) String() string {
	return toStringDirection[d]
}

// Map to store the string representations of the enum
var toStringDirection = map[DirectionType]string{
	DirectionTypeNone:     "none",
	DirectionTypeNorth:    "north",
	DirectionTypeSouth:    "south",
	DirectionTypeEast:     "east",
	DirectionTypeWest:     "west",
	DirectionTypeUp:       "up",
	DirectionTypeDown:     "down",
	DirectionTypeForward:  "forward",
	DirectionTypeBackward: "backward",
}

// Map to store the enum values by string
var toEnumDirection = map[string]DirectionType{
	"none":     DirectionTypeNone,
	"north":    DirectionTypeNorth,
	"south":    DirectionTypeSouth,
	"east":     DirectionTypeEast,
	"west":     DirectionTypeWest,
	"up":       DirectionTypeUp,
	"down":     DirectionTypeDown,
	"forward":  DirectionTypeForward,
	"backward": DirectionTypeBackward,
}

// ActionType - Custom type for action types
type ActionType int

const (
	ActionTypeNone      ActionType = iota // Starts at 0
	ActionTypeGo                          // 1
	ActionTypeMove                        // 2
	ActionTypeLoot                        // 3
	ActionTypeDescribe                    // 4
	ActionTypeTake                        // 5
	ActionTypeKick                        // 6
	ActionTypeLock                        // 7
	ActionTypeUnlock                      // 8
	ActionTypeOpen                        // 9
	ActionTypeLook                        // 10
	ActionTypeClose                       // 11
	ActionTypeBreak                       // 12
	ActionTypeThrow                       // 13
	ActionTypeDrop                        // 14
	ActionTypeInventory                   // 15
	ActionTypeBurn                        // 16
	ActionTypeLight                       // 17
	ActionTypeDamage                      // 18
	ActionTypeHit                         // 19
	ActionTypeAcquire                     // 20
	ActionTypeSniff                       //21
)

func (a ActionType) String() string {
	return toStringAction[a]
}

// Map to store the string representations of the enum
var toStringAction = map[ActionType]string{
	ActionTypeNone:      "none",
	ActionTypeGo:        "go",
	ActionTypeMove:      "move",
	ActionTypeLoot:      "loot",
	ActionTypeDescribe:  "describe",
	ActionTypeTake:      "take",
	ActionTypeKick:      "kick",
	ActionTypeLock:      "lock",
	ActionTypeUnlock:    "unlock",
	ActionTypeOpen:      "open",
	ActionTypeLook:      "look",
	ActionTypeClose:     "close",
	ActionTypeBreak:     "break",
	ActionTypeThrow:     "throw",
	ActionTypeDrop:      "drop",
	ActionTypeInventory: "inventory",
	ActionTypeBurn:      "burn",
	ActionTypeLight:     "light",
	ActionTypeDamage:    "damage",
	ActionTypeHit:       "hit",
	ActionTypeAcquire:   "acquire",
	ActionTypeSniff:     "sniff",
}

// Map to store the enum values by string
var toEnumAction = map[string]ActionType{
	"none":      ActionTypeNone,
	"go":        ActionTypeGo,
	"move":      ActionTypeMove,
	"loot":      ActionTypeLoot,
	"describe":  ActionTypeDescribe,
	"take":      ActionTypeTake,
	"kick":      ActionTypeKick,
	"lock":      ActionTypeLock,
	"unlock":    ActionTypeUnlock,
	"open":      ActionTypeOpen,
	"look":      ActionTypeLook,
	"close":     ActionTypeClose,
	"break":     ActionTypeBreak,
	"rhrow":     ActionTypeThrow,
	"drop":      ActionTypeDrop,
	"inventory": ActionTypeInventory,
	"burn":      ActionTypeBurn,
	"light":     ActionTypeLight,
	"damage":    ActionTypeDamage,
	"hit":       ActionTypeHit,
	"acquire":   ActionTypeAcquire,
	"sniff":     ActionTypeSniff,
}

// ObjectType - Custom type for object types
type ObjectType int

const (
	ObjectTypeNone     ObjectType = iota // Starts at 0
	ObjectTypeFootball                   // 1
	ObjectTypeKey                        // 2
	ObjectTypeKnife                      // 3
	ObjectTypeBottle                     // 4
	ObjectTypeHay                        // 5
	ObjectTypePetrol                     // 6
	ObjectTypeDoor                       // 7
	ObjectTypeWindow                     // 8
	ObjectTypeStairs                     // 9
	ObjectTypeLadder                     // 10
	ObjectTypePath                       // 11
	ObjectTypeTrail                      // 12
	ObjectTypeMatches                    // 13
	ObjectTypeDynamite                   // 14
	ObjectTypeGlue                       // 15
	ObjectTypeBoulder                    // 16
	ObjectTypeTrapdoor                   // 17
)

func (o ObjectType) String() string {
	return toStringObject[o]
}

// Map to store the string representations of the enum
var toStringObject = map[ObjectType]string{
	ObjectTypeNone:     "none",
	ObjectTypeFootball: "football",
	ObjectTypeKey:      "key",
	ObjectTypeKnife:    "knife",
	ObjectTypeBottle:   "bottle",
	ObjectTypeHay:      "hay",
	ObjectTypePetrol:   "petrol",
	ObjectTypeDoor:     "door",
	ObjectTypeWindow:   "window",
	ObjectTypeStairs:   "stairs",
	ObjectTypeLadder:   "ladder",
	ObjectTypePath:     "path",
	ObjectTypeTrail:    "trail",
	ObjectTypeMatches:  "matches",
	ObjectTypeDynamite: "dynamite",
	ObjectTypeGlue:     "glue",
	ObjectTypeBoulder:  "boulder",
	ObjectTypeTrapdoor: "trapdoor",
}

// Map to store the enum values by string
var toEnumObject = map[string]ObjectType{
	"none":     ObjectTypeNone,
	"football": ObjectTypeFootball,
	"key":      ObjectTypeKey,
	"knife":    ObjectTypeKnife,
	"bottle":   ObjectTypeBottle,
	"hay":      ObjectTypeHay,
	"petrol":   ObjectTypePetrol,
	"door":     ObjectTypeDoor,
	"window":   ObjectTypeWindow,
	"stairs":   ObjectTypeStairs,
	"ladder":   ObjectTypeLadder,
	"path":     ObjectTypePath,
	"trail":    ObjectTypeTrail,
	"matches":  ObjectTypeMatches,
	"dynamite": ObjectTypeDynamite,
	"glue":     ObjectTypeGlue,
	"boulder":  ObjectTypeBoulder,
	"trapdoor": ObjectTypeTrapdoor,
}

// MaterialType - Custom type for material types
type MaterialType int

const (
	MaterialTypeNone  MaterialType = iota // Starts at 0
	MaterialTypeWood                      // 1
	MaterialTypeStone                     // 2
	MaterialTypeIron                      // 3
	MaterialTypeShit                      // 4
	MaterialTypeIKEA                      // 5
	MaterialTypeFlesh                     // 6
	MaterialTypeDirt                      // 7
	MaterialTypeMud                       // 8
	MaterialTypeGlass                     // 9
)

func (m MaterialType) String() string {
	return toStringMaterial[m]
}

// Map to store the string representations of the enum
var toStringMaterial = map[MaterialType]string{
	MaterialTypeNone:  "none",
	MaterialTypeWood:  "wood",
	MaterialTypeStone: "stone",
	MaterialTypeIron:  "iron",
	MaterialTypeShit:  "shit",
	MaterialTypeIKEA:  "IKEA",
	MaterialTypeFlesh: "flesh",
	MaterialTypeDirt:  "dirt",
	MaterialTypeMud:   "mud",
	MaterialTypeGlass: "glass",
}

// Map to store the enum values by string
var toEnumMaterial = map[string]MaterialType{
	"none":  MaterialTypeNone,
	"wood":  MaterialTypeWood,
	"stone": MaterialTypeStone,
	"iron":  MaterialTypeIron,
	"shit":  MaterialTypeShit,
	"IKEA":  MaterialTypeIKEA,
	"flesh": MaterialTypeFlesh,
	"dirt":  MaterialTypeDirt,
	"mud":   MaterialTypeMud,
	"glass": MaterialTypeGlass,
}

// TxtDefType - Custom type for text definition types
type TxtDefType int

const (
	TxtDefTypeNone      TxtDefType = iota // Starts at 0
	TxtDefTypeDirObject                   // 1
	TxtDefTypeDir                         // 2
	TxtDefTypePlace                       // 3
	TxtDefTypeObject                      // 4
	TxtDefTypeAction                      // 5
)

func (t TxtDefType) String() string {
	return toStringTxDef[t]
}

// Map to store the string representations of the enum
var toStringTxDef = map[TxtDefType]string{
	TxtDefTypeNone:      "None",
	TxtDefTypeDirObject: "DirObject",
	TxtDefTypeDir:       "Dir",
	TxtDefTypePlace:     "Place",
	TxtDefTypeObject:    "Object",
	TxtDefTypeAction:    "Action",
}

// Map to store the enum values by string
var toEnumTxDef = map[string]TxtDefType{
	"None":      TxtDefTypeNone,
	"DirObject": TxtDefTypeDirObject,
	"Dir":       TxtDefTypeDir,
	"Place":     TxtDefTypePlace,
	"Object":    TxtDefTypeObject,
	"Action":    TxtDefTypeAction,
}

// GrammarType - Custom type for the gramar definition types
type GrammarType int

const (
	GrammarTypeNone              GrammarType = iota // Starts at 0
	GrammarTypeDefinitionArticle                    // 1
	GrammarTypePreposition                          // 2
	GrammarTypeAdverb                               // 3
)

var toStringGrammar = map[GrammarType]string{
	GrammarTypeNone:              "none",
	GrammarTypeDefinitionArticle: "the",
	GrammarTypeAdverb:            "around",
}

var toArrayGrammar = map[GrammarType][]string{
	GrammarTypePreposition: {"to", "at", "with"},
}

var toEnumGrammar = map[string]GrammarType{
	"none":   GrammarTypeNone,
	"the":    GrammarTypeDefinitionArticle,
	"to":     GrammarTypePreposition,
	"at":     GrammarTypePreposition,
	"with":   GrammarTypePreposition,
	"around": GrammarTypeAdverb,
}
