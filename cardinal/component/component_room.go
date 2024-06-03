package component

import (
	"github.com/ArchetypalTech/TheOrugginTrail-ArgusWE/cardinal/enums"
)

type Room struct {
	ID          uint32         `json:"id"`
	Description string         `json:"description"`
	RoomType    enums.RoomType `json:"room_type"`
	ObjectIDs   [32]uint32     `json:"object_ids"`
	DirObjIDs   [32]uint32     `json:"dir_obj_ids"`
}

type RoomStore struct {
	rooms  map[uint32]Room
	nextID uint32
}

func (Room) Name() string {
	return "Room"
}

func (RoomStore) Name() string {
	return "RoomStore"
}

func NewRoomStore() *RoomStore {
	return &RoomStore{
		rooms:  make(map[uint32]Room),
		nextID: 1,
	}
}

func (store *RoomStore) Add(room Room) uint32 {
	room.ID = store.nextID
	store.rooms[store.nextID] = room
	store.nextID++
	return room.ID
}

func (store *RoomStore) Get(id uint32) (Room, bool) {
	room, found := store.rooms[id]
	return room, found
}

func (store *RoomStore) Set(id uint32, room Room) {
	store.rooms[id] = room
}

// Method to update the description of a room
func (store *RoomStore) SetDescription(id uint32, description string) {
	if room, exists := store.rooms[id]; exists {
		room.Description = description
		store.rooms[id] = room
	}
}

// Method to update the room type
func (store *RoomStore) SetRoomType(id uint32, roomType enums.RoomType) {
	if room, exists := store.rooms[id]; exists {
		room.RoomType = roomType
		store.rooms[id] = room
	}
}
