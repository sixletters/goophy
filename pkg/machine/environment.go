package machine


type EnvironmentFrame struct {
	vars map[string]interface{}
}

func NewEnvironmentFrame() *EnvironmentFrame {
	return &EnvironmentFrame{vars: make(map[string]interface{})}
}

func (ef *EnvironmentFrame) setVar(varName string, value interface{}) {
	ef.vars[varName] = value
}

func (ef *EnvironmentFrame) getVar(varName string) (interface{}, bool) {
	value, ok := ef.vars[varName]
	return value, ok
}

type EnvironmentStack struct {
	envFrames []EnvironmentFrame
}

func NewEnvironmentStack() *EnvironmentStack {
	firstFrame := make([]EnvironmentFrame, 0)
	return &EnvironmentStack{
		envFrames: firstFrame,
	}
}

func (env *EnvironmentStack) Extend() {
	env.envFrames = append(env.envFrames, *NewEnvironmentFrame())
}

func (env *EnvironmentStack) Pop() *EnvironmentFrame {
	if len(env.envFrames) == 0 {
		return nil
	}
	popped := env.envFrames[len(env.envFrames)-1]
	env.envFrames = env.envFrames[:len(env.envFrames)-1]
	return &popped
}

func (env *EnvironmentStack) Set(varName string, value interface{}) {
	envFrame := env.envFrames[len(env.envFrames)-1]
	envFrame.setVar(varName, value)
}

func (env *EnvironmentStack) Get(varName string) (interface{}, bool) {
	for i := len(env.envFrames) - 1; i >= 0; i-- {
		if val, ok := env.envFrames[i].getVar(varName); ok {
			return val, true
		}

	}
	return nil, false
}
