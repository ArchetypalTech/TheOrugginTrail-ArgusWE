package component

type Description struct {
	TxtIDs []string `json:"txt_ids"`
}

func (Description) Name() string {
	return "Description"
}
