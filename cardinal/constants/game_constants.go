package constants

import "fmt"

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

// Define the Custom Error Type
type ErrorTypes struct {
	Code    uint8
	Message string
}

func (e *ErrorTypes) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// Define Error Constants
var (
	NOERR = &ErrorTypes{Code: 0, Message: "NO ERROR"}

	ErrDirectionRoutineND         = &ErrorTypes{Code: 122, Message: "Error DirectionRoutine DR"}
	ErrDirectionRoutineNOP        = &ErrorTypes{Code: 123, Message: "Error DirectionRoutine NOP"}
	ErrParserRoutineND            = &ErrorTypes{Code: 124, Message: "Error ParserRoutine ND"}
	ErrParserRoutineNT            = &ErrorTypes{Code: 125, Message: "Error ParserRoutine NT"}
	ErrParserRoutineTKCX          = &ErrorTypes{Code: 126, Message: "Error ParserRoutine TK CX: > MAX TOKS"}
	ErrParserRoutineNOP           = &ErrorTypes{Code: 127, Message: "Error ParserRoutine NOP"}
	ErrParserRoutineTKC1          = &ErrorTypes{Code: 128, Message: "Error ParserRoutine TK C1: < MIN TOKS"}
	ErrNoDirectObject             = &ErrorTypes{Code: 129, Message: "Error No DirectObject"}
	ErrBadLookCommand             = &ErrorTypes{Code: 130, Message: "Error Bad Look Command"}
	ErrNoDirectObjectLook         = &ErrorTypes{Code: 131, Message: "Error No DirectObject in Look"}
	ErrNoDirectObjectTKPR         = &ErrorTypes{Code: 132, Message: "Error No DirectObject in TKPR"}
	ErrSizedArrayOutOfSpace       = &ErrorTypes{Code: 133, Message: "Error Sized Array Out of Space"}
	ErrSizedArrayNotItemsToRemove = &ErrorTypes{Code: 134, Message: "Error Sized Array Not Items to Remove"}
	ErrNoObjectsToHandle          = &ErrorTypes{Code: 135, Message: "Error No Objects to Handle"}

	ErrNoExit                  = &ErrorTypes{Code: 8, Message: "Error No Exit"}
	ErrLookCustomReturnType    = &ErrorTypes{Code: 9, Message: "Look Custom Return Type"}
	ErrActionHandleBadCommand0 = &ErrorTypes{Code: 10, Message: "Action Handle Bad Command 0"}
)
