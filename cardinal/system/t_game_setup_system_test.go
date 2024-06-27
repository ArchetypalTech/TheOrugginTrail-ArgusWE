package system

import (
	"fmt"
	"log"
	"testing"

	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/component"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"
	"github.com/stretchr/testify/assert"
	"pkg.world.dev/world-engine/cardinal" //--->MOCKOUT
)

/////MOCKUP ENGINE//// SEE HOW NECESARRY///DEPENDICY INJECTION

// GameSetupInterface defines the methods to be implemented
type GameSetupInterface interface {
	Init()
	setupWorld()
	setupRooms()
	setupPlain()
	createAction(actionType enums.ActionType, desc string, enabled bool, dBit bool, revert bool, affectsID uint32, affectedByID uint32) uint32
	createDirObject(dirType enums.DirectionType, dstID enums.RoomType, dOType enums.ObjectType, mType enums.MaterialType, desc string, actionObjects [32]uint32) uint32
	createObject(objType enums.ObjectType, mType enums.MaterialType, desc, objName string, actionObjects [32]uint32) uint32
	_textGuid(desc string) string
	createPlace(roomID uint32, roomType enums.RoomType, dObjs [32]uint32, objs [32]uint32, tid string)
}

// MockGameSetup implements the GameSetupInterface for testing purposes
type MockGameSetup struct {
	GameSetupInterface
}

func (m *MockGameSetup) Init()       {}
func (m *MockGameSetup) setupWorld() {}
func (m *MockGameSetup) setupRooms() {}
func (m *MockGameSetup) setupPlain() {}
func (m *MockGameSetup) createActionTest(actionType enums.ActionType, desc string, enabled bool, dBit bool, revert bool, affectsID uint32, affectedByID uint32) uint32 {
	return 1
}
func (m *MockGameSetup) createDirObjectTest(dirType enums.DirectionType, dstID enums.RoomType, dOType enums.ObjectType, mType enums.MaterialType, desc string, actionObjects [32]uint32) uint32 {
	return 1
}
func (m *MockGameSetup) createObjectTest(objType enums.ObjectType, mType enums.MaterialType, desc, objName string, actionObjects [32]uint32) uint32 {
	return 1
}
func (m *MockGameSetup) _textGuidTest(desc string) string {
	return "txt-" + desc
}
func (m *MockGameSetup) createPlaceTest(roomID uint32, roomType enums.RoomType, dObjs [32]uint32, objs [32]uint32, tid string) {
}

// NGameSetupSystem initializes the game setup system
func NGameSetupSystemTest(world cardinal.WorldContext) error {
	// Initialization logic
	return nil
}

/*  MOCKUP ENGINE
// MockEngineContext implements the engine.Context interface
type MockEngineContext struct {
	messageErrors map[types.TxHash]error // Store errors associated with transaction hashes
	transactions  map[types.MessageID]types.TxHash
	currentTick   uint64 // Store the current tick of the engine
}

// MockEngineContext creates a new mock engine context for testing
func NewMockEngineContext() engine.Context {
	// Implement the mock engine context
	return &MockEngineContext{
		messageErrors: make(map[types.TxHash]error),
		transactions:  make(map[types.MessageID]types.TxHash),
		currentTick:   0, // Initialize the current tick
	}
}

// AddMessageError implements the engine.Context interface method
func (m *MockEngineContext) AddMessageError(txHash types.TxHash, err error) {
	m.messageErrors[txHash] = err
}

// / AddTransaction implements the engine.Context interface method
func (m *MockEngineContext) AddTransaction(msgID types.MessageID, msg any, tx *sign.Transaction) (uint64, types.TxHash) {
	// Create a dummy transaction hash
	txHash := types.TxHash(fmt.Sprintf("txhash-%d", msgID))
	// Store the transaction hash
	m.transactions[msgID] = txHash
	// Log the transaction
	fmt.Println("AddTransaction called with:", msgID, msg, tx)
	// Return dummy values
	return 1, txHash
}

// CurrentTick implements the engine.Context interface method
func (m *MockEngineContext) CurrentTick() uint64 {
	// Return the current tick
	return m.currentTick
}

// IncrementTick is a method to simulate the passage of ticks in the mock context
func (m *MockEngineContext) IncrementTick() {
	m.currentTick++
}

// EmitEvent implements the engine.Context interface method
func (m *MockEngineContext) EmitEvent(eventData map[string]interface{}) error {
	// Placeholder method
	return nil
}

// EmitStringEvent implements the engine.Context interface method
func (m *MockEngineContext) EmitStringEvent(eventType string) error {
	// Placeholder method
	return nil
}

// GetComponentByName implements the engine.Context interface method
func (m *MockEngineContext) GetComponentByName(componentName string) (types.ComponentMetadata, error) {
	// Placeholder method
	return nil, nil
}

// GetMessageByType implements the engine.Context interface method
func (m *MockEngineContext) GetMessageByType(messageType reflect.Type) (types.Message, bool) {
	// Placeholder method
	return nil, false
}

// GetSignerForPersonaTag implements the engine.Context interface method
func (m *MockEngineContext) GetSignerForPersonaTag(tag string) (sign.SignerComponent, error) {
	// Placeholder method
	return nil, nil
}
*/

// GameSetup represents the game setup
type GameSetupTest struct {
	worldCtx       cardinal.WorldContext
	RoomStore      *component.RoomStore
	DirObjectStore *component.ObjectStore
	ObjectStore    *component.ObjectStore
	ActionStore    *component.ActionStore
	TxtDefStore    *component.TxtDefStore
}

