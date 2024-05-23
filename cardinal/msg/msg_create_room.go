package msg

type CreateRoomMsg struct {
	RoomID      uint32 `json:"room_id"`
	RoomType    int    `json:"room_type"` // Use enums.RoomType
	Description string `json:"description"`
}

type CreateRoomReply struct {
	Success bool `json:"success"`
}
