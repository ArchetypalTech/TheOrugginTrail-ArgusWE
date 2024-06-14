package component

import (
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"
)

type Object struct {
	ObjectID        uint32 `json:"id"`
	ObjectName      string
	ObjectType      enums.ObjectType    `json:"object_type"`
	MaterialType    enums.MaterialType  `json:"material_type"`
	DirType         enums.DirectionType `json:"dir_type"`
	DestID          enums.RoomType      `json:"dest_id"`
	TxtDefID        string              `json:"txt_def_id"`
	ObjectActionIDs [32]uint32          `json:"object_action_ids"`
}

type ObjectStore struct {
	objects map[uint32]Object
	nextID  uint32
}

func (Object) Name() string {
	return "Object"
}

func (ObjectStore) Name() string {
	return "ObjectStore"
}

func NewObjectStore() *ObjectStore {
	return &ObjectStore{
		objects: make(map[uint32]Object),
		nextID:  1,
	}
}

func (store *ObjectStore) Add(object Object) uint32 {
	object.ObjectID = store.nextID
	store.objects[store.nextID] = object
	store.nextID++
	return object.ObjectID
}

func (store *ObjectStore) Get(id uint32) (Object, bool) {
	object, found := store.objects[id]
	return object, found
}

func (store *ObjectStore) Set(id uint32, object Object) {
	store.objects[id] = object
}
