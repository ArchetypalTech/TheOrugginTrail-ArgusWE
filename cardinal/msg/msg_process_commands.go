package msg

type ProcessCommandsMsg struct {
	PlayerName string   `json:"PlayerName"` // Name of the player.
	Tokens     []string `json:"Tokens"`     // Array of commands that are sent. They are of type string.
}

type ProcessCommandsReply struct {
	Success bool   `json:"Success"` // Indicates whether the move was successful or not.
	Message string `json:"Message"` // Optional message providing additional information.

}
