package environment

// Environment is a linked list of environments
type Environment struct {
	values map[string]interface{}
	parent *Environment
}

// NewEnvironment creates a new environment with an empty values map
func NewEnvironment() *Environment {
	return &Environment{values: make(map[string]interface{}), parent: nil}
}

// Extend creates a new child environment for the current environment
func (env *Environment) Extend() *Environment {
	child := &Environment{values: make(map[string]interface{}), parent: env}
	return child
}

// Set sets the value of a variable in the current environment
func (env *Environment) Set_assign(name string, value interface{}) {
	_, ok := env.values[name]
	if !ok {
		env.parent.Set_assign(name, value)
	} else if env.parent != nil {
		panic("Symbol not found!")
	}
	env.values[name] = value
}

func (env *Environment) Set_declare(name string, value interface{}) {
	env.values[name] = value
}

// Get returns the value of a variable in the current environment or its parent environments
func (env *Environment) Get(name string) (interface{}, bool) {
	value, ok := env.values[name]
	if ok {
		return value, true
	} else if env.parent != nil {
		return env.parent.Get(name)
	} else {
		return nil, false
	}
}

var Unassigned = struct {
	tag string
}{
	tag: "unassigned",
}
