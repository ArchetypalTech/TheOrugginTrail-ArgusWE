package system

import (
	"testing"

	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/constants"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"
	"github.com/stretchr/testify/assert"
)

func TestNewTokeniserSystem(t *testing.T) {
	ts := NewTokeniserSystem()

	assert.NotNil(t, ts.vrbLookup)
	assert.NotNil(t, ts.dirLookup)
	assert.NotNil(t, ts.dirObjLookup)
	assert.NotNil(t, ts.objLookup)
	assert.NotNil(t, ts.grammarLookup)
	assert.NotNil(t, ts.responseLookup)

	// Check if all predefined values are set correctly
	assert.Equal(t, enums.ActionTypeGo, ts.vrbLookup["GO"])
	assert.Equal(t, enums.ObjectTypeFootball, ts.objLookup["FOOTBALL"])
	assert.Equal(t, enums.DirectionTypeNorth, ts.dirLookup["NORTH"])
	assert.Equal(t, enums.DirObjectTypeDoor, ts.dirObjLookup["DOOR"])
	assert.Equal(t, enums.GrammarTypeDefinitionArticle, ts.grammarLookup["THE"])
	assert.ElementsMatch(t, []enums.ActionType{enums.ActionTypeBreak, enums.ActionTypeHit, enums.ActionTypeDamage}, ts.responseLookup[enums.ActionTypeKick])
}

func TestLookupFunctions(t *testing.T) {
	ts := NewTokeniserSystem()

	assert.Equal(t, enums.ObjectTypeFootball, ts.GetObjectType("FOOTBALL"))
	assert.Equal(t, enums.ActionTypeNone, ts.GetActionType(""))
	assert.Equal(t, enums.GrammarTypeNone, ts.GetGrammarType(""))
	assert.Equal(t, enums.DirectionTypeNorth, ts.GetDirectionType("NORTH"))
}

func TestFishTokensKickTheBallToTheWindow(t *testing.T) {
	ts := NewTokeniserSystem()

	tokens := []string{"KICK", "THE", "BALL", "TO", "THE", "WINDOW"}
	expectedVerb := enums.ActionTypeKick
	expectedDObj := enums.ObjectTypeFootball
	expectedIDirObj := enums.DirObjectTypeWindow
	expectedIObj := enums.ObjectTypeNone
	expectedErr := constants.ErrDirectionRoutineND

	result := ts.FishTokens(tokens)

	assert.Equal(t, expectedVerb, result.Verb)
	assert.Equal(t, expectedDObj, result.DirectObject)
	assert.Equal(t, expectedIDirObj, result.IndirectObject)
	assert.Equal(t, expectedIObj, result.IndirectObjNoun)

	// Check if the error is nil or matches the expected error
	if result.ErrCode != nil {
		assert.Equal(t, expectedErr.Code, result.ErrCode.Code)
		assert.Equal(t, expectedErr.Message, result.ErrCode.Message)
	} else {
		assert.Nil(t, result.ErrCode)
	}
}

func TestFishTokensOpenTheDoorToTheKnife(t *testing.T) {
	ts := NewTokeniserSystem()

	tokens := []string{"OPEN", "THE", "DOOR", "WITH", "THE", "IRON", "KNIFE"}
	expectedVerb := enums.ActionTypeOpen
	expectedDObj := enums.ObjectTypeNone
	expectedIDirObj := enums.DirObjectTypeDoor
	expectedIObj := enums.ObjectTypeKnife
	expectedErr := constants.ErrNoDirectObject // Example of using a custom error from constants

	result := ts.FishTokens(tokens)

	assert.Equal(t, expectedVerb, result.Verb)
	assert.Equal(t, expectedDObj, result.DirectObject)
	assert.Equal(t, expectedIDirObj, result.IndirectObject)
	assert.Equal(t, expectedIObj, result.IndirectObjNoun)

	// Check if the error is nil or matches the expected error
	if result.ErrCode != nil {
		assert.Equal(t, expectedErr.Code, result.ErrCode.Code)
		assert.Equal(t, expectedErr.Message, result.ErrCode.Message)
	} else {
		assert.Nil(t, result.ErrCode)
	}
}

func TestGetResponseForVerb(t *testing.T) {
	ts := NewTokeniserSystem()

	assert.ElementsMatch(t, []enums.ActionType{enums.ActionTypeBreak, enums.ActionTypeHit, enums.ActionTypeDamage}, ts.GetResponseForVerb(enums.ActionTypeKick))
	assert.ElementsMatch(t, []enums.ActionType{enums.ActionTypeBurn, enums.ActionTypeLight, enums.ActionTypeDamage}, ts.GetResponseForVerb(enums.ActionTypeBurn))
}
