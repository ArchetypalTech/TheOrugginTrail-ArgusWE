package component

type TxtDef struct {
	TxtDefID   string `json:"txt_def_id"`
	Owner      uint32 `json:"owner"`
	TxtDefType int    `json:"txt_def_type"` // Use enums.TxtDefType
	Value      string `json:"value"`
}

func (TxtDef) Name() string {
	return "TxtDef"
}
