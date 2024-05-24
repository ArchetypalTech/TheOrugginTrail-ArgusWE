/*
The VerbData structure is a fundamental part of the tokeniser system, providing a clear and efficient way to manage action-related data.
It enhances code readability, simplifies data passing, aids in error handling, and allows for future extensibility.
*/

package component

import "github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"

// VerbData represents the data structure for verb-related information
type VerbData struct {
	Verb            enums.ActionType    // The action/verb from the command
	DirectNoun      enums.ObjectType    // The direct object of the action
	IndirectDirNoun enums.DirObjectType // The indirect directional object of the action
	ErrCode         uint8               // Error code for parsing the tokens
}

// Name returns the name of the structure, used for identification
func (VerbData) Name() string {
	return "VerbData"
}
