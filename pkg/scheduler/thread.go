package scheduler

import (
	"cs4215/goophy/pkg/environment"
	util "cs4215/goophy/pkg/util"
)

type ThreadID int32

type Thread struct {
	Os  util.Stack
	Env *environment.Environment
	Pc  int
	Rts util.Stack
}

type ThreadTable map[ThreadID]Thread
