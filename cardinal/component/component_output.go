package component

type Output struct {
	PlayerID uint32 `json:"player_id"`
	Text     string `json:"text"`
}

func (Output) Name() string {
	return "Output"
}
