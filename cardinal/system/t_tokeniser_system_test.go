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
	assert.Equal(t, enums.ObjectTypeDoor, ts.dirObjLookup["DOOR"])
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

func TestFishTokens1(t *testing.T) {
	ts := NewTokeniserSystem()

	tokens := []string{"KICK", "THE", "BALL", "AT", "THE", "WINDOW"}
	expectedVerb := enums.ActionTypeKick
	expectedDObj := enums.ObjectTypeFootball
	expectedIDirObj := enums.ObjectTypeWindow
	expectedErr := constants.NOERR

	result := ts.FishTokens(tokens)

	assert.Equal(t, expectedVerb, result.Verb)
	assert.Equal(t, expectedDObj, result.DirectObject)
	assert.Equal(t, expectedIDirObj, result.IndirectObject)

	// Check if the error is nil or matches the expected error
	if result.ErrCode != nil {
		assert.Equal(t, expectedErr.Code, result.ErrCode.Code)
		assert.Equal(t, expectedErr.Message, result.ErrCode.Message)
	} else {
		assert.Nil(t, result.ErrCode)
	}
}

func TestFishTokens2(t *testing.T) {
	ts := NewTokeniserSystem()

	tokens := []string{"OPEN", "THE", "DOOR", "WITH", "THE", "KNIFE"}
	expectedVerb := enums.ActionTypeOpen
	expectedDObj := enums.ObjectTypeDoor
	expectedIDirObj := enums.ObjectTypeKnife
	expectedErr := constants.NOERR // Example of using a custom error from constants

	result := ts.FishTokens(tokens)

	assert.Equal(t, expectedVerb, result.Verb)
	assert.Equal(t, expectedDObj, result.DirectObject)
	assert.Equal(t, expectedIDirObj, result.IndirectObject)

	// Check if the error is nil or matches the expected error
	if result.ErrCode != nil {
		assert.Equal(t, expectedErr.Code, result.ErrCode.Code)
		assert.Equal(t, expectedErr.Message, result.ErrCode.Message)
	} else {
		assert.Nil(t, result.ErrCode)
	}
}

func TestFishTokens3(t *testing.T) {
	ts := NewTokeniserSystem()

	tokens := []string{"GO", "NORTH"}
	expectedVerb := enums.ActionTypeGo
	expectedDObj := enums.ObjectTypeNone
	expectedIDirObj := enums.ObjectTypeNone
	expectedErr := constants.ErrNoDirectObject // Example of using a custom error from constants

	result := ts.FishTokens(tokens)

	assert.Equal(t, expectedVerb, result.Verb)
	assert.Equal(t, expectedDObj, result.DirectObject)
	assert.Equal(t, expectedIDirObj, result.IndirectObject)

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
