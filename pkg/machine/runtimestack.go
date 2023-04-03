package machine

import "cs4215/goophy/pkg/environment"

// Stackframe for RTS to be used with Stack defined in machine
type stackFrame struct {
	tag string
	E   *environment.Environment
	PC  int
}
