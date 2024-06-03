package enums

/////////////// ENUMS///////////

// RoomType - Custom type for room types
type RoomType int

const (
	RoomTypeNone       RoomType = iota // Starts at 0
	RoomTypeWoodCabin                  // 1
	RoomTypeStore                      // 2
	RoomTypeCavern                     // 3
	RoomTypeStoneCabin                 // 4
	RoomTypeFort                       // 5
	RoomTypeRoom                       // 6
	RoomTypePlain                      // 7
	RoomTypeForge                      // 8
)

func (r RoomType) String() string {
	return toStringRoom[r]
}

// Map to store the string representations of the enum
var toStringRoom = map[RoomType]string{
	RoomTypeNone:       "None",
	RoomTypeWoodCabin:  "WoodCabin",
	RoomTypeStore:      "Store",
	RoomTypeCavern:     "Cavern",
	RoomTypeStoneCabin: "StoneCabin",
	RoomTypeFort:       "Fort",
	RoomTypeRoom:       "Room",
	RoomTypePlain:      "Plain",
	RoomTypeForge:      "Forge",
}

// Map to store the enum values by string
var toEnumRoom = map[string]RoomType{
	"None":       RoomTypeNone,
	"WoodCabin":  RoomTypeWoodCabin,
	"Store":      RoomTypeStore,
	"Cavern":     RoomTypeCavern,
	"StoneCabin": RoomTypeStoneCabin,
	"Fort":       RoomTypeFort,
	"Room":       RoomTypeRoom,
	"Plain":      RoomTypePlain,
	"Forge":      RoomTypeForge,
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
	DirectionTypeNone:     "None",
	DirectionTypeNorth:    "North",
	DirectionTypeSouth:    "South",
	DirectionTypeEast:     "East",
	DirectionTypeWest:     "West",
	DirectionTypeUp:       "Up",
	DirectionTypeDown:     "Down",
	DirectionTypeForward:  "Forward",
	DirectionTypeBackward: "Backward",
}

// Map to store the enum values by string
var toEnumDirection = map[string]DirectionType{
	"None":     DirectionTypeNone,
	"North":    DirectionTypeNorth,
	"South":    DirectionTypeSouth,
	"East":     DirectionTypeEast,
	"West":     DirectionTypeWest,
	"Up":       DirectionTypeUp,
	"Down":     DirectionTypeDown,
	"Forward":  DirectionTypeForward,
	"Backward": DirectionTypeBackward,
}

// DirObjectType - Custom type for direction object types
type DirObjectType int

const (
	DirObjectTypeNone   DirObjectType = iota // Starts at  0
	DirObjectTypeDoor                        // 1
	DirObjectTypeWindow                      // 2
	DirObjectTypeStairs                      // 3
	DirObjectTypeLadder                      // 4
	DirObjectTypePath                        // 5
	DirObjectTypeTrail                       // 6
)

func (d DirObjectType) String() string {
	return toStringDObject[d]
}

// Map to store the string representations of the enum
var toStringDObject = map[DirObjectType]string{
	DirObjectTypeNone:   "None",
	DirObjectTypeDoor:   "Door",
	DirObjectTypeWindow: "Window",
	DirObjectTypeStairs: "Stairs",
	DirObjectTypeLadder: "Ladder",
	DirObjectTypePath:   "Path",
	DirObjectTypeTrail:  "Trail",
}

// Map to store the enum values by string
var toEnumDObject = map[string]DirObjectType{
	"None":   DirObjectTypeNone,
	"Door":   DirObjectTypeDoor,
	"Window": DirObjectTypeWindow,
	"Stairs": DirObjectTypeStairs,
	"Ladder": DirObjectTypeLadder,
	"Path":   DirObjectTypePath,
	"Trail":  DirObjectTypeTrail,
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
)

func (a ActionType) String() string {
	return toStringAction[a]
}

