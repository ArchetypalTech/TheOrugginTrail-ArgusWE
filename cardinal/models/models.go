/*
The VerbData structure is a fundamental part of the tokeniser system, providing a clear and efficient way to manage action-related data.
It enhances code readability, simplifies data passing, aids in error handling, and allows for future extensibility.
*/

package models

import "github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"

type VerbData struct {
	Verb            enums.ActionType    `json:"verb"`
	DirectNoun      enums.ObjectType    `json:"direct_noun"`
	IndirectDirNoun enums.DirObjectType `json:"indirect_dir_noun"`
	ErrCode         uint8               `json:"err_code"`
}

func (vd VerbData) Name() string {
	return "VerbData"
}
