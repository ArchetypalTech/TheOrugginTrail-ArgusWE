package component

import (
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"
)

type TxtDef struct {
	ID          string           `json:"id"`
	TxtDefType  enums.TxtDefType `json:"txt_def_type"`
	Description string           `json:"description"`
}

type TxtDefStore struct {
	TxtDefs map[string]TxtDef
}

func (TxtDef) Name() string {
	return "RTxtDefoom"
}

func (TxtDefStore) Name() string {
	return "TxtDefStore"
}

func NewTxtDefStore() *TxtDefStore {
	return &TxtDefStore{
		TxtDefs: make(map[string]TxtDef),
	}
}

func (store *TxtDefStore) Set(id string, txtDefType enums.TxtDefType, description string) {
	store.TxtDefs[id] = TxtDef{ID: id, TxtDefType: txtDefType, Description: description}
}

func (store *TxtDefStore) Get(id string) (TxtDef, bool) {
	txtDef, found := store.TxtDefs[id]
	return txtDef, found
}
