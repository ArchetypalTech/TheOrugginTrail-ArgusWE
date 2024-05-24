package component

import (
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"
)

type DirObject struct {
	DirObjID        uint32              `json:"dir_obj_id"` // Direction object ID.
	ObjType         enums.DirObjectType `json:"obj_type"`   // The object type: door, window, etc. Use enums.DirObjectType
	DirType         enums.DirectionType `json:"dir_type"`   // The direction it is located: North, South, Up, etc. Use enums.DirectionType
	MatType         enums.MaterialType  `json:"mat_type"`   // The material it is done with: iron, IKEA, etc. Use enums.MaterialType
	DestID          uint32              `json:"dest_id"`    // The destination room ID.
	TxtDefID        string              `json:"txt_def_id"`
	ObjectActionIDs [32]uint32          `json:"object_action_ids"` // The actions to perform on the object: Open, Loct, Breal, etc.
}

func (DirObject) Name() string {
	return "DirObject"
}
