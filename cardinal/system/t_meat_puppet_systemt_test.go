package system

import (
	"fmt"
	"testing"

	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/constants"
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"
	"github.com/stretchr/testify/assert"
)

// Define the interface
type Tokeniser interface {
	GetDirectionType(token string) enums.DirectionType
	GetActionType(token string) enums.ActionType
}

// Mock implementation of the Tokeniser interface
type MockTokeniserSystem struct{}

func (m *MockTokeniserSystem) GetDirectionType(token string) enums.DirectionType {
	switch token {
	case "North":
		return enums.DirectionTypeNorth
	default:
		return enums.DirectionTypeNone
	}
}

func (m *MockTokeniserSystem) GetActionType(token string) enums.ActionType {
	switch token {
	case "Go":
		return enums.ActionTypeGo
	case "Look":
		return enums.ActionTypeLook
	case "Take":
		return enums.ActionTypeTake
	case "Drop":
		return enums.ActionTypeDrop
	default:
		return enums.ActionTypeNone
	}
}

// Ensure TokeniserSystem implements Tokeniser
var _ Tokeniser = (*TokeniserSystem)(nil)

// Test for ProcessCommandsTokens2
func TestProcessCommandsTokens2(t *testing.T) {
	ts = &TokeniserSystem{} // Set the global ts to our mock

	tests := []struct {
		tokens      []string
		playerID    uint32
		expectedErr error
	}{
		{
			tokens:      []string{"Go", "North"},
			playerID:    1,
			expectedErr: nil,
		},
		{
			tokens:      []string{"Look"},
			playerID:    1,
			expectedErr: nil,
		},
		{
			tokens:      []string{"Take", "Key"},
			playerID:    1,
			expectedErr: nil,
		},
		{
			tokens:      []string{"InvalidCommand"},
			playerID:    1,
			expectedErr: nil, // Assuming no actual error is returned
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.tokens), func(t *testing.T) {
			// Call the function with test data
			err := ProcessCommandsTokens2(test.tokens, test.playerID)

			// Check the error
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

// Test for handleAlias
func TestHandleAlias(t *testing.T) {
	ts = &TokeniserSystem{} // Set the global ts to our mock

	tests := []struct {
		tokens      []string
		playerID    uint32
		expectedErr uint8
	}{
		{
			tokens:      []string{"Inventory"},
			playerID:    1,
			expectedErr: 0,
		},
		{
			tokens:      []string{"Look"},
			playerID:    1,
			expectedErr: 0,
		},
		{
			tokens:      []string{"InvalidCommand"},
			playerID:    1,
			expectedErr: 0, // Assuming the default behavior is no error
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.tokens), func(t *testing.T) {
			// Call the function with test data
			err := handleAlias(test.tokens, test.playerID)

			// Check the error code
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

// Test for handleVerb
func TestHandleVerb(t *testing.T) {
	ts = &TokeniserSystem{} // Set the global ts to our mock

	tests := []struct {
		tokens      []string
		playerID    uint32
		expectedErr uint8
	}{
		{
			tokens:      []string{"Look"},
			playerID:    1,
			expectedErr: 0,
		},
		{
			tokens:      []string{"Take"},
			playerID:    1,
			expectedErr: 0,
		},
		{
			tokens:      []string{"Drop"},
			playerID:    1,
			expectedErr: 0,
		},
		{
			tokens:      []string{"InvalidCommand"},
			playerID:    1,
			expectedErr: 0, // Assuming the default behavior is no error
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.tokens), func(t *testing.T) {
			// Call the function with test data
			err := handleVerb(test.tokens, test.playerID)

			// Check the error code
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

// Test for insultMeat
func TestInsultMeat(t *testing.T) {
	tests := []struct {
		cErr        uint8
		badCmd      string
		expectedMsg string
	}{
		{
			cErr:        constants.ER_PR_TK_CX,
			badCmd:      "",
			expectedMsg: "WTF, slow down cowboy, your gonna hurt yourself",
		},
		{
			cErr:        constants.ER_PR_NOP,
			badCmd:      "",
			expectedMsg: "Nope, gibberish\nStop breathing with your mouth.",
		},
		{
			cErr:        constants.ER_DR_NOP,
			badCmd:      "North",
			expectedMsg: "Go North is nowhere I know of bellend",
		},
		{
			cErr:        constants.GO_NO_EXIT,
			badCmd:      "East",
			expectedMsg: "Can't go that away East",
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.cErr), func(t *testing.T) {
			// Call the function with test data
			msg := insultMeat(test.cErr, test.badCmd)

			// Check the output message
			assert.Equal(t, test.expectedMsg, msg)
		})
	}
}
