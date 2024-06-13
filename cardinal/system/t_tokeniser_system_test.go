package system

import (
	"testing"

	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"
	"github.com/stretchr/testify/assert"
)

func TestNewTokeniserSystem(t *testing.T) {
	ts := NewTokeniserSystem()

	assert.NotNil(t, ts.cmdLookup)
	assert.NotNil(t, ts.dirLookup)
	assert.NotNil(t, ts.dirObjLookup)
	assert.NotNil(t, ts.objLookup)
	assert.NotNil(t, ts.grammarLookup)
	assert.NotNil(t, ts.responseLookup)

	// Check if all predefined values are set correctly
	assert.Equal(t, enums.ActionTypeGo, ts.cmdLookup["GO"])
	assert.Equal(t, enums.ObjectTypeFootball, ts.objLookup["FOOTBALL"])
	assert.Equal(t, enums.DirectionTypeNorth, ts.dirLookup["NORTH"])
	assert.Equal(t, enums.DirObjectTypeDoor, ts.dirObjLookup["DOOR"])
	assert.Equal(t, enums.GrammarTypeDefinitionArticle, ts.grammarLookup["THE"])
	assert.ElementsMatch(t, []enums.ActionType{enums.ActionTypeBreak, enums.ActionTypeHit, enums.ActionTypeDamage}, ts.responseLookup[enums.ActionTypeKick])
}

func TestLookupFunctions(t *testing.T) {
	ts := NewTokeniserSystem()

	assert.Equal(t, enums.ObjectTypeFootball, ts.GetObjectType("FOOTBALL"))
	assert.Equal(t, enums.ActionTypeGo, ts.GetActionType("GO"))
	assert.Equal(t, enums.GrammarTypeDefinitionArticle, ts.GetGrammarType("THE"))
	assert.Equal(t, enums.DirectionTypeNorth, ts.GetDirectionType("NORTH"))
}

func TestFishTokensKickTheBallToTheWindow(t *testing.T) {
	ts := NewTokeniserSystem()

	tokens := []string{"KICK", "THE", "BALL", "TO", "THE", "WINDOW"}
	expectedVerb := enums.ActionTypeKick
	expectedObj := enums.ObjectTypeFootball
	expectedDObj := enums.DirObjectTypeWindow
	expectedIObj := enums.ObjectTypeNone
	expectedErr := uint8(0)

	result := ts.FishTokens(tokens)

	assert.Equal(t, expectedVerb, result.Verb)
	assert.Equal(t, expectedObj, result.DirectNoun)
	assert.Equal(t, expectedDObj, result.IndirectDirNoun)
	assert.Equal(t, expectedIObj, result.IndirectObjNoun)
	assert.Equal(t, expectedErr, result.ErrCode)
}

func TestFishTokensOpenTheDoorToTheKnife(t *testing.T) {
	ts := NewTokeniserSystem()

	tokens := []string{"OPEN", "THE", "DOOR", "WITH", "THE", "IRON", "KNIFE"}
	expectedVerb := enums.ActionTypeOpen
	expectedObj := enums.ObjectTypeNone
	expectedDObj := enums.DirObjectTypeDoor
	expectedIObj := enums.ObjectTypeKnife
	expectedErr := uint8(0)

	result := ts.FishTokens(tokens)

	assert.Equal(t, expectedVerb, result.Verb)
	assert.Equal(t, expectedObj, result.DirectNoun)
	assert.Equal(t, expectedDObj, result.IndirectDirNoun)
	assert.Equal(t, expectedIObj, result.IndirectObjNoun)
	assert.Equal(t, expectedErr, result.ErrCode)
}

func TestGetResponseForVerb(t *testing.T) {
	ts := NewTokeniserSystem()

	assert.ElementsMatch(t, []enums.ActionType{enums.ActionTypeBreak, enums.ActionTypeHit, enums.ActionTypeDamage}, ts.GetResponseForVerb(enums.ActionTypeKick))
	assert.ElementsMatch(t, []enums.ActionType{enums.ActionTypeBurn, enums.ActionTypeLight, enums.ActionTypeDamage}, ts.GetResponseForVerb(enums.ActionTypeBurn))
}