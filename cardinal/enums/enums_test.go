package enums

import (
	"testing"
)

// Test for RoomType enum
func TestRoomType(t *testing.T) {
	// Test enum to string conversion
	if RoomTypeNone.String() != "none" {
		t.Errorf("RoomTypeNone.String() = %v; want %v", RoomTypeNone.String(), "none")
	}
	if RoomTypeBarn.String() != "barn" {
		t.Errorf("RoomTypeBarn.String() = %v; want %v", RoomTypeBarn.String(), "barn")
	}
	if RoomTypeStore.String() != "store" {
		t.Errorf("RoomTypeStore.String() = %v; want %v", RoomTypeStore.String(), "store")
	}
	if RoomTypeCavern.String() != "cavern" {
		t.Errorf("RoomTypeCavern.String() = %v; want %v", RoomTypeCavern.String(), "cavern")
	}
	if RoomTypeMountainPath.String() != "mountain path" {
		t.Errorf("RoomTypeMountainPath.String() = %v; want %v", RoomTypeMountainPath.String(), "mountain path")
	}
	if RoomTypeFort.String() != "fort" {
		t.Errorf("RoomTypeFort.String() = %v; want %v", RoomTypeFort.String(), "fort")
	}
	if RoomTypeRoom.String() != "room" {
		t.Errorf("RoomTypeRoom.String() = %v; want %v", RoomTypeRoom.String(), "room")
	}
	if RoomTypePlain.String() != "plain" {
		t.Errorf("RoomTypePlain.String() = %v; want %v", RoomTypePlain.String(), "plain")
	}
	if RoomTypeForge.String() != "forge" {
		t.Errorf("RoomTypeForge.String() = %v; want %v", RoomTypeForge.String(), "forge")
	}

	// Test map lookup
	if toStringRoom[RoomTypeNone] != "none" {
		t.Errorf("toStringRoom[RoomTypeNone] = %v; want %v", toStringRoom[RoomTypeNone], "mone")
	}
	if toStringRoom[RoomTypeBarn] != "barn" {
		t.Errorf("toStringRoom[RoomTypeBarn] = %v; want %v", toStringRoom[RoomTypeBarn], "barn")
	}
	if toStringRoom[RoomTypeStore] != "store" {
		t.Errorf("toStringRoom[RoomTypeStore] = %v; want %v", toStringRoom[RoomTypeStore], "store")
	}
	if toStringRoom[RoomTypeCavern] != "cavern" {
		t.Errorf("toStringRoom[RoomTypeCavern] = %v; want %v", toStringRoom[RoomTypeCavern], "cavern")
	}
	if toStringRoom[RoomTypeMountainPath] != "mountain path" {
		t.Errorf("toStringRoom[RoomTypeMountainPath] = %v; want %v", toStringRoom[RoomTypeMountainPath], "mountain path")
	}
	if toStringRoom[RoomTypeFort] != "fort" {
		t.Errorf("toStringRoom[RoomTypeFort] = %v; want %v", toStringRoom[RoomTypeFort], "fort")
	}
	if toStringRoom[RoomTypeRoom] != "room" {
		t.Errorf("toStringRoom[RoomTypeRoom] = %v; want %v", toStringRoom[RoomTypeRoom], "room")
	}
	if toStringRoom[RoomTypePlain] != "plain" {
		t.Errorf("toStringRoom[RoomTypePlain] = %v; want %v", toStringRoom[RoomTypePlain], "plain")
	}
	if toStringRoom[RoomTypeForge] != "forge" {
		t.Errorf("toStringRoom[RoomTypeForge] = %v; want %v", toStringRoom[RoomTypeForge], "forge")
	}

	// Test reverse map lookup
	if toEnumRoom["none"] != RoomTypeNone {
		t.Errorf("toEnumRoom[\"none\"] = %v; want %v", toEnumRoom["none"], RoomTypeNone)
	}
	if toEnumRoom["barn"] != RoomTypeBarn {
		t.Errorf("toEnumRoom[\"barn\"] = %v; want %v", toEnumRoom["barn"], RoomTypeBarn)
	}
	if toEnumRoom["store"] != RoomTypeStore {
		t.Errorf("toEnumRoom[\"store\"] = %v; want %v", toEnumRoom["store"], RoomTypeStore)
	}
	if toEnumRoom["cavern"] != RoomTypeCavern {
		t.Errorf("toEnumRoom[\"cavern\"] = %v; want %v", toEnumRoom["cavern"], RoomTypeCavern)
	}
	if toEnumRoom["mountain path"] != RoomTypeMountainPath {
		t.Errorf("toEnumRoom[\"mountain path\"] = %v; want %v", toEnumRoom["mountain path"], RoomTypeMountainPath)
	}
	if toEnumRoom["fort"] != RoomTypeFort {
		t.Errorf("toEnumRoom[\"fort\"] = %v; want %v", toEnumRoom["fort"], RoomTypeFort)
	}
	if toEnumRoom["room"] != RoomTypeRoom {
		t.Errorf("toEnumRoom[\"room\"] = %v; want %v", toEnumRoom["room"], RoomTypeRoom)
	}
	if toEnumRoom["plain"] != RoomTypePlain {
		t.Errorf("toEnumRoom[\"plain\"] = %v; want %v", toEnumRoom["plain"], RoomTypePlain)
	}
	if toEnumRoom["forge"] != RoomTypeForge {
		t.Errorf("toEnumRoom[\"forge\"] = %v; want %v", toEnumRoom["forge"], RoomTypeForge)
	}
}

