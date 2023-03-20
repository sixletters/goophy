package machine

// Stackframe for RTS to be used with Stack defined in machine
type stackFrame struct {
	tag string
	E   EnvironmentStack
	PC  int
}
