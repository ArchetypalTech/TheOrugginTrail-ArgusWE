package main

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
)

func (r RoomType) String() string {
	return [...]string{"None", "WoodCabin", "Store", "Cavern", "StoneCabin", "Fort", "Room", "Plain"}[r]
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
	return [...]string{"None", "North", "South", "East", "West", "Up", "Down", "Forward", "Backward"}[d]
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
	return [...]string{"None", "Door", "Window", "Stairs", "Ladder", "Path", "Trail"}[d]
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
	return [...]string{
		"None", "Go", "Move", "Loot", "Describe", "Take", "Kick", "Lock", "Unlock", "Open", "Look", "Close",
		"Break", "Throw", "Drop", "Inventory", "Burn", "Light", "Damage", "Hit", "Acquire",
	}[a]
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
	return [...]string{"None", "Football", "Key", "Knife", "Bottle", "Straw", "Petrol"}[o]
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
	return [...]string{"None", "Wood", "Stone", "Iron", "Shit", "IKEA", "Flesh", "Dirt", "Mud", "Glass"}[m]
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
	return [...]string{"None", "DirObject", "Dir", "Place", "Object", "Action"}[t]
}
