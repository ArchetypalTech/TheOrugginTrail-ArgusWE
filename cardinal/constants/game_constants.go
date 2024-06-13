package constants

// Define the constants for GameConstants
const (
	// MAX_SIZES for functions
	MAX_TOK uint8 = 16
	MIN_TOK uint8 = 2

	// DATA_BITS for packing
	TERRAIN_BITS uint8 = 24 // << 24
	ROOM_BITS    uint8 = 16 // << 16
	OBJECT_BITS  uint8 = 8  // << 8

	// Direction bits
	NORTH_DIR uint8 = 1 // 0x0001
	EAST_DIR  uint8 = 2 // 0x0010
	SOUTH_DIR uint8 = 4 // 0x0100
	WEST_DIR  uint8 = 8 // 0x1000

	SIZED_AR_SIZE uint32 = 32 // the max size of a Sized Array, 31 items + 1 for count
)

// Define the constants for ErrCodes
const (
	ER_DR_ND                        uint8 = 122 // Error DirectionRoutine DR
	ER_DR_NOP                       uint8 = 123
	ER_PR_ND                        uint8 = 124 // Error ParserRoutine DR
	ER_PR_NT                        uint8 = 125
	ER_PR_TK_CX                     uint8 = 126 // > MAX TOKS
	ER_PR_NOP                       uint8 = 127
	ER_PR_TK_C1                     uint8 = 128 // < MIN TOKS
	ER_PR_NO                        uint8 = 129 // Error No DirectObject
	ER_LK_NOP                       uint8 = 130 // Error Bad Look Command
	ER_AR_BNDS                      uint8 = 131 // Error No DirectObject
	ER_TKPR_NO                      uint8 = 132 // Error No DirectObject
	ER_SIZED_AR_OUT_OF_SPACE        uint8 = 133 // when we try and add an item to a size array using the add function, but its full
	ER_SIZED_AR_NOT_ITEMS_TO_REMOVE uint8 = 134 // when we try and remove an item to a sized array using the remove function, but noone is home!
	ER_ACTION_HDL_NO                uint8 = 135 // Error No Objects to handle
)

// some result codes (from game commands)
const (
	GO_NO_EXIT uint8 = 8 // Error DirectionRoutine DR
	// We use a custom return type for LOOK's
	// because we cant return a 0 for no err unlike the rest of our commands
	LK_RT   uint8 = 9
	AH_BC_0 uint8 = 10
)
