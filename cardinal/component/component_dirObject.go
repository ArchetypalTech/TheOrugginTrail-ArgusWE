package component

import (
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"
)

type DirObject struct {
	ID              uint32              `json:"id"`
	DirType         enums.DirectionType `json:"dir_type"`
	DestID          enums.RoomType      `json:"dest_id"`
	ObjType         enums.DirObjectType `json:"obj_type"`
	MatType         enums.MaterialType  `json:"mat_type"`
	TxtDefID        string              `json:"txt_def_id"`
	ObjectActionIDs [32]uint32          `json:"object_action_ids"`
}

type DirObjectStore struct {
	dirObjects map[uint32]DirObject
	nextID     uint32
}

func (DirObject) Name() string {
	return "DirObject"
}

func (DirObjectStore) Name() string {
	return "DirObjectStore"
}

func NewDirObjectStore() *DirObjectStore {
	return &DirObjectStore{
		dirObjects: make(map[uint32]DirObject),
		nextID:     1,
	}
}

func (store *DirObjectStore) Add(dirObject DirObject) uint32 {
	dirObject.ID = store.nextID
	store.dirObjects[store.nextID] = dirObject
	store.nextID++
	return dirObject.ID
}

func (store *DirObjectStore) Get(id uint32) (DirObject, bool) {
	dirObject, found := store.dirObjects[id]
	return dirObject, found
}

func (store *DirObjectStore) Set(id uint32, dirObject DirObject) {
	store.dirObjects[id] = dirObject
}
