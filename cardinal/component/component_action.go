package component

import (
	"sync"

	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"
)

type Action struct {
	ID                 uint32           `json:"id"`
	ActionType         enums.ActionType `json:"action_type"`
	DBitTxt            string           `json:"d_bit_txt"`
	Enabled            bool             `json:"enabled"`
	Revert             bool             `json:"revert"`
	DBit               bool             `json:"d_bit"`
	AffectsActionID    uint32           `json:"affects_action_id"`
	AffectedByActionID uint32           `json:"affected_by_action_id"`
}

type ActionStore struct {
	actions map[uint32]Action
	nextID  uint32
}

var instance *ActionStore
var once sync.Once

func (ActionStore) Name() string {
	return "ActionStore"
}

func (Action) Name() string {
	return "Action"
}

func NewActionStore() *ActionStore {
	once.Do(func() {
		instance = &ActionStore{
			actions: make(map[uint32]Action),
			nextID:  1,
		}
	})
	return instance
}

func GetActionStore() *ActionStore {
	return instance
}

func (store *ActionStore) Add(action Action) uint32 {
	action.ID = store.nextID
	store.actions[store.nextID] = action
	store.nextID++
	return action.ID
}

func (store *ActionStore) Get(id uint32) (Action, bool) {
	action, found := store.actions[id]
	return action, found
}

func (store *ActionStore) Set(id uint32, action Action) {
	store.actions[id] = action
}