// Map to store the string representations of the enum
var toStringAction = map[ActionType]string{
	ActionTypeNone:      "None",
	ActionTypeGo:        "Go",
	ActionTypeMove:      "Move",
	ActionTypeLoot:      "Loot",
	ActionTypeDescribe:  "Describe",
	ActionTypeTake:      "Take",
	ActionTypeKick:      "Kick",
	ActionTypeLock:      "Lock",
	ActionTypeUnlock:    "Unlock",
	ActionTypeOpen:      "Open",
	ActionTypeLook:      "Look",
	ActionTypeClose:     "Close",
	ActionTypeBreak:     "Break",
	ActionTypeThrow:     "Throw",
	ActionTypeDrop:      "Drop",
	ActionTypeInventory: "Inventory",
	ActionTypeBurn:      "Burn",
	ActionTypeLight:     "Light",
	ActionTypeDamage:    "Damage",
	ActionTypeHit:       "Hit",
	ActionTypeAcquire:   "Acquire",
}

// Map to store the enum values by string
var toEnumAction = map[string]ActionType{
	"None":      ActionTypeNone,
	"Go":        ActionTypeGo,
	"Move":      ActionTypeMove,
	"Loot":      ActionTypeLoot,
	"Describe":  ActionTypeDescribe,
	"Take":      ActionTypeTake,
	"Kick":      ActionTypeKick,
	"Lock":      ActionTypeLock,
	"Unlock":    ActionTypeUnlock,
	"Open":      ActionTypeOpen,
	"Look":      ActionTypeLook,
	"Close":     ActionTypeClose,
	"Break":     ActionTypeBreak,
	"Throw":     ActionTypeThrow,
	"Drop":      ActionTypeDrop,
	"Inventory": ActionTypeInventory,
	"Burn":      ActionTypeBurn,
	"Light":     ActionTypeLight,
	"Damage":    ActionTypeDamage,
	"Hit":       ActionTypeHit,
	"Acquire":   ActionTypeAcquire,
}

// ObjectType - Custom type for object types
type ObjectType int

const (
	ObjectTypeNone     ObjectType = iota // Starts at 0
	ObjectTypeFootball                   // 1
	ObjectTypeKey                        // 2
	ObjectTypeKnife                      // 3
	ObjectTypeBottle                     // 4
	ObjectTypeStraw                      // 5
	ObjectTypePetrol                     // 6
)

func (o ObjectType) String() string {
	return toStringObject[o]
}

// Map to store the string representations of the enum
var toStringObject = map[ObjectType]string{
	ObjectTypeNone:     "None",
	ObjectTypeFootball: "Football",
	ObjectTypeKey:      "Key",
	ObjectTypeKnife:    "Knife",
	ObjectTypeBottle:   "Bottle",
	ObjectTypeStraw:    "Straw",
	ObjectTypePetrol:   "Petrol",
}

// Map to store the enum values by string
var toEnumObject = map[string]ObjectType{
	"None":     ObjectTypeNone,
	"Football": ObjectTypeFootball,
	"Key":      ObjectTypeKey,
	"Knife":    ObjectTypeKnife,
	"Bottle":   ObjectTypeBottle,
	"Straw":    ObjectTypeStraw,
	"Petrol":   ObjectTypePetrol,
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
	MaterialTypeNone:  "None",
	MaterialTypeWood:  "Wood",
	MaterialTypeStone: "Stone",
	MaterialTypeIron:  "Iron",
	MaterialTypeShit:  "Shit",
	MaterialTypeIKEA:  "IKEA",
	MaterialTypeFlesh: "Flesh",
	MaterialTypeDirt:  "Dirt",
	MaterialTypeMud:   "Mud",
	MaterialTypeGlass: "Glass",
}

// Map to store the enum values by string
var toEnumMaterial = map[string]MaterialType{
	"None":  MaterialTypeNone,
	"Wood":  MaterialTypeWood,
	"Stone": MaterialTypeStone,
	"Iron":  MaterialTypeIron,
	"Shit":  MaterialTypeShit,
	"IKEA":  MaterialTypeIKEA,
	"Flesh": MaterialTypeFlesh,
	"Dirt":  MaterialTypeDirt,
	"Mud":   MaterialTypeMud,
	"Glass": MaterialTypeGlass,
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
