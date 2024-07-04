package component

import (
	"sync"

	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"
	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/types"
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
	Actions map[uint32]Action
	NextID  uint32
}

var instance *ActionStore
var once sync.Once

func (ActionStore) Name() string {
	return "ActionStore"
}

func (Action) Name() string {
	return "Action"
}

func NewActionStore(world cardinal.WorldContext) *ActionStore {

	actionStoreManagerID, err := cardinal.Create(world, ActionStore{
		Actions: make(map[uint32]Action),
		NextID:  1,
	})
	if err != nil {
		world.Logger().Debug().Msgf("Failed to create actionStore entity: %v", err)
	}

	actionStoreManager, err := cardinal.GetComponent[ActionStore](world, actionStoreManagerID)
	if err != nil {
		world.Logger().Error().Msgf("Error getting ActionStore: %v", err)
	}

	return actionStoreManager

}

func GetActionStore(world cardinal.WorldContext) ActionStore {

	var existingActionStore ActionStore
	err := cardinal.NewSearch().Entity(filter.Exact(filter.Component[ActionStore]())).
		Each(world, func(id types.EntityID) bool {
			actionStore, err := cardinal.GetComponent[ActionStore](world, id)
			if err != nil {
				world.Logger().Error().Msgf("GAS: Error getting ActionStore: %v", err)
				return true
			}

			existingActionStore = *actionStore
			return false
		})
	if err != nil {
		world.Logger().Debug().Msgf("Error updating the Player entity: %v when creating", err)
	}

	return existingActionStore
}

func (store *ActionStore) Add(action Action) uint32 {
	action.ID = store.NextID
	store.Actions[store.NextID] = action
	store.NextID++
	return action.ID
}

func (store *ActionStore) Get(id uint32) (Action, bool) {
	action, found := store.Actions[id]
	return action, found
}

func (store *ActionStore) Set(id uint32, action Action) {
	store.Actions[id] = action
}