// Test for DirectionType enum
func TestDirectionType(t *testing.T) {
	// Test enum to string conversion
	if DirectionTypeNone.String() != "None" {
		t.Errorf("DirectionTypeNone.String() = %v; want %v", DirectionTypeNone.String(), "None")
	}
	if DirectionTypeNorth.String() != "North" {
		t.Errorf("DirectionTypeNorth.String() = %v; want %v", DirectionTypeNorth.String(), "North")
	}
	if DirectionTypeSouth.String() != "South" {
		t.Errorf("DirectionTypeSouth.String() = %v; want %v", DirectionTypeSouth.String(), "South")
	}
	if DirectionTypeEast.String() != "East" {
		t.Errorf("DirectionTypeEast.String() = %v; want %v", DirectionTypeEast.String(), "East")
	}
	if DirectionTypeWest.String() != "West" {
		t.Errorf("DirectionTypeWest.String() = %v; want %v", DirectionTypeWest.String(), "West")
	}
	if DirectionTypeUp.String() != "Up" {
		t.Errorf("DirectionTypeUp.String() = %v; want %v", DirectionTypeUp.String(), "Up")
	}
	if DirectionTypeDown.String() != "Down" {
		t.Errorf("DirectionTypeDown.String() = %v; want %v", DirectionTypeDown.String(), "Down")
	}
	if DirectionTypeForward.String() != "Forward" {
		t.Errorf("DirectionTypeForward.String() = %v; want %v", DirectionTypeForward.String(), "Forward")
	}
	if DirectionTypeBackward.String() != "Backward" {
		t.Errorf("DirectionTypeBackward.String() = %v; want %v", DirectionTypeBackward.String(), "Backward")
	}

	// Test map lookup
	if toStringDirection[DirectionTypeNone] != "None" {
		t.Errorf("toStringDirection[DirectionTypeNone] = %v; want %v", toStringDirection[DirectionTypeNone], "None")
	}
	if toStringDirection[DirectionTypeNorth] != "North" {
		t.Errorf("toStringDirection[DirectionTypeNorth] = %v; want %v", toStringDirection[DirectionTypeNorth], "North")
	}
	if toStringDirection[DirectionTypeSouth] != "South" {
		t.Errorf("toStringDirection[DirectionTypeSouth] = %v; want %v", toStringDirection[DirectionTypeSouth], "South")
	}
	if toStringDirection[DirectionTypeEast] != "East" {
		t.Errorf("toStringDirection[DirectionTypeEast] = %v; want %v", toStringDirection[DirectionTypeEast], "East")
	}
	if toStringDirection[DirectionTypeWest] != "West" {
		t.Errorf("toStringDirection[DirectionTypeWest] = %v; want %v", toStringDirection[DirectionTypeWest], "West")
	}
	if toStringDirection[DirectionTypeUp] != "Up" {
		t.Errorf("toStringDirection[DirectionTypeUp] = %v; want %v", toStringDirection[DirectionTypeUp], "Up")
	}
	if toStringDirection[DirectionTypeDown] != "Down" {
		t.Errorf("toStringDirection[DirectionTypeDown] = %v; want %v", toStringDirection[DirectionTypeDown], "Down")
	}
	if toStringDirection[DirectionTypeForward] != "Forward" {
		t.Errorf("toStringDirection[DirectionTypeForward] = %v; want %v", toStringDirection[DirectionTypeForward], "Forward")
	}
	if toStringDirection[DirectionTypeBackward] != "Backward" {
		t.Errorf("toStringDirection[DirectionTypeBackward] = %v; want %v", toStringDirection[DirectionTypeBackward], "Backward")
	}

	// Test reverse map lookup
	if toEnumDirection["None"] != DirectionTypeNone {
		t.Errorf("toEnumDirection[\"None\"] = %v; want %v", toEnumDirection["None"], DirectionTypeNone)
	}
	if toEnumDirection["North"] != DirectionTypeNorth {
		t.Errorf("toEnumDirection[\"North\"] = %v; want %v", toEnumDirection["North"], DirectionTypeNorth)
	}
	if toEnumDirection["South"] != DirectionTypeSouth {
		t.Errorf("toEnumDirection[\"South\"] = %v; want %v", toEnumDirection["South"], DirectionTypeSouth)
	}
	if toEnumDirection["East"] != DirectionTypeEast {
		t.Errorf("toEnumDirection[\"East\"] = %v; want %v", toEnumDirection["East"], DirectionTypeEast)
	}
	if toEnumDirection["West"] != DirectionTypeWest {
		t.Errorf("toEnumDirection[\"West\"] = %v; want %v", toEnumDirection["West"], DirectionTypeWest)
	}
	if toEnumDirection["Up"] != DirectionTypeUp {
		t.Errorf("toEnumDirection[\"Up\"] = %v; want %v", toEnumDirection["Up"], DirectionTypeUp)
	}
	if toEnumDirection["Down"] != DirectionTypeDown {
		t.Errorf("toEnumDirection[\"Down\"] = %v; want %v", toEnumDirection["Down"], DirectionTypeDown)
	}
	if toEnumDirection["Forward"] != DirectionTypeForward {
		t.Errorf("toEnumDirection[\"Forward\"] = %v; want %v", toEnumDirection["Forward"], DirectionTypeForward)
	}
	if toEnumDirection["Backward"] != DirectionTypeBackward {
		t.Errorf("toEnumDirection[\"Backward\"] = %v; want %v", toEnumDirection["Backward"], DirectionTypeBackward)
	}
}

