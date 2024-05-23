package component

type ActionOutput struct {
	PlayerID uint32   `json:"player_id"`
	TxtIDs   []string `json:"txt_ids"`
}

func (ActionOutput) Name() string {
	return "ActionOutput"
}
