package enums

import (
	"testing"
)

// Test for RoomType enum
func TestRoomType(t *testing.T) {
	// Test enum to string conversion
	if RoomTypeNone.String() != "None" {
		t.Errorf("RoomTypeNone.String() = %v; want %v", RoomTypeNone.String(), "None")
	}
	if RoomTypeWoodCabin.String() != "WoodCabin" {
		t.Errorf("RoomTypeWoodCabin.String() = %v; want %v", RoomTypeWoodCabin.String(), "WoodCabin")
	}
	if RoomTypeStore.String() != "Store" {
		t.Errorf("RoomTypeStore.String() = %v; want %v", RoomTypeStore.String(), "Store")
	}
	if RoomTypeCavern.String() != "Cavern" {
		t.Errorf("RoomTypeCavern.String() = %v; want %v", RoomTypeCavern.String(), "Cavern")
	}
	if RoomTypeStoneCabin.String() != "StoneCabin" {
		t.Errorf("RoomTypeStoneCabin.String() = %v; want %v", RoomTypeStoneCabin.String(), "StoneCabin")
	}
	if RoomTypeFort.String() != "Fort" {
		t.Errorf("RoomTypeFort.String() = %v; want %v", RoomTypeFort.String(), "Fort")
	}
	if RoomTypeRoom.String() != "Room" {
		t.Errorf("RoomTypeRoom.String() = %v; want %v", RoomTypeRoom.String(), "Room")
	}
	if RoomTypePlain.String() != "Plain" {
		t.Errorf("RoomTypePlain.String() = %v; want %v", RoomTypePlain.String(), "Plain")
	}
	if RoomTypeForge.String() != "Forge" {
		t.Errorf("RoomTypeForge.String() = %v; want %v", RoomTypeForge.String(), "Forge")
	}

	// Test map lookup
	if toStringRoom[RoomTypeNone] != "None" {
		t.Errorf("toStringRoom[RoomTypeNone] = %v; want %v", toStringRoom[RoomTypeNone], "None")
	}
	if toStringRoom[RoomTypeWoodCabin] != "WoodCabin" {
		t.Errorf("toStringRoom[RoomTypeWoodCabin] = %v; want %v", toStringRoom[RoomTypeWoodCabin], "WoodCabin")
	}
	if toStringRoom[RoomTypeStore] != "Store" {
		t.Errorf("toStringRoom[RoomTypeStore] = %v; want %v", toStringRoom[RoomTypeStore], "Store")
	}
	if toStringRoom[RoomTypeCavern] != "Cavern" {
		t.Errorf("toStringRoom[RoomTypeCavern] = %v; want %v", toStringRoom[RoomTypeCavern], "Cavern")
	}
	if toStringRoom[RoomTypeStoneCabin] != "StoneCabin" {
		t.Errorf("toStringRoom[RoomTypeStoneCabin] = %v; want %v", toStringRoom[RoomTypeStoneCabin], "StoneCabin")
	}
	if toStringRoom[RoomTypeFort] != "Fort" {
		t.Errorf("toStringRoom[RoomTypeFort] = %v; want %v", toStringRoom[RoomTypeFort], "Fort")
	}
	if toStringRoom[RoomTypeRoom] != "Room" {
		t.Errorf("toStringRoom[RoomTypeRoom] = %v; want %v", toStringRoom[RoomTypeRoom], "Room")
	}
	if toStringRoom[RoomTypePlain] != "Plain" {
		t.Errorf("toStringRoom[RoomTypePlain] = %v; want %v", toStringRoom[RoomTypePlain], "Plain")
	}
	if toStringRoom[RoomTypeForge] != "Forge" {
		t.Errorf("toStringRoom[RoomTypeForge] = %v; want %v", toStringRoom[RoomTypeForge], "Forge")
	}

	// Test reverse map lookup
	if toEnumRoom["None"] != RoomTypeNone {
		t.Errorf("toEnumRoom[\"None\"] = %v; want %v", toEnumRoom["None"], RoomTypeNone)
	}
	if toEnumRoom["WoodCabin"] != RoomTypeWoodCabin {
		t.Errorf("toEnumRoom[\"WoodCabin\"] = %v; want %v", toEnumRoom["WoodCabin"], RoomTypeWoodCabin)
	}
	if toEnumRoom["Store"] != RoomTypeStore {
		t.Errorf("toEnumRoom[\"Store\"] = %v; want %v", toEnumRoom["Store"], RoomTypeStore)
	}
	if toEnumRoom["Cavern"] != RoomTypeCavern {
		t.Errorf("toEnumRoom[\"Cavern\"] = %v; want %v", toEnumRoom["Cavern"], RoomTypeCavern)
	}
	if toEnumRoom["StoneCabin"] != RoomTypeStoneCabin {
		t.Errorf("toEnumRoom[\"StoneCabin\"] = %v; want %v", toEnumRoom["StoneCabin"], RoomTypeStoneCabin)
	}
	if toEnumRoom["Fort"] != RoomTypeFort {
		t.Errorf("toEnumRoom[\"Fort\"] = %v; want %v", toEnumRoom["Fort"], RoomTypeFort)
	}
	if toEnumRoom["Room"] != RoomTypeRoom {
		t.Errorf("toEnumRoom[\"Room\"] = %v; want %v", toEnumRoom["Room"], RoomTypeRoom)
	}
	if toEnumRoom["Plain"] != RoomTypePlain {
		t.Errorf("toEnumRoom[\"Plain\"] = %v; want %v", toEnumRoom["Plain"], RoomTypePlain)
	}
	if toEnumRoom["Forge"] != RoomTypeForge {
		t.Errorf("toEnumRoom[\"Forge\"] = %v; want %v", toEnumRoom["Forge"], RoomTypeForge)
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

// Test for DirObjectType enum
func TestDirObjectType(t *testing.T) {
	// Test enum to string conversion
	if DirObjectTypeNone.String() != "None" {
		t.Errorf("DirObjectTypeNone.String() = %v; want %v", DirObjectTypeNone.String(), "None")
	}
	if DirObjectTypeDoor.String() != "Door" {
		t.Errorf("DirObjectTypeDoor.String() = %v; want %v", DirObjectTypeDoor.String(), "Door")
	}
	if DirObjectTypeWindow.String() != "Window" {
		t.Errorf("DirObjectTypeWindow.String() = %v; want %v", DirObjectTypeWindow.String(), "Window")
	}
	if DirObjectTypeStairs.String() != "Stairs" {
		t.Errorf("DirObjectTypeStairs.String() = %v; want %v", DirObjectTypeStairs.String(), "Stairs")
	}
	if DirObjectTypeLadder.String() != "Ladder" {
		t.Errorf("DirObjectTypeLadder.String() = %v; want %v", DirObjectTypeLadder.String(), "Ladder")
	}
	if DirObjectTypePath.String() != "Path" {
		t.Errorf("DirObjectTypePath.String() = %v; want %v", DirObjectTypePath.String(), "Path")
	}
	if DirObjectTypeTrail.String() != "Trail" {
		t.Errorf("DirObjectTypeTrail.String() = %v; want %v", DirObjectTypeTrail.String(), "Trail")
	}

	// Test map lookup
	if toStringDObject[DirObjectTypeNone] != "None" {
		t.Errorf("toStringDObject[DirObjectTypeNone] = %v; want %v", toStringDObject[DirObjectTypeNone], "None")
	}
	if toStringDObject[DirObjectTypeDoor] != "Door" {
		t.Errorf("toStringDObject[DirObjectTypeDoor] = %v; want %v", toStringDObject[DirObjectTypeDoor], "Door")
	}
	if toStringDObject[DirObjectTypeWindow] != "Window" {
		t.Errorf("toStringDObject[DirObjectTypeWindow] = %v; want %v", toStringDObject[DirObjectTypeWindow], "Window")
	}
	if toStringDObject[DirObjectTypeStairs] != "Stairs" {
		t.Errorf("toStringDObject[DirObjectTypeStairs] = %v; want %v", toStringDObject[DirObjectTypeStairs], "Stairs")
	}
	if toStringDObject[DirObjectTypeLadder] != "Ladder" {
		t.Errorf("toStringDObject[DirObjectTypeLadder] = %v; want %v", toStringDObject[DirObjectTypeLadder], "Ladder")
	}
	if toStringDObject[DirObjectTypePath] != "Path" {
		t.Errorf("toStringDObject[DirObjectTypePath] = %v; want %v", toStringDObject[DirObjectTypePath], "Path")
	}
	if toStringDObject[DirObjectTypeTrail] != "Trail" {
		t.Errorf("toStringDObject[DirObjectTypeTrail] = %v; want %v", toStringDObject[DirObjectTypeTrail], "Trail")
	}

	// Test reverse map lookup
	if toEnumDObject["None"] != DirObjectTypeNone {
		t.Errorf("toEnumDObject[\"None\"] = %v; want %v", toEnumDObject["None"], DirObjectTypeNone)
	}
	if toEnumDObject["Door"] != DirObjectTypeDoor {
		t.Errorf("toEnumDObject[\"Door\"] = %v; want %v", toEnumDObject["Door"], DirObjectTypeDoor)
	}
	if toEnumDObject["Window"] != DirObjectTypeWindow {
		t.Errorf("toEnumDObject[\"Window\"] = %v; want %v", toEnumDObject["Window"], DirObjectTypeWindow)
	}
	if toEnumDObject["Stairs"] != DirObjectTypeStairs {
		t.Errorf("toEnumDObject[\"Stairs\"] = %v; want %v", toEnumDObject["Stairs"], DirObjectTypeStairs)
	}
	if toEnumDObject["Ladder"] != DirObjectTypeLadder {
		t.Errorf("toEnumDObject[\"Ladder\"] = %v; want %v", toEnumDObject["Ladder"], DirObjectTypeLadder)
	}
	if toEnumDObject["Path"] != DirObjectTypePath {
		t.Errorf("toEnumDObject[\"Path\"] = %v; want %v", toEnumDObject["Path"], DirObjectTypePath)
	}
	if toEnumDObject["Trail"] != DirObjectTypeTrail {
		t.Errorf("toEnumDObject[\"Trail\"] = %v; want %v", toEnumDObject["Trail"], DirObjectTypeTrail)
	}
}

// Test for MaterialType enum
func TestMaterialType(t *testing.T) {
	// Test enum to string conversion
	if MaterialTypeNone.String() != "None" {
		t.Errorf("MaterialTypeNone.String() = %v; want %v", MaterialTypeNone.String(), "None")
	}
	if MaterialTypeWood.String() != "Wood" {
		t.Errorf("MaterialTypeWood.String() = %v; want %v", MaterialTypeWood.String(), "Wood")
	}
	if MaterialTypeStone.String() != "Stone" {
		t.Errorf("MaterialTypeStone.String() = %v; want %v", MaterialTypeStone.String(), "Stone")
	}
	if MaterialTypeIron.String() != "Iron" {
		t.Errorf("MaterialTypeIron.String() = %v; want %v", MaterialTypeIron.String(), "Iron")
	}
	if MaterialTypeShit.String() != "Shit" {
		t.Errorf("MaterialTypeShit.String() = %v; want %v", MaterialTypeShit.String(), "Shit")
	}
	if MaterialTypeIKEA.String() != "IKEA" {
		t.Errorf("MaterialTypeIKEA.String() = %v; want %v", MaterialTypeIKEA.String(), "IKEA")
	}
	if MaterialTypeFlesh.String() != "Flesh" {
		t.Errorf("MaterialTypeFlesh.String() = %v; want %v", MaterialTypeFlesh.String(), "Flesh")
	}
	if MaterialTypeDirt.String() != "Dirt" {
		t.Errorf("MaterialTypeDirt.String() = %v; want %v", MaterialTypeDirt.String(), "Dirt")
	}
	if MaterialTypeMud.String() != "Mud" {
		t.Errorf("MaterialTypeMud.String() = %v; want %v", MaterialTypeMud.String(), "Mud")
	}
	if MaterialTypeGlass.String() != "Glass" {
		t.Errorf("MaterialTypeGlass.String() = %v; want %v", MaterialTypeGlass.String(), "Glass")
	}

	// Test map lookup
	if toStringMaterial[MaterialTypeNone] != "None" {
		t.Errorf("toStringMaterial[MaterialTypeNone] = %v; want %v", toStringMaterial[MaterialTypeNone], "None")
	}
	if toStringMaterial[MaterialTypeWood] != "Wood" {
		t.Errorf("toStringMaterial[MaterialTypeWood] = %v; want %v", toStringMaterial[MaterialTypeWood], "Wood")
	}
	if toStringMaterial[MaterialTypeStone] != "Stone" {
		t.Errorf("toStringMaterial[MaterialTypeStone] = %v; want %v", toStringMaterial[MaterialTypeStone], "Stone")
	}
	if toStringMaterial[MaterialTypeIron] != "Iron" {
		t.Errorf("toStringMaterial[MaterialTypeIron] = %v; want %v", toStringMaterial[MaterialTypeIron], "Iron")
	}
	if toStringMaterial[MaterialTypeShit] != "Shit" {
		t.Errorf("toStringMaterial[MaterialTypeShit] = %v; want %v", toStringMaterial[MaterialTypeShit], "Shit")
	}
	if toStringMaterial[MaterialTypeIKEA] != "IKEA" {
		t.Errorf("toStringMaterial[MaterialTypeIKEA] = %v; want %v", toStringMaterial[MaterialTypeIKEA], "IKEA")
	}
	if toStringMaterial[MaterialTypeFlesh] != "Flesh" {
		t.Errorf("toStringMaterial[MaterialTypeFlesh] = %v; want %v", toStringMaterial[MaterialTypeFlesh], "Flesh")
	}
	if toStringMaterial[MaterialTypeDirt] != "Dirt" {
		t.Errorf("toStringMaterial[MaterialTypeDirt] = %v; want %v", toStringMaterial[MaterialTypeDirt], "Dirt")
	}
	if toStringMaterial[MaterialTypeMud] != "Mud" {
		t.Errorf("toStringMaterial[MaterialTypeMud] = %v; want %v", toStringMaterial[MaterialTypeMud], "Mud")
	}
	if toStringMaterial[MaterialTypeGlass] != "Glass" {
		t.Errorf("toStringMaterial[MaterialTypeGlass] = %v; want %v", toStringMaterial[MaterialTypeGlass], "Glass")
	}

	// Test reverse map lookup
	if toEnumMaterial["None"] != MaterialTypeNone {
		t.Errorf("toEnumMaterial[\"None\"] = %v; want %v", toEnumMaterial["None"], MaterialTypeNone)
	}
	if toEnumMaterial["Wood"] != MaterialTypeWood {
		t.Errorf("toEnumMaterial[\"Wood\"] = %v; want %v", toEnumMaterial["Wood"], MaterialTypeWood)
	}
	if toEnumMaterial["Stone"] != MaterialTypeStone {
		t.Errorf("toEnumMaterial[\"Stone\"] = %v; want %v", toEnumMaterial["Stone"], MaterialTypeStone)
	}
	if toEnumMaterial["Iron"] != MaterialTypeIron {
		t.Errorf("toEnumMaterial[\"Iron\"] = %v; want %v", toEnumMaterial["Iron"], MaterialTypeIron)
	}
	if toEnumMaterial["Shit"] != MaterialTypeShit {
		t.Errorf("toEnumMaterial[\"Shit\"] = %v; want %v", toEnumMaterial["Shit"], MaterialTypeShit)
	}
	if toEnumMaterial["IKEA"] != MaterialTypeIKEA {
		t.Errorf("toEnumMaterial[\"IKEA\"] = %v; want %v", toEnumMaterial["IKEA"], MaterialTypeIKEA)
	}
	if toEnumMaterial["Flesh"] != MaterialTypeFlesh {
		t.Errorf("toEnumMaterial[\"Flesh\"] = %v; want %v", toEnumMaterial["Flesh"], MaterialTypeFlesh)
	}
	if toEnumMaterial["Dirt"] != MaterialTypeDirt {
		t.Errorf("toEnumMaterial[\"Dirt\"] = %v; want %v", toEnumMaterial["Dirt"], MaterialTypeDirt)
	}
	if toEnumMaterial["Mud"] != MaterialTypeMud {
		t.Errorf("toEnumMaterial[\"Mud\"] = %v; want %v", toEnumMaterial["Mud"], MaterialTypeMud)
	}
	if toEnumMaterial["Glass"] != MaterialTypeGlass {
		t.Errorf("toEnumMaterial[\"Glass\"] = %v; want %v", toEnumMaterial["Glass"], MaterialTypeGlass)
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