// Test for ObjectType enum
func TestObjectType(t *testing.T) {
	// Test enum to string conversion
	if ObjectTypeNone.String() != "none" {
		t.Errorf("DirObjectTypeNone.String() = %v; want %v", ObjectTypeNone.String(), "none")
	}
	if ObjectTypeDoor.String() != "door" {
		t.Errorf("DirObjectTypeDoor.String() = %v; want %v", ObjectTypeDoor.String(), "door")
	}
	if ObjectTypeWindow.String() != "window" {
		t.Errorf("DirObjectTypeWindow.String() = %v; want %v", ObjectTypeWindow.String(), "window")
	}
	if ObjectTypeStairs.String() != "stairs" {
		t.Errorf("DirObjectTypeStairs.String() = %v; want %v", ObjectTypeStairs.String(), "stairs")
	}
	if ObjectTypeLadder.String() != "ladder" {
		t.Errorf("DirObjectTypeLadder.String() = %v; want %v", ObjectTypeLadder.String(), "ladder")
	}
	if ObjectTypePath.String() != "path" {
		t.Errorf("DirObjectTypePath.String() = %v; want %v", ObjectTypePath.String(), "path")
	}
	if ObjectTypeTrail.String() != "trail" {
		t.Errorf("DirObjectTypeTrail.String() = %v; want %v", ObjectTypeTrail.String(), "trail")
	}

	// Test map lookup
	if toStringObject[ObjectTypeNone] != "none" {
		t.Errorf("toStringDObject[DirObjectTypeNone] = %v; want %v", toStringObject[ObjectTypeNone], "none")
	}
	if toStringObject[ObjectTypeDoor] != "door" {
		t.Errorf("toStringDObject[DirObjectTypeDoor] = %v; want %v", toStringObject[ObjectTypeDoor], "door")
	}
	if toStringObject[ObjectTypeWindow] != "window" {
		t.Errorf("toStringDObject[DirObjectTypeWindow] = %v; want %v", toStringObject[ObjectTypeWindow], "window")
	}
	if toStringObject[ObjectTypeStairs] != "stairs" {
		t.Errorf("toStringDObject[DirObjectTypeStairs] = %v; want %v", toStringObject[ObjectTypeStairs], "stairs")
	}
	if toStringObject[ObjectTypeLadder] != "ladder" {
		t.Errorf("toStringDObject[DirObjectTypeLadder] = %v; want %v", toStringObject[ObjectTypeLadder], "ladder")
	}
	if toStringObject[ObjectTypePath] != "path" {
		t.Errorf("toStringDObject[DirObjectTypePath] = %v; want %v", toStringObject[ObjectTypePath], "path")
	}
	if toStringObject[ObjectTypeTrail] != "trail" {
		t.Errorf("toStringDObject[DirObjectTypeTrail] = %v; want %v", toStringObject[ObjectTypeTrail], "trail")
	}

	// Test reverse map lookup
	if toEnumObject["none"] != ObjectTypeNone {
		t.Errorf("toEnumDObject[\"none\"] = %v; want %v", toEnumObject["none"], ObjectTypeNone)
	}
	if toEnumObject["door"] != ObjectTypeDoor {
		t.Errorf("toEnumDObject[\"door\"] = %v; want %v", toEnumObject["door"], ObjectTypeDoor)
	}
	if toEnumObject["window"] != ObjectTypeWindow {
		t.Errorf("toEnumDObject[\"window\"] = %v; want %v", toEnumObject["window"], ObjectTypeWindow)
	}
	if toEnumObject["stairs"] != ObjectTypeStairs {
		t.Errorf("toEnumDObject[\"stairs\"] = %v; want %v", toEnumObject["stairs"], ObjectTypeStairs)
	}
	if toEnumObject["ladder"] != ObjectTypeLadder {
		t.Errorf("toEnumDObject[\"ladder\"] = %v; want %v", toEnumObject["ladder"], ObjectTypeLadder)
	}
	if toEnumObject["path"] != ObjectTypePath {
		t.Errorf("toEnumDObject[\"path\"] = %v; want %v", toEnumObject["path"], ObjectTypePath)
	}
	if toEnumObject["trail"] != ObjectTypeTrail {
		t.Errorf("toEnumDObject[\"trail\"] = %v; want %v", toEnumObject["trail"], ObjectTypeTrail)
	}
}

