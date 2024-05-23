package component

type Object struct {
	ObjectID        uint32     `json:"object_id"`     // ID of the object.
	ObjectType      int        `json:"object_type"`   // The type it is: Key, Knife, etc. Use enums.ObjectType.
	MaterialType    int        `json:"material_type"` // The material it is done with: Iron, IKEA, etc. Use enums.MaterialType.
	TxtDefID        string     `json:"txt_def_id"`
	ObjectActionIDs [32]uint32 `json:"object_action_ids"`
	Description     string     `json:"description"`
}

func (Object) Name() string {
	return "Object"
}