// NewGameSetup creates a new GameSetup instance
func NewGameSetupTest(world cardinal.WorldContext) *GameSetupTest {
	return &GameSetupTest{
		worldCtx:       world,
		RoomStore:      component.NewRoomStore(),
		DirObjectStore: component.NewObjectStore(),
		ObjectStore:    component.NewObjectStore(),
		ActionStore:    component.NewActionStore(),
		TxtDefStore:    component.NewTxtDefStore(),
	}
}

func (s *GameSetupTest) createDirObjectTest(dirType enums.DirectionType, dstID enums.RoomType, dOType enums.ObjectType, mType enums.MaterialType, desc string, actionObjects []uint32) uint32 {
	txtID := s._textGuidTest(desc)
	s.TxtDefStore.Set(txtID, enums.TxtDefTypeDirObject, desc)
	directionObjData := component.Object{
		DirType:         dirType,
		DestID:          dstID,
		ObjectType:      dOType,
		MaterialType:    mType,
		Description:     txtID,
		ObjectActionIDs: actionObjects,
	}
	dirObjID := s.ObjectStore.Add(directionObjData)
	log.Printf("Directional object created - ID: %d, Type: %s, Destination Room Type: %s, Material: %s, Description: %s", dirObjID, dirType, dstID, mType, desc)
	return dirObjID
}

func (s *GameSetupTest) createObjectTest(objType enums.ObjectType, mType enums.MaterialType, desc, objName string, actionObjects []uint32) uint32 {
	txt := s._textGuidTest(desc)
	s.TxtDefStore.Set(txt, enums.TxtDefTypeObject, desc)
	objData := component.Object{
		ObjectType:      objType,
		MaterialType:    mType,
		Description:     txt,
		ObjectActionIDs: actionObjects,
		ObjectName:      objName,
	}
	objID := s.ObjectStore.Add(objData)
	log.Printf("Object created - ID: %d, Type: %s, Material: %s, Description: %s", objID, objType, mType, desc)
	return objID
}

func (s *GameSetupTest) createActionTest(actionType enums.ActionType, desc string, enabled bool, dBit bool, revert bool, affectsID uint32, affectedByID uint32) uint32 {
	txtID := s._textGuidTest(desc)
	actionData := component.Action{ActionType: actionType, DBitTxt: txtID, Enabled: enabled, Revert: revert, DBit: dBit, AffectsActionID: affectsID, AffectedByActionID: affectedByID}
	actionID := s.ActionStore.Add(actionData)
	return actionID
}

func (s *GameSetupTest) _textGuidTest(desc string) string {
	return "txt-" + desc
}

func (s *GameSetupTest) createPlaceTest(roomID uint32, roomType enums.RoomType, dObjs []component.Object, objs []component.Object, tid string) {
	entityID, err := cardinal.Create(s.worldCtx,
		component.Room{
			ID:          roomID,
			Description: tid,
			RoomType:    roomType,
			Objects:     make(map[int]component.Object),
			DirObjs:     make(map[int]component.Object),
		},
	)
	if err != nil {
		log.Fatalf("Failed to create entity for room with ID %d: %v", roomID, err)
	}
	log.Printf("Room with ID %d created successfully, entity ID: %d", roomID, entityID)
	objDescriptions := make([]string, 0, len(objs))
	for _, objID := range objs {
		objDef, _ := s.TxtDefStore.Get(fmt.Sprintf("%v", objID))
		objDescriptions = append(objDescriptions, objDef.Description)
	}
	log.Printf("Objects added to the room: %v", objDescriptions)
	dirObjDescriptions := make([]string, 0, len(dObjs))
	for _, dirObjID := range dObjs {
		dirObjDef, _ := s.TxtDefStore.Get(fmt.Sprintf("%v", dirObjID))
		dirObjDescriptions = append(dirObjDescriptions, dirObjDef.Description)
	}
	log.Printf("Directional objects added to the room: %v", dirObjDescriptions)
}

// TestGameSetup tests the GameSetup system
func TestGameSetup(t *testing.T) {
	//world := NewMockEngineContext()
	//setup := NewGameSetupTest(world)
	mockSetup := &MockGameSetup{}

	//assert.NotNil(t, setup)
	assert.NotNil(t, mockSetup)

	mockSetup.Init()
	mockSetup.setupWorld()
	mockSetup.setupRooms()
	mockSetup.setupPlain()

	actionID := mockSetup.createActionTest(enums.ActionType(1), "Test Action", true, false, true, 0, 0)
	assert.Equal(t, uint32(1), actionID)

	dirObjID := mockSetup.createDirObjectTest(enums.DirectionType(1), enums.RoomType(1), enums.ObjectType(1), enums.MaterialType(1), "Test DirObject", [32]uint32{})
	assert.Equal(t, uint32(1), dirObjID)

	objID := mockSetup.createObjectTest(enums.ObjectType(1), enums.MaterialType(1), "Test Object", "Test Name", [32]uint32{})
	assert.Equal(t, uint32(1), objID)

	txtID := mockSetup._textGuidTest("Test Desc")
	assert.Equal(t, "txt-Test Desc", txtID)

	mockSetup.createPlaceTest(1, enums.RoomType(1), [32]uint32{}, [32]uint32{}, "Test Place")
}