// Test for MaterialType enum
func TestMaterialType(t *testing.T) {
	// Test enum to string conversion
	if MaterialTypeNone.String() != "none" {
		t.Errorf("MaterialTypeNone.String() = %v; want %v", MaterialTypeNone.String(), "none")
	}
	if MaterialTypeWood.String() != "wood" {
		t.Errorf("MaterialTypeWood.String() = %v; want %v", MaterialTypeWood.String(), "wood")
	}
	if MaterialTypeStone.String() != "stone" {
		t.Errorf("MaterialTypeStone.String() = %v; want %v", MaterialTypeStone.String(), "stone")
	}
	if MaterialTypeIron.String() != "iron" {
		t.Errorf("MaterialTypeIron.String() = %v; want %v", MaterialTypeIron.String(), "iron")
	}
	if MaterialTypeShit.String() != "shit" {
		t.Errorf("MaterialTypeShit.String() = %v; want %v", MaterialTypeShit.String(), "shit")
	}
	if MaterialTypeIKEA.String() != "IKEA" {
		t.Errorf("MaterialTypeIKEA.String() = %v; want %v", MaterialTypeIKEA.String(), "IKEA")
	}
	if MaterialTypeFlesh.String() != "flesh" {
		t.Errorf("MaterialTypeFlesh.String() = %v; want %v", MaterialTypeFlesh.String(), "flesh")
	}
	if MaterialTypeDirt.String() != "dirt" {
		t.Errorf("MaterialTypeDirt.String() = %v; want %v", MaterialTypeDirt.String(), "dirt")
	}
	if MaterialTypeMud.String() != "mud" {
		t.Errorf("MaterialTypeMud.String() = %v; want %v", MaterialTypeMud.String(), "mud")
	}
	if MaterialTypeGlass.String() != "glass" {
		t.Errorf("MaterialTypeGlass.String() = %v; want %v", MaterialTypeGlass.String(), "glass")
	}

	// Test map lookup
	if toStringMaterial[MaterialTypeNone] != "none" {
		t.Errorf("toStringMaterial[MaterialTypeNone] = %v; want %v", toStringMaterial[MaterialTypeNone], "none")
	}
	if toStringMaterial[MaterialTypeWood] != "wood" {
		t.Errorf("toStringMaterial[MaterialTypeWood] = %v; want %v", toStringMaterial[MaterialTypeWood], "wood")
	}
	if toStringMaterial[MaterialTypeStone] != "stone" {
		t.Errorf("toStringMaterial[MaterialTypeStone] = %v; want %v", toStringMaterial[MaterialTypeStone], "stone")
	}
	if toStringMaterial[MaterialTypeIron] != "iron" {
		t.Errorf("toStringMaterial[MaterialTypeIron] = %v; want %v", toStringMaterial[MaterialTypeIron], "iron")
	}
	if toStringMaterial[MaterialTypeShit] != "shit" {
		t.Errorf("toStringMaterial[MaterialTypeShit] = %v; want %v", toStringMaterial[MaterialTypeShit], "shit")
	}
	if toStringMaterial[MaterialTypeIKEA] != "IKEA" {
		t.Errorf("toStringMaterial[MaterialTypeIKEA] = %v; want %v", toStringMaterial[MaterialTypeIKEA], "IKEA")
	}
	if toStringMaterial[MaterialTypeFlesh] != "flesh" {
		t.Errorf("toStringMaterial[MaterialTypeFlesh] = %v; want %v", toStringMaterial[MaterialTypeFlesh], "flesh")
	}
	if toStringMaterial[MaterialTypeDirt] != "dirt" {
		t.Errorf("toStringMaterial[MaterialTypeDirt] = %v; want %v", toStringMaterial[MaterialTypeDirt], "dirt")
	}
	if toStringMaterial[MaterialTypeMud] != "mud" {
		t.Errorf("toStringMaterial[MaterialTypeMud] = %v; want %v", toStringMaterial[MaterialTypeMud], "mud")
	}
	if toStringMaterial[MaterialTypeGlass] != "glass" {
		t.Errorf("toStringMaterial[MaterialTypeGlass] = %v; want %v", toStringMaterial[MaterialTypeGlass], "glass")
	}

	// Test reverse map lookup
	if toEnumMaterial["none"] != MaterialTypeNone {
		t.Errorf("toEnumMaterial[\"none\"] = %v; want %v", toEnumMaterial["none"], MaterialTypeNone)
	}
	if toEnumMaterial["wood"] != MaterialTypeWood {
		t.Errorf("toEnumMaterial[\"wood\"] = %v; want %v", toEnumMaterial["wood"], MaterialTypeWood)
	}
	if toEnumMaterial["stone"] != MaterialTypeStone {
		t.Errorf("toEnumMaterial[\"stone\"] = %v; want %v", toEnumMaterial["stone"], MaterialTypeStone)
	}
	if toEnumMaterial["iron"] != MaterialTypeIron {
		t.Errorf("toEnumMaterial[\"iron\"] = %v; want %v", toEnumMaterial["iron"], MaterialTypeIron)
	}
	if toEnumMaterial["shit"] != MaterialTypeShit {
		t.Errorf("toEnumMaterial[\"shit\"] = %v; want %v", toEnumMaterial["shit"], MaterialTypeShit)
	}
	if toEnumMaterial["IKEA"] != MaterialTypeIKEA {
		t.Errorf("toEnumMaterial[\"IKEA\"] = %v; want %v", toEnumMaterial["IKEA"], MaterialTypeIKEA)
	}
	if toEnumMaterial["flesh"] != MaterialTypeFlesh {
		t.Errorf("toEnumMaterial[\"flesh\"] = %v; want %v", toEnumMaterial["flesh"], MaterialTypeFlesh)
	}
	if toEnumMaterial["dirt"] != MaterialTypeDirt {
		t.Errorf("toEnumMaterial[\"dirt\"] = %v; want %v", toEnumMaterial["dirt"], MaterialTypeDirt)
	}
	if toEnumMaterial["mud"] != MaterialTypeMud {
		t.Errorf("toEnumMaterial[\"mud\"] = %v; want %v", toEnumMaterial["mud"], MaterialTypeMud)
	}
	if toEnumMaterial["glass"] != MaterialTypeGlass {
		t.Errorf("toEnumMaterial[\"glass\"] = %v; want %v", toEnumMaterial["glass"], MaterialTypeGlass)
	}
}

