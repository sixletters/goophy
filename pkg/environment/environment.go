package environment

type StackFrame struct {
	vars map[string]interface{}
}

func NewStackFrame() *StackFrame {
	return &StackFrame{vars: make(map[string]interface{})}
}

func (sf *StackFrame) setVar(varName string, value interface{}) {
	sf.vars[varName] = value
}

func (sf *StackFrame) getVar(varName string) (interface{}, bool) {
	value, ok := sf.vars[varName]
	return value, ok
}

type EnvironmentStack struct {
	stackFrames []StackFrame
}

func NewEnvironmentStack() *EnvironmentStack {
	firstFrame := make([]StackFrame, 0)
	return &EnvironmentStack{
		stackFrames: firstFrame,
	}
}

func (env *EnvironmentStack) Extend() {
	env.stackFrames = append(env.stackFrames, *NewStackFrame())
}

func (env *EnvironmentStack) Pop() {
	env.stackFrames = env.stackFrames[:len(env.stackFrames)-2]
}

func (env *EnvironmentStack) Set(varName string, value interface{}) {
	stackFrame := env.stackFrames[len(env.stackFrames)-1]
	stackFrame.setVar(varName, value)
}

func (env *EnvironmentStack) Get(varName string) (interface{}, bool) {
	for i := len(env.stackFrames) - 1; i >= 0; i-- {
		if val, ok := env.stackFrames[i].getVar(varName); ok {
			return val, true
		}

	}
	return nil, false
}
