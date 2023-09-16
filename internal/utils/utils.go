package utils

import "time"

// Constant declaration for string symbols.
const (
	Empty = ""
	Space = " "
)

// Constant declaration for wait times.
const (
	WaitBlink  = time.Second * 1
	WaitTiny   = time.Second * 2
	WaitMicro  = time.Second * 5
	WaitSmall  = time.Second * 15
	WaitMedium = time.Second * 30
	WaitLarge  = time.Second * 60
)