// Test for TxtDefType enum
func TestTxtDefType(t *testing.T) {
	// Test enum to string conversion
	if TxtDefTypeNone.String() != "None" {
		t.Errorf("TxtDefTypeNone.String() = %v; want %v", TxtDefTypeNone.String(), "None")
	}
	if TxtDefTypeDirObject.String() != "DirObject" {
		t.Errorf("TxtDefTypeDirObject.String() = %v; want %v", TxtDefTypeDirObject.String(), "DirObject")
	}
	if TxtDefTypeDir.String() != "Dir" {
		t.Errorf("TxtDefTypeDir.String() = %v; want %v", TxtDefTypeDir.String(), "Dir")
	}
	if TxtDefTypePlace.String() != "Place" {
		t.Errorf("TxtDefTypePlace.String() = %v; want %v", TxtDefTypePlace.String(), "Place")
	}
	if TxtDefTypeObject.String() != "Object" {
		t.Errorf("TxtDefTypeObject.String() = %v; want %v", TxtDefTypeObject.String(), "Object")
	}
	if TxtDefTypeAction.String() != "Action" {
		t.Errorf("TxtDefTypeAction.String() = %v; want %v", TxtDefTypeAction.String(), "Action")
	}

	// Test map lookup
	if toStringTxDef[TxtDefTypeNone] != "None" {
		t.Errorf("toStringTxDef[TxtDefTypeNone] = %v; want %v", toStringTxDef[TxtDefTypeNone], "None")
	}
	if toStringTxDef[TxtDefTypeDirObject] != "DirObject" {
		t.Errorf("toStringTxDef[TxtDefTypeDirObject] = %v; want %v", toStringTxDef[TxtDefTypeDirObject], "DirObject")
	}
	if toStringTxDef[TxtDefTypeDir] != "Dir" {
		t.Errorf("toStringTxDef[TxtDefTypeDir] = %v; want %v", toStringTxDef[TxtDefTypeDir], "Dir")
	}
	if toStringTxDef[TxtDefTypePlace] != "Place" {
		t.Errorf("toStringTxDef[TxtDefTypePlace] = %v; want %v", toStringTxDef[TxtDefTypePlace], "Place")
	}
	if toStringTxDef[TxtDefTypeObject] != "Object" {
		t.Errorf("toStringTxDef[TxtDefTypeObject] = %v; want %v", toStringTxDef[TxtDefTypeObject], "Object")
	}
	if toStringTxDef[TxtDefTypeAction] != "Action" {
		t.Errorf("toStringTxDef[TxtDefTypeAction] = %v; want %v", toStringTxDef[TxtDefTypeAction], "Action")
	}

	// Test reverse map lookup
	if toEnumTxDef["None"] != TxtDefTypeNone {
		t.Errorf("toEnumTxDef[\"None\"] = %v; want %v", toEnumTxDef["None"], TxtDefTypeNone)
	}
	if toEnumTxDef["DirObject"] != TxtDefTypeDirObject {
		t.Errorf("toEnumTxDef[\"DirObject\"] = %v; want %v", toEnumTxDef["DirObject"], TxtDefTypeDirObject)
	}
	if toEnumTxDef["Dir"] != TxtDefTypeDir {
		t.Errorf("toEnumTxDef[\"Dir\"] = %v; want %v", toEnumTxDef["Dir"], TxtDefTypeDir)
	}
	if toEnumTxDef["Place"] != TxtDefTypePlace {
		t.Errorf("toEnumTxDef[\"Place\"] = %v; want %v", toEnumTxDef["Place"], TxtDefTypePlace)
	}
	if toEnumTxDef["Object"] != TxtDefTypeObject {
		t.Errorf("toEnumTxDef[\"Object\"] = %v; want %v", toEnumTxDef["Object"], TxtDefTypeObject)
	}
	if toEnumTxDef["Action"] != TxtDefTypeAction {
		t.Errorf("toEnumTxDef[\"Action\"] = %v; want %v", toEnumTxDef["Action"], TxtDefTypeAction)
	}
}
