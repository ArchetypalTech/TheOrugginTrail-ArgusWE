package component

import (
	"pkg.world.dev/world-engine/cardinal/types"
)

type Room struct {
	RoomEntityID types.EntityID // Entity ID for the room. This will allows access to it.
	RoomID       uint32         `json:"room_id"`   // Room ID.
	RoomType     main.RoomType  `json:"room_type"` // Type of room.
	Description  string         `json:"description"`
	ObjectIDs    [32]uint32     `json:"object_ids"`  // Objects map that are on the room.
	DirObjIDs    [32]uint32     `json:"dir_obj_ids"` // Objects that allow to go into a direction. They are the "doors".
	Players      [32]uint32     `json:"players"`     // Map of players that are on the room.
}

func (Room) Name() string {
	return "Room"
}
