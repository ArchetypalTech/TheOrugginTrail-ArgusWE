package component

import (
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"
)

type Action struct {
	ActionID           uint32           `json:"action_id"`
	ActionType         enums.ActionType `json:"action_type"`
	DBitTxt            string           `json:"d_bit_txt"`
	Enabled            bool             `json:"enabled"`
	Revert             bool             `json:"revert"`
	DBit               bool             `json:"d_bit"`
	AffectsActionID    uint32           `json:"affects_action_id"`
	AffectedByActionID uint32           `json:"affected_by_action_id"`
}

func (Action) Name() string {
	return "Action"
}
